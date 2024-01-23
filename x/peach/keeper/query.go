package keeper

import (
	"peach/x/peach/types"
)

var _ types.QueryServer = Keeper{}
