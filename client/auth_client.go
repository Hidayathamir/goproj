package client

type AuthClient interface {
	Login(username, password string) string
	Logout(token string) bool
	ValidateToken(token string) bool
}

// compile-time check
var _ AuthClient = (*AuthClientImpl)(nil)

type AuthClientImpl struct {
}

func NewAuthClient() *AuthClientImpl {
	return &AuthClientImpl{}
}

func (c *AuthClientImpl) Login(username, password string) string {
	return "dummy_token_123"
}

func (c *AuthClientImpl) Logout(token string) bool {
	return true
}

func (c *AuthClientImpl) ValidateToken(token string) bool {
	return true
}
