package app

import (
	"bytes"
	"errors"
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

type Notifier interface {
	Notify(event MonitorNotification) error
}

type MonitorNotification struct {
	MonitorID  int64
	Name       string
	URL        string
	Method     string
	Status     string
	PrevStatus string
	CheckedAt  time.Time
	StatusCode int
	LatencyMs  int
	Error      string
}

type EmailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
	To       []string
}

type EmailNotifier struct {
	addr string
	from string
	to   []string
	auth smtp.Auth
}

func NewEmailNotifier(cfg EmailConfig) (*EmailNotifier, error) {
	if strings.TrimSpace(cfg.Host) == "" {
		return nil, errors.New("smtp host is required")
	}
	if strings.TrimSpace(cfg.Port) == "" {
		return nil, errors.New("smtp port is required")
	}
	if strings.TrimSpace(cfg.From) == "" {
		return nil, errors.New("smtp from is required")
	}
	if len(cfg.To) == 0 {
		return nil, errors.New("smtp to is required")
	}

	addr := cfg.Host + ":" + cfg.Port
	var auth smtp.Auth
	if strings.TrimSpace(cfg.Username) != "" || strings.TrimSpace(cfg.Password) != "" {
		auth = smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	}

	return &EmailNotifier{
		addr: addr,
		from: cfg.From,
		to:   cfg.To,
		auth: auth,
	}, nil
}

func (n *EmailNotifier) Notify(event MonitorNotification) error {
	subject, body := buildEmail(event)
	msg := buildMessage(n.from, n.to, subject, body)
	return smtp.SendMail(n.addr, n.auth, n.from, n.to, msg)
}

func buildEmail(event MonitorNotification) (string, string) {
	status := strings.ToUpper(event.Status)
	isDown := event.Status == "down"
	subjectPrefix := "Incident"
	if !isDown {
		subjectPrefix = "Resolved"
	}
	subject := "[Upwatch] " + subjectPrefix + " - " + event.Name

	checkedAt := event.CheckedAt.UTC().Format(time.RFC3339)
	statusCode := "n/a"
	if event.StatusCode > 0 {
		statusCode = fmt.Sprintf("%d", event.StatusCode)
	}
	errLine := "none"
	if strings.TrimSpace(event.Error) != "" {
		errLine = event.Error
	}

	summary := "Service is unavailable."
	impact := "Users may experience errors or timeouts."
	eventLabel := "Incident started"
	if !isDown {
		summary = "Service has recovered."
		impact = "Service is operating normally."
		eventLabel = "Incident resolved"
	}

	var buf bytes.Buffer
	buf.WriteString("Upwatch Incident Notification\n")
	buf.WriteString("=================================\n\n")
	fmt.Fprintf(&buf, "Event: %s\n", eventLabel)
	fmt.Fprintf(&buf, "Status: %s (previous: %s)\n", status, strings.ToUpper(event.PrevStatus))
	fmt.Fprintf(&buf, "Summary: %s\n", summary)
	fmt.Fprintf(&buf, "Impact: %s\n\n", impact)

	buf.WriteString("Service details\n")
	buf.WriteString("---------------\n")
	fmt.Fprintf(&buf, "Monitor: %s\n", event.Name)
	fmt.Fprintf(&buf, "URL: %s\n", event.URL)
	fmt.Fprintf(&buf, "Method: %s\n\n", event.Method)

	buf.WriteString("Detection details\n")
	buf.WriteString("-----------------\n")
	fmt.Fprintf(&buf, "Checked at: %s\n", checkedAt)
	fmt.Fprintf(&buf, "Latency: %d ms\n", event.LatencyMs)
	fmt.Fprintf(&buf, "Status code: %s\n", statusCode)
	fmt.Fprintf(&buf, "Error: %s\n\n", errLine)

	buf.WriteString("Next update\n")
	buf.WriteString("-----------\n")
	if isDown {
		buf.WriteString("We are investigating and will provide updates as more information becomes available.\n")
	} else {
		buf.WriteString("Monitoring will continue. No further action is required.\n")
	}

	return subject, buf.String()
}

func buildMessage(from string, to []string, subject, body string) []byte {
	var buf bytes.Buffer
	buf.WriteString("From: " + from + "\r\n")
	buf.WriteString("To: " + strings.Join(to, ", ") + "\r\n")
	buf.WriteString("Subject: " + subject + "\r\n")
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	buf.WriteString("\r\n")
	buf.WriteString(body)
	return buf.Bytes()
}

func ParseEmailList(value string) []string {
	if strings.TrimSpace(value) == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}
