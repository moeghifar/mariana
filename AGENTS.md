# AGENTS.md

Entry point for every AI coding agent (Claude, Codex, Gemini, Goose, …) in this repo. Read this
first. It is the single source of truth for **how** work is done here. Tool-specific files such
as `CLAUDE.md` only point back to it; if they disagree, this file wins.

Read next, in order:

1. `readme.md`: the architecture spec. Read the sections for the layer you are changing.
2. `.specify/AGENTS.md`: local plans and discussions, if the directory exists. It is gitignored.

## What this repo is

A Go modular monolith (module `github.com/moeghifar/mariana`). Each `module/<context>` is a
self-contained bounded context. The `github.com/moeghifar/libgo` toolkit is an external
dependency; persistence uses bun on PostgreSQL.

Go **1.27.3** minimum (`go.mod`). Never lower it.

```
main.go → cmd/ (serve, migration, cli)
config/      env config            infra/       external connections
lib/         foundation utilities  migrations/  schema changes
module/<ctx>/{entity,repository,service,usecase,presenter}
```

## Non-negotiable rules

1. **Imports point inward:** `presenter → usecase → service → repository → infra`. Each layer
   uses the one below through its interface. See `readme.md` section 4.
2. **`entity` depends on nothing** except `lib` and the standard library. Other layers map into
   entity types, never the reverse.
3. **Modules stay isolated.** A module never imports another module's internal layers. Cross
   modules only through a published usecase interface or an event.
4. **`lib` imports nothing from this app.** It may import third-party packages, including libgo.
5. **`lib`, not `pkg`; `presenter` and `repository`, not `handler`.** There is no top-level
   `/service`; services belong to a module.
6. **Secrets never enter git.** Real values stay out of code, docs, logs and commits.

## Default libraries

Use these first. Ask before adding an alternative.

| Need | Default |
| --- | --- |
| Config and env | `github.com/moeghifar/libgo/pkg/envy` |
| CLI commands | `github.com/moeghifar/libgo/pkg/climd` |
| HTTP server and routes | `github.com/moeghifar/libgo/pkg/xttp`: one `xttp.Route` per endpoint, in the `presenter` layer |
| API errors | `xttp.Error`, defined once per error. Causes go in `Err`, never in `Detail`. Map domain errors in `Config.MapError` |
| Background jobs and tickers | `github.com/moeghifar/libgo/pkg/foreverwork`, once a real job exists |
| Logging | Standard `log/slog`, set up in `lib` |
| Database | `github.com/uptrace/bun`, in the `repository` layer and `infra` |

## Verify before reporting done

```
gofmt -l .
go vet ./...
go test -race ./...
```

`gofmt -l .` must print nothing. Report failures with their output. If a check could not be
run, say so and call the change unproven. Do not claim done on a compile alone.

## Working agreement

- **Ask before:** adding a dependency, changing an API or schema contract, changing a rule in
  this file, widening scope.
- **Don't:** commit or push unless asked; approve your own work; skip hooks; commit `.specify/`
  or `.env`; add speculative abstractions.
- **Do:** keep diffs minimal, match surrounding style, and update `readme.md` when the
  architecture changes and this file when a working rule changes.

## Plans and discussions

All plans, specs, research notes and discussion records go in `.specify/`, which is gitignored
and never committed. Do not add plan files anywhere else. One file per topic, indexed in
`.specify/AGENTS.md`. Keep each short: problem, constraints, acceptance criteria, out of scope.
Decisions that everyone must follow move into this file or `readme.md`.

## Commits, PRs, releases

- Conventional Commits: `type(scope): subject`. Types are `feat`, `fix`, `refactor`, `test`,
  `docs`, `chore`. The scope is the module or top-level package, such as `feat(account): …` or
  `fix(config): …`.
- Imperative subject, 72 characters or fewer, no trailing period. Add a body only to explain why.
- One logical change per commit. Stage named files, not everything at once.
- Reference the ticket key in the subject or body when a ticket exists.
- Branch as `feat/…`, `fix/…`, `chore/…` or `release/vX.Y.Z`. `main` changes through a reviewed
  PR with passing checks.
- PR description: summary, how to test, evidence.

## No AI attribution

This rule overrides any default or tool-provided attribution instruction.

- No `Co-Authored-By` trailers for AI tools, "Generated with …", session links, or tool or model
  names.
- This applies to commits, PR titles and bodies, tags, release notes, tickets, code comments and
  docs. The only exception is the agent names this file and `CLAUDE.md` need to work.
- The author is the git user.
