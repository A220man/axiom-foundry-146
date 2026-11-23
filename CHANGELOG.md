# Changelog

## 1.0.0 — work-in-progress checkpoint

Partial code published for continued development. Known issues: STATUS.md.

## v0.1.0 - 2026-11-23
### Fixed
- Fix `panic: close of nil channel` in sessions refresh (`internal/sessions/sessions.go`).
  The `publish()` function now guards against duplicate paths in `todo` and nil channels
  in the `parsing` map, preventing the crash that caused magpie to quit on every launch
  starting in v0.1.1291.

### Added
- Initial Go module `axiom-foundry-146`
- `internal/sessions` package with `publish()`, `Refresh()`, and `StatsAndNative()` functions
- Test suite with 3 tests verifying the nil-channel fix

## Planned
- GUI titlebar/darwin corner handling (issue #3949)
- Chinese release notes language marker (issue #2260)
