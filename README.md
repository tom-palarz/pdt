# PDT — The Personal Data Toolkit

**PDT (`pdt`)** is a local-first command-line tool that helps you index, query, and organize the data companies hold about you. 

Unlike privacy tools focused solely on deleting your footprint, PDT helps you make sense of your digital history so you can organize your life. It initializes quickly from large offline archives (like Google Takeout `.mbox` files) and stays up to date via incremental API syncs.

## Who It’s For
- Anyone looking to audit every online account tied to their email address so they can lock down their password manager.
- Job seekers tracking every role and company they have applied to across years of emails.
- People looking to surface forgotten subscriptions, map contact touchpoints, or bulk-prune old emails safely.

## Built With
- **Go** — Fast, single-binary execution
- **Cobra** — Structured CLI commands and flags
- **SQLite** — Local, zero-config storage and full-text search
- **Gmail API & Mbox** — Hybrid bulk-import and incremental sync architecture

## Quick Start

### 1. Prerequisites
- Go 1.23 or newer
- (Optional) A Google Takeout `.mbox` export of your Gmail archive
- (Optional) Google Cloud OAuth credentials for live Gmail API syncing

### 2. Installation
```bash
git clone https://github.com
cd pdt
cp .env.example .env
go build -o bin/pdt ./cmd/pdt
```

### 3. Basic Usage (Planned)
```bash
# Seed local database from a large Gmail Takeout Mbox file
./bin/pdt import mbox ~/Downloads/All-mail.mbox

# Sync recent changes via the Gmail REST API
./bin/pdt sync gmail

# Query your data
./bin/pdt query accounts       # Find sites where you likely have an account
./bin/pdt query contacts       # List contacts and interaction counts
./bin/pdt query jobs           # List detected job applications
./bin/pdt query subscriptions  # Find recurring billing & newsletters
./bin/pdt labels               # List Gmail labels and message counts
```

## Roadmap
See [ROADMAP.md](ROADMAP.md) for current milestones, planned integrations (Google Drive, Google Tasks, LinkedIn, Facebook, DEVONthink/Bookends), and the upcoming BubbleTea TUI.

## License
Distributed under the MIT License. See [LICENSE](LICENSE) for details.

