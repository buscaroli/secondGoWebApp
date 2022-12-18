package handlers

import (
	"fmt"
	"net/http"

	"github.com/buscaroli/secondGoWebApp/pkg/config"
	"github.com/buscaroli/secondGoWebApp/pkg/models"
	"github.com/buscaroli/secondGoWebApp/pkg/render"
)

// repository used by the handlers
var Repo *Repository

// repository type
type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// handler method on the Repository struct
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.RequestURI)
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Method)

	// saving the remote IP when the homepage is accessed
	m.App.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// handler method on the Repository struct
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "This is a test"

	// if user has visited the homepage their remote_ip will be displayed
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template and render it
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
