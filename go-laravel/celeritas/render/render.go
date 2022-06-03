package render

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/okeefem2/celeritas/session"
)

type Render struct {
  Renderer string
  RootPath string
	Secure bool
	Port string
	ServerName string
	JetViews *jet.Set
	Session *scs.SessionManager
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap map[string] int
	StringMap map[string]string
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Port string
	ServerName string
	Secure bool
}

func (re *Render) defaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Secure = re.Secure
	td.ServerName = re.ServerName
	td.Port = re.Port
	// This might be useful to codify
	if session.IsAuthenticated(re.Session, r) {
		td.IsAuthenticated = true
	}
	return td
}

func (re *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(re.Renderer) {
		case "go":
			return re.GoPage(w, r, view, data)
		case "jet":
			return re.JetPage(w, r, view, variables, data)
		default:
	}
	return errors.New("no supported rendering engine specified for celeritas, please set the RENDERER environment varaible to one of (go, jet)")
}

func (re *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", re.RootPath, view))

	if err != nil { return err }
	td := &TemplateData{}

	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)

	if err != nil { return err }

	return nil
}

func (re *Render) JetPage(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	var vars jet.VarMap

	// If variables are nil, init the ds otherwise cast
	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}
	td = re.defaultData(td, r);
	if data != nil {
		td = data.(*TemplateData)
	}


	t, err := re.JetViews.GetTemplate(fmt.Sprintf("%s.jet", view))
	// Note that the configured logger isn't accessible here, so that should be done as an improvement
	if err != nil { return err }

	if err = t.Execute(w, vars, td); err != nil {
		return err
	}
	return nil
}
