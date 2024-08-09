package usecase

import (
	"context"

	"github.com/marco-fpereira/to-do-list-client/app/domain/model"
	inputPort "github.com/marco-fpereira/to-do-list-client/app/domain/port/input"
	outputPort "github.com/marco-fpereira/to-do-list-client/app/domain/port/output"
)

type AccountUseCase struct {
	output        outputPort.ToDoServerPort
	authenticator outputPort.AuthenticationPort
}

func NewAccountUseCase(
	output outputPort.ToDoServerPort,
	authenticator outputPort.AuthenticationPort,
) inputPort.AccountPort {
	return &AccountUseCase{
		output:        output,
		authenticator: authenticator,
	}
}

func (a *AccountUseCase) Signup(
	ctx context.Context,
	userCredentials *model.UserCredentialsDomain,
	requestId string,
) error {
	token, err := a.authenticator.GenerateToken()
	if err != nil {
		return err
	}
	return a.output.Signup(ctx, userCredentials, requestId, *token)
}

func (a *AccountUseCase) Login(
	ctx context.Context,
	userCredentials *model.UserCredentialsDomain,
	requestId string,
) (*string, error) {
	token, err := a.authenticator.GenerateToken()
	if err != nil {
		return nil, err
	}
	userId, err := a.output.Login(ctx, userCredentials, requestId, *token)
	if err != nil {
		return nil, err
	}
	return userId, err
}
