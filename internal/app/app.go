package app

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	db       *sql.DB
	checker  *Checker
	auth     AuthConfig
	sessions *SessionManager
}

func New(db *sql.DB, auth AuthConfig) *App {
	if auth.CookieName == "" {
		auth.CookieName = defaultCookieName
	}
	if auth.SessionTTL <= 0 {
		auth.SessionTTL = defaultSessionTTL
	}
	if auth.Username == "" {
		auth.Username = "admin"
	}
	if auth.Password == "" {
		auth.Password = "admin"
	}
	return &App{
		db:       db,
		checker:  NewChecker(db),
		auth:     auth,
		sessions: NewSessionManager(auth.SessionTTL),
	}
}

func (a *App) StartMonitors() error {
	monitors, err := listMonitors(a.db)
	if err != nil {
		return err
	}
	a.checker.StartAll(monitors)
	return nil
}

func (a *App) Router() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	a.routes(r)
	return r
}

func (a *App) routes(r *mux.Router) {
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	r.HandleFunc("/", serveStatusPage).Methods("GET")
	r.HandleFunc("/api/health", a.handleHealth).Methods("GET")
	r.HandleFunc("/api/status", a.handleStatus).Methods("GET")
	r.HandleFunc("/api/status/stream", a.handleStatusStream).Methods("GET")
	r.HandleFunc("/api/incidents", a.handleListIncidents).Methods("GET")
	r.HandleFunc("/api/settings", a.handleGetSettings).Methods("GET")
	r.HandleFunc("/login", a.handleLoginForm).Methods("GET")
	r.HandleFunc("/login", a.handleLogin).Methods("POST")
	r.HandleFunc("/logout", a.handleLogout).Methods("GET")
	r.HandleFunc("/dashboard", a.requireAuth(a.handleDashboard)).Methods("GET")
	r.HandleFunc("/settings", a.requireAuth(a.handleSettingsPage)).Methods("GET")

	r.HandleFunc("/api/monitors", a.requireAuthAPI(a.handleListMonitors)).Methods("GET")
	r.HandleFunc("/api/monitors", a.requireAuthAPI(a.handleCreateMonitor)).Methods("POST")
	r.HandleFunc("/api/monitors/{id}", a.requireAuthAPI(a.handleGetMonitor)).Methods("GET")
	r.HandleFunc("/api/monitors/{id}", a.requireAuthAPI(a.handleUpdateMonitor)).Methods("PUT")
	r.HandleFunc("/api/monitors/{id}", a.requireAuthAPI(a.handleDeleteMonitor)).Methods("DELETE")
	r.HandleFunc("/api/monitors/{id}/events", a.requireAuthAPI(a.handleMonitorEvents)).Methods("GET")
	r.HandleFunc("/api/incidents", a.requireAuthAPI(a.handleCreateIncident)).Methods("POST")
	r.HandleFunc("/api/incidents/{id}", a.requireAuthAPI(a.handleUpdateIncident)).Methods("PUT")
	r.HandleFunc("/api/incidents/{id}", a.requireAuthAPI(a.handleDeleteIncident)).Methods("DELETE")
	r.HandleFunc("/api/settings", a.requireAuthAPI(a.handleUpdateSettings)).Methods("PUT")
}
