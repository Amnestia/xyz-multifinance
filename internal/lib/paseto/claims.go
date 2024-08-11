package paseto

// Payload struct containing auth data
type Payload struct {
	ID        int64  `json:"id"`
	TokenType string `json:"token_type"`
}

// CtxKey context key for auth
type CtxKey string

// claims constant
const (
	AccessToken  = "Access-Token"
	RefreshToken = "Refresh-Token"

	audience = "xyz_multifinance - User"
	issuer   = "xyz_multifinance - Auth"
	jti      = "xyz_multifinance - XYZ"
	footer   = "xyz_multifinance - Auth Token"

	payloadKey   = "Payload"
	tokenTypeKey = "Token-Type"

	AuthData CtxKey = "AuthData"
)
