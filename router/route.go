package router

import (
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"optimus.dev.br/gaia/controller"
)

func registerRoutes(db *gorm.DB, r *gin.Engine, c *client.Client) {
	groups(db, r)
	users(db, r)
	customers(db, r)
	apps(db, r, c)
	auth(db, r)
}

func groups(db *gorm.DB, r *gin.Engine) {
	ctl := controller.NewGroupController(db)
	r.GET("/groups", ctl.Index)
	r.GET("/groups/:id", ctl.Show)
	r.POST("/groups", ctl.Create)
	r.PUT("/groups/:id", ctl.Update)
	r.DELETE("/groups/:id", ctl.Delete)
}

func users(db *gorm.DB, r *gin.Engine) {
	ctl := controller.NewUserController(db)
	r.GET("/users", ctl.Index)
	r.GET("/users/:id", ctl.Show)
	r.POST("/users", ctl.Create)
	r.PUT("/users/:id", ctl.Update)
	r.DELETE("/users/:id", ctl.Delete)
}

func customers(db *gorm.DB, r *gin.Engine) {
	ctl := controller.NewCustomerController(db)
	r.GET("/customers", ctl.Index)
	r.GET("/customers/:id", ctl.Show)
	r.POST("/customers", ctl.Create)
	r.PUT("/customers/:id", ctl.Update)
	r.DELETE("/customers/:id", ctl.Delete)
}

func apps(db *gorm.DB, r *gin.Engine, c *client.Client) {
	ctl := controller.NewAppController(db, c)
	r.GET("/apps", ctl.Index)
	r.GET("/apps/:id", ctl.Show)
	r.POST("/apps", ctl.Create)
	r.PUT("/apps/:id", ctl.Update)
	r.DELETE("/apps/:id", ctl.Delete)
}

func auth(db *gorm.DB, r *gin.Engine) {
	ctl := controller.NewAuthController(db)
	r.POST("/login", ctl.Login)
}
