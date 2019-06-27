package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Query Result Payload for a resolve query
type QueryResResolve struct {
	Shipper   sdk.AccAddress `json:"shipper"`
	Broker    sdk.AccAddress `json:"broker"`
	Price     sdk.Coins      `json:"price"`
	Completed bool           `json:"completed"`
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Shipper: %s
	Broker: %s
	Price: %s`, r.Shipper, r.Broker, r.Price))
}

type QueryResKeys []string

// implement fmt.Stringer
func (n QueryResKeys) String() string {
	return strings.Join(n[:], "\n")
}
