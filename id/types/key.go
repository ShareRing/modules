package types

const (
	// ModuleName defines the module name
	ModuleName = "id"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	DefaultParamspace = ModuleName
)

const (
	QueryPathAddress = "address"
	QueryPathId      = "id"
)

// ID message types
const (
	TypeMsgCreateID = "create_id"
)

// ID events
const (
	EventCreateID = "create_id"
)

const (
	EventAttrIssuer = "issuer"
	EventAttrId     = "id"
	EventAttrOwner  = "owner"
)
