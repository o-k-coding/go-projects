package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist string
	CookieSecure string
	CookieName string
	CookieDomain string
	SessionType string
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	// how long should sessions last
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		// Might be nice to log this
		minutes = 60
	}

	persist = strings.ToLower(c.CookiePersist) == "true"
	secure = strings.ToLower(c.CookieSecure) == "true"

	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Name = c.CookieName
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.Persist = persist
	session.Cookie.Secure = secure
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite#lax
	session.Cookie.SameSite = http.SameSiteLaxMode

	switch strings.ToLower(c.SessionType) {
		case "redis":
		case "mysql", "mariadb":
		case "postgres", "postgresql":
		default:
		// cookie
	}

	return session
}


func IsAuthenticated(s *scs.SessionManager, r *http.Request) bool {
	return s.Exists(r.Context(), "userEmail")
}
