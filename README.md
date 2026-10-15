# Nap

[![CI](https://github.com/music-gang/nap/actions/workflows/ci.yml/badge.svg)](https://github.com/music-gang/nap/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/music-gang/nap)](LICENSE)

Scale-to-zero for [HashiCorp Nomad](https://www.nomadproject.io/).

Nap puts the apps nobody uses to sleep and wakes them on the first request.
It is the `forwardAuth` target of your proxy: each request to an app passes
through Nap, which starts the app if it sleeps and lets the request through
once the app is healthy.

> **Status:** under construction, not usable yet. Built for a self-hosted
> cluster for personal use, next to [Nops](https://github.com/music-gang/nops),
> which shares its philosophy but no code.

## Nomad compatibility

Nap is tested against **Nomad 2.0.7**. Other 2.0.x releases should work.

## Documentation

See the [documentation index](docs/README.md). To contribute, see
[CONTRIBUTING](CONTRIBUTING.md).

## License

Licensed under the [Apache License 2.0](LICENSE). Copyright 2026 Music Gang.
