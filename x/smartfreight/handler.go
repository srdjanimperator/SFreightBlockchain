package smartfreight

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "smartfreight" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetDelivery:
			return handleMsgSetDelivery(ctx, keeper, msg)
		case MsgCompleteDelivery:
			return handleMsgCompleteDelivery(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized smartfreight Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgSetDelivery(ctx sdk.Context, keeper Keeper, msg MsgSetDelivery) sdk.Result {
	fmt.Println("handler")
	if msg.Delivery.Shipper.Equals(msg.Delivery.Broker) {
		return sdk.ErrUnauthorized("Same broker and shipper").Result()
	}
	keeper.SetDelivery(ctx, msg.JobID, msg.Delivery)
	return sdk.Result{}
}

func handleMsgCompleteDelivery(ctx sdk.Context, keeper Keeper, msg MsgCompleteDelivery) sdk.Result {
	d := keeper.GetDeliveryByJobID(ctx, msg.JobID)
	if d.Broker.Empty() { // ?
		return sdk.ErrUnknownRequest("Unknown job ID").Result()
	}
	err := keeper.coinKeeper.SendCoins(ctx, d.Broker, d.Shipper, d.Price)
	if err != nil {
		return sdk.ErrInsufficientCoins("Broker does not have enough coins").Result()
	}
	keeper.SetCompleted(ctx, msg.JobID)
	return sdk.Result{}
}
