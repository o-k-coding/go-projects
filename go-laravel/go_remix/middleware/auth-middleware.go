package middleware

import (
	"net/http"

	"github.com/okeefem2/celeritas/session"
)

func (m *Middleware) AuthRedirectGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !session.IsAuthenticated(m.App.Session, r) {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			next.ServeHTTP(w, r)
		}
	})
}
