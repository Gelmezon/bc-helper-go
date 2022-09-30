package router

import (
	"awesomeProject/support"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	routerList := support.GenericList[BaseRouter]{}
	routerList.Push(BcWearRouter{})
	routerList.Push(BcChatRouter{})
	return include(&routerList)
}

func register(router BaseRouter, engine *gin.Engine) {
	router.Register(engine)
}

func include(routers *support.GenericList[BaseRouter]) *gin.Engine {
	r := gin.Default()
	for _, baseRouter := range routers.GetAll() {
		register(baseRouter, r)
	}
	return r
}
