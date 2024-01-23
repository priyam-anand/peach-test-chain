package peach_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "peach/testutil/keeper"
	"peach/testutil/nullify"
	"peach/x/peach/module"
	"peach/x/peach/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PeachKeeper(t)
	peach.InitGenesis(ctx, k, genesisState)
	got := peach.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
