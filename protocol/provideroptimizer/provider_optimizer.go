package provideroptimizer

import (
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/lavanet/lava/protocol/common"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/utils/score"
)

const (
	CacheMaxCost               = 1000 // each item cost would be 1
	CacheNumCounters           = 1000 // expect 100 items
	INITIAL_DATA_STALENESS     = 24
	HALF_LIFE_TIME             = time.Hour
	MAX_HALF_TIME              = 14 * 24 * time.Hour
	PROBE_UPDATE_WEIGHT        = 0.25
	RELAY_UPDATE_WEIGHT        = 1
	DEFAULT_EXPLORATION_CHANCE = 0.1
)

type ConcurrentBlockStore struct {
	Lock  sync.Mutex
	Time  time.Time
	Block uint64
}

type ProviderOptimizer struct {
	strategy                        Strategy
	providersStorage                *ristretto.Cache
	providerRelayStats              *ristretto.Cache // used to decide on the half time of the decay
	averageBlockTime                time.Duration
	baseWorldLatency                time.Duration
	wantedNumProvidersInConcurrency int
	latestSyncData                  ConcurrentBlockStore
}

type ProviderData struct {
	Availability score.ScoreStore // will be used to calculate the probability of error
	Latency      score.ScoreStore // will be used to calculate the latency score
	Sync         score.ScoreStore // will be used to calculate the sync score for spectypes.LATEST_BLOCK/spectypes.NOT_APPLICABLE requests
	SyncBlock    uint64           // will be used to calculate the probability of block error
}

type Strategy int

const (
	STRATEGY_BALANCED Strategy = iota
	STRATEGY_LATENCY
	STRATEGY_SYNC_FRESHNESS
	STRATEGY_COST
	STRATEGY_PRIVACY
	STRATEGY_ACCURACY
)

func (po *ProviderOptimizer) AppendRelayData(providerAddress string, latency time.Duration, isHangingApi bool, success bool, cu uint64, syncBlock uint64) {
	latestSync, timeSync := po.updateLatestSyncData(syncBlock)
	providerData := po.getProviderData(providerAddress)
	halfTime := po.calculateHalfTime(providerAddress)
	providerData = po.updateProbeEntryAvailability(providerData, success, RELAY_UPDATE_WEIGHT, halfTime)
	if success {
		if latency > 0 {
			baseLatency := po.baseWorldLatency + common.BaseTimePerCU(cu)
			providerData = po.updateProbeEntryLatency(providerData, latency, baseLatency, RELAY_UPDATE_WEIGHT, halfTime)
		}
		if syncBlock > providerData.SyncBlock {
			// do not allow providers to go back
			providerData.SyncBlock = syncBlock
		}
		syncLag := po.calculateSyncLag(latestSync, timeSync, providerData.SyncBlock)
		providerData = po.updateProbeEntrySync(providerData, syncLag, po.averageBlockTime, halfTime)
	}
	po.providersStorage.Set(providerAddress, providerData, 1)
	po.updateRelayTime(providerAddress)
}

func (po *ProviderOptimizer) AppendProbeRelayData(providerAddress string, latency time.Duration, success bool) {
	providerData := po.getProviderData(providerAddress)
	halfTime := po.calculateHalfTime(providerAddress)
	providerData = po.updateProbeEntryAvailability(providerData, success, PROBE_UPDATE_WEIGHT, halfTime)
	if success && latency > 0 {
		// base latency for a probe is the world latency
		providerData = po.updateProbeEntryLatency(providerData, latency, po.baseWorldLatency, PROBE_UPDATE_WEIGHT, halfTime)
	}
	po.providersStorage.Set(providerAddress, providerData, 1)
}

