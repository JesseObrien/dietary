package main

import (
	"github.com/unrolled/render"
	"net/http"
)

type Handlers struct {
	Render *render.Render
}

func SiteHandlers() *Handlers {
	r := render.New(render.Options{
		Layout: "layout",
	})
	return &Handlers{Render: r}
}

func (h Handlers) index(w http.ResponseWriter, req *http.Request) {
	h.Render.HTML(w, http.StatusOK, "index", nil)
}

func (h Handlers) signup(w http.ResponseWriter, req *http.Request) {
	h.Render.HTML(w, http.StatusOK, "accounts/signup", nil)
}

/**
 * Account handlers
 */
func (h Handlers) accountIndex(w http.ResponseWriter, req *http.Request) {
	h.Render.HTML(w, http.StatusOK, "accounts/index", map[string]interface{}{"name": "Jesse O'Brien"})
}
