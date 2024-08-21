package main

import (
	"Asynq/client"
	"Asynq/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go server.StartServer()
	go client.StartClient()

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	<-channel

}
