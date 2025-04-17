package main

import (
	"log"
	"os"

	"github.com/MonkieeBoi/LinGo/internal/components"

	"gioui.org/app"
	"gioui.org/layout"
)

type C = layout.Context
type D = layout.Dimensions

func main() {
	go func() {
		window := new(app.Window)
		if err := components.NewAppWindow(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
