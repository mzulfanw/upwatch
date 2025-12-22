package app

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Checker struct {
	db      *sql.DB
	client  *http.Client
	mu      sync.Mutex
	runners map[int64]context.CancelFunc
}

func NewChecker(db *sql.DB) *Checker {
	return &Checker{
		db:      db,
		client:  &http.Client{},
		runners: make(map[int64]context.CancelFunc),
	}
}

func (c *Checker) StartAll(monitors []Monitor) {
	for _, monitor := range monitors {
		c.StartMonitor(monitor)
	}
}

func (c *Checker) StartMonitor(monitor Monitor) {
	if monitor.IntervalSec <= 0 {
		return
	}
	c.StopMonitor(monitor.ID)
	ctx, cancel := context.WithCancel(context.Background())
	c.mu.Lock()
	c.runners[monitor.ID] = cancel
	c.mu.Unlock()

	go c.runMonitor(ctx, monitor)
}

func (c *Checker) StopMonitor(id int64) {
	c.mu.Lock()
	cancel, ok := c.runners[id]
	if ok {
		delete(c.runners, id)
	}
	c.mu.Unlock()
	if ok {
		cancel()
	}
}

func (c *Checker) runMonitor(ctx context.Context, monitor Monitor) {
	c.checkOnce(monitor)
	ticker := time.NewTicker(time.Duration(monitor.IntervalSec) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.checkOnce(monitor)
		}
	}
}

func (c *Checker) checkOnce(monitor Monitor) {
	start := time.Now()
	timeout := time.Duration(monitor.TimeoutSec) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, monitor.Method, monitor.URL, nil)
	if err != nil {
		log.Printf("monitor %d request: %v", monitor.ID, err)
		return
	}

	resp, err := c.client.Do(req)
	latencyMs := int(time.Since(start).Milliseconds())

	success := false
	statusCode := 0
	errMsg := ""

	if err != nil {
		errMsg = err.Error()
	} else {
		statusCode = resp.StatusCode
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
		if statusCode >= 200 && statusCode < 400 {
			success = true
		} else {
			errMsg = fmt.Sprintf("unexpected status %d", statusCode)
		}
	}

	if latencyMs < 0 {
		latencyMs = 0
	}
	if err := recordCheck(c.db, monitor.ID, success, statusCode, latencyMs, errMsg); err != nil {
		log.Printf("monitor %d record: %v", monitor.ID, err)
	}
}
