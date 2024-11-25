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
)

type CustomerController struct {
	Controller
	repo *repository.CustomerRepository
}

func NewCustomerController(db *gorm.DB) *CustomerController {
	return &CustomerController{
		repo: repository.NewCustomerRepository(db),
	}
}

func (c *CustomerController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.repo.FindAll(ctx))
}

func (c *CustomerController) Show(ctx *gin.Context) {
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

func (c *CustomerController) Create(ctx *gin.Context) {
	var request request.CustomerRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if c.repo.ExistsByName(ctx, strings.ToUpper(request.Name)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name already exists"})
		return
	}

	data := model.Customer{
		Name:     strings.ToUpper(request.Name),
		Phone:    request.Phone,
		Email:    strings.ToLower(request.Email),
		Document: request.Document,
		Address:  request.Address,
		Contacts: request.Contacts,
		Active:   request.Active,
	}

	if err := c.repo.Create(ctx, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, data)
}

func (c *CustomerController) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var request request.CustomerRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if c.repo.ExistsByNameAndIDNot(ctx, strings.ToUpper(request.Name), id) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name already exists"})
		return
	}

	data := model.Customer{
		ID:       id,
		Name:     strings.ToUpper(request.Name),
		Phone:    request.Phone,
		Email:    strings.ToLower(request.Email),
		Document: request.Document,
		Address:  request.Address,
		Contacts: request.Contacts,
		Active:   request.Active,
	}
	if err := c.repo.Update(ctx, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *CustomerController) Delete(ctx *gin.Context) {
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
