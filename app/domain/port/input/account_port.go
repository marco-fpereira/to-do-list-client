package input

import (
	"context"

	"github.com/marco-fpereira/to-do-list-client/app/domain/model"
)

type AccountPort interface {
	Signup(
		ctx context.Context,
		userCredentials *model.UserCredentialsDomain,
		requestId string,
	) error

	Login(
		ctx context.Context,
		userCredentials *model.UserCredentialsDomain,
		requestId string,
	) (*string, error)
}
