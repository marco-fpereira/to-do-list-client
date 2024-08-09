package output

type Claim struct {
	ClaimName  string
	ClaimValue string
}

type AuthenticationPort interface {
	GenerateToken() (*string, error)
	GenerateTokenWithClaim(...Claim) (*string, error)
}
