package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Whois is a struct that contains all the metadata of a name
type Delivery struct {
	Shipper   sdk.AccAddress `json:"shipper"`
	Broker    sdk.AccAddress `json:"broker"`
	Price     sdk.Coins      `json:"price"`
	Completed bool           `json:"completed"`
}

// Returns a new Whois with the minprice as the price
func NewDelivery(broker sdk.AccAddress, shipper sdk.AccAddress, price sdk.Coins) Delivery {
	return Delivery{shipper, broker, price, false}
}

// implement fmt.Stringer
func (d Delivery) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
Broker: %s
Shipper: %s
Price: %s`, d.Broker, d.Shipper, d.Price))
}
