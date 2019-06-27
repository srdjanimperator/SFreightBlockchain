package smartfreight

import (
	"encoding/json"
	"smartfreight/x/smartfreight/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the nameservice Querier
const (
	QueryResolve = "resolve"
	QueryNames   = "keys"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			d, err := queryResolve(ctx, path[1:], req, keeper)
			if err != nil {
				return []byte{}, err
			}
			bytes, _ := json.Marshal(&d)
			return bytes, nil
		case QueryNames:
			return queryNames(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

// nolint: unparam
func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (Delivery, sdk.Error) {
	jobID, err := strconv.ParseInt(path[0], 10, 64)

	if err != nil { // TODO: fix return statements
		return Delivery{}, sdk.ErrUnknownRequest("JobID must be number")
	}

	d := keeper.GetDeliveryByJobID(ctx, jobID)

	if d.Broker.Empty() { // ?
		return Delivery{}, sdk.ErrUnknownRequest("could not resolve job delivery")
	}

	if err != nil {
		panic("could not marshal result to JSON")
	}

	return d, nil
}

func queryNames(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var keysList types.QueryResKeys

	iterator := keeper.GetDeliveryIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		keysList = append(keysList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, keysList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
