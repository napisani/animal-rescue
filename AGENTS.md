# AGENTS: animal-rescue (Go)

CLI for curating [pet](https://github.com/knqyf263/pet) snippets from multiple
sources (local TOML files, package.json, Makefiles). See `README.md` for
user-facing usage.

## Build & Test Commands

Prefer the `make` targets over invoking `go` directly.

- Build: `make build` (binary `animal-rescue`; embeds version/commit/date ldflags).
- Test: `make test` (`go test -v ./...`).
- Install: `make install` (`go install` into `$GOBIN`).
- Clean: `make clean`.
- Nix dev shell: `nix develop` (also `shell.nix`/`default.nix` for legacy nix).

## Layout

- `*.go` — flat package (main package), all source files in root.
- Module path: `github.com/napisani/animal-rescue`.

## Code Style & Conventions

- **Language**: Go. Standard `gofmt` formatting.
- **Errors**: return explicit errors; no panics.
- **Snippet sources**: each source type (local pet, Makefile, npm/package.json,
  additional pet config) is isolated behind the `SnippetSource` interface in
  `source.go`. Keep backends cleanly separated.

## Release

- `.goreleaser.yaml` drives release builds; `.drone.yml` is the upstream CI (in
  the standalone repo).
- This subproject is auto-published from the monorepo to
  [`napisani/animal-rescue`](https://github.com/napisani/animal-rescue) — work
  in the monorepo, never commit to the published repo directly (see root
  `AGENTS.md`).