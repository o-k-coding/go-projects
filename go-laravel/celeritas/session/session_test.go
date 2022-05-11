package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alexedwards/scs/v2"
)

func TestSession_InitSession(t *testing.T) {
	sessionConfig := &Session{
		CookieLifetime: "100",
		CookiePersist: "true",
		CookieName: "test",
		CookieDomain: "localhost",
		SessionType: "cookie",
	}

	var sm *scs.SessionManager

	s := sessionConfig.InitSession()

	var sessionKind reflect.Kind
	var sessionType reflect.Type
	rv := reflect.ValueOf(s) // get runtime value

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		fmt.Println("In for loop: ", rv.Kind(), rv.Type(), rv)
		sessionKind = rv.Kind()
		sessionType = rv.Type()

		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Error("Invalid type or kind: ", rv.Kind(), "type: ", rv.Type())
	}

	expectedSessionKind := reflect.ValueOf(sm).Kind()
	if sessionKind !=  expectedSessionKind{
		t.Error("wrong kind returned testing cookie session. Expected", expectedSessionKind, "and got", sessionKind)
	}

	expectedSessionType := reflect.ValueOf(sm).Type()
	if sessionType !=  expectedSessionType {
		t.Error("wrong type returned testing cookie session. Expected", expectedSessionType, "and got", sessionType)
	}
}
