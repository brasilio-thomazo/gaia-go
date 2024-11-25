package controller

import (
	"net/http"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"optimus.dev.br/gaia/model"
	"optimus.dev.br/gaia/repository"
	"optimus.dev.br/gaia/request"
	"optimus.dev.br/gaia/response"
)

type AppController struct {
	repo *repository.AppRepository
	cli  *client.Client
}

var (
	networkName = "client-lan"
)

func NewAppController(db *gorm.DB, cli *client.Client) *AppController {
	return &AppController{repo: repository.NewAppRepository(db), cli: cli}
}

func (c *AppController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.repo.FindAll(ctx))
}

func (c *AppController) Show(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := c.repo.FindByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	container, err := c.cli.ContainerInspect(ctx, data.ContainerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response.AppRespone{App: *data, Container: container})
}

func (c *AppController) Create(ctx *gin.Context) {
	var request request.AppCreateRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	host := &container.HostConfig{
		PortBindings: nat.PortMap{},
	}

	if request.Listening {
		for _, appPort := range request.Ports {
			port, err := nat.NewPort(appPort.Protocol, appPort.Port)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			host.PortBindings[port] = []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: appPort.HostPort,
				},
			}
		}
	}

	net := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			networkName: {},
		},
	}

	container := &container.Config{
		Image: request.Image,
		Cmd:   request.Cmd,
		Env:   request.Env,
	}

	res, err := c.cli.ContainerCreate(ctx.Request.Context(), container, host, net, nil, request.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data := model.App{
		ID:          id,
		Name:        request.Name,
		ContainerID: res.ID,
		CustomerID:  request.CustomerID,
		Image:       request.Image,
		Env:         request.Env,
		Cmd:         request.Cmd,
		Ports:       request.Ports,
		Replicas:    request.Replicas,
		Listening:   request.Listening,
		Active:      request.Active,
	}

	if err := c.repo.Create(ctx, &data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, data)
}

func (c *AppController) Update(ctx *gin.Context) {}

func (c *AppController) Delete(ctx *gin.Context) {}
