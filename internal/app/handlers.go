package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func (a *App) buildStatusPayload() (StatusResponse, error) {
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
	return StatusResponse{
		Counts:   counts,
		Monitors: respMonitors,
		Updated:  time.Now().UTC().Format(time.RFC3339),
	}, nil
}

func (a *App) handleStatus(w http.ResponseWriter, r *http.Request) {
	payload, err := a.buildStatusPayload()
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

	send := func() {
		payload, err := a.buildStatusPayload()
		if err != nil {
			fmt.Fprintf(w, "event: error\ndata: %q\n\n", err.Error())
			flusher.Flush()
			return
		}
		data, err := json.Marshal(payload)
		if err != nil {
			fmt.Fprintf(w, "event: error\ndata: %q\n\n", err.Error())
			flusher.Flush()
			return
		}
		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()
	}

	send()
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			return
		case <-ticker.C:
			send()
		}
	}
}
