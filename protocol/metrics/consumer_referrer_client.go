package metrics

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lavanet/lava/utils"
)

const (
	referrerName = "referrer"
)

type ReferrerSender interface {
	AppendReferrer(referrer ReferrerRequest)
}

type ConsumerReferrerClient struct {
	*QueueSender
}

func NewReferrerRequest(referrerId string) ReferrerRequest {
	return ReferrerRequest{
		Name:       referrerName,
		ReferrerId: referrerId,
		Count:      1,
	}
}

type ReferrerRequest struct {
	ReferrerId string `json:"referer-id"`
	Name       string `json:"name"`
	Count      uint64 `json:"count"`
}

func (rr ReferrerRequest) String() string {
	rr.Name = reportName
	bytes, err := json.Marshal(rr)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func NewConsumerReferrerClient(endpointAddress string, interval ...time.Duration) *ConsumerReferrerClient {
	if endpointAddress == "" {
		utils.LavaFormatInfo("Running with referrer Server Disabled")
		return nil
	}

	cuc := &ConsumerReferrerClient{
		QueueSender: NewQueueSender(endpointAddress, "ConsumerReferrer", ConsumerReferrerClient{}.aggregation, interval...),
	}
	return cuc
}

func (cuc *ConsumerReferrerClient) AppendReferrer(referrer ReferrerRequest) {
	if cuc == nil {
		return
	}
	cuc.appendQueue(referrer)
}

func (cuc ConsumerReferrerClient) aggregation(aggregate []fmt.Stringer) []fmt.Stringer {
	referrers := map[string]ReferrerRequest{}
	aggregated := []fmt.Stringer{}
	for _, valueToAggregate := range aggregate {
		referrerRequest, ok := valueToAggregate.(ReferrerRequest)
		if !ok {
			// it's something else in the queue
			aggregated = append(aggregated, valueToAggregate)
			continue
		}
		if referrerReq, ok := referrers[referrerRequest.ReferrerId]; ok {
			referrerReq.Count += 1
			referrers[referrerRequest.ReferrerId] = referrerReq
		} else {
			referrers[referrerRequest.ReferrerId] = referrerRequest
		}
	}
	for _, referrerReq := range referrers {
		aggregated = append(aggregated, referrerReq)
	}
	return aggregated
}
