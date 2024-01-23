package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "peach/testutil/keeper"
	"peach/x/peach/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.PeachKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
