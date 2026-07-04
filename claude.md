# Diosys

## Project Overview

This repository consists of two applications:

- `api/` : Backend REST API written in Go
- `web/` : Angular web Application (public company website + `/admin` CMS)

Diosys is a B2B technology agency website (custom web/app/AI development). The
web consumes APIs exposed by the backend. Content is stored in MySQL
(schema `diosys_main`) and images are stored on Google Drive.

> Heritage: this project was rebuilt ("rombak total") from a single-person
> portfolio app into the Diosys agency platform. The portfolio schema now lives
> per developer; the site aggregates multiple developers' portfolios.

Brand / design system: `G:\My Drive\diosys-company\task\brand.md` — premium dark
"quiet luxury" theme (Matte Charcoal `#0F1115`, Soft Lavender `#A5B4FC` CTA,
Inter font). There is **no** light/dark toggle; follow the theme inherently.
Logo: `G:\My Drive\diosys-company\diosys-logo.svg`. All UI/content/code in English.

---

## Project Structure

```text
ucup-porto-app/
├── api/
└── web/
```

### Backend

Location: `api/`

Tech Stack:
- Go + Gin (HTTP)
- sqlx + MySQL (schema `diosys_main`)
- golang-migrate (embedded migrations, auto-run on startup)
- golang-jwt v5 (session tokens)
- Google Identity Services / `google.golang.org/api/idtoken` (Google sign-in)
- Google Drive API v3 (image storage)
- viper (configuration)

Architecture: layered, module-per-feature, following the `go-core-api` design
pattern (`handler -> service -> repository`), simplified to a single database
with manual constructor dependency injection.

Responsibilities:
- Authentication: **Google sign-in only**, strict single-email whitelist
  (`ADMIN_EMAIL`); no registration, no passwords. Issues a JWT after verifying
  the Google ID token.
- Developer management: each developer is an `ms_user` with a public profile at
  `/[username]` aggregating their summary, experience, education, certificates,
  skills, languages, and projects.
- Rich project portfolios (features, technology tags, gallery images, live /
  repo links, status, featured flag).
- Agency content: services, client inbox (contact form messages), system
  settings (e.g. dynamic WhatsApp number/message).
- Image upload to Google Drive.
- Public read-only API endpoints consumed by the website.

See `api/README.md` for full endpoint and setup documentation.

### web

Location: `web/`

Tech Stack:
- Angular
- TypeScript
- Vite
- SCSS

Responsibilities:
- Public Diosys website (`/`): company overview + services, company portfolio
  (aggregated from developers), developer directory, contact form + dynamic
  WhatsApp CTA.
- Developer profile pages at `/[username]` (e.g. `/yusuf-wijaya`).
- CMS admin UI (`/admin`, Google-login protected): developers, projects, inbox,
  settings.
- Consuming backend APIs, responsive design, user interaction.

---

## Development Rules

### General

- Always understand existing code before modifying it.
- Prefer extending existing modules instead of creating duplicate implementations.
- Keep code simple and maintainable.
- Follow existing project patterns.
- Avoid unnecessary refactoring outside the requested scope.

### Naming Convention

- Use meaningful variable names.
- Avoid abbreviations unless already established in the codebase.
- Use English for code, comments, and documentation.

### Comments

- Only add comments for non-obvious business logic.
- Do not add redundant comments that merely repeat the code.

---

## web Rules

Location: `web/`

### Angular

- Prefer standalone components if the project already uses them.
- Keep components focused on a single responsibility.
- Move reusable logic into services.
- Use TypeScript strict typing whenever possible.
- Avoid using `any` unless absolutely necessary.

### UI

- Follow the `brand.md` design system; no light/dark toggle.
- Maintain responsive design.
- Reuse existing styling patterns.
- Avoid inline styles unless necessary.
- Prefer SCSS classes over style attributes.

### API Integration

- Centralize API calls in services.
- Avoid direct HTTP calls inside components.
- Handle loading, success, and error states properly.

---

## Backend Rules

Location: `api/`

### Architecture

- Follow the `go-core-api` layered pattern: each feature lives under
  `modules/<feature>/` with `*_model`, `*_dto`, `*_repository`, `*_service`,
  `*_handler` and a `router.go`.
- Keep the dependency direction `handler -> service -> repository`.
- Wire dependencies with manual constructors (`New...`) inside each module's
  `router.go`.
- Define interfaces for repository, service and handler layers.

### API Development

- Follow existing endpoint conventions and the standard response envelope
  (`path/timestamp/status/code/message/result/errors`) via `http_helper`.
- Validate all incoming request data; return errors through `error_helper`.
- Public read endpoints live under `/api/public`; protected CMS endpoints under
  `/api/cms` (JWT). Developer-scoped CMS resources are nested under
  `/api/cms/developers/:userID/...`.
- Auth: `POST /api/auth/google` (verify Google ID token + whitelist → JWT),
  `GET /api/auth/me`.

### Database

- Schema `diosys_main`. Column names in **camelCase**; primary keys named
  `<entity>ID` (e.g. `userID`, `projectID`).
- Table prefixes: **`ms_`** master, **`map_`** mapping/junction, **`lk_`**
  lookup.
- Normalize nested lists into child tables with `ON DELETE CASCADE`.
- Use sqlx through the repository layer; wrap multi-statement writes in a
  transaction.
- Schema changes go through new files in `api/migrations/` (golang-migrate).

### Images / Google Drive

- Store only the Google Drive file ID (`*GdriveID`) and file name in the
  database; derive the public URL at read time via `gdrive_helper`.
- Root folder ID: `1SGW0xhGLAeNjFAbDz5O0-fm8Pk_9Lo4r`.

### Security

- Never expose secrets or credentials (Google OAuth / JWT secrets stay in env).
- Validate and sanitize user input.
- Apply authorization checks where required (JWT middleware on CMS routes; the
  Google sign-in whitelist gates all CMS access).

---

## Testing Checklist

Before completing a task:

### web

```bash
cd web
npm install
npm run start
```

Ensure:
- Application builds successfully
- No TypeScript errors
- No console errors

### Backend

```bash
cd api
go mod tidy
go build ./...
go vet ./...
go run .
```

Ensure:
- The project builds and `go vet` is clean
- Migrations apply successfully on startup
- Affected endpoints respond as expected

---

## Git Rules

- Make minimal and focused changes.
- Do not modify unrelated files.
- Preserve existing code style.
- Keep commits scoped to a single purpose.

---

## When Making Changes

Always:

1. Read related files first.
2. Understand current implementation.
3. Propose the simplest solution.
4. Update only necessary files.
5. Verify no existing functionality is broken.
