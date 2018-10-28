package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joshuakwan/hydra/client"
	"github.com/joshuakwan/hydra/models"
	"github.com/robfig/cron"
)

var scheduleVar string
var serverVar string

func main() {
	flag.StringVar(&scheduleVar, "c", "@every 10s", "set the cron schedule")
	flag.StringVar(&serverVar, "s", "http://127.0.0.1:8080", "set the API server URL")
	flag.Parse()
	c := cron.New()
	client := client.NewClient(serverVar)

	fmt.Printf("launch the cron trigger with the schedule %s\n", scheduleVar)
	c.AddFunc(scheduleVar, func() {
		current := time.Now().Unix()
		fmt.Printf("Hello %s\n", strconv.FormatInt(current, 10))
		fmt.Println("creating event")
		event := models.Event{
			Type:      models.IF,
			Source:    "cron_trigger",
			Message:   "hello",
			Timestamp: current,
		}
		if err := client.CreateEvent(&event); err != nil {
			log.Fatal(err)
		}

	})

	c.Start()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		c.Stop()

		os.Exit(1)
	}()

	for {

	}
}
