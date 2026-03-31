# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2026-03-31

### Changed
- Migrated from `github.com/gouniverse/csvutils` to `github.com/dracory/csvutils`
- Renamed all Go files to snake_case

## [0.1.0] - 2023-11-18

### Added
- Initial release with `ToArrays` and `ToMaps` functions
- `ToArrays` - reads CSV file into `[][]string`
- `ToMaps` - reads CSV file into `[]map[string]string` with header support
- Header trimming and name replacement support in `ToMaps`

[Unreleased]: https://github.com/dracory/csvutils/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/dracory/csvutils/releases/tag/v0.2.0
[0.1.0]: https://github.com/dracory/csvutils/releases/tag/v0.1.0
