# Contributing

Thank you for your interest in contributing to the project! This document provides guidelines and instructions for contributing.

## Code of Conduct

Please read and follow our [Code of Conduct](CODE_OF_CONDUCT.md) in all your interactions with the project.

## Development Workflow

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/cool-new-feature`)
3. Run tests and verify your changes:

   ```bash
   ./scripts/test.sh
   ./scripts/test.sh --coverage
   golangci-lint run
   ```

4. Commit your changes (`git commit -m 'Add a cool new feature'`)
5. Push to the branch (`git push origin feature/cool-new-feature`)
6. Open a Pull Request

## Pull Request Process

1. Ensure your PR description clearly describes the problem and solution
2. Update the README.md if necessary
3. The PR must pass all CI checks
4. A maintainer will review your PR and merge it once approved

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: Version 1.24.6 or later
- **golangci-lint**: For code quality checks

## Development Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/Dobefu/DLiteScript.git
   cd DLiteScript
   ```

2. Install Go dependencies:

   ```bash
   go mod download
   ```

3. Verify your setup by running tests:

   ```bash
   ./scripts/test.sh
   ```

## Code Style

- Follow the [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- Use `gofmt` for Go code formatting
- Use `golangci-lint` for code style checks
- Write tests for new features and bug fixes

## Testing

- Run the test suite before submitting a PR:

  ```bash
  ./scripts/test.sh
  ```

- Ensure test coverage doesn't decrease (run `./scripts/test.sh --coverage` to check)
- Add tests for new features

## Security

If you discover a security vulnerability, please use GitHub's private vulnerability reporting feature instead of opening a public issue. This helps protect our users while we work on a fix.

## Questions?

Feel free to open an issue if you have any questions about contributing to the project.
