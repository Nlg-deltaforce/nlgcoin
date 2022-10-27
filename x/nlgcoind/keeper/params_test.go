package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "nlgcoind/testutil/keeper"
	"nlgcoind/x/nlgcoind/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NlgcoindKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
