# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

## Changed

- Update go to 1.17
- Update github.com/giantswarm/microerror to v0.4.0
- Update  github.com/giantswarm/micrologger to v0.6.0
- Update  golang.org/x/sync to v0.0.0-20210220032951-036812b2e83c
- Update  k8s.io/api to v0.22.2
- Update  k8s.io/apimachinery to v0.22.2
- Update  k8s.io/client-go to v0.22.2
- Update giantswarm/architect orb to v4.14.2

## [3.1.1] - 2021-02-17

### Fixed

- Restore `AllCerts` list of `cert-operator`-managed certificates.

## [3.1.0] - 2020-09-28

### Added

- Added support for secrets in other namespaces than `default`.
- Added etcd client certificates for Prometheus.

## [3.0.0] - 2020-08-10

### Changed

- Update Kubernetes dependencies to v1.18.5.

## [2.0.0] 2020-05-07

### Changed

- Modify certs searching for HA Masters.
- Prepare project structure for new major release.



## [0.2.0] 2020-03-24

### Changed

- Switch from dep to Go modules.
- Use architect orb.



## [0.1.0] 2020-03-18

### Added

- First release.



[Unreleased]: https://github.com/giantswarm/certs/compare/v3.1.1...HEAD
[3.1.1]: https://github.com/giantswarm/certs/compare/v3.1.0...v3.1.1
[3.1.0]: https://github.com/giantswarm/certs/compare/v3.0.0...v3.1.0
[3.0.0]: https://github.com/giantswarm/certs/compare/v2.0.0...v3.0.0
[2.0.0]: https://github.com/giantswarm/certs/compare/v0.2.0...v2.0.0
[0.2.0]: https://github.com/giantswarm/certs/compare/v0.1.0...v0.2.0

[0.1.0]: https://github.com/giantswarm/certs/releases/tag/v0.1.0
