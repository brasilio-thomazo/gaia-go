package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"optimus.dev.br/gaia/model"
	"optimus.dev.br/gaia/repository"
	"optimus.dev.br/gaia/request"
	"optimus.dev.br/gaia/security"
)

type UserController struct {
	Controller
	repo *repository.UserRepository
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		repo: repository.NewUserRepository(db),
	}
}

func (c *UserController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.repo.FindAll(ctx))
}

func (c *UserController) Show(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := c.repo.FindByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *UserController) Create(ctx *gin.Context) {
	var request request.UserRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required", "field": "name"})
		return
	}

	if request.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username is required", "field": "username"})
		return
	}

	if len(request.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 6 characters", "field": "password"})
		return
	}

	if request.Password != request.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "passwords do not match", "field": "password_confirm"})
		return
	}

	if c.repo.ExistsByUsername(ctx, strings.ToLower(request.Username)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username already exists", "field": "username"})
		return
	}

	if c.repo.ExistsByEmail(ctx, strings.ToLower(request.Email)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email already exists", "field": "email"})
		return
	}

	hash, err := security.BcryptHashPassword(request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := model.User{
		Name:     strings.ToUpper(request.Name),
		Phone:    request.Phone,
		JobTitle: request.JobTitle,
		Email:    strings.ToLower(request.Email),
		Username: strings.ToLower(request.Username),
		Password: hash,
		Visible:  true,
		Editable: true,
		Locked:   request.Locked,
	}

	if err := c.repo.Create(ctx, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, data)
}

func (c *UserController) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var request request.UserRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required", "field": "name"})
		return
	}

	if request.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username is required", "field": "username"})
		return
	}

	if request.Password != "" && len(request.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 6 characters", "field": "password"})
		return
	}

	if request.Password != "" && request.Password != request.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "passwords do not match", "field": "password_confirm"})
		return
	}

	if c.repo.ExistsByUsernameAndIDNot(ctx, strings.ToLower(request.Username), id) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username already exists", "field": "username"})
		return
	}

	if c.repo.ExistsByEmailAndIDNot(ctx, strings.ToLower(request.Email), id) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email already exists", "field": "email"})
		return
	}

	data := model.User{
		ID:       id,
		Name:     strings.ToUpper(request.Name),
		Phone:    request.Phone,
		JobTitle: request.JobTitle,
		Email:    strings.ToLower(request.Email),
		Username: strings.ToLower(request.Username),
		Visible:  true,
		Editable: true,
		Locked:   request.Locked,
	}
	if request.Password != "" {
		hash, err := security.BcryptHashPassword(request.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		data.Password = hash
	}
	if err := c.repo.Update(ctx, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *UserController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.repo.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
