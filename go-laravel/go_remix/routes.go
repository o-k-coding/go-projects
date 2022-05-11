package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/okeefem2/go_remix/data"
)

func (app *application) routes() *chi.Mux {
	// middleware comes first before routes

	// routes come next

	app.App.Routes.Get("/", app.Handlers.Home)
	app.App.Routes.Get("/go-page", app.Handlers.GoPage)
	app.App.Routes.Get("/jet-page", app.Handlers.JetPage)
	app.App.Routes.Get("/session-page", app.Handlers.SessionPage)
	app.App.Routes.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		u := data.User{
			FirstName: "Crokus",
			LastName: "Younghand",
			Email: "cutter@phoenix.com",
			Password: "password",
			Active: 1,
		}

		id, err := app.Models.Users.Insert(&u)
		if err != nil {
			app.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%d", id)
	})

	app.App.Routes.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := app.Models.Users.GetAll(nil)
		if err != nil {
			app.App.ErrorLog.Println(err)
			return
		}

		for _, x := range users {
			fmt.Fprintf(w, "%d", x.ID)
		}
	})

	app.App.Routes.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		u, err := app.Models.Users.Get(id)

		if err != nil {
			app.App.ErrorLog.Println(err)
			return
		}

		fmt.Fprintf(w, "%s %s %d", u.FirstName, u.LastName, u.ID)
	})

	app.App.Routes.Put("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))

		// TODO would actually accept data from the body
		u, err := app.Models.Users.Get(id)

		u.UpdatedAt = time.Now()

		if err != nil {
			app.App.ErrorLog.Println(err)
			return
		}

		app.Models.Users.Update(u)

		fmt.Fprintf(w, "%s %s %d", u.FirstName, u.LastName, u.ID)
	})

	// static routes

	fileServer := http.FileServer(http.Dir("./public"))
	app.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return app.App.Routes
}
