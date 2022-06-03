package main

import (
	"github.com/okeefem2/celeritas"
	"github.com/okeefem2/go_remix/data"
	"github.com/okeefem2/go_remix/handlers"
	"github.com/okeefem2/go_remix/middleware"
)

type application struct {
	// This is called app, but it is really more of the framework config... has all of the "out of the box" stuff
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
	Middleware *middleware.Middleware
	Models data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
