# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.7] - 2024-01-17

### Updated

- The `servemux.Param` type to export the `Name` and `Value` fields so that
  they can be used for testing when the packages are installed.

## [0.0.6] - 2024-01-17

### Fixed

- Removed unnecessary logs.

## [0.0.5] - 2024-01-17

### Updated

- Added context defaults for the additional packages, such as `auth`'s
  whitelisted and `servemux`'s params.

## [0.0.4] - 2024-01-17

- Updated the `ContextMiddleware` and `ContextHandlerFunc` to use the context
  and not the value pointed to by the context. This is to follow best practices
  of using the context value directly.

## [0.0.3] - 2024-01-10

### Updated

- Updated the parsing of the route context to ensure the correct context is
  passed to the handler.

## [0.0.2] - 2024-01-10

### Updated

- The pattern match to allow for numbers in the URL path parameters.

## [0.0.1] - 2024-01-10

### Updated

- The pattern match to allow for `-` and `_` in the URL path parameters.

## [0.0.0] - 2023-11-29

### Added

- The initial implementation for the following packages.
- `auth` package.
- `openapi` package.
- `servemux` package.
