package enum

type AuthType string

const (
	// authorization_code
	AUTH_TYPE_AUTHORIZATION_CODE AuthType = "authorization_code"
	// refresh_token
	AUTH_TYPE_REFRESH_TOKEN AuthType = "refresh_token"
)
