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

type GroupController struct {
	Controller
	repo *repository.GroupRepository
}

func NewGroupController(db *gorm.DB) *GroupController {
	return &GroupController{
		repo: repository.NewGroupRepository(db),
	}
}

func (c *GroupController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.repo.FindAll(ctx))
}

func (c *GroupController) Show(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
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

func (c *GroupController) Create(ctx *gin.Context) {
	var request request.GroupRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if c.repo.ExistsByName(ctx, strings.ToLower(request.Name)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name already exists"})
		return
	}

	data := model.Group{
		Name:        strings.ToLower(request.Name),
		Permissions: request.Permissions,
		Visible:     true,
		Editable:    true,
		Locked:      request.Locked,
	}

	if err := c.repo.Create(ctx, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, data)
}

func (c *GroupController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var request request.GroupRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if c.repo.ExistsByNameAndIDNot(ctx, strings.ToLower(request.Name), id) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name already exists"})
		return
	}

	data := model.Group{
		ID:          id,
		Name:        strings.ToLower(request.Name),
		Permissions: request.Permissions,
		Locked:      request.Locked,
	}
	if err := c.repo.Update(ctx, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *GroupController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
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
