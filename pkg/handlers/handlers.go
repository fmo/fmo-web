package handlers

import (
	"github.com/fmo/fmo-web/pkg/config"
	"github.com/fmo/fmo-web/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "home.page.tmpl")
}
