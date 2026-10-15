# Philosophy and invariants

Nap is scale-to-zero for a self-hosted Nomad cluster for personal use. It
grows with the need: a feature is built when the cluster needs it, never for a
scale it doesn't have yet. Each choice is the right one, not just the
simplest, and stays simple enough to grow.

Nap shares this philosophy with [Nops](https://github.com/music-gang/nops) but
no code. The two meet only in Nomad's state.

## Invariants

No change may break these rules. Each one has a test that tries to break it.

1. **Only groups that opt in.** Nap touches a group only if its meta has
   `nap_enabled = "true"` and it has an enabled `scaling` block with
   `min = 0`. *Why:* a cluster runs apps that must never sleep, and the
   jobspec is where a person says which ones can.

2. **Only the count.** Nap never stops, registers or edits a job: the only
   write it makes is a group's count. If Nomad refuses a scale, Nap retries
   and reports it, and never forces it. *Why:* `nomad job stop` removes the
   proxy route, so the next request would find nothing to wake the app, and a
   register would collide with whoever manages the spec.

3. **When in doubt, stay awake.** Nap puts a group to sleep only after seeing
   it idle for the whole period. An error, invalid meta, or a restart never
   puts a group to sleep, and after a restart every timer starts over. *Why:*
   a sleeping app that should run is an outage, an awake app that could sleep
   only costs resources.

4. **Ready means healthy.** Nap lets a request through only when the group's
   allocations are healthy, plus the margin the proxy needs to notice them.
   *Why:* Nomad registers a service before its health check passes, and a
   request let through then gets a 502.

## Fail loud

Nap says so rather than carry on with uncertain state, and its safe path is
to leave the app awake:

- It never ignores a Nomad error. Every error tells you what failed and for
  which target.
- Invalid meta means the group is never put to sleep, and an ERROR in the log.
- The state lives in memory. Losing it leaves every app awake, which is the
  safe situation, so Nap persists nothing.

## What Nap leaves to Nops

Nops decides what a job is. Nap decides only whether a group runs. They don't
call each other, so a person can use one without the other. How they behave
together is in [architecture](architecture.md#running-with-nops).
