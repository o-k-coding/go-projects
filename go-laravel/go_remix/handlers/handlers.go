package handlers

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/okeefem2/celeritas"
	"github.com/okeefem2/go_remix/data"
)

// I feel like this is just
type Handlers struct {
	App *celeritas.Celeritas
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) GoPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(w, r, "go-template", nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) JetPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(w, r, "jet-template", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) SessionPage(w http.ResponseWriter, r *http.Request) {
	myData := "may the force be with you"
	h.App.Session.Put(r.Context(), "data", myData)
	myValue := h.App.Session.GetString(r.Context(), "data")

	pageData := make(jet.VarMap)
	pageData.Set("data", myValue)
	err := h.App.Render.JetPage(w, r, "session", pageData, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
