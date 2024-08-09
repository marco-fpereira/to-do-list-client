package input

// Implement Route, Hanlder/Controller, Service and Repository
// https://medium.com/@wahyubagus1910/build-scalable-restful-api-with-golang-gin-gonic-framework-43793c730d10

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marco-fpereira/to-do-list-client/app/adapters/input/dto"
	"github.com/marco-fpereira/to-do-list-client/app/adapters/utils"
	"github.com/marco-fpereira/to-do-list-client/app/domain/model"
	"github.com/marco-fpereira/to-do-list-client/app/domain/port/input"
)

type AccountAdapterInterafece interface {
	Signup(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type AccountAdapter struct {
	Account input.AccountPort
}

func NewAccountAdapter(
	accountPort input.AccountPort,
) AccountAdapterInterafece {
	return &AccountAdapter{
		Account: accountPort,
	}
}

func (a *AccountAdapter) Signup(ctx *gin.Context) {
	var accountDetailsDTO dto.AccountDetailsDTO
	if err := ctx.BindJSON(&accountDetailsDTO); err != nil {
		ctx.Error(err)
	}
	userCredentialsDomain := model.UserCredentialsDomain{
		Username: accountDetailsDTO.Username,
		Password: accountDetailsDTO.Password,
	}
	err := a.Account.Signup(ctx, &userCredentialsDomain, utils.GetRequestId(ctx))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"data": map[string]string{
			"message": "User successfully signed up",
		},
	})
}

func (a *AccountAdapter) Login(ctx *gin.Context) {
	var accountDetailsDTO dto.AccountDetailsDTO
	if err := ctx.BindJSON(&accountDetailsDTO); err != nil {
		ctx.Error(err)
	}
	userCredentialsDomain := model.UserCredentialsDomain{
		Username: accountDetailsDTO.Username,
		Password: accountDetailsDTO.Password,
	}
	userId, err := a.Account.Login(ctx, &userCredentialsDomain, utils.GetRequestId(ctx))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"data": map[string]string{
			"userId": *userId,
		},
	})
}
