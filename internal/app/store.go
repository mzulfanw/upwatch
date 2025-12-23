package app

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

func InitDB(db *sql.DB) error {
	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		return err
	}
	schema := `
CREATE TABLE IF NOT EXISTS monitors (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	url TEXT NOT NULL,
	method TEXT NOT NULL DEFAULT 'GET',
	interval_sec INTEGER NOT NULL DEFAULT 60,
	timeout_sec INTEGER NOT NULL DEFAULT 10,
	status TEXT NOT NULL DEFAULT 'unknown',
	last_status_code INTEGER,
	last_error TEXT,
	last_latency_ms INTEGER,
	last_checked_at INTEGER,
	created_at INTEGER NOT NULL,
	updated_at INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS check_events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	monitor_id INTEGER NOT NULL,
	checked_at INTEGER NOT NULL,
	success INTEGER NOT NULL,
	status_code INTEGER,
	latency_ms INTEGER,
	error TEXT,
	FOREIGN KEY(monitor_id) REFERENCES monitors(id) ON DELETE CASCADE
);
CREATE INDEX IF NOT EXISTS idx_check_events_monitor ON check_events(monitor_id, checked_at DESC);
CREATE TABLE IF NOT EXISTS incidents (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	status TEXT NOT NULL,
	message TEXT NOT NULL,
	started_at INTEGER NOT NULL,
	resolved_at INTEGER,
	created_at INTEGER NOT NULL,
	updated_at INTEGER NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_incidents_started ON incidents(started_at DESC);
CREATE TABLE IF NOT EXISTS settings (
	id INTEGER PRIMARY KEY CHECK (id = 1),
	brand_name TEXT NOT NULL,
	brand_tagline TEXT NOT NULL,
	status_title TEXT NOT NULL,
	status_subtitle TEXT NOT NULL,
	updated_at INTEGER NOT NULL
);
`
	_, err := db.Exec(schema)
	return err
}

