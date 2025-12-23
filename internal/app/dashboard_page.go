package app

import "net/http"

func serveDashboardPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(dashboardPageHTML))
}

const dashboardPageHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Upwatch Dashboard</title>
  <meta name="description" content="Admin dashboard for managing monitors and incidents." />
  <meta name="robots" content="noindex,nofollow" />
  <meta property="og:title" content="Upwatch Dashboard" />
  <meta property="og:description" content="Admin dashboard for managing monitors and incidents." />
  <meta property="og:type" content="website" />
  <meta property="og:site_name" content="Upwatch" />
  <meta name="twitter:title" content="Upwatch Dashboard" />
  <meta name="twitter:description" content="Admin dashboard for managing monitors and incidents." />
  <link rel="icon" type="image/png" href="/assets/upwatch.png" />
  <meta property="og:image" content="/assets/upwatch.png" />
  <meta name="twitter:card" content="summary_large_image" />
  <meta name="twitter:image" content="/assets/upwatch.png" />
  <style>
    @import url("https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;600;700&display=swap");
    :root {
      color-scheme: light dark;
      --bg: #141515;
      --bg-2: #141515;
      --surface: #202122;
      --surface-2: #27292a;
      --text: #f2f2f0;
      --muted: #b3b3ad;
      --up: #7fb58a;
      --down: #d77a6a;
      --unknown: #d9b56d;
      --accent: #c9a27d;
      --accent-2: #b9856b;
      --danger: #d77a6a;
      --grid: rgba(255, 255, 255, 0.08);
    }
    [data-theme="light"] {
      color-scheme: light;
      --bg: #f6f6f4;
      --bg-2: #ededeb;
      --surface: #ffffff;
      --surface-2: #f4f4f2;
      --text: #2a241f;
      --muted: #6f6f6a;
      --accent: #e99673;
      --accent-2: #de7c69;
      --danger: #c06355;
      --grid: rgba(40, 33, 28, 0.08);
    }
    * { box-sizing: border-box; }
    body {
      margin: 0;
      min-height: 100vh;
      font-family: "JetBrains Mono", "SFMono-Regular", Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
      color: var(--text);
      background: var(--bg);
      position: relative;
      overflow-x: hidden;
    }
    [data-theme="light"] body {
      background: var(--bg);
    }
    body::before,
    body::after {
      display: none;
    }
    .shell {
      max-width: 1200px;
      margin: 0 auto;
      padding: 3rem 2rem 4.5rem;
      position: relative;
      z-index: 1;
    }
    .topbar {
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 1.5rem;
      margin-bottom: 3rem;
      flex-wrap: wrap;
    }
    .brand {
      display: flex;
      align-items: center;
      gap: 1rem;
    }
    .brand-logo {
      width: 64px;
      height: 64px;
      border-radius: 16px;
      overflow: hidden;
      flex: 0 0 auto;
    }
    .brand-logo img {
      width: 100%;
      height: 100%;
      object-fit: contain;
      display: block;
    }
    .brand-text {
      display: grid;
      gap: 0.35rem;
    }
    .brand-mark {
      font-size: clamp(2rem, 4vw, 3.2rem);
      font-weight: 700;
      letter-spacing: -0.01em;
    }
    .brand-sub {
      color: var(--muted);
      font-size: 0.85rem;
      line-height: 1.6;
    }
    .actions {
      display: flex;
      align-items: center;
      gap: 0.75rem;
      flex-wrap: wrap;
    }
    .btn {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      padding: 0.6rem 1.2rem;
      border-radius: 999px;
      border: 1px solid transparent;
      font-weight: 600;
      font-size: 0.85rem;
      line-height: 1;
      font-family: inherit;
      min-height: 42px;
      cursor: pointer;
      text-decoration: none;
      transition: transform 0.2s ease, border-color 0.2s ease;
    }
    .btn-icon {
      width: 42px;
      height: 42px;
      padding: 0;
      position: relative;
    }
    .btn-primary {
      background: linear-gradient(120deg, var(--accent), var(--accent-2));
      color: white;
    }
    .btn-ghost {
      background: rgba(255, 255, 255, 0.08);
      color: var(--text);
      border-color: rgba(255, 255, 255, 0.14);
    }
    [data-theme="light"] .btn-ghost {
      background: rgba(16, 32, 80, 0.06);
      border-color: rgba(16, 32, 80, 0.12);
    }
    .btn:hover {
      transform: translateY(-1px);
    }
    .theme-icon {
      width: 18px;
      height: 18px;
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%) scale(1) rotate(0deg);
      transition: opacity 0.2s ease, transform 0.2s ease;
    }
    .theme-icon.sun {
      opacity: 1;
    }
    .theme-icon.moon {
      opacity: 0;
      transform: translate(-50%, -50%) scale(0.8) rotate(35deg);
    }
    [data-theme="light"] .theme-icon.sun {
      opacity: 0;
      transform: translate(-50%, -50%) scale(0.8) rotate(-35deg);
    }
    [data-theme="light"] .theme-icon.moon {
      opacity: 1;
      transform: translate(-50%, -50%) scale(1) rotate(0deg);
    }
    .sr-only {
      position: absolute;
      width: 1px;
      height: 1px;
      padding: 0;
      margin: -1px;
      overflow: hidden;
      clip: rect(0, 0, 0, 0);
      border: 0;
    }
    .panel {
      background: var(--surface);
      border: 1px solid var(--card-border);
      border-radius: 20px;
      padding: 2rem;
    }
    [data-theme="light"] .panel {
      background: var(--surface);
    }
    .panel + .panel {
      margin-top: 2.5rem;
    }
    .panel + section {
      margin-top: 2.5rem;
    }
    .panel-title {
      margin: 0 0 0.5rem;
      font-size: 1.3rem;
      line-height: 1.4;
    }
    .panel-subtitle {
      margin: 0 0 1.25rem;
      color: var(--muted);
      font-size: 0.9rem;
      line-height: 1.6;
    }
    .insights {
      display: grid;
      gap: 1.5rem;
      grid-template-columns: minmax(0, 1fr);
      margin-bottom: 2rem;
    }
    .hero {
      display: grid;
      gap: 1.4rem;
    }
    .eyebrow {
      text-transform: uppercase;
      letter-spacing: 0.18em;
      font-size: 0.7rem;
      color: var(--muted);
    }
    .stat-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(140px, 1fr));
      gap: 1rem;
    }
    .stat-card {
      background: var(--bg);
      border: 1px solid rgba(255, 255, 255, 0.08);
      border-radius: 14px;
      padding: 0.9rem 1rem;
    }
    [data-theme="light"] .stat-card {
      background: rgba(16, 32, 80, 0.06);
      border-color: rgba(16, 32, 80, 0.08);
    }
    .stat-label {
      font-size: 0.75rem;
      color: var(--muted);
    }
    .stat-value {
      font-size: 1.6rem;
      font-weight: 700;
      margin-top: 0.3rem;
    }
    .stat-value.up { color: var(--up); }
    .stat-value.down { color: var(--down); }
    .stat-value.unknown { color: var(--unknown); }

    .chart-grid {
      display: grid;
      gap: 1.5rem;
      grid-template-columns: minmax(220px, 0.9fr) minmax(280px, 1.1fr);
    }
    .chart-card {
      background: var(--bg);
      border: 1px solid rgba(255, 255, 255, 0.08);
      border-radius: 16px;
      padding: 1rem 1.2rem;
      min-height: 260px;
      display: flex;
      flex-direction: column;
      gap: 0.8rem;
    }
    [data-theme="light"] .chart-card {
      background: rgba(16, 32, 80, 0.05);
      border-color: rgba(16, 32, 80, 0.08);
    }
    .chart-card h3 {
      margin: 0;
      font-size: 1.1rem;
    }
    .chart-row {
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 0.8rem;
      flex-wrap: wrap;
    }
    .chart-legend {
      display: grid;
      gap: 0.4rem;
      font-size: 0.85rem;
      color: var(--muted);
    }
    .legend-item {
      display: flex;
      align-items: center;
      gap: 0.5rem;
    }
    .legend-swatch {
      width: 10px;
      height: 10px;
      border-radius: 50%;
    }
    canvas {
      width: 100%;
      height: 220px;
      display: block;
    }
    .chart-meta {
      display: flex;
      gap: 0.5rem;
      flex-wrap: wrap;
    }
    .chip {
      padding: 0.35rem 0.75rem;
      border-radius: 999px;
      background: rgba(255, 255, 255, 0.08);
      border: 1px solid rgba(255, 255, 255, 0.12);
      font-size: 0.8rem;
    }
    [data-theme="light"] .chip {
      background: rgba(16, 32, 80, 0.06);
      border-color: rgba(16, 32, 80, 0.12);
    }

    form {
      display: grid;
      gap: 0.8rem;
      grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
      align-items: end;
    }
    label {
      font-size: 0.85rem;
      color: var(--muted);
    }
    input, select, textarea {
      width: 100%;
      padding: 0.65rem 0.75rem;
      border-radius: 10px;
      border: 1px solid rgba(255, 255, 255, 0.15);
      background: var(--bg);
      color: var(--text);
    }
    textarea {
      min-height: 96px;
      resize: vertical;
    }
    [data-theme="light"] input,
    [data-theme="light"] select,
    [data-theme="light"] textarea {
      background: rgba(16, 32, 80, 0.06);
      border-color: rgba(16, 32, 80, 0.12);
    }
    .field-full {
      grid-column: 1 / -1;
    }
    .grid {
      display: flex;
      flex-direction: column;
      gap: 1.25rem;
      margin-top: 1.75rem;
    }
    .card {
      background: rgba(13, 16, 45, 0.85);
      border: 1px solid rgba(255, 255, 255, 0.08);
      border-radius: 18px;
      padding: 1.35rem;
      position: relative;
      overflow: hidden;
      animation: fadeUp 0.6s ease both;
      display: flex;
      flex-direction: column;
    }
    [data-theme="light"] .card {
      background: rgba(255, 255, 255, 0.9);
      border-color: rgba(16, 32, 80, 0.12);
    }
    .card::after {
      content: "";
      position: absolute;
      inset: 0;
      background: radial-gradient(circle at top right, rgba(91, 226, 255, 0.12), transparent 55%);
      pointer-events: none;
    }
    .card[data-status="up"] { border-color: rgba(42, 211, 165, 0.45); }
    .card[data-status="down"] { border-color: rgba(255, 100, 124, 0.45); }
    .card[data-status="unknown"] { border-color: rgba(247, 200, 91, 0.45); }
    .card h3 {
      margin: 0 0 0.4rem;
      font-size: 1.05rem;
    }
    .meta {
      font-size: 0.85rem;
      color: var(--muted);
      line-height: 1.4;
    }
    .status {
      display: inline-flex;
      align-items: center;
      gap: 0.4rem;
      font-weight: 600;
      margin-top: 0.7rem;
    }
    .dot {
      width: 10px;
      height: 10px;
      border-radius: 50%;
      display: inline-block;
    }
    .dot.up { background: var(--up); box-shadow: 0 0 10px rgba(42, 211, 165, 0.6); }
    .dot.down { background: var(--down); box-shadow: 0 0 10px rgba(255, 100, 124, 0.6); }
    .dot.unknown { background: var(--unknown); box-shadow: 0 0 10px rgba(247, 200, 91, 0.6); }
    .card-actions {
      margin-top: 0.9rem;
    }
    .danger {
      background: transparent;
      color: var(--danger);
      border: 1px solid rgba(255, 100, 124, 0.5);
      padding: 0.4rem 0.75rem;
      border-radius: 999px;
      cursor: pointer;
      font-size: 0.85rem;
      line-height: 1;
      font-family: inherit;
      position: relative;
      z-index: 1;
    }
    .btn-small {
      padding: 0.35rem 0.8rem;
      border-radius: 999px;
      font-size: 0.85rem;
      line-height: 1;
    }
    .incident-grid {
      margin-top: 0;
    }
    .card.incident-card {
      border: 1px solid rgba(215, 122, 106, 0.55);
    }
    [data-theme="light"] .card.incident-card {
      border: 1px solid rgba(192, 99, 85, 0.55);
    }
    .card.incident-card::after {
      background: radial-gradient(circle at top right, rgba(215, 122, 106, 0.18), transparent 55%);
    }
    .incident-top {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
      gap: 0.35rem;
    }
    .incident-title {
      font-weight: 600;
      font-size: 1.02rem;
    }
    .incident-meta {
      color: var(--muted);
      font-size: 0.85rem;
      margin-top: 0.3rem;
    }
    .incident-message {
      margin: 0.75rem 0 0;
      line-height: 1.5;
    }
    .incident-badge {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      padding: 0.3rem 0.7rem;
      border-radius: 999px;
      font-size: 0.7rem;
      text-transform: uppercase;
      letter-spacing: 0.12em;
      border: 1px solid transparent;
    }
    .incident-badge.investigating {
      color: var(--down);
      border-color: rgba(215, 122, 106, 0.5);
      background: rgba(215, 122, 106, 0.12);
    }
    .incident-badge.identified {
      color: var(--down);
      border-color: rgba(215, 122, 106, 0.5);
      background: rgba(215, 122, 106, 0.12);
    }
    .incident-badge.monitoring {
      color: var(--unknown);
      border-color: rgba(217, 181, 109, 0.5);
      background: rgba(217, 181, 109, 0.12);
    }
    .incident-badge.resolved {
      color: var(--up);
      border-color: rgba(127, 181, 138, 0.5);
      background: rgba(127, 181, 138, 0.12);
    }
    .incident-badge.maintenance,
    .incident-badge.scheduled {
      color: var(--unknown);
      border-color: rgba(217, 181, 109, 0.5);
      background: rgba(217, 181, 109, 0.12);
    }
    .incident-actions {
      display: flex;
      flex-wrap: wrap;
      gap: 0.5rem;
      margin-top: 0.75rem;
    }
    .footer {
      margin-top: 3.5rem;
      text-align: center;
      color: var(--muted);
      font-size: 0.75rem;
      letter-spacing: 0.18em;
      text-transform: uppercase;
    }
    .footer-link {
      color: inherit;
      text-decoration: none;
      border-bottom: 1px solid transparent;
      padding-bottom: 0.1rem;
    }
    .footer-link:hover {
      border-bottom-color: currentColor;
    }
    .error {
      color: var(--danger);
      font-size: 0.9rem;
      margin-top: 0.8rem;
      min-height: 1rem;
    }
    .section-head {
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 1.25rem;
      flex-wrap: wrap;
      margin-top: 2.5rem;
      margin-bottom: 0.75rem;
    }
    .section-title {
      margin: 0;
    }
    .muted {
      color: var(--muted);
      margin: 0.25rem 0 0;
      line-height: 1.5;
    }

    @keyframes fadeUp {
      from { opacity: 0; transform: translateY(10px); }
      to { opacity: 1; transform: translateY(0); }
    }

    @media (max-width: 900px) {
      .chart-grid {
        grid-template-columns: 1fr;
      }
      .actions {
        width: 100%;
        justify-content: flex-start;
      }
      .shell {
        padding: 2.5rem 1.25rem 3.5rem;
      }
      .panel {
        padding: 1.6rem;
      }
    }
  </style>
