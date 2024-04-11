package keeper

import (
	"invoice/x/invoice/types"
)

var _ types.QueryServer = Keeper{}
