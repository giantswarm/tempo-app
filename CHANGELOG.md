# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Fix HTTPRoute template.

## [0.10.1] - 2026-02-12

### Changed

- Change team annotation in `Chart.yaml` to OpenContainers format (`io.giantswarm.application.team`).

## [0.10.0] - 2026-02-03

### Added

- Add Crossplane AWS support for automated S3 bucket provisioning.

## [0.9.0] - 2026-01-28

### Changed

- Add Gateway API and Envoy Gateway resources

## [0.8.1] - 2025-12-12

### Changed

- Upgrade Tempo chart from 1.53.2 to 1.58.0

## [0.8.0] - 2025-11-07

### Changed

- Upgrade Tempo chart from 1.51.1 1.53.2

## [0.7.0] - 2025-10-27

### Changed

- Upgrade Tempo chart from 1.48.1 to 1.51.1
  - Upgrade Tempo from 2.8.2 to 2.9.0 
- Upgrade Tempo Vulture from 0.9.1 to 0.10.0

## [0.6.1] - 2025-10-15

### Fixed

- Fix `metrics-generator` kind to `StatefulSet` instead of `Statefulset`

## [0.6.0] - 2025-10-09

### Changed

- Enable Persistence for Tempo

## [0.5.0] - 2025-10-06

### Added

- Add ingress template for HTTP and gRPC traffic with configurable backends and auto-generated TLS configuration.

## [0.4.0] - 2025-09-29

### Added

- Add Tempo Vulture as a dependency.

### Changed

- Update upstream chart to 1.48.0
- Update app to 2.8.2

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

- Fix `_helpers` to pass ci.

## [0.1.0] - 2023-11-09

### Changed

- Initialize tempo app

[Unreleased]: https://github.com/giantswarm/tempo-app/compare/v0.10.1...HEAD
[0.10.1]: https://github.com/giantswarm/tempo-app/compare/v0.10.0...v0.10.1
[0.10.0]: https://github.com/giantswarm/tempo-app/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/giantswarm/tempo-app/compare/v0.8.1...v0.9.0
[0.8.1]: https://github.com/giantswarm/tempo-app/compare/v0.8.0...v0.8.1
[0.8.0]: https://github.com/giantswarm/tempo-app/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/giantswarm/tempo-app/compare/v0.6.1...v0.7.0
[0.6.1]: https://github.com/giantswarm/tempo-app/compare/v0.6.0...v0.6.1
[0.6.0]: https://github.com/giantswarm/tempo-app/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/tempo-app/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/giantswarm/tempo-app/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/tempo-app/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/tempo-app/compare/v0.1.2...v0.2.0
[0.1.2]: https://github.com/giantswarm/tempo-app/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/giantswarm/tempo-app/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/giantswarm/tempo-app/releases/tag/v0.1.0
