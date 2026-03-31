# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2026-03-31

### Added
- Initial release with `ToArrays` and `ToMaps` functions
- `ToArrays` - reads CSV file into `[][]string`
- `ToMaps` - reads CSV file into `[]map[string]string` with header support
- Header trimming and name replacement support in `ToMaps`
- Test suite with 100% coverage

[Unreleased]: https://github.com/dracory/csvutils/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/dracory/csvutils/releases/tag/v1.0.0
