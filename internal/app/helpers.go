package app

import (
	"database/sql"
	"time"
)

func monitorResponse(m Monitor) MonitorResponse {
	return MonitorResponse{
		ID:             m.ID,
		Name:           m.Name,
		URL:            m.URL,
		Method:         m.Method,
		IntervalSec:    m.IntervalSec,
		TimeoutSec:     m.TimeoutSec,
		Status:         m.Status,
		LastStatusCode: nullInt(m.LastStatusCode),
		LastError:      nullString(m.LastError),
		LastLatencyMs:  nullInt(m.LastLatencyMs),
		LastCheckedAt:  nullTime(m.LastCheckedAt),
		CreatedAt:      time.Unix(m.CreatedAt, 0).UTC().Format(time.RFC3339),
		UpdatedAt:      time.Unix(m.UpdatedAt, 0).UTC().Format(time.RFC3339),
	}
}

func eventResponse(e CheckEvent) CheckEventResponse {
	return CheckEventResponse{
		ID:         e.ID,
		MonitorID:  e.MonitorID,
		CheckedAt:  time.Unix(e.CheckedAt, 0).UTC().Format(time.RFC3339),
		Success:    e.Success,
		StatusCode: nullInt(e.StatusCode),
		LatencyMs:  nullInt(e.LatencyMs),
		Error:      nullString(e.Error),
	}
}

func nullInt(value sql.NullInt64) *int {
	if !value.Valid {
		return nil
	}
	v := int(value.Int64)
	return &v
}

func nullString(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}
	v := value.String
	return &v
}

func nullTime(value sql.NullInt64) *string {
	if !value.Valid {
		return nil
	}
	v := time.Unix(value.Int64, 0).UTC().Format(time.RFC3339)
	return &v
}
