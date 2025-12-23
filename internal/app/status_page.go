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
  <meta name="description" content="Live system status and incident history." />
  <link rel="canonical" href="" id="canonicalLink" />
  <meta property="og:title" content="Upwatch Status" />
  <meta property="og:description" content="Live system status and incident history." />
  <meta property="og:type" content="website" />
  <meta property="og:site_name" content="Upwatch" />
  <meta property="og:url" content="" id="ogUrl" />
  <meta name="twitter:title" content="Upwatch Status" />
  <meta name="twitter:description" content="Live system status and incident history." />
  <link rel="icon" type="image/png" href="/assets/upwatch.png" />
  <meta property="og:image" content="/assets/upwatch.png" />
  <meta name="twitter:card" content="summary_large_image" />
  <meta name="twitter:image" content="/assets/upwatch.png" />
  <style>
    @import url("https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;600;700&display=swap");
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
      box-shadow: var(--shadow);
    }
    [data-theme="light"] .panel {
      background: var(--surface);
    }
    .panel + .panel {
      margin-top: 2.5rem;
    }
    section + section {
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
      margin-top: 0.5rem;
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
      margin-top: 0;
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
    .incident-grid {
      margin-top: 0;
    }
    .card.incident-card {
      border: 1px solid rgba(240, 138, 124, 0.55);
    }
    [data-theme="light"] .card.incident-card {
      border: 1px solid rgba(226, 111, 99, 0.55);
    }
    .incident-card::after {
      background: radial-gradient(circle at top right, rgba(240, 138, 124, 0.18), transparent 55%);
    }
    .incident-top {
      display: flex;
      align-items: flex-start;
      justify-content: space-between;
      gap: 1rem;
    }
    .incident-title {
      font-weight: 600;
      font-size: 1.05rem;
      margin: 0;
    }
    .incident-meta {
      color: var(--muted);
      font-size: 0.85rem;
      margin-top: 0.35rem;
    }
    .incident-message {
      margin: 0.75rem 0 0;
      line-height: 1.5;
    }
    .incident-badge {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      padding: 0.35rem 0.7rem;
      border-radius: 999px;
      font-size: 0.7rem;
      text-transform: uppercase;
      letter-spacing: 0.12em;
      border: 1px solid transparent;
    }
    .incident-badge.investigating {
      color: var(--down);
      border-color: rgba(240, 138, 124, 0.45);
      background: rgba(240, 138, 124, 0.12);
    }
    .incident-badge.identified {
      color: var(--down);
      border-color: rgba(240, 138, 124, 0.45);
      background: rgba(240, 138, 124, 0.12);
    }
    .incident-badge.monitoring {
      color: var(--unknown);
      border-color: rgba(242, 194, 125, 0.45);
      background: rgba(242, 194, 125, 0.12);
    }
    .incident-badge.resolved {
      color: var(--up);
      border-color: rgba(127, 199, 164, 0.45);
      background: rgba(127, 199, 164, 0.12);
    }
    .incident-badge.maintenance,
    .incident-badge.scheduled {
      color: var(--unknown);
      border-color: rgba(242, 194, 125, 0.45);
      background: rgba(242, 194, 125, 0.12);
    }
    .incident-empty {
      color: var(--muted);
      font-size: 0.9rem;
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
        <div class="brand-logo">
          <img src="/assets/upwatch.png" alt="Upwatch logo" />
        </div>
        <div class="brand-text">
          <div class="brand-mark" id="brandName">Upwatch Status</div>
          <div class="brand-sub" id="brandTagline">Public status feed for tracked services.</div>
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
      </div>
    </div>

    <section class="panel">
      <div class="eyebrow">Live overview</div>
      <h2 class="panel-title" id="statusTitle">Current system health</h2>
      <p class="panel-subtitle" id="statusSubtitle">Updated in real time from the stream.</p>
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

    <section>
      <div class="section-head">
        <div>
          <h2 class="section-title">Past incidents</h2>
          <p class="muted">Maintenance windows and resolved disruptions.</p>
        </div>
      </div>
      <div class="grid incident-grid" id="incidentList"></div>
    </section>
    <footer class="footer">Powered by <a class="footer-link" id="footerBrandLink" href="https://github.com/mzulfanw/upwatch" target="_blank" rel="noopener">Upwatch</a></footer>
  </div>

  <script>
    const gridEl = document.getElementById("monitorGrid");
    const updatedAtEl = document.getElementById("updatedAt");
    const streamStatusEl = document.getElementById("streamStatus");
    const themeToggle = document.getElementById("themeToggle");
    const incidentListEl = document.getElementById("incidentList");
    const brandNameEl = document.getElementById("brandName");
    const brandTaglineEl = document.getElementById("brandTagline");
    const statusTitleEl = document.getElementById("statusTitle");
    const statusSubtitleEl = document.getElementById("statusSubtitle");
    const footerBrandEl = document.getElementById("footerBrandLink");
    const canonicalLinkEl = document.getElementById("canonicalLink");
    const ogUrlEl = document.getElementById("ogUrl");
    const metaDescriptionEl = document.querySelector('meta[name="description"]');
    const ogTitleEl = document.querySelector('meta[property="og:title"]');
    const ogDescriptionEl = document.querySelector('meta[property="og:description"]');
    const twitterTitleEl = document.querySelector('meta[name="twitter:title"]');
    const twitterDescriptionEl = document.querySelector('meta[name="twitter:description"]');

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

    const setText = (el, value) => {
      if (!el || value === undefined || value === null) return;
      const trimmed = String(value).trim();
      if (trimmed === "") {
        el.textContent = "";
        el.style.display = "none";
        return;
      }
      el.textContent = trimmed;
      el.style.display = "";
    };

    const setMeta = (el, value) => {
      if (!el || !value) return;
      el.setAttribute("content", value);
    };

    const updateMeta = (title, description) => {
      if (title) {
        document.title = title;
        setMeta(ogTitleEl, title);
        setMeta(twitterTitleEl, title);
      }
      if (description) {
        setMeta(metaDescriptionEl, description);
        setMeta(ogDescriptionEl, description);
        setMeta(twitterDescriptionEl, description);
      }
    };

    const applySettings = (settings) => {
      if (!settings) return;
      setText(brandNameEl, settings.brand_name);
      setText(brandTaglineEl, settings.brand_tagline);
      setText(statusTitleEl, settings.status_title);
      setText(statusSubtitleEl, settings.status_subtitle);
      const name = settings.brand_name && String(settings.brand_name).trim() !== ""
        ? settings.brand_name.trim()
        : "Upwatch Status";
      const description = settings.brand_tagline && String(settings.brand_tagline).trim() !== ""
        ? settings.brand_tagline.trim()
        : "Live system status and incident history.";
      updateMeta(name, description);
      if (footerBrandEl) {
        footerBrandEl.textContent = name;
      }
    };

    const loadSettings = async () => {
      try {
        const response = await fetch("/api/settings");
        if (!response.ok) {
          return;
        }
        const settings = await response.json();
        applySettings(settings);
      } catch (err) {
      }
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

    const incidentWindow = (incident) => {
      const started = formatTime(incident.started_at);
      if (incident.resolved_at) {
        return "Started " + started + " · Resolved " + formatTime(incident.resolved_at);
      }
      return "Started " + started + " · Ongoing";
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

    const renderIncidents = (incidents) => {
      incidentListEl.innerHTML = "";
      if (!incidents || incidents.length === 0) {
        const empty = document.createElement("div");
        empty.className = "incident-empty";
        empty.textContent = "No incidents reported yet.";
        incidentListEl.appendChild(empty);
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

        card.appendChild(top);
        card.appendChild(message);
        incidentListEl.appendChild(card);
      });
    };

    const loadIncidents = async () => {
      try {
        const response = await fetch("/api/incidents?limit=20");
        if (!response.ok) {
          throw new Error("bad response");
        }
        const incidents = await response.json();
        renderIncidents(incidents);
      } catch (err) {
        incidentListEl.innerHTML = "";
        const empty = document.createElement("div");
        empty.className = "incident-empty";
        empty.textContent = "Unable to load incidents.";
        incidentListEl.appendChild(empty);
      }
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
      if (data.history) {
        state.history = data.history;
      } else {
        updateHistory(state.monitors);
      }
      renderMonitors(state.monitors);
      updatedAtEl.textContent = "Updated: " + formatTime(data.updated_at);
    };

    let streamSource = null;
    let streamTimer = null;
    const streamRetryMs = 5000;

    const startStream = () => {
      if (!("EventSource" in window)) {
        streamStatusEl.textContent = "Stream: unsupported";
        return;
      }
      if (streamTimer) {
        clearTimeout(streamTimer);
        streamTimer = null;
      }
      if (streamSource) {
        streamSource.close();
        streamSource = null;
      }
      streamStatusEl.textContent = "Stream: connecting";
      streamSource = new EventSource("/api/status/stream");
      streamSource.onopen = () => {
        streamStatusEl.textContent = "Stream: live";
      };
      streamSource.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data);
          applyStatus(data);
        } catch (err) {
        }
      };
      streamSource.onerror = () => {
        streamStatusEl.textContent = "Stream: offline";
        if (streamSource) {
          streamSource.close();
          streamSource = null;
        }
        if (!streamTimer) {
          streamTimer = window.setTimeout(startStream, streamRetryMs);
        }
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
    const pageUrl = window.location.origin + window.location.pathname;
    if (canonicalLinkEl) {
      canonicalLinkEl.setAttribute("href", pageUrl);
    }
    if (ogUrlEl) {
      ogUrlEl.setAttribute("content", pageUrl);
    }
    loadSettings();
    window.setInterval(loadSettings, 60000);
    startStream();
    loadIncidents();
    window.setInterval(loadIncidents, 60000);
  </script>
</body>
</html>`
