package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hirokryptor/xyzchain/x/xyzchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strconv"
)

// GetPoolCount get the total number of pool
func (k Keeper) GetPoolCount(ctx sdk.Context) int64 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolCountKey))
	byteKey := types.KeyPrefix(types.PoolCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetPoolCount set the total number of pool
func (k Keeper) SetPoolCount(ctx sdk.Context, count int64)  {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolCountKey))
	byteKey := types.KeyPrefix(types.PoolCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreatePool(ctx sdk.Context, msg types.MsgCreatePool) {
	// Create the pool
    count := k.GetPoolCount(ctx)
    var pool = types.Pool{
        Creator: msg.Creator,
        Id:      strconv.FormatInt(count, 10),
        Token1: msg.Token1,
        Token2: msg.Token2,
        Total1: msg.Total1,
        Total2: msg.Total2,
    }

    store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
    key := types.KeyPrefix(types.PoolKey + pool.Id)
    value := k.cdc.MustMarshalBinaryBare(&pool)
    store.Set(key, value)

    // Update pool count
    k.SetPoolCount(ctx, count+1)
}

func (k Keeper) UpdatePool(ctx sdk.Context, pool types.Pool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	b := k.cdc.MustMarshalBinaryBare(&pool)
	store.Set(types.KeyPrefix(types.PoolKey + pool.Id), b)
}

func (k Keeper) GetPool(ctx sdk.Context, key string) types.Pool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	var pool types.Pool
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PoolKey + key)), &pool)
	return pool
}

func (k Keeper) HasPool(ctx sdk.Context, id string) bool {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	return store.Has(types.KeyPrefix(types.PoolKey + id))
}

func (k Keeper) GetPoolOwner(ctx sdk.Context, key string) string {
    return k.GetPool(ctx, key).Creator
}

// DeletePool deletes a pool
func (k Keeper) DeletePool(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	store.Delete(types.KeyPrefix(types.PoolKey + key))
}

func (k Keeper) GetAllPool(ctx sdk.Context) (msgs []types.Pool) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PoolKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Pool
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
        msgs = append(msgs, msg)
	}

    return
}
