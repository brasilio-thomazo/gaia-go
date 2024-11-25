package response

import (
	"github.com/docker/docker/api/types"
	"optimus.dev.br/gaia/model"
)

type AppRespone struct {
	model.App
	Container types.ContainerJSON `json:"container"`
}
