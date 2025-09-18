# Go + React Todo Starter

A minimal, **ready-to-run** Todo app:

- **Backend:** Go (Fiber) + GORM + SQLite (no external DB needed for dev)
- **Frontend:** React + Vite + TypeScript + Tailwind + React Query
- **Features:** list, add, toggle status, delete, filter by status/search

## Quick Start

### 1) Backend
```bash
cd server
go mod tidy
go run ./cmd/server
# Server runs on http://localhost:8080
```

### 2) Frontend
```bash
cd web
npm install
npm run dev
# Vite runs on http://localhost:5173
```

The frontend proxies API calls to `http://localhost:8080`.

## API (quick reference)
- `GET    /api/tasks?status=&q=`
- `POST   /api/tasks` { title, description?, dueDate?, priority? }
- `PUT    /api/tasks/:id` { any fields }
- `PATCH  /api/tasks/:id/status` { status }
- `DELETE /api/tasks/:id`

## Environment
Duplicate `.env.example` to `.env` in `/server` if you want to override defaults.

## GitHub repo setup
From the project root:
```bash
git init
git add .
git commit -m "chore: initial commit (Go + React todo starter)"
# create a new repo on GitHub, then:
git branch -M main
git remote add origin https://github.com/<your-username>/<repo-name>.git
git push -u origin main
```

## Notes
- SQLite file is created next to the server binary (default `todo.db`).
- You can later switch to Postgres by replacing the GORM driver and DSN.
- CORS is enabled for dev; lock it down for prod.
