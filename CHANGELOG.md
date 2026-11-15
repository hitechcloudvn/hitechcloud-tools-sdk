# Changelog

> **English** | [Tiếng Việt](CHANGELOG.vi.md)

---

All notable changes to the HiTechCloud Tools SDK will be documented in this file.

## [1.1.0] - 2026-11-16

### Changed
- Synchronized SDK sources with the latest HiTechCloud API Tools definitions
- Stopped tracking build artifacts in the repository
- Fixed CI/CD workflows (release.yml, ci.yml), Java/PHP/C# client issues and npm scope for GitHub Packages

### Added
- Bilingual Vietnamese-English documentation across the whole repository
  (every doc has a `.vi.md` mirror).
- GitLab CI pipeline mirroring the GitHub Actions tests, with GitLab
  Releases published on `v*` tags.

## [1.0.0] - 2024-12-01

### Added
- Initial release with 7 language SDKs (Python, Node.js, PHP, Go, Ruby, Java, C#)
- 455 endpoints across 15 resource groups
- OpenAPI 3.0 specification
- Docker support
- CI/CD pipelines
- Comprehensive documentation
