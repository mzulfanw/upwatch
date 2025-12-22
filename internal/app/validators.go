package app

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeMonitorInput(input MonitorInput) (MonitorInput, error) {
	input.Name = strings.TrimSpace(input.Name)
	input.URL = strings.TrimSpace(input.URL)
	input.Method = strings.ToUpper(strings.TrimSpace(input.Method))

	if input.Name == "" {
		return MonitorInput{}, errors.New("name is required")
	}
	if input.URL == "" {
		return MonitorInput{}, errors.New("url is required")
	}
	if err := validateURL(input.URL); err != nil {
		return MonitorInput{}, err
	}
	if input.Method == "" {
		input.Method = "GET"
	}
	if err := validateMethod(input.Method); err != nil {
		return MonitorInput{}, err
	}
	if input.IntervalSec == 0 {
		input.IntervalSec = defaultIntervalSec
	}
	if input.TimeoutSec == 0 {
		input.TimeoutSec = defaultTimeoutSec
	}
	if input.IntervalSec <= 0 {
		return MonitorInput{}, errors.New("interval_sec must be greater than 0")
	}
	if input.TimeoutSec <= 0 {
		return MonitorInput{}, errors.New("timeout_sec must be greater than 0")
	}
	return input, nil
}

func validateURL(raw string) error {
	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		return errors.New("invalid url")
	}
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return errors.New("url must start with http or https")
	}
	if parsed.Host == "" {
		return errors.New("invalid url")
	}
	return nil
}

func validateMethod(method string) error {
	switch method {
	case "GET", "HEAD":
		return nil
	default:
		return errors.New("method must be GET or HEAD")
	}
}
