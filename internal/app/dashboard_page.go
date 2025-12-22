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
  <style>
    :root {
      color-scheme: light dark;
      --bg: #1c1916;
      --bg-2: #221f1b;
      --surface: #27231f;
      --surface-2: #181513;
      --card-border: rgba(255, 255, 255, 0.08);
      --text: #f2ede7;
      --muted: #b8aca1;
      --up: #7fb58a;
      --down: #d77a6a;
      --unknown: #d9b56d;
      --accent: #c9a27d;
      --accent-2: #b9856b;
      --danger: #d77a6a;
      --shadow: 0 24px 60px rgba(0, 0, 0, 0.55);
      --grid: rgba(255, 255, 255, 0.08);
    }
    [data-theme="light"] {
      color-scheme: light;
      --bg: #f4efe9;
      --bg-2: #e9e1d8;
      --surface: #fbf7f2;
      --surface-2: #f1e8de;
      --card-border: rgba(40, 33, 28, 0.12);
      --text: #2a241f;
      --muted: #6f6256;
      --accent: #b98b63;
      --accent-2: #c07a66;
      --danger: #c06355;
      --shadow: 0 24px 60px rgba(40, 33, 28, 0.12);
      --grid: rgba(40, 33, 28, 0.08);
    }
    * { box-sizing: border-box; }
    body {
      margin: 0;
      min-height: 100vh;
      font-family: "Space Grotesk", "IBM Plex Sans", "Segoe UI", sans-serif;
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
      display: grid;
      gap: 0.35rem;
    }
    .brand-mark {
      font-size: clamp(2rem, 4vw, 3.2rem);
      font-weight: 700;
      letter-spacing: -0.02em;
    }
    .brand-sub {
      color: var(--muted);
      font-size: 0.95rem;
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
      cursor: pointer;
      text-decoration: none;
      transition: transform 0.2s ease, box-shadow 0.2s ease, border-color 0.2s ease;
    }
    .btn-primary {
      background: linear-gradient(120deg, var(--accent), var(--accent-2));
      color: #0b1029;
      box-shadow: 0 14px 30px rgba(91, 226, 255, 0.3);
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
      box-shadow: 0 10px 24px rgba(0, 0, 0, 0.2);
    }
    .panel {
      background: var(--surface);
      border: 1px solid var(--card-border);
      border-radius: 20px;
      padding: 2rem;
      box-shadow: var(--shadow);
    }
    [data-theme="light"] .panel {
      background: var(--surface);
    }
    .panel + .panel {
      margin-top: 2.5rem;
    }
    .panel-title {
      margin: 0 0 0.4rem;
      font-size: 1.4rem;
    }
    .panel-subtitle {
      margin: 0;
      color: var(--muted);
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
      background: rgba(7, 10, 30, 0.55);
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
      background: rgba(7, 10, 30, 0.55);
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
    input, select {
      width: 100%;
      padding: 0.65rem 0.75rem;
      border-radius: 10px;
      border: 1px solid rgba(255, 255, 255, 0.15);
      background: rgba(13, 16, 45, 0.7);
      color: var(--text);
    }
    [data-theme="light"] input,
    [data-theme="light"] select {
      background: rgba(16, 32, 80, 0.06);
      border-color: rgba(16, 32, 80, 0.12);
    }
    .grid {
      display: grid;
      gap: 1.25rem;
      grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
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
      font-size: 0.8rem;
      position: relative;
      z-index: 1;
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
    }
    .section-title {
      margin: 0;
    }
    .muted {
      color: var(--muted);
      margin: 0.25rem 0 0;
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
        <div class="brand-mark">Upwatch Dashboard</div>
        <div class="brand-sub">Self hosted uptime command center.</div>
      </div>
      <div class="actions">
        <button class="btn btn-ghost" id="themeToggle" type="button">Light mode</button>
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

    const state = {
      monitors: [],
      counts: { up: 0, down: 0, unknown: 0, total: 0 },
      events: []
    };

    const getVar = (name, fallback) => {
      const value = getComputedStyle(document.documentElement).getPropertyValue(name).trim();
      return value || fallback;
    };

    const setTheme = (theme) => {
      document.documentElement.setAttribute("data-theme", theme);
      themeToggle.textContent = theme === "light" ? "Dark mode" : "Light mode";
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

    focusSelect.addEventListener("change", loadEventsForFocus);
    refreshBtn.addEventListener("click", loadMonitors);
    themeToggle.addEventListener("click", () => {
      const current = document.documentElement.getAttribute("data-theme") || "dark";
      setTheme(current === "light" ? "dark" : "light");
    });
    window.addEventListener("resize", () => {
      drawStatusChart(state.counts);
      drawLatencyChart(state.events);
    });

    initTheme();
    loadMonitors();
  </script>
</body>
</html>`
