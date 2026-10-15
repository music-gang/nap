# Architecture

How Nap works inside. The rules every part keeps are in
[philosophy](philosophy.md).

## Components

| Package | Role |
|---|---|
| `cmd/nap` | The entry point. |
| `internal/version` | The build's version. |

## Running with Nops

Nap and [Nops](https://github.com/music-gang/nops) don't integrate. They
coordinate through Nomad's state only, and the fixes that make this work are
in Nops.
