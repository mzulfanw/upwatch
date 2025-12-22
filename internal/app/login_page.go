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
  <style>
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
      font-size: 1.1rem;
      font-weight: 600;
    }
    .brand-sub {
      color: var(--muted);
      font-size: 0.8rem;
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
    [data-theme="light"] .btn {
      border-color: rgba(16, 32, 80, 0.12);
      background: rgba(16, 32, 80, 0.06);
    }
    .btn:hover {
      transform: translateY(-1px);
      box-shadow: 0 10px 24px rgba(0, 0, 0, 0.2);
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
      font-size: 2rem;
    }
    p {
      margin: 0 0 1.5rem;
      color: var(--muted);
    }
    label {
      display: block;
      font-size: 0.9rem;
      margin-bottom: 0.4rem;
    }
    input {
      width: 100%;
      padding: 0.7rem 0.8rem;
      border-radius: 12px;
      border: 1px solid rgba(255, 255, 255, 0.15);
      background: var(--bg);
      color: var(--text);
      margin-bottom: 1rem;
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
      color: #0b1029;
      font-weight: 700;
      cursor: pointer;
    }
    .error {
      margin-top: 1rem;
      color: var(--danger);
      font-size: 0.9rem;
      min-height: 1rem;
    }
  </style>
</head>
<body>
  <div class="shell">
    <div class="topbar">
      <div>
        <div class="brand">Upwatch Admin</div>
        <div class="brand-sub">Secure login for monitor control.</div>
      </div>
      <button class="btn" id="themeToggle" type="button">Light mode</button>
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
  </div>
  <script>
    const params = new URLSearchParams(window.location.search);
    const errorEl = document.getElementById("error");
    const themeToggle = document.getElementById("themeToggle");

    const setTheme = (theme) => {
      document.documentElement.setAttribute("data-theme", theme);
      themeToggle.textContent = theme === "light" ? "Dark mode" : "Light mode";
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
