# App

Монорепозиторий приложения: два изолированных приложения — серверное и клиентское.

## Структура

```
App/
├── backend/     # Серверная логика на Go
│   ├── go.mod   # Go-модуль app/backend
│   └── main.go  # Точка входа: HTTP-сервер, эндпоинт GET /healthz
├── frontend/    # Клиент на TypeScript + React (Vite)
│   ├── package.json
│   └── src/     # main.tsx (вход), App.tsx (корневой компонент)
├── adr/         # Архитектурные решения (ADR)
├── openspec/    # Спецификации (SSOT) и change-процесс OpenSpec
├── AGENTS.md    # Инструкции для AI-агентов
└── README.md
```

Каждая часть самодостаточна: собственные зависимости и собственные команды, выполняемые из своего каталога.

## Разработка: OpenSpec и SDD

Проект ведётся по подходу SDD (spec-driven development) с помощью OpenSpec (CLI: `openspec`):

- Спецификации в `openspec/specs/` — единственный источник истины (SSOT) о поведении системы.
- Архитектурные решения фиксируются в ADR (`adr/`, конвенции схемы [`spec-driven-with-adr`](https://github.com/intent-driven-dev/openspec-schemas/tree/main/openspec/schemas/spec-driven-with-adr)): ADR — решение и обоснование (why), спека — поведение (what). Принятые ADR неизменяемы, обязательны и автоматически подгружаются агенту (см. `opencode.jsonc`).
- Изменение поведения оформляется как change в `openspec/changes/` (артефакты: proposal → specs → design → tasks) и только затем реализуется в коде.
- Сложные задачи сначала исследуются (`/opsx-explore`); конвенции артефактов закреплены в `rules` файла `openspec/config.yaml`; capability-спеки организуются по частям: `specs/backend/<capability>/`, `specs/frontend/<capability>/`.
- После реализации change проверяют и архивируют (`openspec archive`) — его дельты синхронизируются с основными спеками.

## Требования

- **Go ≥ 1.27** — https://go.dev/dl/
- **Node.js ≥ 24 (LTS)** и npm — рекомендуется через [nvm](https://github.com/nvm-sh/nvm): `nvm install --lts`

## Backend

```bash
cd backend
go build ./...   # сборка
go run .         # запуск (порт по умолчанию 8080, переопределяется: PORT=9090 go run .)
```

Проверка работоспособности: `curl http://localhost:8080/healthz` → `200 OK`.

## Frontend

```bash
cd frontend
npm install      # установка зависимостей (один раз)
npm run dev      # dev-сервер → http://localhost:5173
npm run build    # производственная сборка в dist/
```