// returns a sub set of selected providers according to their scores, perturbation factor will be added to each score in order to randomly select providers that are not always on top
func (po *ProviderOptimizer) ChooseProvider(allAddresses []string, ignoredProviders map[string]struct{}, cu uint64, requestedBlock int64, perturbationPercentage float64) (addresses []string) {
	returnedProviders := make([]string, 1) // location 0 is always the best score
	latencyScore := math.MaxFloat64        // smaller = better i.e less latency
	syncScore := math.MaxFloat64           // smaller = better i.e less sync lag
	for _, providerAddress := range allAddresses {
		if _, ok := ignoredProviders[providerAddress]; ok {
			// ignored provider, skip it
			continue
		}
		providerData := po.getProviderData(providerAddress)

		// latency score
		latencyScoreCurrent := po.calculateLatencyScore(providerData, cu, requestedBlock) // smaller == better i.e less latency
		// latency perturbation
		latencyScoreCurrent = pertrubWithNormalGaussian(latencyScoreCurrent, perturbationPercentage)

		// sync score
		syncScoreCurrent := float64(0)
		if requestedBlock < 0 {
			// means user didn't ask for a specific block and we want to give him the best
			syncScoreCurrent = po.calculateSyncScore(providerData.Sync) // smaller == better i.e less sync lag
			// sync perturbation
			syncScoreCurrent = pertrubWithNormalGaussian(syncScoreCurrent, perturbationPercentage)
		}

		// we want the minimum latency and sync diff
		if po.isBetterProviderScore(latencyScore, latencyScoreCurrent, syncScore, syncScoreCurrent) || len(returnedProviders) == 0 {
			if len(returnedProviders) > 0 && po.shouldExplore(len(returnedProviders)) {
				// we are about to overwrite position 0, and this provider needs a chance to be in exploration
				returnedProviders = append(returnedProviders, returnedProviders[0])
			}
			returnedProviders[0] = providerAddress // best provider is always on position 0
			latencyScore = latencyScoreCurrent
			syncScore = syncScoreCurrent
			continue
		}
		if po.shouldExplore(len(returnedProviders)) {
			returnedProviders = append(returnedProviders, providerAddress)
		}
	}

	return returnedProviders
}

// calculate the expected average time until this provider catches up with the given latestSync block
func (po *ProviderOptimizer) calculateSyncLag(latestSync uint64, timeSync time.Time, providerBlock uint64) time.Duration {
	if latestSync < providerBlock {
		return 0
	}
	timeLag := time.Since(timeSync)                                            // received the latest block at time X, this provider provided the entry at time Y, which is X-Y time after
	blocksGap := time.Duration(latestSync-providerBlock) * po.averageBlockTime // the provider is behind by X blocks, so is expected to catch up in averageBlockTime * X
	timeLag += blocksGap
	return timeLag
}

func (po *ProviderOptimizer) updateLatestSyncData(providerLatestBlock uint64) (uint64, time.Time) {
	po.latestSyncData.Lock.Lock()
	defer po.latestSyncData.Lock.Unlock()
	latestBlock := po.latestSyncData.Block
	if latestBlock < providerLatestBlock {
		// saved latest block is older, so update
		po.latestSyncData.Block = latestBlock
		po.latestSyncData.Time = time.Now()
	}
	return po.latestSyncData.Block, po.latestSyncData.Time
}

func (po *ProviderOptimizer) shouldExplore(currentNumProvders int) bool {
	if currentNumProvders >= po.wantedNumProvidersInConcurrency {
		return false
	}
	explorationChance := DEFAULT_EXPLORATION_CHANCE
	switch po.strategy {
	case STRATEGY_LATENCY:
		return true // we want a lot of parallel tries on latency
	case STRATEGY_ACCURACY:
		return true
	case STRATEGY_COST:
		explorationChance = 0.01
	case STRATEGY_PRIVACY:
		return false // only one at a time
	}
	return rand.Float64() < explorationChance
}

func (po *ProviderOptimizer) isBetterProviderScore(latencyScore float64, latencyScoreCurrent float64, syncScore float64, syncScoreCurrent float64) bool {
	var latencyWeight float64
	switch po.strategy {
	case STRATEGY_LATENCY:
		latencyWeight = 0.9
	case STRATEGY_SYNC_FRESHNESS:
		latencyWeight = 0.2
	case STRATEGY_PRIVACY:
		// pick at random regardless of score
		if rand.Intn(2) == 0 {
			return true
		}
	default:
		latencyWeight = 0.8
	}
	if syncScoreCurrent == 0 {
		return latencyScore > latencyScoreCurrent
	}
	return latencyScore*latencyWeight+syncScore*(1-latencyWeight) > latencyScoreCurrent*latencyWeight+syncScoreCurrent*(1-latencyWeight)
}

func (po *ProviderOptimizer) calculateSyncScore(SyncScore score.ScoreStore) float64 {
	var historicalSyncLatency time.Duration
	if SyncScore.Denom == 0 {
		historicalSyncLatency = 0
	} else {
		historicalSyncLatency = time.Duration(SyncScore.Num/SyncScore.Denom) * po.averageBlockTime // give it units of block time
	}
	return historicalSyncLatency.Seconds()
}