</head>
<body>
  <div class="shell">
    <div class="topbar">
      <div class="brand">
        <div class="brand-logo">
          <img src="/assets/upwatch.png" alt="Upwatch logo" />
        </div>
        <div class="brand-text">
          <div class="brand-mark">Upwatch Dashboard</div>
          <div class="brand-sub">Self hosted uptime command center.</div>
        </div>
      </div>
      <div class="actions">
        <button class="btn btn-ghost btn-icon" id="themeToggle" type="button" aria-label="Switch to light mode" title="Switch to light mode">
          <svg class="theme-icon sun" viewBox="0 0 24 24" aria-hidden="true" focusable="false">
            <circle cx="12" cy="12" r="4" fill="none" stroke="currentColor" stroke-width="1.8" />
            <line x1="12" y1="2.5" x2="12" y2="5.2" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" />
            <line x1="12" y1="18.8" x2="12" y2="21.5" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" />
            <line x1="2.5" y1="12" x2="5.2" y2="12" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" />
            <line x1="18.8" y1="12" x2="21.5" y2="12" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" />
            <line x1="4.6" y1="4.6" x2="6.6" y2="6.6" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" />
            <line x1="17.4" y1="17.4" x2="19.4" y2="19.4" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" />
            <line x1="4.6" y1="19.4" x2="6.6" y2="17.4" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" />
            <line x1="17.4" y1="6.6" x2="19.4" y2="4.6" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" />
          </svg>
          <svg class="theme-icon moon" viewBox="0 0 24 24" aria-hidden="true" focusable="false">
            <path d="M21 14a8.5 8.5 0 1 1-11-11 7 7 0 0 0 11 11z" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
          <span class="sr-only">Toggle theme</span>
        </button>
        <a class="btn btn-ghost" href="/settings">Settings</a>
        <button class="btn btn-ghost" id="refreshBtn" type="button">Refresh</button>
        <a class="btn btn-primary" href="/logout">Log out</a>
      </div>
    </div>

    <section class="insights">
      <div class="panel hero">
        <div>
          <div class="eyebrow">System pulse</div>
          <h2 class="panel-title">Monitor health at a glance</h2>
          <p class="panel-subtitle">Live signal from every endpoint you track.</p>
        </div>
        <div class="stat-grid">
          <div class="stat-card">
            <div class="stat-label">Total monitors</div>
            <div class="stat-value" id="countTotal">0</div>
          </div>
          <div class="stat-card">
            <div class="stat-label">Up</div>
            <div class="stat-value up" id="countUp">0</div>
          </div>
          <div class="stat-card">
            <div class="stat-label">Down</div>
            <div class="stat-value down" id="countDown">0</div>
          </div>
          <div class="stat-card">
            <div class="stat-label">Unknown</div>
            <div class="stat-value unknown" id="countUnknown">0</div>
          </div>
        </div>
      </div>

      <div class="panel">
        <div class="chart-grid">
          <div class="chart-card">
            <h3>Health overview</h3>
            <div class="meta">Current status distribution.</div>
            <canvas id="statusChart"></canvas>
            <div class="chart-legend">
              <div class="legend-item"><span class="legend-swatch" style="background: var(--up);"></span>Up <span id="legendUp">0</span></div>
              <div class="legend-item"><span class="legend-swatch" style="background: var(--down);"></span>Down <span id="legendDown">0</span></div>
              <div class="legend-item"><span class="legend-swatch" style="background: var(--unknown);"></span>Unknown <span id="legendUnknown">0</span></div>
            </div>
          </div>
          <div class="chart-card">
            <div class="chart-row">
              <div>
                <h3>Response timeline</h3>
                <div class="meta">Recent checks for selected monitor.</div>
              </div>
              <select id="focusMonitor"></select>
            </div>
            <canvas id="latencyChart"></canvas>
            <div class="chart-meta">
              <div class="chip" id="latencyHint">No monitor selected</div>
              <div class="chip" id="latencyLast">Last latency: n/a</div>
              <div class="chip" id="latencyUptime">Uptime: n/a</div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section class="panel">
      <h2 class="panel-title">Add monitor</h2>
      <p class="panel-subtitle">Provide the URL and timing details.</p>
      <form id="monitorForm">
        <div>
          <label for="name">Name</label>
          <input id="name" name="name" type="text" required />
        </div>
        <div>
          <label for="url">URL</label>
          <input id="url" name="url" type="url" placeholder="https://example.com" required />
        </div>
        <div>
          <label for="method">Method</label>
          <select id="method" name="method">
            <option value="GET">GET</option>
            <option value="HEAD">HEAD</option>
          </select>
        </div>
        <div>
          <label for="interval">Interval (sec)</label>
          <input id="interval" name="interval" type="number" min="5" value="60" />
        </div>
        <div>
          <label for="timeout">Timeout (sec)</label>
          <input id="timeout" name="timeout" type="number" min="1" value="10" />
        </div>
        <div>
          <button class="btn btn-primary" type="submit">Add monitor</button>
        </div>
      </form>
      <div class="error" id="formError"></div>
    </section>

    <section class="panel">
      <h2 class="panel-title">Incidents</h2>
      <p class="panel-subtitle">Log outages and maintenance updates.</p>
      <form id="incidentForm">
        <div>
          <label for="incidentTitle">Title</label>
          <input id="incidentTitle" name="incidentTitle" type="text" required />
        </div>
        <div>
          <label for="incidentStatus">Status</label>
          <select id="incidentStatus" name="incidentStatus">
            <option value="investigating">Investigating</option>
            <option value="identified">Identified</option>
            <option value="monitoring">Monitoring</option>
            <option value="maintenance">Maintenance</option>
            <option value="scheduled">Scheduled</option>
            <option value="resolved">Resolved</option>
          </select>
        </div>
        <div>
          <label for="incidentStarted">Started at (optional)</label>
          <input id="incidentStarted" name="incidentStarted" type="datetime-local" />
        </div>
        <div>
          <label for="incidentResolved">Resolved at (optional)</label>
          <input id="incidentResolved" name="incidentResolved" type="datetime-local" />
        </div>
        <div class="field-full">
          <label for="incidentMessage">Message</label>
          <textarea id="incidentMessage" name="incidentMessage" rows="4" required></textarea>
        </div>
        <div>
          <button class="btn btn-primary" type="submit">Publish incident</button>
        </div>
      </form>
      <div class="error" id="incidentError"></div>
    </section>

    <section>
      <div class="section-head">
        <div>
          <h2 class="section-title">Past incidents</h2>
          <p class="muted">Resolved disruptions and scheduled work.</p>
        </div>
      </div>
      <div class="grid incident-grid" id="incidentGrid"></div>
    </section>

    <section>
      <div class="section-head">
        <div>
          <h2 class="section-title">Active monitors</h2>
          <p class="muted">Live status for all tracked endpoints.</p>
        </div>
        <div class="chip" id="lastRefresh">Last refresh: never</div>
      </div>
      <div class="grid" id="monitorGrid"></div>
    </section>
    <footer class="footer">Powered by <a class="footer-link" href="https://github.com/mzulfanw/upwatch" target="_blank" rel="noopener">Upwatch</a></footer>
  </div>

  <script>
    const gridEl = document.getElementById("monitorGrid");
    const formEl = document.getElementById("monitorForm");
    const errorEl = document.getElementById("formError");
    const refreshBtn = document.getElementById("refreshBtn");
    const statusCanvas = document.getElementById("statusChart");
    const latencyCanvas = document.getElementById("latencyChart");
    const focusSelect = document.getElementById("focusMonitor");
    const lastRefreshEl = document.getElementById("lastRefresh");
    const themeToggle = document.getElementById("themeToggle");

    const countTotalEl = document.getElementById("countTotal");
    const countUpEl = document.getElementById("countUp");
    const countDownEl = document.getElementById("countDown");
    const countUnknownEl = document.getElementById("countUnknown");
    const legendUpEl = document.getElementById("legendUp");
    const legendDownEl = document.getElementById("legendDown");
    const legendUnknownEl = document.getElementById("legendUnknown");
    const latencyHintEl = document.getElementById("latencyHint");
    const latencyLastEl = document.getElementById("latencyLast");
    const latencyUptimeEl = document.getElementById("latencyUptime");
    const incidentFormEl = document.getElementById("incidentForm");
    const incidentErrorEl = document.getElementById("incidentError");
    const incidentGridEl = document.getElementById("incidentGrid");
    const incidentTitleEl = document.getElementById("incidentTitle");
    const incidentStatusEl = document.getElementById("incidentStatus");
    const incidentStartedEl = document.getElementById("incidentStarted");
    const incidentResolvedEl = document.getElementById("incidentResolved");
    const incidentMessageEl = document.getElementById("incidentMessage");

    const state = {
      monitors: [],
      counts: { up: 0, down: 0, unknown: 0, total: 0 },
      events: [],
      incidents: []
    };

    const getVar = (name, fallback) => {
      const value = getComputedStyle(document.documentElement).getPropertyValue(name).trim();
      return value || fallback;
    };

    const updateThemeToggle = (theme) => {
      const next = theme === "light" ? "dark" : "light";
      const label = "Switch to " + next + " mode";
      themeToggle.setAttribute("aria-label", label);
      themeToggle.setAttribute("title", label);
    };

    const setTheme = (theme) => {
      document.documentElement.setAttribute("data-theme", theme);
      updateThemeToggle(theme);
      localStorage.setItem("upwatch-theme", theme);
      drawStatusChart(state.counts);
      drawLatencyChart(state.events);
    };

    const initTheme = () => {
      const stored = localStorage.getItem("upwatch-theme");
      if (stored === "light" || stored === "dark") {
        setTheme(stored);
        return;
      }
      const prefersLight = window.matchMedia && window.matchMedia("(prefers-color-scheme: light)").matches;
      setTheme(prefersLight ? "light" : "dark");
    };

    const statusLabel = (status) => {
      if (status === "up") return "Up";
      if (status === "down") return "Down";
      return "Unknown";
    };

    const incidentStatusLabel = (status) => {
      if (!status) return "Unknown";
      return status.replace(/_/g, " ").replace(/\b\w/g, (char) => char.toUpperCase());
    };

    const formatTime = (value) => {
      if (!value) return "never";
      const date = new Date(value);
      if (Number.isNaN(date.getTime())) return "never";
      return date.toLocaleString();
    };

    const formatLatency = (value) => {
      if (value === null || value === undefined) return "n/a";
      return value + " ms";
    };

    const incidentWindow = (incident) => {
      const started = formatTime(incident.started_at);
      if (incident.resolved_at) {
        return "Started " + started + " · Resolved " + formatTime(incident.resolved_at);
      }
      return "Started " + started + " · Ongoing";
    };

    const toISO = (value) => {
      if (!value) return "";
      const date = new Date(value);
      if (Number.isNaN(date.getTime())) return "";
      return date.toISOString();
    };

    const setupCanvas = (canvas) => {
      const rect = canvas.getBoundingClientRect();
      const ratio = window.devicePixelRatio || 1;
      canvas.width = Math.max(1, Math.floor(rect.width * ratio));
      canvas.height = Math.max(1, Math.floor(rect.height * ratio));
      const ctx = canvas.getContext("2d");
      ctx.setTransform(ratio, 0, 0, ratio, 0, 0);
      return { ctx, width: rect.width, height: rect.height };
    };

    const drawStatusChart = (counts) => {
      const total = counts.total || 0;
      const { ctx, width, height } = setupCanvas(statusCanvas);
      const cx = width / 2;
      const cy = height / 2;
      const radius = Math.min(width, height) / 2 - 16;
      const upColor = getVar("--up", "#2ad3a5");
      const downColor = getVar("--down", "#ff647c");
      const unknownColor = getVar("--unknown", "#f7c85b");
      const textColor = getVar("--text", "#f6f3ff");
      const mutedColor = getVar("--muted", "#9da8ff");
      const gridColor = getVar("--grid", "rgba(255, 255, 255, 0.08)");

      ctx.clearRect(0, 0, width, height);
      ctx.lineWidth = 18;
      ctx.strokeStyle = gridColor;
      ctx.beginPath();
      ctx.arc(cx, cy, radius, 0, Math.PI * 2);
      ctx.stroke();

      if (total > 0) {
        const segments = [
          { value: counts.up, color: upColor },
          { value: counts.down, color: downColor },
          { value: counts.unknown, color: unknownColor }
        ];
        let start = -Math.PI / 2;
        segments.forEach((segment) => {
          if (!segment.value) return;
          const angle = (segment.value / total) * Math.PI * 2;
          ctx.strokeStyle = segment.color;
          ctx.beginPath();
          ctx.arc(cx, cy, radius, start, start + angle);
          ctx.stroke();
          start += angle;
        });
      }

      ctx.fillStyle = textColor;
      ctx.font = "600 20px Space Grotesk, sans-serif";
      ctx.textAlign = "center";
      ctx.textBaseline = "middle";
      ctx.fillText(total + " monitors", cx, cy - 6);
      ctx.fillStyle = mutedColor;
      ctx.font = "12px IBM Plex Sans, sans-serif";
      ctx.fillText("total", cx, cy + 14);
    };

    const drawLatencyChart = (events) => {
      const { ctx, width, height } = setupCanvas(latencyCanvas);
      const gridColor = getVar("--grid", "rgba(255, 255, 255, 0.08)");
      const accentColor = getVar("--accent", "#5be2ff");
      const downColor = getVar("--down", "#ff647c");
      const mutedColor = getVar("--muted", "#9da8ff");
      ctx.clearRect(0, 0, width, height);

      if (!events || events.length === 0) {
        ctx.fillStyle = mutedColor;
        ctx.font = "12px IBM Plex Sans, sans-serif";
        ctx.textAlign = "center";
        ctx.textBaseline = "middle";
        ctx.fillText("No events yet", width / 2, height / 2);
        return;
      }

      const padding = { top: 18, right: 16, bottom: 26, left: 32 };
      const plotWidth = width - padding.left - padding.right;
      const plotHeight = height - padding.top - padding.bottom;
      const values = events.map((event) => event.latency_ms || 0);
      const maxValue = Math.max(50, ...values);

      ctx.strokeStyle = gridColor;
      ctx.lineWidth = 1;
      for (let i = 0; i <= 3; i++) {
        const y = padding.top + (plotHeight * i) / 3;
        ctx.beginPath();
        ctx.moveTo(padding.left, y);
        ctx.lineTo(width - padding.right, y);
        ctx.stroke();
      }

      ctx.beginPath();
      values.forEach((value, index) => {
        const x = padding.left + (plotWidth * index) / Math.max(1, values.length - 1);
        const y = padding.top + plotHeight - (value / maxValue) * plotHeight;
        if (index === 0) {
          ctx.moveTo(x, y);
        } else {
          ctx.lineTo(x, y);
        }
      });
      ctx.strokeStyle = accentColor;
      ctx.lineWidth = 2;
      ctx.stroke();

      events.forEach((event, index) => {
        if (event.success) return;
        const x = padding.left + (plotWidth * index) / Math.max(1, values.length - 1);
        const y = padding.top + plotHeight - ((event.latency_ms || 0) / maxValue) * plotHeight;
        ctx.fillStyle = downColor;
        ctx.beginPath();
        ctx.arc(x, y, 4, 0, Math.PI * 2);
        ctx.fill();
      });
    };

    const updateCounts = (monitors) => {
      const counts = { up: 0, down: 0, unknown: 0, total: monitors.length };
      monitors.forEach((monitor) => {
        if (monitor.status === "up") {
          counts.up++;
        } else if (monitor.status === "down") {
          counts.down++;
        } else {
          counts.unknown++;
        }
      });
      state.counts = counts;
      countTotalEl.textContent = counts.total;
      countUpEl.textContent = counts.up;
      countDownEl.textContent = counts.down;
      countUnknownEl.textContent = counts.unknown;
      legendUpEl.textContent = counts.up;
      legendDownEl.textContent = counts.down;
      legendUnknownEl.textContent = counts.unknown;
      drawStatusChart(counts);
    };

    const updateFocusOptions = (monitors) => {
      const current = focusSelect.value;
      focusSelect.innerHTML = "";
      monitors.forEach((monitor) => {
        const option = document.createElement("option");
        option.value = monitor.id;
        option.textContent = monitor.name;
        focusSelect.appendChild(option);
      });

      if (monitors.length === 0) {
        focusSelect.innerHTML = "";
        focusSelect.disabled = true;
        state.events = [];
        latencyHintEl.textContent = "No monitors available";
        latencyLastEl.textContent = "Last latency: n/a";
        latencyUptimeEl.textContent = "Uptime: n/a";
        drawLatencyChart([]);
        return;
      }

      focusSelect.disabled = false;
      if (current && monitors.some((monitor) => String(monitor.id) === current)) {
        focusSelect.value = current;
      } else {
        focusSelect.value = monitors[0].id;
      }
    };

    const loadEventsForFocus = async () => {
      const id = focusSelect.value;
      if (!id) {
        state.events = [];
        drawLatencyChart([]);
        return;
      }
      latencyHintEl.textContent = "Loading events...";
      const response = await fetch("/api/monitors/" + id + "/events?limit=40");
      if (!response.ok) {
        latencyHintEl.textContent = "Unable to load events";
        state.events = [];
        drawLatencyChart([]);
        return;
      }
      const events = await response.json();
      state.events = events.slice().reverse();
      drawLatencyChart(state.events);

      const total = events.length;
      const upCount = events.filter((event) => event.success).length;
      const downCount = total - upCount;
      const lastLatency = total ? events[0].latency_ms : null;

      latencyHintEl.textContent = "Last " + total + " checks, down " + downCount;
      latencyLastEl.textContent = "Last latency: " + formatLatency(lastLatency);
      latencyUptimeEl.textContent = total ? "Uptime: " + Math.round((upCount / total) * 100) + "%" : "Uptime: n/a";
    };

    const renderMonitors = (monitors) => {
      gridEl.innerHTML = "";
      monitors.forEach((monitor, index) => {
        const card = document.createElement("div");
        card.className = "card";
        card.style.animationDelay = (index * 0.04) + "s";
        card.setAttribute("data-status", monitor.status || "unknown");
        card.innerHTML =
          "<h3>" + monitor.name + "</h3>" +
          "<div class=\"meta\">" + monitor.url + "</div>" +
          "<div class=\"meta\">Every " + monitor.interval_sec + "s / Timeout " + monitor.timeout_sec + "s</div>" +
          "<div class=\"meta\">Last check: " + formatTime(monitor.last_checked_at) + "</div>" +
          "<div class=\"meta\">Latency: " + formatLatency(monitor.last_latency_ms) + "</div>" +
          "<div class=\"status\">" +
            "<span class=\"dot " + monitor.status + "\"></span>" +
            statusLabel(monitor.status) +
          "</div>" +
          "<div class=\"card-actions\">" +
            "<button class=\"danger\" data-id=\"" + monitor.id + "\">Delete</button>" +
          "</div>";
        gridEl.appendChild(card);
      });

      document.querySelectorAll(".danger").forEach((button) => {
        button.addEventListener("click", async (event) => {
          const id = event.target.getAttribute("data-id");
          await deleteMonitor(id);
        });
      });
    };

    const renderIncidents = (incidents) => {
      incidentGridEl.innerHTML = "";
      if (!incidents || incidents.length === 0) {
        const empty = document.createElement("div");
        empty.className = "incident-meta";
        empty.textContent = "No incidents logged yet.";
        incidentGridEl.appendChild(empty);
        return;
      }
      incidents.forEach((incident) => {
        const card = document.createElement("div");
        card.className = "card incident-card";

        const top = document.createElement("div");
        top.className = "incident-top";

        const details = document.createElement("div");
        const title = document.createElement("div");
        title.className = "incident-title";
        title.textContent = incident.title || "Untitled incident";
        const meta = document.createElement("div");
        meta.className = "incident-meta";
        meta.textContent = incidentWindow(incident);
        details.appendChild(title);
        details.appendChild(meta);

        const badge = document.createElement("div");
        const statusClass = incident.status ? incident.status.toLowerCase() : "unknown";
        badge.className = "incident-badge " + statusClass;
        badge.textContent = incidentStatusLabel(statusClass);

        top.appendChild(details);
        top.appendChild(badge);

        const message = document.createElement("div");
        message.className = "incident-message";
        message.textContent = incident.message || "";

        const actions = document.createElement("div");
        actions.className = "incident-actions";
        if (statusClass !== "resolved") {
          const resolveBtn = document.createElement("button");
          resolveBtn.className = "btn btn-ghost btn-small";
          resolveBtn.textContent = "Resolve";
          resolveBtn.addEventListener("click", async () => {
            await updateIncidentStatus(incident.id, "resolved");
          });
          actions.appendChild(resolveBtn);
        }
        const deleteBtn = document.createElement("button");
        deleteBtn.className = "danger btn-small";
        deleteBtn.textContent = "Delete";
        deleteBtn.addEventListener("click", async () => {
          await deleteIncident(incident.id);
        });
        actions.appendChild(deleteBtn);

        card.appendChild(top);
        card.appendChild(message);
        card.appendChild(actions);
        incidentGridEl.appendChild(card);
      });
    };

    const loadIncidents = async () => {
      incidentErrorEl.textContent = "";
      const response = await fetch("/api/incidents?limit=20");
      if (!response.ok) {
        incidentErrorEl.textContent = "Unable to load incidents.";
        renderIncidents([]);
        return;
      }
      const incidents = await response.json();
      state.incidents = incidents;
      renderIncidents(incidents);
    };

    const updateIncidentStatus = async (id, status) => {
      const payload = { status: status };
      if (status === "resolved") {
        payload.resolved_at = new Date().toISOString();
      }
      const response = await fetch("/api/incidents/" + id, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });
      if (!response.ok) {
        const data = await response.json();
        incidentErrorEl.textContent = data.error || "Failed to update incident.";
        return;
      }
      loadIncidents();
    };

    const deleteIncident = async (id) => {
      const response = await fetch("/api/incidents/" + id, { method: "DELETE" });
      if (!response.ok) {
        incidentErrorEl.textContent = "Failed to delete incident.";
        return;
      }
      loadIncidents();
    };

    const loadMonitors = async () => {
      const response = await fetch("/api/monitors");
      if (!response.ok) {
        return;
      }
      const monitors = await response.json();
      state.monitors = monitors;
      updateCounts(monitors);
      updateFocusOptions(monitors);
      renderMonitors(monitors);
      await loadEventsForFocus();
      lastRefreshEl.textContent = "Last refresh: " + new Date().toLocaleTimeString();
    };

    const loadAll = async () => {
      await loadMonitors();
      await loadIncidents();
    };

    const deleteMonitor = async (id) => {
      const response = await fetch("/api/monitors/" + id, { method: "DELETE" });
      if (!response.ok) {
        return;
      }
      loadMonitors();
    };

    formEl.addEventListener("submit", async (event) => {
      event.preventDefault();
      errorEl.textContent = "";

      const payload = {
        name: formEl.name.value.trim(),
        url: formEl.url.value.trim(),
        method: formEl.method.value,
        interval_sec: Number(formEl.interval.value) || 60,
        timeout_sec: Number(formEl.timeout.value) || 10
      };

      const response = await fetch("/api/monitors", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });

      if (!response.ok) {
        const data = await response.json();
        errorEl.textContent = data.error || "Failed to add monitor.";
        return;
      }

      formEl.reset();
      formEl.method.value = "GET";
      formEl.interval.value = "60";
      formEl.timeout.value = "10";
      loadMonitors();
    });

    incidentFormEl.addEventListener("submit", async (event) => {
      event.preventDefault();
      incidentErrorEl.textContent = "";

      const payload = {
        title: incidentTitleEl.value.trim(),
        status: incidentStatusEl.value,
        message: incidentMessageEl.value.trim(),
        started_at: toISO(incidentStartedEl.value),
        resolved_at: toISO(incidentResolvedEl.value)
      };

      if (!payload.title || !payload.message) {
        incidentErrorEl.textContent = "Title and message are required.";
        return;
      }

      const response = await fetch("/api/incidents", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });

      if (!response.ok) {
        const data = await response.json();
        incidentErrorEl.textContent = data.error || "Failed to publish incident.";
        return;
      }

      incidentFormEl.reset();
      incidentStatusEl.value = "investigating";
      loadIncidents();
    });

    focusSelect.addEventListener("change", loadEventsForFocus);
    refreshBtn.addEventListener("click", loadAll);
    themeToggle.addEventListener("click", () => {
      const current = document.documentElement.getAttribute("data-theme") || "dark";
      setTheme(current === "light" ? "dark" : "light");
    });
    window.addEventListener("resize", () => {
      drawStatusChart(state.counts);
      drawLatencyChart(state.events);
    });

    initTheme();
    loadAll();
  </script>
</body>
</html>`
