# Data Anonymizer

![CI](https://github.com/Qyroxen/Data-Anonymizer/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Data-Anonymizer/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Data-Anonymizer?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Data-Anonymizer)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Data-Anonymizer)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Data-Anonymizer?style=social)](https://github.com/Qyroxen/Data-Anonymizer/stargazers)

## What is it?

Data Anonymizer is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Data-Anonymizer.git
cd Data-Anonymizer
go build -o dataanonymizer .

# Run
./dataanonymizer --help
```

## CLI Usage

```bash
# Basic usage
./dataanonymizer

# With flags
./dataanonymizer --verbose --output json

# Get help
./dataanonymizer --help
```

## Examples

```bash
# Example 1
./dataanonymizer example1

# Example 2
./dataanonymizer example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o dataanonymizer .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Data-Anonymizer/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Data-Anonymizer?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Data-Anonymizer/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Data-Anonymizer?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Data-Anonymizer/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Data-Anonymizer" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Data-Anonymizer/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Data-Anonymizer" alt="Pull Requests">
  </a>
</p>
