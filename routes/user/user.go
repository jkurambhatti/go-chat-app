package user

import (
	"net/http"

	templates "github.com/auth0/auth0-golang/examples/regular-web-app/routes"

	"github.com/auth0/auth0-golang/examples/regular-web-app/app"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := app.GlobalSessions.SessionStart(w, r)
	defer session.SessionRelease(w)

	templates.RenderTemplate(w, r, "user", session.Get("profile"))
}
