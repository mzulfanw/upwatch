package app

import "net/http"

func serveLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(loginPageHTML))
}

const loginPageHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Upwatch Login</title>
  <meta name="description" content="Secure admin login for Upwatch." />
  <meta name="robots" content="noindex,nofollow" />
  <meta property="og:title" content="Upwatch Login" />
  <meta property="og:description" content="Secure admin login for Upwatch." />
  <meta property="og:type" content="website" />
  <meta property="og:site_name" content="Upwatch" />
  <meta name="twitter:title" content="Upwatch Login" />
  <meta name="twitter:description" content="Secure admin login for Upwatch." />
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
      min-height: 100vh;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      padding: 2.5rem 1.5rem 4rem;
      position: relative;
      z-index: 1;
      gap: 2rem;
    }
    .topbar {
      width: 100%;
      max-width: 480px;
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 1rem;
    }
    .brand {
      display: flex;
      align-items: center;
      gap: 0.85rem;
    }
    .brand-logo {
      width: 56px;
      height: 56px;
      border-radius: 14px;
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
      font-size: 0.95rem;
      font-weight: 600;
      letter-spacing: 0.04em;
      text-transform: uppercase;
    }
    .brand-sub {
      color: var(--muted);
      font-size: 0.75rem;
      letter-spacing: 0.02em;
    }
    .btn {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      padding: 0.5rem 1rem;
      border-radius: 999px;
      border: 1px solid rgba(255, 255, 255, 0.14);
      background: rgba(255, 255, 255, 0.08);
      color: var(--text);
      font-weight: 600;
      cursor: pointer;
      text-decoration: none;
      transition: transform 0.2s ease;
    }
    .btn-icon {
      width: 42px;
      height: 42px;
      padding: 0;
      position: relative;
    }
    [data-theme="light"] .btn {
      border-color: rgba(16, 32, 80, 0.12);
      background: rgba(16, 32, 80, 0.06);
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
    .card {
      width: 100%;
      max-width: 480px;
      background: var(--surface);
      border: 1px solid var(--card-border);
      border-radius: 20px;
      padding: 2rem;
      box-shadow: var(--shadow);
    }
    [data-theme="light"] .card {
      background: var(--surface);
    }
    h1 {
      margin: 0 0 0.5rem;
      font-size: 1.9rem;
      letter-spacing: -0.01em;
    }
    p {
      margin: 0 0 1.75rem;
      color: var(--muted);
      line-height: 1.6;
      font-size: 0.9rem;
    }
    label {
      display: block;
      font-size: 0.72rem;
      margin-bottom: 0.45rem;
      letter-spacing: 0.08em;
      text-transform: uppercase;
      color: var(--muted);
    }
    input {
      width: 100%;
      padding: 0.7rem 0.8rem;
      border-radius: 12px;
      border: 1px solid rgba(255, 255, 255, 0.15);
      background: var(--bg);
      color: var(--text);
      margin-bottom: 1.1rem;
      font-family: inherit;
      font-size: 0.95rem;
    }
    [data-theme="light"] input {
      background: rgba(16, 32, 80, 0.06);
      border-color: rgba(16, 32, 80, 0.12);
    }
    .primary {
      width: 100%;
      padding: 0.75rem;
      border-radius: 999px;
      border: none;
      background: linear-gradient(120deg, var(--accent), var(--accent-2));
      color: white;
      font-weight: 700;
      cursor: pointer;
      font-family: inherit;
      letter-spacing: 0.03em;
    }
    .error {
      margin-top: 1rem;
      color: var(--danger);
      font-size: 0.9rem;
      min-height: 1rem;
    }
    .footer {
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
          <div class="brand-mark">Upwatch Admin</div>
          <div class="brand-sub">Secure login for monitor control.</div>
        </div>
      </div>
      <button class="btn btn-icon" id="themeToggle" type="button" aria-label="Switch to light mode" title="Switch to light mode">
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

    <div class="card">
      <h1>Sign in</h1>
      <p>Enter your admin credentials to manage monitors.</p>
      <form method="post" action="/login">
        <label for="username">Username</label>
        <input id="username" name="username" type="text" autocomplete="username" required />
        <label for="password">Password</label>
        <input id="password" name="password" type="password" autocomplete="current-password" required />
        <button class="primary" type="submit">Sign in</button>
        <div class="error" id="error"></div>
      </form>
    </div>
    <footer class="footer">Powered by <a class="footer-link" href="https://github.com/mzulfanw/upwatch" target="_blank" rel="noopener">Upwatch</a></footer>
  </div>
  <script>
    const params = new URLSearchParams(window.location.search);
    const errorEl = document.getElementById("error");
    const themeToggle = document.getElementById("themeToggle");

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

    if (params.get("error")) {
      errorEl.textContent = "Invalid credentials.";
    }

    themeToggle.addEventListener("click", () => {
      const current = document.documentElement.getAttribute("data-theme") || "dark";
      setTheme(current === "light" ? "dark" : "light");
    });

    initTheme();
  </script>
</body>
</html>`
