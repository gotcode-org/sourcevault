# SourceVault

**A next-generation, Git-first code collaboration forge. Audit-first, designed
for absolute data sovereignty, and built for a decentralized future.**

![Status](https://img.shields.io/badge/Status-In%20Development-yellow)
![Language](https://img.shields.io/badge/Language-Go-blue)
![License](https://img.shields.io/badge/License-AGPL--3.0-green)

---

## What is SourceVault?

SourceVault is a sovereign code forge built to disrupt the centralized
"Forge-as-a-Service" model. Unlike traditional platforms (like GitHub, GitLab,
or Gitea) where your project's issues, PRs, and user data are trapped in a
proprietary SQL database, SourceVault uses an entirely **Git-First
Architecture**.

It treats an internal Git repository as the authoritative System Registry for
all system state, using the database purely as a high-performance query cache.
If you ever leave SourceVault, you take your entire project history—including
all metadata—with you via a simple `git clone`.

> **Note:** SourceVault is currently in active development.
> Keep your eyes out for the official `v0.1.0` release! Also note that
> SourceVault is **not** a fork of existing projects like Gitea or GitLab. It
> is a completely clean-room implementation built from scratch. While we draw
> deep insights and architectural inspiration from the giants that came before
> us, SourceVault puts a radically different spin on code collaboration by
> keeping the source-of-truth purely within Git.

## Core Architecture & Features

### 📦 Git-First Persistence
Every administrative action, user creation, or issue comment is committed
directly to the internal Git registry. Your infrastructure is inherently
portable, immutable, and fully transparent.

### 🧩 Two-Binary Architecture
Strictly partitioned logic between:
- `sourcevaultd`: The background server daemon.
- `sourcevault`: The unified administration CLI.

### 🔒 Vault-Style Security
Built-in SSH Certificate Authority (CA) with a HashiCorp Vault-style
"Seal/Unseal" mechanism to protect sensitive cryptographic signing keys in
memory.

### 📋 Audit-First Compliance
Because the source of truth is Git, every action leaves a cryptographically
signed trail. It is designed from the ground up to meet the highest standards
for enterprise and government security (FIPS ready).

### 🌐 Federated by Design
Built for a decentralized future. SourceVault aims to allow independent,
sovereign forges to seamlessly collaborate, sync, and merge patches without a
single central authority.

### 🏛️ Hexagonal Architecture (CQRS)
Engineered using strict Hexagonal Architecture (Ports and Adapters) and Command
Query Responsibility Segregation (CQRS) to ensure high cohesion, low coupling,
and pure business logic.

## License

SourceVault is licensed under the AGPL-3.0 License.
