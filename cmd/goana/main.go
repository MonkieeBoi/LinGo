package main

import (
	"log"
	"os"

	"github.com/MonkieeBoi/goana/internal/app"
	"github.com/MonkieeBoi/goana/internal/db"

	gapp "gioui.org/app"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initalise database.")
	}
	go func() {
		window := new(gapp.Window)
		if err := app.NewAppWindow(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	gapp.Main()
}
