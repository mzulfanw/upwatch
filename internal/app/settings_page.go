package app

import "net/http"

func (a *App) handleSettingsPage(w http.ResponseWriter, r *http.Request) {
	serveSettingsPage(w, r)
}

func serveSettingsPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(settingsPageHTML))
}

const settingsPageHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Upwatch Settings</title>
  <meta name="description" content="Branding and copy settings for the Upwatch status page." />
  <meta name="robots" content="noindex,nofollow" />
  <meta property="og:title" content="Upwatch Settings" />
  <meta property="og:description" content="Branding and copy settings for the Upwatch status page." />
  <meta property="og:type" content="website" />
  <meta property="og:site_name" content="Upwatch" />
  <meta name="twitter:title" content="Upwatch Settings" />
  <meta name="twitter:description" content="Branding and copy settings for the Upwatch status page." />
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
      --accent: #c9a27d;
      --accent-2: #b9856b;
      --danger: #d77a6a;
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
    .shell {
      max-width: 980px;
      margin: 0 auto;
      padding: 3rem 2rem 4rem;
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
      font-size: clamp(1.8rem, 3vw, 2.6rem);
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
    .btn-primary {
      background: linear-gradient(120deg, var(--accent), var(--accent-2));
      color: #0b1029;
    }
    .btn:hover {
      transform: translateY(-1px);
      box-shadow: 0 10px 24px rgba(0, 0, 0, 0.2);
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
    .theme-icon.sun { opacity: 1; }
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
    form {
      display: grid;
      gap: 1rem;
      grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
    }
    label {
      font-size: 0.72rem;
      margin-bottom: 0.45rem;
      letter-spacing: 0.08em;
      text-transform: uppercase;
      color: var(--muted);
    }
    input, textarea {
      width: 100%;
      padding: 0.7rem 0.8rem;
      border-radius: 12px;
      border: 1px solid rgba(255, 255, 255, 0.15);
      background: var(--bg);
      color: var(--text);
      font-family: inherit;
      font-size: 0.95rem;
    }
    textarea {
      min-height: 120px;
      resize: vertical;
    }
    [data-theme="light"] input,
    [data-theme="light"] textarea {
      background: rgba(16, 32, 80, 0.06);
      border-color: rgba(16, 32, 80, 0.12);
    }
    .field {
      display: flex;
      flex-direction: column;
    }
    .field-full {
      grid-column: 1 / -1;
    }
    .form-actions {
      display: flex;
      align-items: center;
      gap: 1rem;
      flex-wrap: wrap;
      grid-column: 1 / -1;
    }
    .hint {
      color: var(--muted);
      font-size: 0.85rem;
    }
    .status {
      margin-top: 1rem;
      font-size: 0.9rem;
      min-height: 1rem;
    }
    .status.error { color: var(--danger); }
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
    @media (max-width: 900px) {
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
          <div class="brand-mark">Brand settings</div>
          <div class="brand-sub">Update public labels shown on the status page.</div>
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
        <a class="btn btn-ghost" href="/dashboard">Back to dashboard</a>
      </div>
    </div>

    <section class="panel">
      <h1 class="panel-title">Status page copy</h1>
      <p class="panel-subtitle">Leave a field empty if you want to hide it on the public page.</p>
      <form id="settingsForm">
        <div class="field">
          <label for="brandName">Brand name</label>
          <input id="brandName" name="brandName" type="text" />
        </div>
        <div class="field">
          <label for="brandTagline">Tagline</label>
          <input id="brandTagline" name="brandTagline" type="text" />
        </div>
        <div class="field">
          <label for="statusTitle">Status headline</label>
          <input id="statusTitle" name="statusTitle" type="text" />
        </div>
        <div class="field-full">
          <label for="statusSubtitle">Status subtitle</label>
          <textarea id="statusSubtitle" name="statusSubtitle"></textarea>
        </div>
        <div class="form-actions">
          <button class="btn btn-primary" type="submit">Save settings</button>
          <div class="hint">Changes appear on the public status page within about a minute.</div>
        </div>
      </form>
      <div class="status" id="saveStatus"></div>
    </section>
    <footer class="footer">Powered by <a class="footer-link" href="https://github.com/mzulfanw/upwatch" target="_blank" rel="noopener">Upwatch</a></footer>
  </div>

  <script>
    const themeToggle = document.getElementById("themeToggle");
    const formEl = document.getElementById("settingsForm");
    const statusEl = document.getElementById("saveStatus");

    const brandNameEl = document.getElementById("brandName");
    const brandTaglineEl = document.getElementById("brandTagline");
    const statusTitleEl = document.getElementById("statusTitle");
    const statusSubtitleEl = document.getElementById("statusSubtitle");

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

    const setStatus = (message, isError) => {
      statusEl.textContent = message;
      statusEl.className = isError ? "status error" : "status";
    };

    const loadSettings = async () => {
      setStatus("", false);
      const response = await fetch("/api/settings");
      if (!response.ok) {
        setStatus("Unable to load settings.", true);
        return;
      }
      const settings = await response.json();
      brandNameEl.value = settings.brand_name || "";
      brandTaglineEl.value = settings.brand_tagline || "";
      statusTitleEl.value = settings.status_title || "";
      statusSubtitleEl.value = settings.status_subtitle || "";
    };

    formEl.addEventListener("submit", async (event) => {
      event.preventDefault();
      setStatus("", false);

      const payload = {
        brand_name: brandNameEl.value.trim(),
        brand_tagline: brandTaglineEl.value.trim(),
        status_title: statusTitleEl.value.trim(),
        status_subtitle: statusSubtitleEl.value.trim()
      };

      const response = await fetch("/api/settings", {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload)
      });

      if (!response.ok) {
        const data = await response.json();
        setStatus(data.error || "Failed to save settings.", true);
        return;
      }

      const data = await response.json();
      setStatus("Saved at " + new Date(data.updated_at).toLocaleTimeString(), false);
    });

    themeToggle.addEventListener("click", () => {
      const current = document.documentElement.getAttribute("data-theme") || "dark";
      setTheme(current === "light" ? "dark" : "light");
    });

    initTheme();
    loadSettings();
  </script>
</body>
</html>`
