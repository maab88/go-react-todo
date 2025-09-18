# Go + React Todo Starter (Monorepo)

This repository contains:
- `/server` – Go (Fiber + GORM + SQLite) API
- `/web` – React + Vite + TypeScript frontend

## Run locally

### Start the API
```bash
cd server
go mod tidy
go run ./cmd/server
```
API on http://localhost:8080

### Start the frontend
```bash
cd web
npm install
npm run dev
```
Web on http://localhost:5173

The frontend is configured to proxy `/api` to the API server.

## Create a GitHub repo

From the project root:
```bash
git init
git add .
git commit -m "chore: initial commit (Go + React todo starter)"
git branch -M main
git remote add origin https://github.com/<your-username>/<repo-name>.git
git push -u origin main
```

## Next steps

- Add due dates, priorities, and editing UI.
- Add tests (Go + Playwright).
- Swap SQLite for Postgres if needed.
