# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.0] - 2025-09-17

### Changed

- Push app to collections.

### Fixed

- Fix streaming support.

## [0.2.0] - 2025-09-15

### Changed

- Configure `gsoci.azurecr.io` as the default container image registry.
- Upgrade tempo-distributed chart from 1.21.0 to 1.38.2
- Enable memcached and configure its image.
- Support streaming of queries.
- Configure `tempo.traces` field in the chart.

## [0.1.2] - 2023-12-19

### Changed

- Upgrade app dependency to 1.7.3

## [0.1.1] - 2023-11-09

### Fixed

- Fix _helpers to pass ci.

## [0.1.0] - 2023-11-09

### Changed

- Initialize tempo app

[Unreleased]: https://github.com/giantswarm/tempo-app/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/giantswarm/tempo-app/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/tempo-app/compare/v0.1.2...v0.2.0
[0.1.2]: https://github.com/giantswarm/tempo-app/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/giantswarm/tempo-app/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/giantswarm/tempo-app/releases/tag/v0.1.0
