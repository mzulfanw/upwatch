# Upwatch

![Upwatch logo](assets/upwatch.png)

Minimal uptime monitor with a public status page, admin dashboard, incidents, and SSE live updates.

## Features
- Public status page with live monitor data and incidents
- Admin dashboard to manage monitors and incidents
- Settings page for branding text on the status page
- SQLite storage (easy to run anywhere)
- SSE stream for real-time status updates

## Quick start (Docker)
```bash
docker pull mzulfanw/upwatch:latest
docker run -d \
  -e ADMIN_USER=admin \
  -e ADMIN_PASSWORD=admin \
  -p 8080:8080 \
  mzulfanw/upwatch:latest
```

Default port is `8080`. Open:
- `http://localhost:8080/` (public status)
- `http://localhost:8080/login` (admin login)

## Docker (build locally)
```bash
docker build -t upwatch .
docker run -d \
  -e ADMIN_USER=admin \
  -e ADMIN_PASSWORD=admin \
  -p 8080:8080 \
  upwatch
```

## Docker Compose
```bash
docker compose up -d
```

Make sure `.env` exists (see `.env.example`).

## Configuration
Environment variables:
- `PORT` (default `8080`)
- `DB_PATH` (default `data/upwatch.db`)
- `ADMIN_USER` (required)
- `ADMIN_PASSWORD` (required)
- `SESSION_TTL` (default `24h`)

## Settings page
Open `http://localhost:8080/settings` after login to update:
- Brand name
- Tagline
- Status headline
- Status subtitle

Changes apply to the public status page automatically.

## Data persistence
SQLite is stored at `DB_PATH`. For Docker, mount a volume:
```bash
docker run -d \
  -e ADMIN_USER=admin \
  -e ADMIN_PASSWORD=admin \
  -v $(pwd)/data:/data \
  -e DB_PATH=/data/upwatch.db \
  -p 8080:8080 \
  mzulfanw/upwatch:latest
```

## Local development (optional)
```bash
go run ./cmd
```

## Notes
- Static assets are served from `/assets` (see `assets/upwatch.png`).
- SSE endpoint: `/api/status/stream`.
