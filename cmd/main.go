package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mzulfanw/upwatch/internal/app"
	_ "modernc.org/sqlite"
)

func main() {
	port := getenv("PORT", "8080")
	dbPath := getenv("DB_PATH", "data/upwatch.db")
	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		log.Fatalf("create data dir: %v", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	db.SetMaxOpenConns(1)

	if err := app.InitDB(db); err != nil {
		log.Fatalf("init db: %v", err)
	}

	adminUser := getenvRequired("ADMIN_USER")
	adminPass := getenvRequired("ADMIN_PASSWORD")
	sessionTTL := getenvDuration("SESSION_TTL", 24*time.Hour)

	application := app.New(db, app.AuthConfig{
		Username:   adminUser,
		Password:   adminPass,
		CookieName: "upwatch_session",
		SessionTTL: sessionTTL,
	})
	if err := application.StartMonitors(); err != nil {
		log.Fatalf("load monitors: %v", err)
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           application.Router(),
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      0,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("Upwatch listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func getenv(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}

func getenvRequired(key string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		log.Fatalf("%s is required", key)
	}
	return value
}

func getenvDuration(key string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return duration
}