func (po *ProviderOptimizer) calculateLatencyScore(providerData ProviderData, cu uint64, requestedBlock int64) float64 {
	baseLatency := po.baseWorldLatency + common.BaseTimePerCU(cu)/2 // divide by two because the returned time is for timeout not for average
	timeoutDuration := common.GetTimePerCu(cu)
	var historicalLatency time.Duration
	if providerData.Latency.Denom == 0 {
		historicalLatency = baseLatency
	} else {
		historicalLatency = baseLatency * time.Duration(providerData.Latency.Num/providerData.Latency.Denom)
	}
	if historicalLatency > timeoutDuration {
		// can't have a bigger latency than timeout
		historicalLatency = timeoutDuration
	}
	probabilityBlockError := po.CalculateProbabilityOfBlockError(requestedBlock, providerData)
	probabilityOfTimeout := po.CalculateProbabilityOfTimeout(providerData.Availability)
	probabilityOfNoError := (1 - probabilityBlockError) * (1 - probabilityOfTimeout)

	costBlockError := historicalLatency.Seconds() + baseLatency.Seconds()
	costTimeout := timeoutDuration.Seconds() + baseLatency.Seconds()
	costSuccess := historicalLatency.Seconds()

	return probabilityBlockError*costBlockError + probabilityOfTimeout*costTimeout + probabilityOfNoError*costSuccess
}

func (po *ProviderOptimizer) CalculateProbabilityOfTimeout(availabilityScore score.ScoreStore) float64 {
	probabilityTimeout := float64(0)
	if availabilityScore.Denom > 0 { // shouldn't happen since we have default values but protect just in case
		mean := availabilityScore.Num / availabilityScore.Denom
		// bernoulli distribution assumption means probability of '1' is the mean, success is 1
		return 1 - mean
	}
	return probabilityTimeout
}

func (po *ProviderOptimizer) CalculateProbabilityOfBlockError(requestedBlock int64, providerData ProviderData) float64 {
	probabilityBlockError := float64(0)
	if requestedBlock > 0 && providerData.SyncBlock < uint64(requestedBlock) {
		// requested a specific block, so calculate a probability of provider having that block
		averageBlockTime := po.averageBlockTime.Seconds()
		blockDistanceRequired := uint64(requestedBlock) - providerData.SyncBlock
		timeSinceSyncReceived := time.Since(providerData.Sync.Time).Seconds()
		eventRate := timeSinceSyncReceived / averageBlockTime                                   // a new block every average block time, numerator is time passed
		probabilityBlockError = 1 - probValueAfterRepetitions(blockDistanceRequired, eventRate) // we need greater than or equal not less than so complementary probability
	}
	return probabilityBlockError
}

func (po *ProviderOptimizer) getProviderData(providerAddress string) ProviderData {
	var providerData ProviderData

	storedVal, found := po.providersStorage.Get(providerAddress)
	if found {
		var ok bool

		providerData, ok = storedVal.(ProviderData)
		if !ok {
			utils.LavaFormatFatal("invalid usage of optimizer provider storage", nil, utils.Attribute{Key: "storedVal", Value: storedVal})
		}
	} else {
		providerData = ProviderData{
			Availability: score.NewScoreStore(1, 2, time.Now().Add(-1*INITIAL_DATA_STALENESS*time.Hour)), // default value of half score
			Latency:      score.NewScoreStore(2, 1, time.Now().Add(-1*INITIAL_DATA_STALENESS*time.Hour)), // default value of half score (twice the time)
			Sync:         score.NewScoreStore(2, 1, time.Now().Add(-1*INITIAL_DATA_STALENESS*time.Hour)), // default value of half score (twice the time)
			SyncBlock:    0,
		}
	}
	return providerData
}

func (po *ProviderOptimizer) updateProbeEntrySync(providerData ProviderData, sync time.Duration, baseSync time.Duration, halfTime time.Duration) ProviderData {
	newScore := score.NewScoreStore(sync.Seconds(), baseSync.Seconds(), time.Now())
	oldScore := providerData.Sync
	providerData.Sync = score.CalculateTimeDecayFunctionUpdate(oldScore, newScore, halfTime, RELAY_UPDATE_WEIGHT)
	return providerData
}

func (po *ProviderOptimizer) updateProbeEntryAvailability(providerData ProviderData, success bool, weight float64, halfTime time.Duration) ProviderData {
	newNumerator := float64(1)
	if !success {
		// if we failed we need the score update to be 0
		newNumerator = 0
	}
	oldScore := providerData.Availability
	newScore := score.NewScoreStore(newNumerator, 1, time.Now()) // denom is 1, entry time is now
	providerData.Availability = score.CalculateTimeDecayFunctionUpdate(oldScore, newScore, halfTime, weight)
	return providerData
}

