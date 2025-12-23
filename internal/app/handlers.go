package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (a *App) handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (a *App) handleListMonitors(w http.ResponseWriter, r *http.Request) {
	monitors, err := listMonitors(a.db)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list monitors")
		return
	}
	resp := make([]MonitorResponse, 0, len(monitors))
	for _, monitor := range monitors {
		resp = append(resp, monitorResponse(monitor))
	}
	writeJSON(w, http.StatusOK, resp)
}

func (a *App) handleGetMonitor(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid monitor id")
		return
	}
	monitor, err := getMonitor(a.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "monitor not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to load monitor")
		return
	}
	writeJSON(w, http.StatusOK, monitorResponse(monitor))
}

func (a *App) handleCreateMonitor(w http.ResponseWriter, r *http.Request) {
	var input MonitorInput
	if err := decodeJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	monitor, err := createMonitor(a.db, input)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	a.checker.StartMonitor(monitor)
	writeJSON(w, http.StatusCreated, monitorResponse(monitor))
}

func (a *App) handleUpdateMonitor(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid monitor id")
		return
	}
	var input MonitorUpdate
	if err := decodeJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	monitor, err := updateMonitor(a.db, id, input)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "monitor not found")
			return
		}
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	a.checker.StartMonitor(monitor)
	writeJSON(w, http.StatusOK, monitorResponse(monitor))
}

func (a *App) handleDeleteMonitor(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid monitor id")
		return
	}
	if err := deleteMonitor(a.db, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "monitor not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete monitor")
		return
	}
	a.checker.StopMonitor(id)
	writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func (a *App) handleMonitorEvents(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid monitor id")
		return
	}
	limit := parseLimit(r.URL.Query().Get("limit"))
	events, err := listEvents(a.db, id, limit)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to load events")
		return
	}
	resp := make([]CheckEventResponse, 0, len(events))
	for _, event := range events {
		resp = append(resp, eventResponse(event))
	}
	writeJSON(w, http.StatusOK, resp)
}

func (a *App) handleListIncidents(w http.ResponseWriter, r *http.Request) {
	limit := parseLimit(r.URL.Query().Get("limit"))
	if limit > maxIncidentsLimit {
		limit = maxIncidentsLimit
	}
	incidents, err := listIncidents(a.db, limit)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to load incidents")
		return
	}
	resp := make([]IncidentResponse, 0, len(incidents))
	for _, incident := range incidents {
		resp = append(resp, incidentResponse(incident))
	}
	writeJSON(w, http.StatusOK, resp)
}

func (a *App) handleCreateIncident(w http.ResponseWriter, r *http.Request) {
	var input IncidentInput
	if err := decodeJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	incident, err := createIncident(a.db, input)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, incidentResponse(incident))
}

func (a *App) handleUpdateIncident(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid incident id")
		return
	}
	var input IncidentUpdate
	if err := decodeJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	incident, err := updateIncident(a.db, id, input)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "incident not found")
			return
		}
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, incidentResponse(incident))
}

func (a *App) handleDeleteIncident(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid incident id")
		return
	}
	if err := deleteIncident(a.db, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writeError(w, http.StatusNotFound, "incident not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete incident")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "deleted"})
}

func (a *App) handleGetSettings(w http.ResponseWriter, r *http.Request) {
	settings, err := getSettings(a.db)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to load settings")
		return
	}
	writeJSON(w, http.StatusOK, settingsResponse(settings))
}

func (a *App) handleUpdateSettings(w http.ResponseWriter, r *http.Request) {
	var input SettingsUpdate
	if err := decodeJSON(r, &input); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	settings, err := updateSettings(a.db, input)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, settingsResponse(settings))
}

func (a *App) buildStatusPayload(includeHistory bool) (StatusResponse, error) {
	monitors, err := listMonitors(a.db)
	if err != nil {
		return StatusResponse{}, err
	}
	counts := StatusCounts{Total: len(monitors)}
	respMonitors := make([]MonitorResponse, 0, len(monitors))
	for _, monitor := range monitors {
		switch monitor.Status {
		case "up":
			counts.Up++
		case "down":
			counts.Down++
		default:
			counts.Unknown++
		}
		respMonitors = append(respMonitors, monitorResponse(monitor))
	}
	resp := StatusResponse{
		Counts:   counts,
		Monitors: respMonitors,
		Updated:  time.Now().UTC().Format(time.RFC3339),
	}
	if includeHistory {
		history, err := a.buildStatusHistory(monitors)
		if err != nil {
			return StatusResponse{}, err
		}
		resp.History = history
	}
	return resp, nil
}

func (a *App) handleStatus(w http.ResponseWriter, r *http.Request) {
	payload, err := a.buildStatusPayload(true)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to load status")
		return
	}
	writeJSON(w, http.StatusOK, payload)
}

func (a *App) handleStatusStream(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		writeError(w, http.StatusInternalServerError, "streaming unsupported")
		return
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)

	write := func(format string, args ...any) bool {
		if _, err := fmt.Fprintf(w, format, args...); err != nil {
			return false
		}
		flusher.Flush()
		return true
	}

	send := func(includeHistory bool) bool {
		payload, err := a.buildStatusPayload(includeHistory)
		if err != nil {
			return write("event: error\ndata: %q\n\n", err.Error())
		}
		data, err := json.Marshal(payload)
		if err != nil {
			return write("event: error\ndata: %q\n\n", err.Error())
		}
		return write("data: %s\n\n", data)
	}

	if !write("retry: 5000\n\n") {
		return
	}
	if !send(true) {
		return
	}

	statusTicker := time.NewTicker(15 * time.Second)
	keepAliveTicker := time.NewTicker(10 * time.Second)
	defer statusTicker.Stop()
	defer keepAliveTicker.Stop()

	for {
		select {
		case <-r.Context().Done():
			return
		case <-statusTicker.C:
			if !send(false) {
				return
			}
		case <-keepAliveTicker.C:
			if !write(": ping\n\n") {
				return
			}
		}
	}
}

func (a *App) buildStatusHistory(monitors []Monitor) (map[string][]StatusHistoryPoint, error) {
	history := make(map[string][]StatusHistoryPoint, len(monitors))
	for _, monitor := range monitors {
		events, err := listEvents(a.db, monitor.ID, statusHistoryLimit)
		if err != nil {
			return nil, err
		}
		points := make([]StatusHistoryPoint, 0, len(events))
		for i := len(events) - 1; i >= 0; i-- {
			event := events[i]
			if event.LatencyMs.Valid {
				value := int(event.LatencyMs.Int64)
				points = append(points, StatusHistoryPoint{V: &value})
			} else {
				points = append(points, StatusHistoryPoint{V: nil})
			}
		}
		history[strconv.FormatInt(monitor.ID, 10)] = points
	}
	return history, nil
}
