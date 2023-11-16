package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "no-file", true, "randering non-existing go template"},
	{"jet_page", "jet", "home", false, "error rendering jet template"},
	{"jet_page_no_template", "go", "no-file", true, "randering non-existing jet template"},
	{"invalid_renderer", "unknown", "home", true, "endering with non-existent template engine"},
}

func TestRender_Page(t *testing.T) {

	for _, e := range pageData {

		r, err := http.NewRequest("GET", "/some-url", nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()

		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./testdata"

		err = testRenderer.Page(w, r, e.template, nil, nil)
		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s:", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s ", e.name, e.errorMessage, err.Error())
			}
		}

	}
	/* durch tabletest ersetzt
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existent template", err)
	}

	testRenderer.Renderer = "jet"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering jet page", err)
	}

	err = testRenderer.Page(w, r, "no-file", nil, nil)
	if err == nil {
		t.Error("Error rendering non-existing jet template", err)
	}

	testRenderer.Renderer = "unknown"
	err = testRenderer.Page(w, r, "home", nil, nil)
	if err == nil {
		t.Error("Error rendering with unknown rendering engine", err)
	}
	*/
}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.GoPage(w, r, "home", nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}

	testRenderer.Renderer = "jet"
	testRenderer.RootPath = "./testdata"

	err = testRenderer.JetPage(w, r, "home", nil, nil)
	if err != nil {
		t.Error("Error rendering page", err)
	}

}
