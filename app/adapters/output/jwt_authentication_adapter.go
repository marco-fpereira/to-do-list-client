package output

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	portOutput "github.com/marco-fpereira/to-do-list-client/app/domain/port/output"
)

type JwtAuthenticationAdapter struct {
	secretKey []byte
}

func NewJwtAuthenticationAdapter() portOutput.AuthenticationPort {
	return &JwtAuthenticationAdapter{
		secretKey: []byte(os.Getenv("JWT-SECRET-KEY")),
	}
}

func (ja *JwtAuthenticationAdapter) GenerateToken() (*string, error) {
	tokenString, err := ja.getSignedToken(
		jwt.MapClaims{"exp": time.Now().Add(time.Second * 600).Unix()},
	)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (ja *JwtAuthenticationAdapter) GenerateTokenWithClaim(
	claims ...portOutput.Claim,
) (*string, error) {
	jwtClaims := jwt.MapClaims{"exp": time.Now().Add(time.Second * 600).Unix()}
	for _, claim := range claims {
		jwtClaims[claim.ClaimName] = claim.ClaimValue
	}
	tokenString, err := ja.getSignedToken(jwtClaims)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func (ja *JwtAuthenticationAdapter) getSignedToken(jwtClaims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString(ja.secretKey)
}
