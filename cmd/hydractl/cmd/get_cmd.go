package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/joshuakwan/hydra/client"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func handleGet(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "events", "ev":
		getEvents()
	case "rules", "ru":
		fmt.Println("not implemented")
	case "actions", "ac":
		fmt.Println("not implemented")
	}
}

func getEvents() {
	client := client.NewClient("http://127.0.0.1:8080")
	events, err := client.ListEvents()
	checkError(err)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetHeader([]string{"Timestamp", "Human Readable Date", "Source", "Message", "Type"})

	for _, e := range events {
		table.Append([]string{fmt.Sprintf("%d", e.Timestamp), fmt.Sprintf("%s", time.Unix(e.Timestamp, 0)), e.Source, e.Message, string(e.Type)})
	}

	table.Render()
}
