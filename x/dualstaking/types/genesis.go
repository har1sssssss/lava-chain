package types

import (
	"github.com/lavanet/lava/common"
	commontypes "github.com/lavanet/lava/common/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params:        DefaultParams(),
		DelegationsFS: *common.DefaultGenesis(),
		DelegatorsFS:  *common.DefaultGenesis(),
		UnbondingsTS:  []commontypes.RawMessage{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
