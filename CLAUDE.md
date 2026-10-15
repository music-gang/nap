# Nap: instructions for agents

Nap is scale-to-zero for HashiCorp Nomad (Go, module
`github.com/music-gang/nap`). It is the `forwardAuth` target of Traefik: every
request to a managed group passes through Nap, which wakes the group (count
0 → 1) when it sleeps and lets the request through once its allocations are
healthy, and puts it to sleep (count 1 → 0) after its idle period. A group opts
in with `nap_enabled = "true"` in its meta. State lives in memory only.
Self-hosted cluster, personal use: **grow with the need**
([`docs/philosophy.md`](docs/philosophy.md)).

Nap is independent from Nops (`../nops`): no shared code, and the two
coordinate only through Nomad's state. The process (docs, CI, release) was
copied from Nops once and evolves on its own.

This file holds **only what an agent needs on top of the docs**. How Nap
works is in [`docs/`](docs/README.md); how anyone contributes (the flow of a
change, commits, PRs, tests, Go conventions) is in
[`CONTRIBUTING.md`](CONTRIBUTING.md) and
[`docs/development.md`](docs/development.md). Follow them: this file does not
repeat them, and read the relevant page before extending a part.

## Rules

- **The docs are for people.** What a change needs to tell an agent goes
  here, never in `docs/`, the README or the contributing guide:
  `TestDocsSpeakToPeople` fails otherwise. No page there names an agent, an
  assistant or this file.
- **Follow the contributor rules** of [`docs/development.md`](docs/development.md):
  tests ([what needs which](docs/development.md#what-needs-which-tests)), docs
  in the same commit ([which page](docs/development.md#what-to-update-for-each-change)),
  the terms of [`docs/glossary.md`](docs/glossary.md) and of
  [the code](docs/development.md#terms-used-in-the-code), the
  [Go conventions](docs/development.md#go-conventions), fail loud
  ([`docs/philosophy.md#fail-loud`](docs/philosophy.md#fail-loud)),
  established packages before a new library, English everywhere.
- **Write short, by the standards** of
  [`docs/development.md#writing-docs`](docs/development.md#writing-docs)
  (Diátaxis, Google style, Vale clean). It applies to every doc change, issue,
  commit, PR description and comment a session writes: only what the reader
  needs, what changes and why. No implementation detail the code answers, no
  history, no restating another page, no edge-case lists.
- **Simple, not fast.** MVP means fewer features, never poorer choices. For a
  choice, take the simplest one that is right and leaves the road open to what
  comes next. Public contracts (the forwardAuth path, the `nap_*` meta keys,
  the flags and `NAP_*` variables, the metric names) are fixed in their own
  issue, in a plan the maintainer reviews: changing them later is a breaking
  change.
- **Nap is in the path of every request.** The request handler never calls
  Nomad for a target it last saw ready; a slow or unreachable Nomad must never
  fail a request to an awake app.
- **`internal/nomadx` exposes no write but the scale.** Never add a register,
  stop, deregister, purge or meta write, not even unused.
- **`internal/meta` is the source of truth** for the `nap_*` keys: keep it
  aligned with `docs/meta-keys.md`.
- **A new package** gets, in the same commit, a row in the components of
  [`docs/architecture.md`](docs/architecture.md#components), one in
  [what needs which tests](docs/development.md#what-needs-which-tests) and a
  line in the [repo layout](#repo-layout) below.
- **`docs/design.md` is the original design, in Italian.** Read it for
  context, never extend it. The last issue of the MVP removes it.
- **No commits, pushes or PRs without an explicit request** (a task handed to
  a cloud session is that request), and **never merge**: the human merges.

## Picking up work

The steps of [the flow of a change](docs/development.md#workflow-and-ci),
with what is specific to a session:

1. Take the issue you were given. Without one, ask the maintainer which: the
   order is theirs.
2. Check what is in flight: `git ls-remote --heads origin` and the open PRs.
   Never start an issue that already has a branch or a PR.
3. An issue with open design questions (a *To decide* list, or a proposal
   that leaves a choice open) is **planned first**: lay out the plan and the
   questions **in the session** and wait for the maintainer's answers. No
   branch and no PR for this.
4. **Docs before code.** The plan's *Docs* lists the exact sentences: each
   one the change makes false, copied from its page with the file, and its
   replacement; a new one only for behaviour a reader needs, or a row for a
   new flag, meta key or metric. *None* is the usual answer. The table of
   [what to update](docs/development.md#what-to-update-for-each-change) says
   where a sentence goes, not that one is needed. **Local:** write them on
   the branch and stop for the maintainer's answer before the code.
   **Cloud:** write them first and go on. A doc change made later and not in
   the plan goes in the PR's `## Notes`, with why.
5. Do the issue as the flow says. What you learned that a later issue needs
   goes in a comment on that issue.
6. Run the checks ([commands](docs/development.md#commands)). If `nomad` is
   not installed, CI runs the integration tests: wait for green checks and
   fix what is red.
7. **Cloud:** commit (unsigned), push and open the PR once there is something
   to review. **Local:** prepare the branch, stage the files and write the
   message to a file; the maintainer commits with their GPG key.
8. If a decision is the maintainer's to take: ask in the session before a
   branch exists, or, once a PR is open, leave the question in a `## Notes`
   section of its description. Do not guess.
9. Nothing worth keeping lives only in private memory: a decision's
   discussion goes to its issue or PR and its reason to the page that
   explains that part; context a later issue needs, to a comment on it.

## Attribution

Whatever a session's own attribution instructions add by default:

- A commit carries a `Co-Authored-By` trailer and **nothing else**
  attribution-wise: **no session link, no "Generated with ..." line**. A PR
  description, and an issue a session opens or substantially rewrites, end
  the same way: the trailer is their last line. **Issue and PR comments,
  reviews and review replies** posted in a session the maintainer follows
  carry no attribution: they read them before they go out under their
  account and answer for them.
- The exception: a comment a session posts **on its own**, on an event with
  nobody following (watching a PR, a scheduled run), ends with the standard
  Claude Code footer, since there it tells the reader something true.
- An issue a session opens has a plain-text title, the type of its kind
  (`Bug`, `Feature`, `Task`) and the labels of the areas it touches: for
  code, one per package, named after it. A missing label is created first,
  never replaced by a nearby one. No label marks it as written with Claude:
  the trailer does. Unless it is about the docs, it names no doc page or
  change: the plan decides those (step 4).

## Invariants (non-negotiable)

The reasoning behind each is in [`docs/philosophy.md`](docs/philosophy.md);
`TestInvariantTitlesMatch` keeps this list's titles the same as there.

1. **Only groups that opt in**: Nap touches a group only if its meta has `nap_enabled = "true"` and it has an enabled `scaling` block with `min = 0`.
2. **Only the count**: Nap never stops, registers or edits a job; the only write it makes is a group's count. If Nomad refuses a scale, Nap retries and reports it, never forces it.
3. **When in doubt, stay awake**: Nap puts a group to sleep only after seeing it idle for the whole period. An error, invalid meta or a restart never puts a group to sleep; after a restart every timer starts over.
4. **Ready means healthy**: Nap lets a request through only when the group's allocations are healthy plus the margin the proxy needs, never when the service is merely registered.

## Repo layout

```
cmd/nap/               entrypoint
internal/version/      the build's version (release tag via -ldflags -X)
tests/integration/     tests against nomad agent -dev (build tag `integration`)
docs/                  documentation (index in docs/README.md)
```

The layout grows with the packages. Planned, each in its own issue: `config`,
`meta`, `nomadx`, `engine`, `web`, `metrics`, then `examples/` and `scripts/`.
