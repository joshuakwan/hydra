package main

// Logic:
// Listens on the events
// Evalute an event against the rules
// Execute the rule -> action
import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	hydra_client "github.com/joshuakwan/hydra/client"
	"github.com/joshuakwan/hydra/models"

	"github.com/antonmedv/expr"
)

var client *hydra_client.Client

func main() {
	client = hydra_client.NewClient("http://127.0.0.1:8080")

	rules, _ := client.ListRules()

	go func() {
		ech := client.WatchEvents()
		for e := range ech {
			fmt.Printf("%s\n", e)
			evaluate(e, rules)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch

		os.Exit(1)
	}()

	for {

	}
}

func evaluate(event models.Event, rules []*models.Rule) {
	for _, rule := range rules {
		p, err := expr.Parse(rule.If, expr.Env(models.Event{}))
		if err != nil {
			fmt.Println(err)
			continue
		}
		out, err := expr.Run(p, event)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if out.(bool) {
			fmt.Println("Evaluation succeeded")
		} else {
			fmt.Println("Evaluation failed")
		}
	}
}
