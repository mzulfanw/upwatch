package app

import "net/http"

func serveStatusPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(statusPageHTML))
}

const statusPageHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Upwatch Status</title>
  <style>
    :root {
      color-scheme: dark;
      --bg: #141515;
      --bg-2: #1b1c1c;
      --surface: #202122;
      --surface-2: #27292a;
      --text: #f2f2f0;
      --muted: #b3b3ad;
      --up: #7fc7a4;
      --down: #f08a7c;
      --unknown: #f2c27d;
      --accent: #f0a07a;
      --accent-2: #e67f6a;
      --danger: #f08a7c;
      --grid: rgba(255, 255, 255, 0.06);
    }
    [data-theme="light"] {
      color-scheme: light;
      --bg: #f6f6f4;
      --bg-2: #ededeb;
      --surface: #ffffff;
      --surface-2: #f4f4f2;
      --text: #242424;
      --muted: #6f6f6a;
      --accent: #e99673;
      --accent-2: #de7c69;
      --danger: #e26f63;
      --grid: rgba(60, 60, 60, 0.08);
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
      transition: transform 0.2s ease, border-color 0.2s ease;
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
      margin-top: 1.6rem;
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

    .grid {
      display: grid;
      gap: 1.25rem;
      grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
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
    .sparkline-wrap {
      margin-top: 1.1rem;
      background: rgba(7, 10, 30, 0.4);
      border-radius: 14px;
      padding: 0.7rem;
      border: 1px solid rgba(255, 255, 255, 0.08);
    }
    [data-theme="light"] .sparkline-wrap {
      background: rgba(16, 32, 80, 0.05);
      border-color: rgba(16, 32, 80, 0.08);
    }
    .sparkline {
      width: 100%;
      height: 120px;
      display: block;
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

    @keyframes fadeUp {
      from { opacity: 0; transform: translateY(10px); }
      to { opacity: 1; transform: translateY(0); }
    }

    @media (max-width: 900px) {
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
        <div class="brand-mark">Upwatch Status</div>
        <div class="brand-sub">Public status feed for tracked services.</div>
      </div>
      <div class="actions">
        <button class="btn btn-ghost" id="themeToggle" type="button">Light mode</button>
      </div>
    </div>

    <section class="panel">
      <div class="eyebrow">Live overview</div>
      <h2 class="panel-title">Current system health</h2>
      <p class="panel-subtitle">Updated in real time from the stream.</p>
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
    </section>

    <section>
      <div class="section-head">
        <div>
          <h2 class="section-title">Active monitors</h2>
          <p class="muted">Live status for all tracked endpoints.</p>
        </div>
        <div class="chip" id="streamStatus">Stream: connecting</div>
        <div class="chip" id="updatedAt">Updated: never</div>
      </div>
      <div class="grid" id="monitorGrid"></div>
    </section>
  </div>

  <script>
    const gridEl = document.getElementById("monitorGrid");
    const updatedAtEl = document.getElementById("updatedAt");
    const streamStatusEl = document.getElementById("streamStatus");
    const themeToggle = document.getElementById("themeToggle");

    const countTotalEl = document.getElementById("countTotal");
    const countUpEl = document.getElementById("countUp");
    const countDownEl = document.getElementById("countDown");
    const countUnknownEl = document.getElementById("countUnknown");

    const state = {
      counts: { up: 0, down: 0, unknown: 0, total: 0 },
      monitors: [],
      history: {}
    };

    const maxPoints = 40;

    const getVar = (name, fallback) => {
      const value = getComputedStyle(document.documentElement).getPropertyValue(name).trim();
      return value || fallback;
    };

    const setTheme = (theme) => {
      document.documentElement.setAttribute("data-theme", theme);
      themeToggle.textContent = theme === "light" ? "Dark mode" : "Light mode";
      localStorage.setItem("upwatch-theme", theme);
      requestAnimationFrame(drawAllSparklines);
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

    const statusLabel = (status) => {
      if (status === "up") return "Up";
      if (status === "down") return "Down";
      return "Unknown";
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
    };

    const updateHistory = (monitors) => {
      const next = {};
      monitors.forEach((monitor) => {
        const id = String(monitor.id);
        const existing = state.history[id] || [];
        const value = monitor.last_latency_ms === null || monitor.last_latency_ms === undefined
          ? null
          : Number(monitor.last_latency_ms);
        const nextSeries = existing.concat([{ v: value }]).slice(-maxPoints);
        next[id] = nextSeries;
      });
      state.history = next;
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
          "<div class=\"sparkline-wrap\">" +
            "<canvas class=\"sparkline\" data-id=\"" + monitor.id + "\" data-status=\"" + monitor.status + "\"></canvas>" +
          "</div>";
        gridEl.appendChild(card);
      });
      requestAnimationFrame(drawAllSparklines);
    };

    const drawSparkline = (canvas, series, status) => {
      const { ctx, width, height } = setupCanvas(canvas);
      const gridColor = getVar("--grid", "rgba(255, 255, 255, 0.08)");
      const mutedColor = getVar("--muted", "#9da8ff");
      const upColor = getVar("--up", "#2ad3a5");
      const downColor = getVar("--down", "#ff647c");
      const unknownColor = getVar("--unknown", "#f7c85b");
      const isDown = status === "down";
      const isUnknown = status === "unknown";
      const lineColor = isDown ? downColor : (isUnknown ? unknownColor : upColor);

      ctx.clearRect(0, 0, width, height);

      if (!series || series.length === 0) {
        ctx.fillStyle = mutedColor;
        ctx.font = "12px IBM Plex Sans, sans-serif";
        ctx.textAlign = "center";
        ctx.textBaseline = "middle";
        ctx.fillText("No data", width / 2, height / 2);
        return;
      }

      const values = series.map((point) => point.v).filter((value) => value !== null && value !== undefined);
      if (values.length === 0) {
        ctx.fillStyle = mutedColor;
        ctx.font = "12px IBM Plex Sans, sans-serif";
        ctx.textAlign = "center";
        ctx.textBaseline = "middle";
        ctx.fillText("Waiting for checks", width / 2, height / 2);
        return;
      }

      const minValue = Math.min(...values);
      const maxValue = Math.max(...values);
      const range = Math.max(1, maxValue - minValue);
      const padding = { top: 10, right: 12, bottom: 14, left: 12 };
      const plotWidth = width - padding.left - padding.right;
      const plotHeight = height - padding.top - padding.bottom;

      ctx.setLineDash([]);
      ctx.strokeStyle = gridColor;
      ctx.lineWidth = 1;
      for (let i = 0; i <= 2; i++) {
        const y = padding.top + (plotHeight * i) / 2;
        ctx.beginPath();
        ctx.moveTo(padding.left, y);
        ctx.lineTo(width - padding.right, y);
        ctx.stroke();
      }

      ctx.beginPath();
      let started = false;
      series.forEach((point, index) => {
        if (point.v === null || point.v === undefined) {
          started = false;
          return;
        }
        const x = padding.left + (plotWidth * index) / Math.max(1, series.length - 1);
        const y = padding.top + plotHeight - ((point.v - minValue) / range) * plotHeight;
        if (!started) {
          ctx.moveTo(x, y);
          started = true;
        } else {
          ctx.lineTo(x, y);
        }
      });
      ctx.strokeStyle = lineColor;
      ctx.lineWidth = 2.4;
      if (isDown) {
        ctx.setLineDash([6, 4]);
      } else {
        ctx.setLineDash([]);
      }
      ctx.stroke();

      ctx.setLineDash([]);
      ctx.lineTo(padding.left + plotWidth, padding.top + plotHeight);
      ctx.lineTo(padding.left, padding.top + plotHeight);
      ctx.closePath();
      ctx.save();
      ctx.globalAlpha = isDown ? 0.28 : 0.22;
      ctx.fillStyle = lineColor;
      ctx.fill();
      ctx.restore();
    };

    const drawAllSparklines = () => {
      const nodes = document.querySelectorAll(".sparkline");
      nodes.forEach((canvas) => {
        const id = canvas.getAttribute("data-id");
        const status = canvas.getAttribute("data-status") || "unknown";
        drawSparkline(canvas, state.history[id] || [], status);
      });
    };

    const applyStatus = (data) => {
      state.monitors = data.monitors || [];
      updateCounts(state.monitors);
      updateHistory(state.monitors);
      renderMonitors(state.monitors);
      updatedAtEl.textContent = "Updated: " + formatTime(data.updated_at);
    };

    const startStream = () => {
      if (!("EventSource" in window)) {
        streamStatusEl.textContent = "Stream: unsupported";
        return;
      }
      streamStatusEl.textContent = "Stream: connecting";
      const source = new EventSource("/api/status/stream");
      source.onopen = () => {
        streamStatusEl.textContent = "Stream: live";
      };
      source.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data);
          applyStatus(data);
        } catch (err) {
        }
      };
      source.onerror = () => {
        streamStatusEl.textContent = "Stream: offline";
      };
    };

    themeToggle.addEventListener("click", () => {
      const current = document.documentElement.getAttribute("data-theme") || "dark";
      setTheme(current === "light" ? "dark" : "light");
    });

    window.addEventListener("resize", () => {
      drawAllSparklines();
    });

    initTheme();
    startStream();
  </script>
</body>
</html>`
