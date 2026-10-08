# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial repository scaffolding and open-source governance documents.
- `AGENTS.md` to define architectural boundaries and AI agent workflows.
- `SECURITY.md` and `CONTRIBUTING.md` policies.
- Standard `LICENSE`, `DCO`, and `MAINTAINERS` files.
- Go module initialization (`gotcode.org/sourcevault`).
- `Makefile` with `tidy`, `fmt`, `vet`, `lint`, `sec`, `test`, and `build` targets.
- Cobra CLI scaffolding for `sourcevault` and `sourcevaultd` with custom lipgloss UI styling.
- Hexagonal CQRS architecture implementation for the `version` command and system information retrieval.

### Changed
- Refactored `Makefile` `-ldflags` injection to target `internal/system` variables instead of CLI `main` variables, fully decoupling version state from the Cobra entrypoints.

### Fixed
- Removed accidentally committed binary files (`bin/sourcevault`, `bin/sourcevaultd`) from the Git index and added a `.gitignore` to prevent future tracking.
