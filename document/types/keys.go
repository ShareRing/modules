package types

const (
	// ModuleName defines the module name
	ModuleName = "doc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	DefaultParamspace = ModuleName
)
const (
	TypeMsgCreateDoc        = "create_doc"
	TypeMsgCreateDocInBatch = "create_doc_batch"
	TypeMsgUpdateDoc        = "update_doc"
	TypeMsgRevokeDoc        = "revoke_doc"
)

const (
	QueryByProof  = "proof"
	QueryByHolder = "holder"
)
const (
	EventTypeCreateDoc = "create_doc"
	EventTypeUpdateDoc = "update_doc"
	EventTypeRevokeDoc = "revoke_doc"

	EventAttrHolder = "holder"
	EventAttrProof  = "proof"
	EventAttrIssuer = "issuer"
	EventAttrData   = "data"
)

var (
	StateKeySep     = "|"
	DocDetailPrefix = []byte{0x1}
	DocBasicPrefix  = []byte{0x2}
	DocRevokeFlag   = 0xffff
)
