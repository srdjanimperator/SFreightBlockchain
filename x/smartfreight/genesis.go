package smartfreight

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	DeliveryRecords []Delivery `json:"delivery_records"`
}

func NewGenesisState(DeliveryRecords []Delivery) GenesisState {
	return GenesisState{DeliveryRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.DeliveryRecords {
		if record.Broker == nil {
			return fmt.Errorf("Invalid Delivery. Error: Missing Broker")
		}
		if record.Shipper == nil {
			return fmt.Errorf("Invalid Delivery. Error: Missing Shipper")
		}
		if record.Price == nil {
			return fmt.Errorf("Invalid Delivery. Error: Missing Price")
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		DeliveryRecords: []Delivery{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Delivery
	iterator := k.GetDeliveryIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		jobId := string(iterator.Key())
		jobIdInt, _ := strconv.ParseInt(jobId, 10, 64)
		var d Delivery
		d = k.GetDeliveryByJobID(ctx, jobIdInt)
		records = append(records, d)
	}
	return GenesisState{DeliveryRecords: records}
}
