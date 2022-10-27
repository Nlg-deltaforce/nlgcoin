package keeper

import (
	"nlgcoind/x/nlgcoind/types"
)

var _ types.QueryServer = Keeper{}
