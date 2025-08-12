# Dependency Management

[![Go Version](https://img.shields.io/badge/go%20version-go1.24-blue)](https://golang.org/doc/go1.24)
[![Last Commit](https://img.shields.io/github/last-commit/EmAchieng/go-module-guardian)](https://github.com/EmAchieng/go-module-guardian/commits)
[![Issues](https://img.shields.io/github/issues/EmAchieng/go-module-guardian)](https://github.com/EmAchieng/go-module-guardian/issues)
[![Pull Requests](https://img.shields.io/github/issues-pr/EmAchieng/go-module-guardian)](https://github.com/EmAchieng/go-module-guardian/pulls)
[![Code Style: gofmt](https://img.shields.io/badge/code%20style-gofmt-brightgreen.svg)](https://golang.org/doc/go1.22#gofmt)
[![License](https://img.shields.io/github/license/EmAchieng/go-module-guardian)](./LICENSE)
[![Slides with Marp](https://img.shields.io/badge/slides-marp-blue?logo=marp)](https://marp.app/)
[![Docker](https://img.shields.io/badge/docker-ready-blue?logo=docker)](https://www.docker.com/)
[![govulncheck](https://img.shields.io/badge/security-govulncheck-blue)](https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck)
[![CI](https://github.com/EmAchieng/go-module-guardian/actions/workflows/ci.yml/badge.svg)](https://github.com/EmAchieng/go-module-guardian/actions/workflows/ci.yml)


Welcome!  
This project is a practical demonstration for my talk at **GopherCon UK 2025**.  
It focuses on Go module dependency management, cleanup, security scanning, and CI automation in a minimal project.

_Slides for this talk ([slides.md](./slides.md)) were prepared using [Marp](https://marp.app/) for Markdown-based presentations._

---

##  Getting Started

**Requirements:**  
- Go 1.22+

**Clone and Run:**
```sh
git clone git@github.com:EmAchieng/go-module-guardian.git
cd module-guardian
go mod tidy
go run main.go
```
---

## Run locally

```bash
go run main.go
```

## Docker

```bash
docker build -t module-guardian .
docker run -p 8080:8080 module-guardian
```
Visit [http://localhost:8080/users](http://localhost:8080/users) to see the demo endpoint.

---

## Dependency Management Concepts

- **go.mod & go.sum:**  
  Track all dependencies and their versions.  
  Ensure reproducible builds and security.

- **Unused Dependency Example:**  
  `github.com/google/uuid` is present in `go.mod` but not used in code.  
  Run `go mod tidy` to remove it and keep your project lean.

- **Pinning, Replacing, and Updating:**  
  - Pin: `go mod edit -require=github.com/gin-gonic/gin@v1.10.0`
  - Replace: `go mod edit -replace=github.com/gin-gonic/gin=../local/gin`
  - Update: `go get -u ./...`

- **Audit & List:**  
  - List all dependencies: `go list -m all`
  - Audit dependency tree: `go mod graph`
  - Understand why a dependency is present: `go mod why <module>`

- **Security Scanning:**  
  Use `govulncheck` to scan for vulnerabilities:
  ```sh
  go install golang.org/x/vuln/cmd/govulncheck@latest
  govulncheck ./...
  ```

- **Vendor When Necessary:**  
  `go mod vendor` creates a local copy of dependencies for maximum control.

- **Automation:**  
  CI workflow runs tidy, tests, security scan, and generates a dependency list.

---

## Example Workflow

1. **Add an unused dependency to `go.mod`**
2. **Run `go mod tidy`**  
   See it removed from `go.mod` and `go.sum`
3. **Run `govulncheck`**  
   Scan for vulnerabilities in all dependencies
4. **Automate with CI**  
   See `.github/workflows/ci.yml` for a full automation pipeline

---

## Thoughts?

- **Go modules** (`go.mod`, `go.sum`) are essential for reproducible builds and security.
- **Regular pruning** (`go mod tidy`) keeps your project lean.
- **Automation** (CI/CD) ensures module health without manual effort.
- **Security tools** like `govulncheck` help guard against vulnerabilities.
- **Document and test** all dependency changes for maintainability.

---

## License

MIT License — see [LICENSE](LICENSE) for details.

---

##  Questions?

Feel free to connect or fork this repo for your own experiments!

---

_Slides for this demo: [slides.md](./slides.md)_