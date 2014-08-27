package main

import (
	"github.com/unrolled/render"
	"net/http"
)

type Handlers struct {
	Render       *render.Render
	LayoutRender *render.Render
}

func BaseHandlers() *Handlers {
	lr := render.New(render.Options{
		Layout: "layout",
	})

	r := render.New(render.Options{})

	return &Handlers{LayoutRender: lr, Render: r}
}

func (h Handlers) landing(w http.ResponseWriter, req *http.Request) {
	h.Render.HTML(w, http.StatusOK, "landing", nil)
}

func (h Handlers) signup(w http.ResponseWriter, req *http.Request) {
	h.LayoutRender.HTML(w, http.StatusOK, "accounts/signup", nil)
}

/**
 * Account handlers
 */
func (h Handlers) accountIndex(w http.ResponseWriter, req *http.Request) {
	h.LayoutRender.HTML(w, http.StatusOK, "accounts/index", map[string]interface{}{"name": "Jesse O'Brien"})
}
