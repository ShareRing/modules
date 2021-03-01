<!--
order: 3
-->
# Messages

## MsgCreateDoc
```go
type MsgCreateDoc struct {
	Data   string         `json:"data"`
	Holder string         `json:"holder"`
	Issuer sdk.AccAddress `json:"issuer"`
	Proof  string         `json:"proof"`
}
```
## MsgCreateDocBatch
```go

type MsgCreateDocBatch struct {
	Data   []string       `json:"data"`
	Issuer sdk.AccAddress `json:"issuer"`
	Holder []string       `json:"holder"`
	Proof  []string       `json:"proof"`
}
```

## MsgRevokeDoc
```go
type MsgRevokeDoc struct {
	Holder string         `json:"holder"`
	Issuer sdk.AccAddress `json:"issuer"`
	Proof  string         `json:"proof"`
}
```

## MsgUpdateDoc
```go

type MsgUpdateDoc struct {
	Data   string         `json:"data"`
	Holder string         `json:"holder"`
	Issuer sdk.AccAddress `json:"issuer"`
	Proof  string         `json:"proof"`
}
```