# Development and testing

How to change Nap: tests, docs, the flow of a change, and Go conventions.

## What needs which tests

| Area | Tests |
|---|---|
| `internal/version` | Unit tests of the fallback from the build info. |
| `cmd/nap` | Unit tests of the flags. |
| The docs | `go test ./...` checks the test names and links in the pages, the Nomad version in the README, and that every copy of the invariant titles matches [philosophy](philosophy.md#invariants). |

A behaviour that depends on Nomad itself gets an integration test. Each
invariant gets an adversarial test that tries to break it.

## Integration

Integration tests live in `tests/integration/` behind the `integration` build
tag. They run against `nomad agent -dev` with ACLs, using
`NAP_TEST_NOMAD_ADDR` and `NAP_TEST_NOMAD_TOKEN`; without the address they skip.

### Nomad version

CI tests against `NOMAD_VERSION` in `.github/workflows/ci.yml`; the README
states the same version, and a test fails if they differ. To move to a new
Nomad, change `NOMAD_VERSION` and `NOMAD_SHA256`, update the README, and run
the integration tests.

## Coverage

At least 80% on `engine` and `meta`, once they exist. No target on `web` and
`cmd`.

## Writing docs

The docs follow common open source practice:

- **Structure:** [Diátaxis](https://diataxis.fr). Each page is a tutorial, a
  how-to guide, reference or explanation, and holds only that kind of content.
  The [index](README.md) groups pages the same way.
- **Style:** the [Google developer documentation style guide](https://developers.google.com/style).
  Vale checks it in CI ([commands](#commands)).
- **Content:** what a reader needs to use or change Nap. Leave out what the
  code already answers. State each fact on one page and link to it from the
  others.
- **Commits, issues and pull requests** follow
  [Conventional Commits](#commit-messages) and the repository's templates,
  and say what changes and why.

## When a piece of work is "done"

1. `go build`, `go vet`, `staticcheck`, `go test -race` and Vale pass.
2. Changes to `engine` or `nomadx` pass the integration tests.
3. Every behaviour change has a test.
4. The docs and `examples/` match the change.

## What to update for each change

| Change | Update |
|---|---|
| Meta key | `internal/meta` and `meta-keys.md` |
| Flag or environment variable | `internal/config` and `configuration.md` |
| Wake or sleep behaviour | `architecture.md` |
| Metric | `metrics.md` |
| Term | `glossary.md` |
| Invariant | `philosophy.md` |
| Doc page | the [index](README.md) |
| Nomad version | `ci.yml` and the README |

## Commands

```sh
go test -race -cover ./...
go tool staticcheck -tags integration ./...
vale sync && vale docs README.md CONTRIBUTING.md

# integration: a dev agent with ACLs, in another terminal
printf 'acl {\n  enabled = true\n}\n' > agent.hcl
nomad agent -dev -config=agent.hcl &
export NAP_TEST_NOMAD_ADDR=http://127.0.0.1:4646
export NAP_TEST_NOMAD_TOKEN=$(curl -s -X POST $NAP_TEST_NOMAD_ADDR/v1/acl/bootstrap | jq -r .SecretID)
go test -tags integration -race -count=1 ./tests/integration/...
```

## Workflow and CI

1. Pick an open issue with no branch or pull request. Settle open design
   questions in the issue first.
2. Branch from `main` as `type/short-description`.
3. Make the change with its tests and docs.
4. Open a pull request titled `type(scope): subject`. Fill in the template with
   *What changes*, *Why*, and *Notes* for open questions, and add `Closes #N`.
5. After review, push new commits. Don't amend or force-push.
6. The maintainer squash-merges and writes the final commit message.

`main` requires a pull request with these checks green:

| Check | Runs |
|---|---|
| `test` | `gofmt`, `go mod tidy`, build, `go vet`, `go test -race` |
| `lint` | `staticcheck`, Vale |
| `integration` | The integration tests against a Nomad dev agent |
| `pr-title` | The pull request title is a Conventional Commit |

`govulncheck` and Dependabot also run. Actions are pinned to full commit SHAs.

Issues use the forms in `.github/ISSUE_TEMPLATE/`: a plain title, the type
(`Bug`, `Feature` or `Task`), and a label for each area it touches.

## Commit messages

[Conventional Commits](https://www.conventionalcommits.org/) in the Angular
style:

```
<type>(<scope>): <subject>

<body: what changed and why>

<footer: BREAKING CHANGE, issue references, Co-Authored-By>
```

- **Types:** `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`,
  `build`, `ci`, `chore`, `revert`.
- **Subject:** imperative, lowercase, no trailing period, 72 characters at
  most. It becomes a line of the release notes, so write it for someone
  running Nap.
- **Body:** wrapped at 72 characters, plain text. Say what changed and why.
- **Breaking change:** `!` after the scope and a `BREAKING CHANGE:` footer
  that says what changed and what to do.

## Terms used in the code

The words of [the glossary](glossary.md), plus:

| Term | Meaning |
|---|---|
| **Nap** / **`nap`** | **Nap** is the product; **`nap`** in code font is the binary, image, module, or a name in code. |
| **engine** | `internal/engine`: the state of the targets, and the wake and sleep loops. |
| **sentinel error** | An exported `Err...` value checked with `errors.Is`. |
| **fake / stub** | A fake is an in-memory stand-in; a stub is an `httptest` server with canned answers. |
| **in flight** | An issue with a branch or an open pull request. |

## Go conventions

- English everywhere: code, comments, docs, commits.
- `gofmt`, `go vet` and `staticcheck`.
- Short, singular package names. No `util` or `common`.
- `context.Context` first in every function that does I/O.
- Wrap errors with context: `fmt.Errorf("scale %s: %w", target, err)`.
  Export sentinels only when callers need them.
- Consumers define small interfaces.
- No `init()` with side effects, no global state; inject time.
- Log with `log/slog`.
- Every flag has its `NAP_*` variable; every interval and timeout is
  configurable.
- Prefer an established package to writing a new library.
