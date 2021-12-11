package controller

import (
	"books/domain"
	"books/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authentication *service.AuthenticationService
	jwt *service.JwtService
}


func NewAuthController(authentication *service.AuthenticationService) *AuthController{
	return &AuthController{
		authentication: authentication,
		jwt: service.NewJwtService(),
	}
}

func (c *AuthController) Authenticate(ctx *gin.Context){ 
	authDTO := domain.User{}
	
	err := ctx.BindJSON(&authDTO)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "could not extract credentials: " + err.Error())
		return
	}

	user, err := c.authentication.AuthenticateUser(authDTO.Username, authDTO.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "could not authenticate " + authDTO.Username + ": " + err.Error())
		return
	}

	token, err := c.jwt.GeneratoToken(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "could not generate token: " + err.Error())
		return
	}
	ctx.JSON(http.StatusOK, token)

}