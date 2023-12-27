package main

import (
	"fmt"
	"time"


	"github.com/rivo/tview"
)


// Set refresh interval to 5 seconds
const refreshInterval = 500 * time.Millisecond

var (
	view *tview.Modal
	app *tview.Application
)


func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("Current time is 15:04:05"))
}

func updateTIme() {
	for {
		time.Sleep(refreshInterval)
		app.QueueUpdateDraw(func() {
		view.SetText(currentTimeString())
		})
	}
}

func main() {
	app = tview.NewApplication()
	view = tview.NewModal().
		SetText(currentTimeString()).
		AddButtons([]string{"Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			}
		})

		go updateTIme()
		if err := app.SetRoot(view, false).Run(); err != nil {
			panic(err)
		}
}