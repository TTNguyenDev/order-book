package keeper

import (
	"orderbook-interchange/x/dex/types"
)

var _ types.QueryServer = Keeper{}
