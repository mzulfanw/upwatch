package app

import (
	"log"
	"net/http"
)

func (a *App) handleLoginForm(w http.ResponseWriter, r *http.Request) {
	if a.isAuthenticated(r) {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}
	serveLoginPage(w, r)
}

func (a *App) handleLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Redirect(w, r, "/login?error=1", http.StatusFound)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	if !a.credentialsMatch(username, password) {
		http.Redirect(w, r, "/login?error=1", http.StatusFound)
		return
	}

	token, expires, err := a.sessions.Create(username)
	if err != nil {
		log.Printf("session create: %v", err)
		writeError(w, http.StatusInternalServerError, "failed to create session")
		return
	}

	a.setSessionCookie(w, token, expires, r.TLS != nil)
	http.Redirect(w, r, "/dashboard", http.StatusFound)
}

func (a *App) handleLogout(w http.ResponseWriter, r *http.Request) {
	if token, ok := a.sessionTokenFromRequest(r); ok {
		a.sessions.Delete(token)
	}
	a.clearSessionCookie(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (a *App) handleDashboard(w http.ResponseWriter, r *http.Request) {
	serveDashboardPage(w, r)
}
