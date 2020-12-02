package keeper

import (
	"github.com/hirokryptor/xyzchain/x/xyzchain/types"
)

var _ types.QueryServer = Keeper{}
