package router

import "github.com/gin-gonic/gin"

type BaseRouter interface {
	Register(engine *gin.Engine)
	list(c *gin.Context)
	page(c *gin.Context)
	create(c *gin.Context)
	update(c *gin.Context)
	remove(c *gin.Context)
}
