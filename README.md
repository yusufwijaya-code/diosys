# Diosys

Diosys is a B2B technology agency website (custom web, mobile, and AI
development) with a public company site and an admin CMS. Content is managed
through the CMS, stored in MySQL, and images are stored on Google Drive.

> This project was rebuilt from a single-person portfolio app into the Diosys
> agency platform. Portfolio data is now per developer, and the public site
> aggregates multiple developers' portfolios.

## Applications

| Path | App | Stack |
| --- | --- | --- |
| [`api/`](api/) | Backend REST API | Go, Gin, sqlx, MySQL, golang-migrate, JWT, Google Sign-in, Google Drive |
| [`frontend/`](frontend/) | Public site + `/admin` CMS | Angular, TypeScript, SCSS |

```
ucup-porto-app/
├── api/        # Go REST API (public + CMS endpoints)
└── frontend/   # Angular app: public Diosys site (/) and CMS (/admin)
```

## Architecture

- The **public website** (`/`) reads content from the public API
  (`/api/public/*`): company services, the aggregated company portfolio, the
  developer directory, and per-developer profiles at `/[username]`.
- The **CMS** (`/admin`) authenticates with Google sign-in and manages
  developers, projects, the client inbox, and settings through the protected API
  (`/api/cms/*`).
- The **API** follows the `go-core-api` layered pattern
  (`handler -> service -> repository`), stores data in the MySQL `diosys_main`
  schema, and uploads images to Google Drive (storing only the file ID + name).

## Brand & design

The visual design follows `G:\My Drive\diosys-company\task\brand.md` — a premium
dark "quiet luxury" theme (Matte Charcoal `#0F1115` background, Soft Lavender
`#A5B4FC` CTAs, Inter font). There is no light/dark toggle. Logo:
`G:\My Drive\diosys-company\diosys-logo.svg`.

## Getting started

### Prerequisites

- Go 1.24+
- Node.js 20+ and npm
- MySQL with a `diosys_main` schema (created empty; tables are created by the API
  migrations on startup)
- A Google OAuth Client ID (for sign-in and Google Drive)

### 1. Backend

```bash
cd api
go mod tidy
go run .
```

The API connects to MySQL, runs migrations automatically, seeds the whitelisted
admin/developer, and starts on port `3005`. Configure it via `api/app.env` (or a
local override `api/app.local.env`). See [`api/README.md`](api/README.md) for
full documentation.

### 2. Frontend

```bash
cd frontend
npm install
npm run start
```

The app serves the public Diosys site at `/`, developer profiles at
`/[username]`, and the CMS at `/admin`.

## Authentication

- **Google sign-in only** — there is no registration and no passwords.
- Only the whitelisted email (`ADMIN_EMAIL`, default `yusufwijaya3@gmail.com`)
  may sign in to the CMS. All other Google accounts are rejected.
- The frontend obtains a Google ID token (Google Identity Services) and posts it
  to `POST /api/auth/google`; the backend verifies it, checks the whitelist, and
  returns a JWT used for subsequent CMS calls.

## Conventions

- **Code & comments:** English only.
- **Database:** schema `diosys_main`; camelCase columns; primary keys named
  `<entity>ID` (e.g. `userID`, `projectID`); table prefixes `ms_` (master),
  `map_` (mapping), `lk_` (lookup); nested lists normalized into child tables.
- **API responses:** standard envelope
  (`path/timestamp/status/code/message/result/errors`).
- **Images:** stored on Google Drive (root folder
  `1SGW0xhGLAeNjFAbDz5O0-fm8Pk_9Lo4r`); the database keeps only the Drive file ID
  and file name, and the public URL is derived at read time.
