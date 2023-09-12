package rewardserver_test

import (
	"fmt"
	"testing"
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/protocol/rpcprovider/rewardserver"
	"github.com/lavanet/lava/testutil/common"
	"github.com/lavanet/lava/utils/sigs"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	db := rewardserver.NewMemoryDB("specId")
	rs := rewardserver.NewRewardDB()
	err := rs.AddDB(db)
	require.NoError(t, err)

	ctx := sdk.WrapSDKContext(sdk.NewContext(nil, tmproto.Header{}, false, nil))
	proof := common.BuildRelayRequest(ctx, "providerAddr", []byte{}, uint64(0), "specId", nil)

	cpe := &rewardserver.ConsumerProofEntity{
		ConsumerAddr: "consumerAddr",
		ConsumerKey:  "consumerKey",
		Proof:        proof,
	}
	err = rs.Save(cpe)
	require.NoError(t, err)

	rewards, err := rs.FindAll()
	require.NoError(t, err)

	require.Equal(t, 1, len(rewards))
}

func TestSaveBatch(t *testing.T) {
	db := rewardserver.NewMemoryDB("specId")
	rs := rewardserver.NewRewardDB()
	err := rs.AddDB(db)
	require.NoError(t, err)

	ctx := sdk.WrapSDKContext(sdk.NewContext(nil, tmproto.Header{}, false, nil))
	proof := common.BuildRelayRequest(ctx, "providerAddr", []byte{}, uint64(0), "specId", nil)

	cpes := []*rewardserver.ConsumerProofEntity{}
	for i := 0; i < 100; i++ {
		cpes = append(cpes, &rewardserver.ConsumerProofEntity{
			ConsumerAddr: fmt.Sprintf("consumerAddr%d", i),
			ConsumerKey:  fmt.Sprintf("consumerKey%d", i),
			Proof:        proof,
		})
	}
	err = rs.BatchSave(cpes)
	require.NoError(t, err)

	rewards, err := rs.FindAll()
	require.NoError(t, err)

	require.Equal(t, 1, len(rewards))
}

func TestFindAll(t *testing.T) {
	db := rewardserver.NewMemoryDB("specId")
	rs := rewardserver.NewRewardDB()
	err := rs.AddDB(db)
	require.NoError(t, err)

	ctx := sdk.WrapSDKContext(sdk.NewContext(nil, tmproto.Header{}, false, nil))
	proof := common.BuildRelayRequest(ctx, "providerAddr", []byte{}, uint64(0), "specId", nil)

	cpe := &rewardserver.ConsumerProofEntity{
		ConsumerAddr: "consumerAddr",
		ConsumerKey:  "consumerKey" + "specId",
		Proof:        proof,
	}
	err = rs.Save(cpe)
	require.NoError(t, err)

	rewards, err := rs.FindAll()
	require.NoError(t, err)
	require.Equal(t, 1, len(rewards))
}

func TestFindOne(t *testing.T) {
	db := rewardserver.NewMemoryDB("specId")
	rs := rewardserver.NewRewardDB()
	err := rs.AddDB(db)
	require.NoError(t, err)

	ctx := sdk.WrapSDKContext(sdk.NewContext(nil, tmproto.Header{}, false, nil))
	proof := common.BuildRelayRequest(ctx, "providerAddr", []byte{}, uint64(0), "specId", nil)
	proof.Epoch = 1

	cpe := &rewardserver.ConsumerProofEntity{
		ConsumerAddr: "consumerAddr",
		ConsumerKey:  "consumerKey",
		Proof:        proof,
	}
	err = rs.Save(cpe)
	require.NoError(t, err)

	reward, err := rs.FindOne(uint64(proof.Epoch), "consumerAddr", "consumerKey", proof.SessionId)
	require.NoError(t, err)
	require.NotNil(t, reward)
}

func TestDeleteClaimedRewards(t *testing.T) {
	db := rewardserver.NewMemoryDB("specId")
	rs := rewardserver.NewRewardDB()
	err := rs.AddDB(db)
	require.NoError(t, err)

	privKey, addr := sigs.GenerateFloatingKey()
	ctx := sdk.WrapSDKContext(sdk.NewContext(nil, tmproto.Header{}, false, nil))

	proof := common.BuildRelayRequest(ctx, "providerAddr", []byte{}, uint64(0), "specId", nil)
	proof.Epoch = 1

	sig, err := sigs.Sign(privKey, *proof)
	require.NoError(t, err)
	proof.Sig = sig

	consumerRewardsKey := "consumerKey"
	cpe := &rewardserver.ConsumerProofEntity{
		ConsumerAddr: addr.String(),
		ConsumerKey:  consumerRewardsKey,
		Proof:        proof,
	}
	err = rs.Save(cpe)
	require.NoError(t, err)

	err = rs.DeleteClaimedRewards(uint64(proof.Epoch), addr.String(), proof.SessionId, consumerRewardsKey)
	require.NoError(t, err)

	rewards, err := rs.FindAll()
	require.NoError(t, err)
	require.Equal(t, 0, len(rewards))
}

func TestDeleteEpochRewards(t *testing.T) {
	db := rewardserver.NewMemoryDB("specId")
	rs := rewardserver.NewRewardDB()
	err := rs.AddDB(db)
	require.NoError(t, err)

	privKey, addr := sigs.GenerateFloatingKey()
	ctx := sdk.WrapSDKContext(sdk.NewContext(nil, tmproto.Header{}, false, nil))

	proof := common.BuildRelayRequest(ctx, "providerAddr", []byte{}, uint64(0), "specId", nil)
	proof.Epoch = 1

	sig, err := sigs.Sign(privKey, *proof)
	require.NoError(t, err)
	proof.Sig = sig

	cpe := &rewardserver.ConsumerProofEntity{
		ConsumerAddr: addr.String(),
		ConsumerKey:  "consumerKey",
		Proof:        proof,
	}
	err = rs.Save(cpe)
	require.NoError(t, err)

	err = rs.DeleteEpochRewards(uint64(proof.Epoch))
	require.NoError(t, err)

	rewards, err := rs.FindAll()
	require.NoError(t, err)
	require.Equal(t, 0, len(rewards))
}

func TestRewardsWithTTL(t *testing.T) {
	db := rewardserver.NewMemoryDB("spec")
	// really really short TTL to make sure the rewards are not queryable
	ttl := 1 * time.Nanosecond
	rs := rewardserver.NewRewardDBWithTTL(ttl)
	err := rs.AddDB(db)
	require.NoError(t, err)

	privKey, addr := sigs.GenerateFloatingKey()
	ctx := sdk.WrapSDKContext(sdk.NewContext(nil, tmproto.Header{}, false, nil))

	proof := common.BuildRelayRequest(ctx, "provider", []byte{}, uint64(0), "spec", nil)
	proof.Epoch = 1

	sig, err := sigs.Sign(privKey, *proof)
	require.NoError(t, err)
	proof.Sig = sig

	cpe := &rewardserver.ConsumerProofEntity{
		ConsumerAddr: addr.String(),
		ConsumerKey:  "consumerKey",
		Proof:        proof,
	}
	err = rs.Save(cpe)
	require.NoError(t, err)

	rewards, err := rs.FindAll()
	require.NoError(t, err)
	require.Equal(t, 0, len(rewards))
}
