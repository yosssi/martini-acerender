package acerender

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/yosssi/ace"
)

const defaultContentType = render.ContentHTML + "; charset=utf-8"

// Render is an interface for parsing Ace templates and redering HTML.
type Render interface {
	HTML(status int, name string, v interface{}, opts *ace.Options)
}

// render represents a renderer of Ace templates.
type renderer struct {
	http.ResponseWriter
	req  *http.Request
	Base string
}

// HTML parses the Ace templates and renders HTML to the response writer.
func (r *renderer) HTML(status int, name string, v interface{}, opts *ace.Options) {
	var basePath, innerPath string

	paths := strings.Split(name, ":")

	basePath = paths[0]

	if len(paths) > 1 {
		innerPath = paths[1]
	}

	tpl, err := ace.Load(basePath, innerPath, opts)

	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)

	if err := tpl.Execute(buf, v); err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return
	}

	r.Header().Set(render.ContentType, defaultContentType)
	r.WriteHeader(status)
	io.Copy(r, buf)
}

func (r *renderer) Ace(status int, name string, v interface{}) {

	options := &ace.Options{
		BaseDir: r.Base,
	}

	r.HTML(status, name, v, options)

}

func (r *renderer) AceOk(name string, v interface{}) {
	r.Ace(http.StatusOK, name, v)
}

func (r *renderer) AceNotFound(name string, v interface{}) {
	r.Ace(http.StatusNotFound, name, v)
}

// Renderer is a Martini middleware that maps a render.Render service into the Martini handler chain.
func Renderer(_ *Options) martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		c.MapTo(
			&renderer{
				ResponseWriter: res,
				req:            req,
			},
			(*Render)(nil),
		)
	}
}
