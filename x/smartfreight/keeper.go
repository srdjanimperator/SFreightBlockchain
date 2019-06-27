package smartfreight

import (
	"crypto/sha1"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	coinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

func intToSha(n int64) string {
	s := strconv.FormatInt(n, 10)
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	s2 := fmt.Sprintf("%x\n", bs)
	return s2
}

// Sets the entire Delivery metadata struct for a name
func (k Keeper) SetDelivery(ctx sdk.Context, jobId int64, delivery Delivery) {
	fmt.Println(delivery)
	key := intToSha(jobId)
	if delivery.Shipper.Empty() || delivery.Broker.Empty() || delivery.Price.Empty() {
		fmt.Println("Prazan")
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(delivery))
}

func (k Keeper) GetDeliveryByJobID(ctx sdk.Context, jobId int64) Delivery {
	key := intToSha(jobId)
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(key)) {
		fmt.Println("EMPTY BYTES")
		return Delivery{}
	}
	dBytes := store.Get([]byte(key))
	var d Delivery
	k.cdc.MustUnmarshalBinaryBare(dBytes, &d)
	return d
}

func (k Keeper) SetCompleted(ctx sdk.Context, jobId int64) {
	key := intToSha(jobId)
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(key)) {
		return
	}
	dBytes := store.Get([]byte(key))
	var d Delivery
	k.cdc.MustUnmarshalBinaryBare(dBytes, &d)
	d.Completed = true
	store.Set([]byte(key), k.cdc.MustMarshalBinaryBare(d))
}

func (k Keeper) GetDeliveryIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
