package acerender

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
	"github.com/yosssi/ace-proxy"
)

func Test_renderer_HTML_parseError(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r := &renderer{
		ResponseWriter: res,
		req:            req,
		p:              proxy.New(nil),
	}

	r.HTML(http.StatusOK, "not_exist_template", nil, nil)

	if res.Code != http.StatusInternalServerError {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusInternalServerError)
	}
}

func Test_renderer_HTML_executeError(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r := &renderer{
		ResponseWriter: res,
		req:            req,
		p:              proxy.New(nil),
	}

	r.HTML(http.StatusOK, "test/0002", "test", nil)

	if res.Code != http.StatusInternalServerError {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusInternalServerError)
	}
}

func Test_renderer_HTML(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	r := &renderer{
		ResponseWriter: res,
		req:            req,
		p:              proxy.New(nil),
	}

	r.HTML(http.StatusOK, "test/0003:test/0004", nil, nil)

	if res.Code != http.StatusOK {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusOK)
	}
}

func TestRenderer(t *testing.T) {
	m := martini.Classic()
	m.Use(Renderer(nil))
	m.Get("/", func(r Render) {
		r.HTML(http.StatusOK, "test/0001", nil, nil)
	})
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	m.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("invalid HTTP status code [actual: %d][expected: %d]", res.Code, http.StatusOK)
	}
}
