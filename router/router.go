package router

import (
	"gorm.io/gorm"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func NewRouter(db *gorm.DB, c *client.Client) *gin.Engine {
	r := gin.Default()
	registerRoutes(db, r, c)
	return r
}
