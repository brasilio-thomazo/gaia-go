package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"optimus.dev.br/gaia/repository"
	"optimus.dev.br/gaia/request"
	"optimus.dev.br/gaia/response"
	"optimus.dev.br/gaia/security"
)

type AuthController struct {
	repo *repository.UserRepository
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		repo: repository.NewUserRepository(db),
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var request request.AuthRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.repo.FindByUsername(ctx, request.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized username or password not match"})
		return
	}
	if !security.BcryptCheckPassword(request.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized username or password not match"})
		return
	}
	ctx.JSON(http.StatusOK, &response.AuthResponse{
		User:  *user,
		Token: "",
	})
}
