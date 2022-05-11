package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func commonHttpTestSetup(t *testing.T) (*http.Request, http.ResponseWriter) {
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil { t.Error(err) }
	w := httptest.NewRecorder()
	return r, w
}


// Table testing
var pageData = []struct {
	name string
	renderer string
	template string
	errorExpected bool
	errorMessage string
}{
	// This is cool, first time I have seen this shorthand syntax for structs in Go
	{ "go_page_exists", "go", "home", false, "Error rendering go template", },
	{ "go_page_not_exists", "go", "non-existant", true, "No error rendering non existant go template", },
	{ "jet_page_exists", "jet", "home", false, "Error rendering jet template", },
	{ "jet_page_not_exists", "jet", "non-existant", true, "No error rendering non existant jet template", },
	{ "no_engine", "", "home", true, "No error rendering using no engine", },
	{ "unsupported_engine", "remix", "home", true, "No error rendering using unsupported engine", },
}

func TestRender_Page_Non_Table(t *testing.T) {
	// Setup
	r, w := commonHttpTestSetup(t)
	testRenderer.RootPath = "./testdata"

	// Test go template
	testRenderer.Renderer = "go"
	err := testRenderer.Page(w, r, "home", nil, nil)
	if err != nil { t.Error("Error rendering go template", err) }

	// Test non existant go file
	err = testRenderer.Page(w, r, "non-existant", nil, nil)
	if err == nil { t.Error("No error rendering non existant go template", err) }

	// Test jet template
	testRenderer.Renderer = "jet"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil { t.Error("Error rendering jet template", err) }

		// Test non existant jet file
	err = testRenderer.Page(w, r, "non-existant", nil, nil)
	if err == nil { t.Error("No error rendering non existant jet template", err) }


	// Test no renderer set
	testRenderer.Renderer = ""
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err == nil { t.Error("No error rendering unsupported engine", err) }
}

func TestRender_Page_Table(t *testing.T) {
	// Setup
	r, w := commonHttpTestSetup(t)
	testRenderer.RootPath = "./testdata"

	for _, data := range pageData {
		testRenderer.Renderer = data.renderer
		err := testRenderer.Page(w, r, data.template, nil, nil)
		// This is technically an xor statement too of errorExpected xor err == nil
		if data.errorExpected && err == nil  {
			// Expected error and didn't get one
			t.Errorf("%s: %s", data.name, data.errorMessage)
		} else if !data.errorExpected && err != nil {
			// Didnm't expect an error but got one
			t.Errorf("%s: %s; %s", data.name, data.errorMessage, err.Error())
		}
	}
}