// update latency data, base latency is the latency for the api defined in the spec
func (po *ProviderOptimizer) updateProbeEntryLatency(providerData ProviderData, latency time.Duration, baseLatency time.Duration, weight float64, halfTime time.Duration) ProviderData {
	newScore := score.NewScoreStore(latency.Seconds(), baseLatency.Seconds(), time.Now())
	oldScore := providerData.Latency
	providerData.Latency = score.CalculateTimeDecayFunctionUpdate(oldScore, newScore, halfTime, weight)
	return providerData
}

func (po *ProviderOptimizer) updateRelayTime(providerAddress string) {
	times := po.getRelayStatsTimes(providerAddress)
	if len(times) == 0 {
		po.providerRelayStats.Set(providerAddress, []time.Time{time.Now()}, 1)
		return
	}
	times = append(times, time.Now())
	po.providerRelayStats.Set(providerAddress, times, 1)
}

func (po *ProviderOptimizer) calculateHalfTime(providerAddress string) time.Duration {
	halfTime := HALF_LIFE_TIME
	relaysHalfTime := po.getRelayStatsTimeDiff(providerAddress)
	if relaysHalfTime > halfTime {
		halfTime = relaysHalfTime
	}
	if halfTime > MAX_HALF_TIME {
		halfTime = MAX_HALF_TIME
	}
	return halfTime
}

func (po *ProviderOptimizer) getRelayStatsTimeDiff(providerAddress string) time.Duration {
	times := po.getRelayStatsTimes(providerAddress)
	if len(times) == 0 {
		return 0
	}
	return time.Since(times[(len(times)-1)/2])
}

func (po *ProviderOptimizer) getRelayStatsTimes(providerAddress string) []time.Time {
	storedVal, found := po.providerRelayStats.Get(providerAddress)
	if found {
		times, ok := storedVal.([]time.Time)
		if !ok {
			utils.LavaFormatFatal("invalid usage of optimizer relay stats cache", nil, utils.Attribute{Key: "storedVal", Value: storedVal})
		}
		return times
	}
	return nil
}

func NewProviderOptimizer(strategy Strategy, averageBlockTIme time.Duration, baseWorldLatency time.Duration, wantedNumProvidersInConcurrency int) *ProviderOptimizer {
	cache, err := ristretto.NewCache(&ristretto.Config{NumCounters: CacheNumCounters, MaxCost: CacheMaxCost, BufferItems: 64, IgnoreInternalCost: true})
	if err != nil {
		utils.LavaFormatFatal("failed setting up cache for queries", err)
	}
	relayCache, err := ristretto.NewCache(&ristretto.Config{NumCounters: CacheNumCounters, MaxCost: CacheMaxCost, BufferItems: 64, IgnoreInternalCost: true})
	if err != nil {
		utils.LavaFormatFatal("failed setting up cache for queries", err)
	}
	if strategy == STRATEGY_PRIVACY {
		// overwrite
		wantedNumProvidersInConcurrency = 1
	}
	return &ProviderOptimizer{strategy: strategy, providersStorage: cache, averageBlockTime: averageBlockTIme, baseWorldLatency: baseWorldLatency, providerRelayStats: relayCache, wantedNumProvidersInConcurrency: wantedNumProvidersInConcurrency}
}

// calculate the probability a random variable with a poisson distribution
// poisson distribution calculates the probability of K events, in this case the probability enough blocks pass and the request will be accessible in the block
func probValueAfterRepetitions(occurrences uint64, lambda float64) float64 {
	if occurrences > 60 {
		// large values of occurences lose precision so we will use a normal distribution approximation instead
		return logPoisson(occurrences, lambda)
	}
	// calculate probability of observing k events
	prob := (math.Pow(lambda, float64(occurrences)) * math.Exp(-lambda)) / math.Gamma(float64(occurrences)+1)
	return prob
}

func logPoisson(occurrences uint64, lambda float64) float64 {
	logGamma, _ := math.Lgamma(float64(occurrences) + 1)
	logLambda := math.Log(lambda)
	logOcc := math.Log(float64(occurrences))
	logProb := logOcc + logLambda - logGamma
	return math.Exp(logProb)
}

func pertrubWithNormalGaussian(orig float64, percentage float64) float64 {
	return orig + rand.NormFloat64()*percentage*orig
}
