package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thinkerajay/dsp/routes"
	"log"
)

const ADDR string = ":32142"

var router *gin.Engine

func Configure() {
	router = gin.Default()
	routes.AddRoutes(router)
}

func gracefulShutDown(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("exiting server !")
	}
}

func Start(ctx context.Context) {
	log.Println("running server at ")

	go func() {
		err := router.Run(ADDR)
		if err != nil {
			log.Fatalln("cannot start the server due to ", err)
			return
		}
	}()
	gracefulShutDown(ctx)

}
