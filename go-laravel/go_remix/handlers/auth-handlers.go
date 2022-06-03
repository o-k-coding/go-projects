package handlers

import "net/http"

func (h *Handlers) UserLogin(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(w, r, "login", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) PostUserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() // Parse data into the r.Form struct
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")
	matches, err := h.Models.Users.PasswordMatches(email, password)
	if err != nil || !matches {
		w.Write([]byte("Error validating user"))
		return
	}

	h.App.Session.Put(r.Context(), "userEmail", email)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) UserLogout(w http.ResponseWriter, r *http.Request) {
	// Not sure why the renew token bit, but ok...
	h.App.Session.RenewToken(r.Context())
	h.App.Session.Remove(r.Context(), "userEmail")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}
