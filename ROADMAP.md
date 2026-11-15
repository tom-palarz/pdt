# PDT Roadmap

## Phase 1: Foundation & Gmail Core (v0.1.0)
- [ ] CLI scaffolding with Cobra (`init`, `import`, `sync`, `query`, `labels`, `prune`)
- [ ] Local SQLite schema for messages, headers, labels, and sync cursors
- [ ] Streaming `.mbox` parser capable of handling multi-gigabyte archives
- [ ] Gmail REST API OAuth2 authentication and incremental history sync (`historyId`)

## Phase 2: Life-Organization Queries (v0.2.0)
- [ ] **Account Discovery:** Heuristic detection of signup, password reset, and welcome emails
- [ ] **Contact Graph:** Extract names, emails, phone numbers, vCards, and interaction frequency ("touches")
- [ ] **Job Tracker:** Identify application confirmations, recruiter outreach, and interview invites
- [ ] **Subscription Audit:** Flag recurring receipts and mailing lists
- [ ] **Bulk Pruning:** Safe dry-run and batch deletion/archiving rules via Gmail API

## Phase 3: Interactive TUI (v0.3.0)
- [ ] Interactive terminal dashboard built with **BubbleTea** and **Lipgloss**
- [ ] Visual contact timeline and subscription cancellation/pruning workflows

## Phase 4: Expanded Data Connectors (v0.4.0+)
- [ ] Google Drive metadata & file index
- [ ] Google Tasks import & sync
- [ ] LinkedIn data archive parser (connections, messages, job applications)
- [ ] Facebook data archive parser
- [ ] Local research & document tag integration (DEVONthink and/or Bookends metadata)

