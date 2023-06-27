/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/spf13/cobra"

	"github.com/sleepypioneer/community-assistant-cli/utils"
)

// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Displays a dashboard of the meetup stats",
	Run: func(cmd *cobra.Command, args []string) {
		eventID := cmd.Flag("event-id").Value.String()
		meetupStats := utils.ScrapeMeetupPage(eventID)

		terminalLayer, err := tcell.New(tcell.ColorMode(terminalapi.ColorMode256), tcell.ClearStyle(cell.ColorYellow, cell.ColorBlack))
		if err != nil {
			fmt.Printf("error running dashboard: %v", err)
			os.Exit(1)
		}
		defer terminalLayer.Close()

		t, err := text.New()
		if err != nil {
			fmt.Printf("error running dashboard: %v", err)
			os.Exit(1)
		}

		t.Write(fmt.Sprintf("Name: %v\nDate: %v\nAttendees: %v", meetupStats.Title, meetupStats.Time, meetupStats.Attendees))

		containerLayer, _ := container.New(
			terminalLayer,
			container.Border(linestyle.Light),
			container.PlaceWidget(t),
		)

		ctx, cancel := context.WithCancel(context.Background())

		quitter := func(k *terminalapi.Keyboard) {
			if k.Key == 'q' || k.Key == 'Q' {
				cancel()
			}
		}

		if err := termdash.Run(ctx, terminalLayer, containerLayer, termdash.KeyboardSubscriber(quitter)); err != nil {
			fmt.Printf("error running dashboard: %v", err)
			os.Exit(1)
		}

		return
	},
}

func init() {
	RootCmd.AddCommand(dashboardCmd)
	dashboardCmd.Flags().StringP("event-id", "e", "", "ID of the meetup page")
}
