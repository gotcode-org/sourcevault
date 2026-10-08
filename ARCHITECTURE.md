# SourceVault Architecture

This document outlines the high-level architecture and core design philosophies
behind SourceVault. If you are contributing to the codebase, understanding
these patterns is strictly required to maintain the system's security,
portability, and integrity.

## 1. The Git-First Paradigm

SourceVault rejects the traditional "Forge-as-a-Service" model where your
project data is locked inside a proprietary SQL database.

Instead, SourceVault relies entirely on a **Git-First Architecture** for all
codebase-related data. Your repository, issues, milestones, pull requests,
and all project metadata are serialized and committed directly into your
project's Git repository.

This means your project is inherently:
- **Portable:** Taking your project (code + issues + PRs) elsewhere is as
  simple as a `git clone`. (Future CLI tools will allow you to read and manage
  issues/PRs locally directly from your cloned repo).
- **Immutable:** Every project action leaves a cryptographically signed,
  tamper-evident audit trail.

**System vs. Project State:**
It is important to distinguish between project data and system data. While
project data lives in the Git repository, global system data (user accounts,
instance permissions, global IDs) is managed separately via the global event
stream. We explicitly avoid storing system data in user repositories to prevent
ID collisions and conflicts when a repository is exported and imported into a
different SourceVault instance.

## 2. Command Query Responsibility Segregation (CQRS)

Because reading directly from Git commits for every API request or UI render
would be prohibitively slow, SourceVault employs strict **CQRS**.

```text
    +----------------+
    | Client Request |
    +----------------+
            |
            v
       [ Handler ]
        /         \
    (Write)     (Read)
      /             \
      v               v
    [Command]       [Query]
      |               |
    (Mutates)       (Reads)
      |               |
      v               v
    +---------+     +---------+
    |   Git   |....>| SQLite  |
    | (Truth) | Sync| (Cache) |
    +---------+     +---------+
```

- **Commands (Writes):** Commands contain business logic that mutates the
  primary Source of Truth. They write directly to the Git repository or the
  Event Outbox.
- **Queries (Reads):** Queries contain no business logic and *never* mutate
  state. They simply query the local **SQLite Cache**.
- **The Read Model:** The SQLite database is purely an ephemeral query cache.
  It is populated by listening to the Git commit event stream. If the SQLite
  database is deleted, SourceVault can seamlessly rebuild it from the Git
  history.

## 3. Flat Hexagonal Architecture

SourceVault avoids deep, nested architectural folders (like standard DDD
`domain/`, `application/`, `infrastructure/` directories) to prevent "folder
soup." Instead, we use a flat, **Package-by-Feature** layout built on Hexagonal
Architecture (Ports and Adapters).

Every feature lives in isolation within `internal/<feature>/`. There is no
global `pkg/` directory.

### Standard File Layout
Within any given feature package, you will find a standardized set of files
separating the core logic from the outside world:

- `domain.go`: Contains the core business structs and pure domain logic.
- `ports.go`: Interfaces defining the required adapters (e.g., storage, network
  clients). Defined by the consumer.
- `commands.go`: The CQRS Command handlers (mutating state).
- `queries.go`: The CQRS Query handlers (reading state).
- `adapter_*.go`: The concrete implementations of the ports (e.g.,
  `adapter_sqlite.go`, `adapter_git.go`).

```text
  [ internal/<feature>/ ]
  
  commands.go \
               +--> domain.go
  queries.go  /
  
  commands.go - - > ports.go
  queries.go  - - > ports.go
  
  adapter_*.go - -> ports.go (Implements)
```

*Crucially: The core structs, commands, and queries NEVER directly invoke `os`
  file operations, `database/sql`, or network clients. They only ever interact
  with the locally defined interfaces (`ports.go`).*

## 4. Two-Binary Topology

To ensure security and separation of concerns, the project is split into two
distinct binaries:

1. `sourcevaultd`: The background server daemon handling Git serving, web UI,
   and APIs.
2. `sourcevault`: The unified CLI tool for operators and administrators.

Both entrypoints live under `cmd/` and are built using `spf13/cobra`. **No
business logic is permitted in the `cmd/` packages.** Their sole responsibility
is argument parsing, dependency injection (wiring adapters to ports), and
executing the feature handlers from `internal/`.

## 5. Security & Cryptography

SourceVault is designed to be FIPS-ready and to operate securely in zero-trust
environments.

- **Seal/Unseal:** Inspired by HashiCorp Vault, SourceVault uses a seal/unseal
  mechanism. Sensitive cryptographic material (like signing keys) is never
  stored in plaintext on disk and must be unsealed into memory upon daemon
  startup.
- **Internal SSH CA:** SourceVault features a built-in SSH Certificate
  Authority to manage secure access without relying on traditional long-lived,
  centrally-managed SSH keys.
- **mTLS by Default:** Because SourceVault is built for a decentralized future,
  node-to-node replication and syncing occur over standard `net/rpc` or `gRPC`,
  and Mutual TLS (mTLS) is strictly mandatory.
