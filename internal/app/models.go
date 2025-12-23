package app

import "database/sql"

const (
	defaultIntervalSec = 60
	defaultTimeoutSec  = 10
	maxEventsLimit     = 1000
	statusHistoryLimit = 40
	maxIncidentsLimit  = 100
	defaultIncidentStatus = "investigating"
	defaultBrandName   = "Upwatch Status"
	defaultBrandTagline = "Public status feed for tracked services."
	defaultStatusTitle = "Current system health"
	defaultStatusSubtitle = "Updated in real time from the stream."
)

type Monitor struct {
	ID             int64
	Name           string
	URL            string
	Method         string
	IntervalSec    int
	TimeoutSec     int
	Status         string
	LastStatusCode sql.NullInt64
	LastError      sql.NullString
	LastLatencyMs  sql.NullInt64
	LastCheckedAt  sql.NullInt64
	CreatedAt      int64
	UpdatedAt      int64
}

type MonitorInput struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Method      string `json:"method"`
	IntervalSec int    `json:"interval_sec"`
	TimeoutSec  int    `json:"timeout_sec"`
}

type MonitorUpdate struct {
	Name        *string `json:"name"`
	URL         *string `json:"url"`
	Method      *string `json:"method"`
	IntervalSec *int    `json:"interval_sec"`
	TimeoutSec  *int    `json:"timeout_sec"`
}

type MonitorResponse struct {
	ID             int64   `json:"id"`
	Name           string  `json:"name"`
	URL            string  `json:"url"`
	Method         string  `json:"method"`
	IntervalSec    int     `json:"interval_sec"`
	TimeoutSec     int     `json:"timeout_sec"`
	Status         string  `json:"status"`
	LastStatusCode *int    `json:"last_status_code"`
	LastError      *string `json:"last_error"`
	LastLatencyMs  *int    `json:"last_latency_ms"`
	LastCheckedAt  *string `json:"last_checked_at"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

type CheckEvent struct {
	ID         int64
	MonitorID  int64
	CheckedAt  int64
	Success    bool
	StatusCode sql.NullInt64
	LatencyMs  sql.NullInt64
	Error      sql.NullString
}

type CheckEventResponse struct {
	ID         int64   `json:"id"`
	MonitorID  int64   `json:"monitor_id"`
	CheckedAt  string  `json:"checked_at"`
	Success    bool    `json:"success"`
	StatusCode *int    `json:"status_code"`
	LatencyMs  *int    `json:"latency_ms"`
	Error      *string `json:"error"`
}

type StatusResponse struct {
	Counts   StatusCounts      `json:"counts"`
	Monitors []MonitorResponse `json:"monitors"`
	History  map[string][]StatusHistoryPoint `json:"history,omitempty"`
	Updated  string            `json:"updated_at"`
}

type StatusCounts struct {
	Up      int `json:"up"`
	Down    int `json:"down"`
	Unknown int `json:"unknown"`
	Total   int `json:"total"`
}

type StatusHistoryPoint struct {
	V *int `json:"v"`
}

type Settings struct {
	BrandName      string
	BrandTagline   string
	StatusTitle   string
	StatusSubtitle string
	UpdatedAt      int64
}

type SettingsUpdate struct {
	BrandName      *string `json:"brand_name"`
	BrandTagline   *string `json:"brand_tagline"`
	StatusTitle   *string `json:"status_title"`
	StatusSubtitle *string `json:"status_subtitle"`
}

type SettingsResponse struct {
	BrandName      string `json:"brand_name"`
	BrandTagline   string `json:"brand_tagline"`
	StatusTitle   string `json:"status_title"`
	StatusSubtitle string `json:"status_subtitle"`
	UpdatedAt      string `json:"updated_at"`
}

type Incident struct {
	ID         int64
	Title      string
	Status     string
	Message    string
	StartedAt  int64
	ResolvedAt sql.NullInt64
	CreatedAt  int64
	UpdatedAt  int64
}

type IncidentInput struct {
	Title      string `json:"title"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	StartedAt  string `json:"started_at"`
	ResolvedAt string `json:"resolved_at"`
}

type IncidentUpdate struct {
	Title      *string `json:"title"`
	Status     *string `json:"status"`
	Message    *string `json:"message"`
	StartedAt  *string `json:"started_at"`
	ResolvedAt *string `json:"resolved_at"`
}

type IncidentResponse struct {
	ID         int64   `json:"id"`
	Title      string  `json:"title"`
	Status     string  `json:"status"`
	Message    string  `json:"message"`
	StartedAt  string  `json:"started_at"`
	ResolvedAt *string `json:"resolved_at"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}
