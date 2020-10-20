// Step 8: add event and attribute
// NOTE: add event const for handlerMsg

package types

// logistic module event types
const (
	// TODO: Create your event types
	// EventType<Action>    		= "action"
	EventTypeInitDeal   = "InitDeal"
	EventTypeTransport  = "Transport"
	EventTypeUpdateTemp = "UpdateTemp"
	EventTypeReceive    = "Receive"

	// TODO: Create keys fo your events, the values will be derivided from the msg
	// AttributeKeyAddress  		= "address"
	AttributeOwner       = "owner"
	AttributeTransporter = "transporter"
	AttributeCustomer    = "customer"
	AttributePrice       = "price"
	AttributeMaxTemp     = "maxTemp"
	AttributeMinTemp     = "minTemp"
	AttributeCancelable  = "cancelable"
	AttributeUpdateTemp  = "updateTemp"
	AttributeOrderID     = "orderid"

	// TODO: Some events may not have values for that reason you want to emit that something happened.
	// AttributeValueDoubleSign = "double_sign"

	AttributeValueCategory = ModuleName
)