func listMonitors(db *sql.DB) ([]Monitor, error) {
	rows, err := db.Query(`
SELECT id, name, url, method, interval_sec, timeout_sec, status,
	last_status_code, last_error, last_latency_ms, last_checked_at, created_at, updated_at
FROM monitors ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var monitors []Monitor
	for rows.Next() {
		var m Monitor
		if err := rows.Scan(
			&m.ID,
			&m.Name,
			&m.URL,
			&m.Method,
			&m.IntervalSec,
			&m.TimeoutSec,
			&m.Status,
			&m.LastStatusCode,
			&m.LastError,
			&m.LastLatencyMs,
			&m.LastCheckedAt,
			&m.CreatedAt,
			&m.UpdatedAt,
		); err != nil {
			return nil, err
		}
		monitors = append(monitors, m)
	}
	return monitors, rows.Err()
}

func getMonitor(db *sql.DB, id int64) (Monitor, error) {
	var m Monitor
	err := db.QueryRow(`
SELECT id, name, url, method, interval_sec, timeout_sec, status,
	last_status_code, last_error, last_latency_ms, last_checked_at, created_at, updated_at
FROM monitors WHERE id = ?`, id).Scan(
		&m.ID,
		&m.Name,
		&m.URL,
		&m.Method,
		&m.IntervalSec,
		&m.TimeoutSec,
		&m.Status,
		&m.LastStatusCode,
		&m.LastError,
		&m.LastLatencyMs,
		&m.LastCheckedAt,
		&m.CreatedAt,
		&m.UpdatedAt,
	)
	return m, err
}

func createMonitor(db *sql.DB, input MonitorInput) (Monitor, error) {
	normalized, err := normalizeMonitorInput(input)
	if err != nil {
		return Monitor{}, err
	}
	now := time.Now().Unix()
	result, err := db.Exec(`
INSERT INTO monitors (name, url, method, interval_sec, timeout_sec, status, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, 'unknown', ?, ?)`,
		normalized.Name,
		normalized.URL,
		normalized.Method,
		normalized.IntervalSec,
		normalized.TimeoutSec,
		now,
		now,
	)
	if err != nil {
		return Monitor{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Monitor{}, err
	}
	return getMonitor(db, id)
}

func updateMonitor(db *sql.DB, id int64, input MonitorUpdate) (Monitor, error) {
	existing, err := getMonitor(db, id)
	if err != nil {
		return Monitor{}, err
	}

	updated := existing
	if input.Name != nil {
		updated.Name = strings.TrimSpace(*input.Name)
		if updated.Name == "" {
			return Monitor{}, errors.New("name is required")
		}
	}
	if input.URL != nil {
		updated.URL = strings.TrimSpace(*input.URL)
		if err := validateURL(updated.URL); err != nil {
			return Monitor{}, err
		}
	}
	if input.Method != nil {
		updated.Method = strings.ToUpper(strings.TrimSpace(*input.Method))
		if err := validateMethod(updated.Method); err != nil {
			return Monitor{}, err
		}
	}
	if input.IntervalSec != nil {
		if *input.IntervalSec <= 0 {
			return Monitor{}, errors.New("interval_sec must be greater than 0")
		}
		updated.IntervalSec = *input.IntervalSec
	}
	if input.TimeoutSec != nil {
		if *input.TimeoutSec <= 0 {
			return Monitor{}, errors.New("timeout_sec must be greater than 0")
		}
		updated.TimeoutSec = *input.TimeoutSec
	}

	now := time.Now().Unix()
	_, err = db.Exec(`
UPDATE monitors
SET name = ?, url = ?, method = ?, interval_sec = ?, timeout_sec = ?, updated_at = ?
WHERE id = ?`,
		updated.Name,
		updated.URL,
		updated.Method,
		updated.IntervalSec,
		updated.TimeoutSec,
		now,
		id,
	)
	if err != nil {
		return Monitor{}, err
	}
	updated.UpdatedAt = now
	return updated, nil
}

func deleteMonitor(db *sql.DB, id int64) error {
	result, err := db.Exec("DELETE FROM monitors WHERE id = ?", id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func listEvents(db *sql.DB, monitorID int64, limit int) ([]CheckEvent, error) {
	rows, err := db.Query(`
SELECT id, monitor_id, checked_at, success, status_code, latency_ms, error
FROM check_events
WHERE monitor_id = ?
ORDER BY checked_at DESC
LIMIT ?`, monitorID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []CheckEvent
	for rows.Next() {
		var e CheckEvent
		var success int
		if err := rows.Scan(
			&e.ID,
			&e.MonitorID,
			&e.CheckedAt,
			&success,
			&e.StatusCode,
			&e.LatencyMs,
			&e.Error,
		); err != nil {
			return nil, err
		}
		e.Success = success == 1
		events = append(events, e)
	}
	return events, rows.Err()
}

func recordCheck(db *sql.DB, monitorID int64, success bool, statusCode int, latencyMs int, errMsg string) error {
	now := time.Now().Unix()
	successVal := 0
	if success {
		successVal = 1
	}
	var statusCodeVal interface{} = nil
	if statusCode > 0 {
		statusCodeVal = statusCode
	}
	var errVal interface{} = nil
	if errMsg != "" {
		errVal = errMsg
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
INSERT INTO check_events (monitor_id, checked_at, success, status_code, latency_ms, error)
VALUES (?, ?, ?, ?, ?, ?)`, monitorID, now, successVal, statusCodeVal, latencyMs, errVal)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	status := "down"
	if success {
		status = "up"
	}
	_, err = tx.Exec(`
UPDATE monitors
SET status = ?, last_status_code = ?, last_error = ?, last_latency_ms = ?, last_checked_at = ?
WHERE id = ?`, status, statusCodeVal, errVal, latencyMs, now, monitorID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
