package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerajay/dsp/handlers"
)

func AddRoutes(router *gin.Engine){
	router.Handle("POST", "/main", handlers.MainHandler)
}
