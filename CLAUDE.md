# PDT (Personal Data Toolkit) — Claude Code Instructions

## Project Overview
PDT is a local-first Go CLI tool (`pdt`) that ingests personal data exports (starting with Gmail Mbox + Gmail REST API) into a local SQLite database so users can query their digital footprint (accounts, contacts, job applications, subscriptions, labels) and bulk-prune old data.

## Tech Stack
- **Language:** Go 1.23+
- **CLI Framework:** `Cobra (https://github.com/spf13/cobra)`
- **Database:** Local SQLite (`modernc.org/sqlite` for pure-Go CGo-free builds)
- **Future TUI:** `Bubble Tea (https://github.com/charmbracelet/bubbletea)`

## Architecture & Conventions
- Keep `cmd/pdt/main.go` minimal; place command definitions in `internal/cli/`.
- Stream Mbox files line-by-line or message-by-message; never load an entire `.mbox` file into memory (archives can exceed 20GB).
- Store all user data locally in `~/.pdt/pdt.db` by default.
- Never log email bodies, OAuth tokens, or personal data to stdout/stderr unless explicitly requested by a query command.

## Security Rules
- NEVER hardcode API keys, OAuth client secrets, or tokens.
- Ensure `.env`, `credentials.json`, `token.json`, `*.mbox`, and `*.db` remain gitignored.

