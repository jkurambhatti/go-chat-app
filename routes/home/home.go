package home

import (
	"html/template"
	"net/http"
	"os"

	templates "github.com/auth0/auth0-golang/examples/regular-web-app/routes"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Auth0ClientId     string
		Auth0ClientSecret string
		Auth0Domain       string
		Auth0CallbackURL  template.URL
	}{
		os.Getenv("AUTH0_CLIENT_ID"),
		os.Getenv("AUTH0_CLIENT_SECRET"),
		os.Getenv("AUTH0_DOMAIN"),
		template.URL(os.Getenv("AUTH0_CALLBACK_URL")),
	}

	templates.RenderTemplate(w, r, "home", data)
}
