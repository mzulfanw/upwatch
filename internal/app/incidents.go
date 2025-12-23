package app

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

var incidentStatuses = map[string]struct{}{
	"investigating": {},
	"identified":    {},
	"monitoring":    {},
	"resolved":      {},
	"maintenance":   {},
	"scheduled":     {},
}

func listIncidents(db *sql.DB, limit int) ([]Incident, error) {
	rows, err := db.Query(`
SELECT id, title, status, message, started_at, resolved_at, created_at, updated_at
FROM incidents
ORDER BY started_at DESC
LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incidents []Incident
	for rows.Next() {
		var incident Incident
		if err := rows.Scan(
			&incident.ID,
			&incident.Title,
			&incident.Status,
			&incident.Message,
			&incident.StartedAt,
			&incident.ResolvedAt,
			&incident.CreatedAt,
			&incident.UpdatedAt,
		); err != nil {
			return nil, err
		}
		incidents = append(incidents, incident)
	}
	return incidents, rows.Err()
}

func getIncident(db *sql.DB, id int64) (Incident, error) {
	var incident Incident
	err := db.QueryRow(`
SELECT id, title, status, message, started_at, resolved_at, created_at, updated_at
FROM incidents
WHERE id = ?`, id).Scan(
		&incident.ID,
		&incident.Title,
		&incident.Status,
		&incident.Message,
		&incident.StartedAt,
		&incident.ResolvedAt,
		&incident.CreatedAt,
		&incident.UpdatedAt,
	)
	return incident, err
}

func createIncident(db *sql.DB, input IncidentInput) (Incident, error) {
	normalized, err := normalizeIncidentInput(input)
	if err != nil {
		return Incident{}, err
	}
	now := time.Now().Unix()
	startedAt := now
	if normalized.StartedAt != "" {
		parsed, err := parseIncidentTime(normalized.StartedAt)
		if err != nil {
			return Incident{}, err
		}
		startedAt = parsed
	}
	resolvedAt, resolvedAtVal, err := resolveIncidentTime(normalized.ResolvedAt)
	if err != nil {
		return Incident{}, err
	}
	if normalized.Status == "resolved" && !resolvedAt.Valid {
		resolvedAt = sql.NullInt64{Int64: now, Valid: true}
		resolvedAtVal = now
	}

	result, err := db.Exec(`
INSERT INTO incidents (title, status, message, started_at, resolved_at, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)`,
		normalized.Title,
		normalized.Status,
		normalized.Message,
		startedAt,
		resolvedAtVal,
		now,
		now,
	)
	if err != nil {
		return Incident{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Incident{}, err
	}

	return Incident{
		ID:         id,
		Title:      normalized.Title,
		Status:     normalized.Status,
		Message:    normalized.Message,
		StartedAt:  startedAt,
		ResolvedAt: resolvedAt,
		CreatedAt:  now,
		UpdatedAt:  now,
	}, nil
}

func updateIncident(db *sql.DB, id int64, input IncidentUpdate) (Incident, error) {
	existing, err := getIncident(db, id)
	if err != nil {
		return Incident{}, err
	}
	now := time.Now().Unix()
	updated, err := applyIncidentUpdate(existing, input, now)
	if err != nil {
		return Incident{}, err
	}

	var resolvedAtVal interface{} = nil
	if updated.ResolvedAt.Valid {
		resolvedAtVal = updated.ResolvedAt.Int64
	}

	_, err = db.Exec(`
UPDATE incidents
SET title = ?, status = ?, message = ?, started_at = ?, resolved_at = ?, updated_at = ?
WHERE id = ?`,
		updated.Title,
		updated.Status,
		updated.Message,
		updated.StartedAt,
		resolvedAtVal,
		now,
		id,
	)
	if err != nil {
		return Incident{}, err
	}
	updated.UpdatedAt = now
	return updated, nil
}

func deleteIncident(db *sql.DB, id int64) error {
	result, err := db.Exec(`DELETE FROM incidents WHERE id = ?`, id)
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

func normalizeIncidentInput(input IncidentInput) (IncidentInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Message = strings.TrimSpace(input.Message)
	input.Status = strings.ToLower(strings.TrimSpace(input.Status))

	if input.Title == "" {
		return IncidentInput{}, errors.New("title is required")
	}
	if input.Message == "" {
		return IncidentInput{}, errors.New("message is required")
	}
	if input.Status == "" {
		input.Status = defaultIncidentStatus
	}
	if _, ok := incidentStatuses[input.Status]; !ok {
		return IncidentInput{}, errors.New("invalid status")
	}
	return input, nil
}

func applyIncidentUpdate(existing Incident, input IncidentUpdate, now int64) (Incident, error) {
	if input.Title != nil {
		title := strings.TrimSpace(*input.Title)
		if title == "" {
			return Incident{}, errors.New("title is required")
		}
		existing.Title = title
	}
	if input.Message != nil {
		message := strings.TrimSpace(*input.Message)
		if message == "" {
			return Incident{}, errors.New("message is required")
		}
		existing.Message = message
	}
	if input.Status != nil {
		status := strings.ToLower(strings.TrimSpace(*input.Status))
		if status == "" {
			status = defaultIncidentStatus
		}
		if _, ok := incidentStatuses[status]; !ok {
			return Incident{}, errors.New("invalid status")
		}
		existing.Status = status
		if status != "resolved" && input.ResolvedAt == nil {
			existing.ResolvedAt = sql.NullInt64{}
		}
	}
	if input.StartedAt != nil {
		parsed, err := parseIncidentTime(*input.StartedAt)
		if err != nil {
			return Incident{}, err
		}
		existing.StartedAt = parsed
	}
	if input.ResolvedAt != nil {
		resolvedAt, _, err := resolveIncidentTime(*input.ResolvedAt)
		if err != nil {
			return Incident{}, err
		}
		existing.ResolvedAt = resolvedAt
	}
	if existing.Status == "resolved" && !existing.ResolvedAt.Valid {
		existing.ResolvedAt = sql.NullInt64{Int64: now, Valid: true}
	}
	return existing, nil
}

func parseIncidentTime(value string) (int64, error) {
	raw := strings.TrimSpace(value)
	if raw == "" {
		return 0, errors.New("time is required")
	}
	parsed, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return 0, errors.New("invalid time format")
	}
	return parsed.Unix(), nil
}

func resolveIncidentTime(value string) (sql.NullInt64, interface{}, error) {
	raw := strings.TrimSpace(value)
	if raw == "" {
		return sql.NullInt64{}, nil, nil
	}
	parsed, err := parseIncidentTime(raw)
	if err != nil {
		return sql.NullInt64{}, nil, err
	}
	return sql.NullInt64{Int64: parsed, Valid: true}, parsed, nil
}
