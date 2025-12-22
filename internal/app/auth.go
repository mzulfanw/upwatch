package app

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"net/http"
	"sync"
	"time"
)

const (
	defaultSessionTTL  = 24 * time.Hour
	defaultCookieName  = "upwatch_session"
	sessionTokenLength = 32
)

type AuthConfig struct {
	Username   string
	Password   string
	CookieName string
	SessionTTL time.Duration
}

type session struct {
	Username  string
	ExpiresAt time.Time
}

type SessionManager struct {
	mu       sync.Mutex
	sessions map[string]session
	ttl      time.Duration
}

func NewSessionManager(ttl time.Duration) *SessionManager {
	if ttl <= 0 {
		ttl = defaultSessionTTL
	}
	return &SessionManager{
		sessions: make(map[string]session),
		ttl:      ttl,
	}
}

func (sm *SessionManager) Create(username string) (string, time.Time, error) {
	token, err := generateToken(sessionTokenLength)
	if err != nil {
		return "", time.Time{}, err
	}
	expires := time.Now().Add(sm.ttl)

	sm.mu.Lock()
	sm.sessions[token] = session{Username: username, ExpiresAt: expires}
	sm.mu.Unlock()

	return token, expires, nil
}

func (sm *SessionManager) Get(token string) (session, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sess, ok := sm.sessions[token]
	if !ok {
		return session{}, false
	}
	if time.Now().After(sess.ExpiresAt) {
		delete(sm.sessions, token)
		return session{}, false
	}
	return sess, true
}

func (sm *SessionManager) Delete(token string) {
	sm.mu.Lock()
	delete(sm.sessions, token)
	sm.mu.Unlock()
}

func generateToken(length int) (string, error) {
	buf := make([]byte, length)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

func (a *App) isAuthenticated(r *http.Request) bool {
	token, ok := a.sessionTokenFromRequest(r)
	if !ok {
		return false
	}
	_, ok = a.sessions.Get(token)
	return ok
}

func (a *App) sessionTokenFromRequest(r *http.Request) (string, bool) {
	cookie, err := r.Cookie(a.auth.CookieName)
	if err != nil {
		return "", false
	}
	if cookie.Value == "" {
		return "", false
	}
	return cookie.Value, true
}

func (a *App) requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !a.isAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next(w, r)
	}
}

func (a *App) requireAuthAPI(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !a.isAuthenticated(r) {
			writeError(w, http.StatusUnauthorized, "unauthorized")
			return
		}
		next(w, r)
	}
}

func (a *App) credentialsMatch(username string, password string) bool {
	if len(username) != len(a.auth.Username) || len(password) != len(a.auth.Password) {
		return false
	}
	userOK := subtle.ConstantTimeCompare([]byte(username), []byte(a.auth.Username)) == 1
	passOK := subtle.ConstantTimeCompare([]byte(password), []byte(a.auth.Password)) == 1
	return userOK && passOK
}

func (a *App) setSessionCookie(w http.ResponseWriter, token string, expires time.Time, secure bool) {
	cookie := http.Cookie{
		Name:     a.auth.CookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   secure,
		Expires:  expires,
	}
	http.SetCookie(w, &cookie)
}

func (a *App) clearSessionCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     a.auth.CookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	}
	http.SetCookie(w, &cookie)
}
