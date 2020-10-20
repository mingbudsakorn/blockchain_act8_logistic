package types

const (
	// ModuleName is the name of the module
	ModuleName = "logistic"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

// Step 5: add deal prefix which is use in KVStore
const (
	DealPrefix = "deal-"
)
