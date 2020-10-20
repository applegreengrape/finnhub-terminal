package main

import (
	"fmt"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/applegreengrape/fintech-terminal/yahoo"
)

func main() {
	if err := ui.Init(); err != nil {
		fmt.Println(err)
	}
	defer ui.Close()

	l := yahoo.GetStockPrice()
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(0, 0, 45, 15)

	// Render the first screen
	ui.Render(l)

	// Create a ticker for refreshing the screen every 5s
	tickEvents := time.NewTicker(5 * time.Second).C

	// Start listening for ui events
	uiEvents := ui.PollEvents()

	// Run the event loop
	for {

		select {

		// Handle UI events
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				l.ScrollDown()
			case "k", "<Up>":
				l.ScrollUp()
			}

			// re-render after handling user input
			ui.Render(l)

		// Handle render events
		case <-tickEvents:
			ui.Render(l)
		}
	}
}
