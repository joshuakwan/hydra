package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joshuakwan/hydra/models"
	"github.com/robfig/cron"
)

var scheduleVar string

func main() {
	flag.StringVar(&scheduleVar, "s", "@every 10s", "set the cron schedule")
	flag.Parse()
	c := cron.New()
	fmt.Printf("launch the cron trigger with the schedule %s\n", scheduleVar)
	c.AddFunc(scheduleVar, func() {
		current := time.Now().Unix()
		fmt.Printf("Hello %s\n", strconv.FormatInt(current, 10))
		fmt.Println("creating event")
		models.CreateIFEvent("cron_trigger", "hello", current)
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
