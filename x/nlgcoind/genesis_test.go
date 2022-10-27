package nlgcoind_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "nlgcoind/testutil/keeper"
	"nlgcoind/testutil/nullify"
	"nlgcoind/x/nlgcoind"
	"nlgcoind/x/nlgcoind/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NlgcoindKeeper(t)
	nlgcoind.InitGenesis(ctx, *k, genesisState)
	got := nlgcoind.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
