package main

import (
	"log"
	"os"

	"github.com/MonkieeBoi/LinGo/internal/components"
	"github.com/MonkieeBoi/LinGo/internal/db"

	"gioui.org/app"
	"gioui.org/layout"
)

type C = layout.Context
type D = layout.Dimensions

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initalise database.")
	}
	go func() {
		window := new(app.Window)
		if err := components.NewAppWindow(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
