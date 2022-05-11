package main

import (
	"github.com/okeefem2/celeritas"
	"github.com/okeefem2/go_remix/data"
	"github.com/okeefem2/go_remix/handlers"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
	Models data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
