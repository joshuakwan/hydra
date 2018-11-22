package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joshuakwan/hydra/models"

	"github.com/joshuakwan/hydra/client"
)

var c *client.Client
var w *models.Worker

func main() {
	c = client.NewClient("http://127.0.0.1:8080")

	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	w = &models.Worker{Name: name, Status: models.UP, Address: "127.0.0.1", Registered: time.Now().Unix()}

	err = c.RegisterWorker(w)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		os.Exit(1)
	}()

	go heartbeat()

	for {
	}
}

func heartbeat() {
	for {
		time.Sleep(time.Second * 5)
		fmt.Println("sending heartbeat")
		c.ReportWorker(w)
	}
}
