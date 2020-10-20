// Step 4: create Deal struct and define necessary enum
package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// enum of deal state
type StateType string

const (
	Created   StateType = "Created"
	InTransit           = "InTransit"
	Complete            = "Complete"
	Cancel				= "Cancel"
)

type Deal struct {
	Owner       sdk.AccAddress `json:"owner" yaml:"owner"`
	Transporter sdk.AccAddress `json:"transporter" yaml:"transporter"`
	Customer    sdk.AccAddress `json:"customer" yaml:"customer"`
	Price       sdk.Coins      `json:"price" yaml:"price"`
	OrderID     string         `json:"orderid" yaml:"orderid"`
	MaxTemp     int            `json:"maxTemp" yaml:"maxTemp"`
	MinTemp     int            `json:"minTemp" yaml:"minTemp"`
	Cancelable  bool           `json:"cancelable" yaml:"cancelable"`
	State       StateType      `json:"state" yaml:"state"`
}

func (d Deal) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
	Transporter: %s
	Customer: %s
	Price: %s
	OrderID: %s
	MaxTemp: %d
	MinTemp: %d,
	Cancelable: %t,
	State: %s`,
		d.Owner,
		d.Transporter,
		d.Customer,
		d.Price,
		d.OrderID,
		d.MaxTemp,
		d.MinTemp,
		d.Cancelable,
		string(d.State),
	))
}
