package main

import (
	"log"
	"os"

	"github.com/okeefem2/celeritas"
	"github.com/okeefem2/go_remix/data"
	"github.com/okeefem2/go_remix/handlers"
	"github.com/okeefem2/go_remix/middleware"
)

func initApplication() *application {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)

	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "go_remix"

	appHandlers := &handlers.Handlers{
		App: cel,
	}

	appMiddleware := &middleware.Middleware{
		App: cel,
	}

	app := &application{
		App:      cel,
		Handlers: appHandlers,
		Middleware: appMiddleware,
	}

	app.Models = data.New(app.App.DB.Pool)
	app.Handlers.Models = app.Models // This feels weird to me btw.
	app.App.Routes = app.routes()

	return app
}
