# Portfolio API (Go)

Backend REST API for the portfolio app. It stores portfolio content in MySQL
(schema `portfolio`), exposes public read endpoints for the website and
protected CMS endpoints (JWT) for content management, and stores images on
Google Drive.

The architecture follows the layered, module-per-feature pattern of
`go-core-api` (`handler -> service -> repository`), simplified to a single
database with manual constructor dependency injection.

## Tech stack

- Go + Gin (HTTP)
- sqlx + go-sql-driver/mysql (database)
- golang-migrate (schema migrations, embedded and auto-run on startup)
- golang-jwt/jwt v5 (authentication)
- Google Drive API v3 + OAuth2 refresh token (image storage)
- viper (configuration)

## Project structure

```
api/
├── main.go                 # entrypoint: config, db, migrate, seed admin, run server
├── router.go               # SetupRouter: registers all module routes
├── migrate.go              # embedded migration runner
├── app.env                 # configuration (DB, JWT, Google Drive)
├── migrations/             # SQL migrations
├── config/                 # config loader
├── database/               # sqlx connection
├── constants/              # response codes
├── middlewares/            # jwt, cors, panic recovery
├── base/helpers/           # http, error, context, jwt, password, gdrive helpers
└── modules/
    ├── auth/               # login, me
    ├── user/               # admin accounts (+ default admin seeding)
    ├── personal_info/      # single profile + photo upload
    ├── summary/            # about + stats/facts/skills (relational children)
    ├── experience/         # + technologies, responsibilities
    ├── education/          # + achievements
    ├── certificate/
    ├── skill/
    ├── language/
    └── project/            # + technologies, thumbnail upload
```

Each module folder contains `*_model`, `*_dto`, `*_repository`, `*_service`,
`*_handler` and a `router.go`.

## Configuration

Edit `app.env` (or create `app.local.env` to override it locally — it is
git-ignored). Defaults target a local XAMPP/MySQL install.

| Key | Description |
| --- | --- |
| `APP_PORT` | HTTP port (default `3005`) |
| `APP_FRONTEND_URL` | Allowed CORS origin |
| `DB_HOST` / `DB_NAME` / `DB_USERNAME` / `DB_PASSWORD` | MySQL connection (`portfolio`) |
| `JWT_SECRET` / `JWT_TOKEN_LIFESPAN_MINUTES` | JWT signing and lifespan |
| `GOOGLE_DRIVE_CLIENT_ID` / `_CLIENT_SECRET` / `_REFRESH_TOKEN` / `_FOLDER_ID` | Google Drive credentials and root folder |

## Running

```bash
cd api
go mod tidy
go run .
```

On startup the API:
1. Connects to MySQL and applies pending migrations automatically.
2. Seeds a default admin account if no users exist.
3. Initializes the Google Drive client (disabled gracefully if credentials are missing).

### Default admin

| Username | Password |
| --- | --- |
| `admin` | `admin123` |

> Change this password after the first login in production.

## API overview

All responses use the envelope:

```json
{ "path": "...", "timestamp": "...", "status": "ok|fail", "code": "200",
  "message": "...", "result": {}, "errors": [] }
```

### Auth
- `POST /api/auth/login` — `{ "username", "password" }` → access token + user
- `GET  /api/auth/me` — current user (Bearer token)

### Public (read-only, no auth) — `GET /api/public/...`
`personal-info`, `summary`, `experiences`, `educations`, `certificates`,
`skills`, `languages`, `projects`

### CMS (Bearer token) — `/api/cms/...`
- `personal-info`: `GET`, `PUT`, `POST /photo` (multipart `file`)
- `summary`: `GET`, `PUT` (replaces nested stats/facts/skills)
- `experiences` / `educations` / `certificates` / `skills` / `languages`:
  `GET`, `GET /:id`, `POST`, `PUT /:id`, `DELETE /:id`
- `projects`: the above + `POST /:id/thumbnail` (multipart `file`)

## Database conventions

- Schema: `portfolio`
- Columns: camelCase (e.g. `createdDate`)
- Primary keys: `<entity>ID` (e.g. `userID`, `projectID`)
- List children are normalized into their own tables with `ON DELETE CASCADE`
- Images: only the Google Drive file ID (`*GdriveID`) and file name are stored;
  the public URL is derived at read time.
