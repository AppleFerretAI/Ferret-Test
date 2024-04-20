package ferret_test

import (
	"testing"

	keepertest "Ferret/testutil/keeper"
	"Ferret/testutil/nullify"
	ferret "Ferret/x/ferret/module"
	"Ferret/x/ferret/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FerretKeeper(t)
	ferret.InitGenesis(ctx, k, genesisState)
	got := ferret.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
