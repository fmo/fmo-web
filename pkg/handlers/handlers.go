package handlers

import (
	"github.com/fmo/fmo-web/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTempalte(w, "home.page.tmpl")
}
