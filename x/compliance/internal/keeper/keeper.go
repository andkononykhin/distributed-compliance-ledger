package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

// Gets the entire ModelInfo metadata struct for a ModelInfoID
func (k Keeper) GetModelInfo(ctx sdk.Context, vid int16, pid int16) types.ModelInfo {
	if !k.IsModelInfoPresent(ctx, vid, pid) {
		panic("ModelInfo does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(id(vid, pid)))

	var device types.ModelInfo

	k.cdc.MustUnmarshalBinaryBare(bz, &device)

	return device
}

// Sets the entire ModelInfo metadata struct for a ModelInfoID
func (k Keeper) SetModelInfo(ctx sdk.Context, device types.ModelInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(id(device.VID, device.PID)), k.cdc.MustMarshalBinaryBare(device))
}

// Deletes the ModelInfo from the store
func (k Keeper) DeleteModelInfo(ctx sdk.Context, vid int16, pid int16) {
	if !k.IsModelInfoPresent(ctx, vid, pid) {
		panic("ModelInfo does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(id(vid, pid)))
}

// Iterate over all ModelInfos
func (k Keeper) IterateModelInfos(ctx sdk.Context, process func(info types.ModelInfo) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, nil)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var modelInfo types.ModelInfo

		k.cdc.MustUnmarshalBinaryBare(val, &modelInfo)

		if process(modelInfo) {
			return
		}

		iter.Next()
	}
}

// Check if the ModelInfo is present in the store or not
func (k Keeper) IsModelInfoPresent(ctx sdk.Context, vid int16, pid int16) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(id(vid, pid)))
}

func (k Keeper) CountTotal(ctx sdk.Context) int {
	store := ctx.KVStore(k.storeKey)
	res := 0

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		res++
	}

	return res
}

func id(vid int16, pid int16) string {
	return fmt.Sprintf("%d:%d", vid, pid)
}