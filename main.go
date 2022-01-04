package main

import (
	"context"
	"github.com/thinkerajay/dsp/server"
	"log"
	"os"
	"os/signal"
)


func main(){
	server.Configure()
	ctrlC := make(chan os.Signal, 1)
	signal.Notify(ctrlC, os.Interrupt)


	

	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctrlC:
			log.Println("Received interrupt performing cleanup !")
			cancelFunc()
		}
	}()
	server.Start(ctx)
}
