package app

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

func defaultSettings() Settings {
	return Settings{
		BrandName:      defaultBrandName,
		BrandTagline:   defaultBrandTagline,
		StatusTitle:    defaultStatusTitle,
		StatusSubtitle: defaultStatusSubtitle,
		UpdatedAt:      time.Now().Unix(),
	}
}

func getSettings(db *sql.DB) (Settings, error) {
	var settings Settings
	err := db.QueryRow(`
SELECT brand_name, brand_tagline, status_title, status_subtitle, updated_at
FROM settings
WHERE id = 1`,
	).Scan(
		&settings.BrandName,
		&settings.BrandTagline,
		&settings.StatusTitle,
		&settings.StatusSubtitle,
		&settings.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return defaultSettings(), nil
	}
	if err != nil {
		return Settings{}, err
	}
	return settings, nil
}

func updateSettings(db *sql.DB, input SettingsUpdate) (Settings, error) {
	if !hasSettingsUpdate(input) {
		return Settings{}, errors.New("no settings provided")
	}
	current, err := getSettings(db)
	if err != nil {
		return Settings{}, err
	}
	updated := applySettingsUpdate(current, input)
	now := time.Now().Unix()
	updated.UpdatedAt = now

	_, err = db.Exec(`
INSERT INTO settings (id, brand_name, brand_tagline, status_title, status_subtitle, updated_at)
VALUES (1, ?, ?, ?, ?, ?)
ON CONFLICT(id) DO UPDATE SET
	brand_name = excluded.brand_name,
	brand_tagline = excluded.brand_tagline,
	status_title = excluded.status_title,
	status_subtitle = excluded.status_subtitle,
	updated_at = excluded.updated_at`,
		updated.BrandName,
		updated.BrandTagline,
		updated.StatusTitle,
		updated.StatusSubtitle,
		now,
	)
	if err != nil {
		return Settings{}, err
	}
	return updated, nil
}

func applySettingsUpdate(current Settings, input SettingsUpdate) Settings {
	if input.BrandName != nil {
		current.BrandName = strings.TrimSpace(*input.BrandName)
	}
	if input.BrandTagline != nil {
		current.BrandTagline = strings.TrimSpace(*input.BrandTagline)
	}
	if input.StatusTitle != nil {
		current.StatusTitle = strings.TrimSpace(*input.StatusTitle)
	}
	if input.StatusSubtitle != nil {
		current.StatusSubtitle = strings.TrimSpace(*input.StatusSubtitle)
	}
	return current
}

func hasSettingsUpdate(input SettingsUpdate) bool {
	return input.BrandName != nil ||
		input.BrandTagline != nil ||
		input.StatusTitle != nil ||
		input.StatusSubtitle != nil
}
