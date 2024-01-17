package dex_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "orderbook-interchange/testutil/keeper"
	"orderbook-interchange/testutil/nullify"
	"orderbook-interchange/x/dex/module"
	"orderbook-interchange/x/dex/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, k, genesisState)
	got := dex.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
