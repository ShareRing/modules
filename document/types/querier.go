package types

// query endpoints supported by the Id Querier
const (
	QueryInfo = "info"
)

// QueryIdByAddressParams defines the params for querying an account id information.
type QueryDocByProofParams struct {
	Proof []byte
}

func NewQueryDocByProofParams(proof []byte) QueryDocByProofParams {
	return QueryDocByProofParams{Proof: proof}
}
