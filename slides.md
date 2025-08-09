---
marp: true
theme: gaia
backgroundColor: #f9f6f2
color: #2c3e3e
---

# Go Full-Stack: Server-Side Rendering vs. SPAs

_A journey through the art of maintaining clean, efficient, and secure Go modules_

<span style="
  font-size:1em;
  color:#5d6d7e;
  display: block; /* Make it a block element to control its width */
  width: 100%; /* Make it span the full width of the slide's content area */
  position: relative; /* **THIS IS CRITICAL for absolute positioning of its child** */
  padding-right: 60px; /* IMPORTANT: Create space for the image to avoid text overlap */
  box-sizing: border-box; /* Ensures padding is included in the width */
">
  by Emily Achieng
  DevOps & Software Engineer
</span>


---

# Agenda

1. **Understanding the Guardians**  
   Meet go.mod and go.sum - the unsung heroes

2. **Spotting the Symptoms**  
   Recognizing when your modules need attention

3. **Dependency Detox: The Cleanup**  
   Practical techniques for dependency detox
---

4. **The Dependency Diet Challenge**  
   Trim your Go modules for a healthier project

5. **Security Patrols**  
   Guarding against vulnerabilities

6. **Future-Proofing**  
   Establishing sustainable practices

7. **Automation**  
   Integrating tooling and workflows for module management

---

# Understanding the Guardians

**go.mod**  
<span style="color:#2980b9;">The project manager who keeps track of what dependencies you need and which versions work best</span>

**go.sum**  
<span style="color:#27ae60;">The security guard who verifies nothing has been tampered with by storing cryptographic hashes</span>

_Together, they ensure consistency, security, and reliability across all environments._

---

# The Unsung Heroes: go.mod & go.sum 

<img src="https://media.giphy.com/media/5GoVLqeAOo6PK/giphy.gif" alt="Superhero flying meme" style="height:220px; display:block; margin-left:auto; margin-right:auto;" />


*Your project’s superhero for consistency and security.*

*These files aren't just boring configs - they're the foundation of reproducible builds and your first line of defense against supply chain attacks.*

---

# The Messy Apartment Syndrome

```go
// go.mod (cluttered)
require (
    github.com/pkg/errors v0.9.1
    github.com/sirupsen/logrus v1.8.1
    github.com/stretchr/testify v1.7.0
    github.com/old/dependency v0.1.0 // unused
    github.com/another/unused v2.0.0 // unused
)
```

<span style="font-size:1.05em; color:#34495e;">
A cluttered <code>go.mod</code> is like a messy apartment, unused dependencies cause maintenance headaches!
</span>

---

# Spotting the Symptoms  

**Slower Builds**  

_When your coffee break extends to a lunch break waiting for builds_

```go
// Build takes forever...
go build
// ...still building...
```
---

# Spotting the Symptoms  (cont.)

**Ballooning Binary Size**  

_Your executable suddenly needs its own storage plan_
```bash
# Check binary size
ls -lh myapp
# -rwxr-xr-x  1 user  staff   120M Aug  5 12:34 myapp
```

---

# Spotting the Symptoms  (cont.)

**Mysterious Errors**  

_"It worked yesterday!", "It worked on my machine!" becomes your team's new catchphrase_

```go
panic: cannot find package "github.com/old/dependency"
```

---

**Dependency Conflicts**  

_When packages argue like siblings about which version should be used_

```go
go: conflicting versions for module github.com/pkg/errors
```

---

# Spotting the Symptoms  (cont.)
### Recognizing when your modules need attention

**Dependency Bloat**  
_Like a cluttered garage, unused dependencies take up space and make it hard to find what you need._

```go
// go.mod (bloated)
require (
    github.com/pkg/errors v0.9.1
    github.com/unused/dependency v1.0.0 // unused
)
```

---

# Spotting the Symptoms (cont.)
**Security Vulnerabilities**  

_Outdated dependencies are like leaving your windows unlocked - they create an easy entry point for attackers._

```go
// go.mod (outdated)
require (
    github.com/old/vulnerable v0.1.0 // known vulnerability
)
```

---

**Build Inconsistencies**  
_Version conflicts._

```go
go: module github.com/pkg/errors@v0.9.1 found (v0.9.1), but does not contain package "github.com/pkg/errors/extra"
```

---

# Dependency Detox: The Cleanup

**Identify Unused Packages**

`go mod tidy` — Your first line of defense

```bash
go mod tidy
# Removes unused dependencies from go.mod and go.sum
```

_Think of it as spring cleaning for your codebase, sometimes painful but always worth it!_

---

# Dependency Detox: The Cleanup (cont.)

**Audit Indirect Dependencies**

You can use `go mod graph` to map your dependency tree

```bash
go mod graph
# Shows all direct and indirect dependencies in your project
```

---

# Dependency Detox: The Cleanup (cont.)

**Resolve Version Conflicts**

Use `go mod why` to understand why a module is needed

```bash
go mod why github.com/pkg/errors
# Explains which packages import this dependency
```

---

# Dependency Detox: The Cleanup (cont.)

**List All Dependencies**

Use `go list -m all` to see everything your module depends on.

```bash
go list -m all
# Shows all direct and indirect dependencies for your project
```

---

# Dependency Detox: The Cleanup (cont.)

**Pinning Versions**

Use `go mod edit -require` to pin a dependency to a specific version.

```bash
go mod edit -require=example.com/pkg@v1.2.3
# Ensures your project uses exactly v1.2.3 of example.com/pkg
```
---

# Dependency Detox: The Cleanup (cont.)

**Replacing Dependencies**

Use `go mod edit -replace` to swap out dependencies for local or forked versions.

```bash
go mod edit -replace=example.com/pkg=../local/pkg
# Use a local copy

go mod edit -replace=example.com/pkg=example.com/fork@v1.0.0
# Use a forked version
```
---
# The Dependency Diet Challenge

**Standard Library First**

Ask yourself:  
<span style="font-size:1.15em; color:#2980b9;">"Do I really need that package?"</span>  before adding new dependencies.

```go
// Instead of adding a third-party package for string manipulation:
import "strings" // Standard library does the job!

// Example usage:
s := "hello world"
upper := strings.ToUpper(s) // "HELLO WORLD"
```

_The Go standard library covers most needs—avoid unnecessary bloat!_


---

# The Dependency Diet Challenge (cont.)

**Package Evaluation**

Before adding a new dependency, check:

- Maintenance Status  
  Is the package actively maintained? Recent commits? Open issues?

- API Stability  
  Are there frequent breaking changes? Is the API well-documented?

--- 

# The Dependency Diet Challenge (cont.)

**Package Evaluation**

- Community Support  
  Are there contributors, stars, and helpful discussions?
```go
// Good sign: recent commits, clear docs, active issues
// Bad sign: last update years ago, lots of unresolved bugs
```

_Choose dependencies that are stable, well-supported, and actively maintained!_


---

# The Dependency Diet Challenge (cont.)

**Regular Pruning**

Schedule regular dependency reviews as part of your development cycle.

Schedule a regular reminder to run `go mod tidy` and review your `go.mod` file for unused or outdated dependencies.

_Keep your project lean and healthy by pruning dependencies regularly!_


---

# Security Patrols

**The Threat Landscape**

Your dependencies are doors to your application, each one needs proper locks.

---

# Security Patrols (cont.)

**Supply Chain Attacks**  
When trusted packages get compromised

**Known Vulnerabilities**  
Security bugs waiting to be exploited


---

# Security Patrols (cont.)

**Vulnerability Scanning**

Use `govulncheck` to identify known vulnerabilities in your dependencies.

```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...
# Scans your code and dependencies for known vulnerabilities
```

---

# Security Patrols (cont.)

**Vendor When Necessary**

`go mod vendor` creates a local copy of dependencies when you need absolute control.

```bash
go mod vendor
# Copies all dependencies into the vendor directory
```

---

# Security Patrols (cont.)

**Checksum Database**

Let Go's checksum database (`sum.golang.org`) protect you from supply chain attacks.

---

# Future-Proofing

<span style="font-size:1.15em; color:#16a085;"><b>Establishing sustainable practices</b></span>

Build habits that keep your Go modules healthy and your project maintainable for the long run.

---

# Sustainable Module Practices

**Schedule Regular Audits**

Set calendar reminders for dependency review days—think of it as routine maintenance for your codebase.

---

# Sustainable Module Practices (cont.)

**Automate Checks**

Integrate module checks into your CI/CD pipeline so every build keeps your dependencies in check—automation means no forgotten reviews!

---

# Sustainable Module Practices (cont.)

**Document Decisions**

Keep a dependency decision log for future reference.  
This helps your team understand why packages were added or removed, making onboarding and troubleshooting easier.

---

# Sustainable Module Practices (cont.)

**Test Update Impact**

Always run comprehensive tests before accepting dependency updates.  
Testing ensures new versions don’t break your app or introduce bugs.

---


# Sustainable Module Practices (cont.)

<span style="font-size:1.1em; color:#34495e;">
Creating habits is the key to long-term module health.
</span>

---

# Automation: Your Module's Best Friend

**Automation Intensity**

From manual effort to full process maturity—let automation handle your module hygiene!

---

# Automation Workflow

**Early Stage:**  
Manual `go mod tidy` and dependency review

**CI Pipeline:**  
Automatically run `go mod tidy` and fail the build if changes are detected

```yaml
# .github/workflows/ci.yml (snippet)
- name: Tidy modules
  run: go mod tidy
- name: Check for changes
  run: git diff --exit-code go.mod go.sum
```

---

# Automation Workflow (cont.)

**Dependency Updates:**  
Schedule automated updates with tools like Renovate or Dependabot

**Security Scanning:**  
Run `govulncheck` regularly in CI

```yaml
- name: Security scan
  run: govulncheck ./...
```

---

# Automation Workflow (cont.)

**Documentation:**  
Auto-generate dependency lists for transparency

```bash
go list -m all > DEPENDENCIES.md
```

---
# Automation Workflow (cont.)

<span style="font-size:1.1em; color:#34495e;">
Remember: The best module hygiene is the kind you don't have to think about.
</span>


---
# Takeaways

- **Go modules are the backbone of reliable builds**  
  Use `go.mod` and `go.sum` to ensure consistency and security.

- **Spot issues early**  
  Watch for slow builds, bloated binaries, and mysterious errors.

- **Keep dependencies lean and healthy**  
  Regularly prune unused packages and prefer the standard library.

---

- **Automate**  
  Integrate checks, updates, and security scans into your CI/CD pipeline.

- **Document and test**  
  Log dependency decisions and always test before updating.

- **Sustainable practices matter**  
  Build habits for long-term module health and project maintainability.

---

## Thank You!

Questions?  
Feel free to connect:  
<a href="https://www.linkedin.com/in/emmilliarchi/">
  <img src="data:image/svg+xml;base64,PHN2ZyByb2xlPSJpbWciIHZpZXdCb3g9IjAgMCAyNCAyNCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48dGl0bGU+TGlua2VkSW4gY29ub3IgaWNvbjwvdGl0bGU+PHBhdGggZD0iTTIwLjQ0NyAxOC43NThWMjQuMDhoLTQuNTA3di03LjQ4MWMwLTEuNzgyLS42MzktMi45OTItMi4yNDQtMi45OTItMS4yMjcgMC0xLjk1Mi44NzUtMi4yNzkgMS43MjUtLjExNS4zMDUtMC4wNy42NC0wLjA3MS45NzV2Ny43NjNoLTQuNTA3VjkuMDI1aDQuNTA3djEuOTJjLjc0MS0xLjIxMSAxLjY5Ni0yLjE0MiAzLjcwOC0yLjE0MiAyLjY4OSAwIDQuNzA3IDEuODU1IDQuNzA3IDUuODYyek01LjAyNiAwQzIuMjExIDAgMCAyLjEyNyAwIDQuNzI3UzIuMTEgNC43MjYgNS4wMjYgNC43MjYgNC43MjYtMi4xMjUgNS4wMjYtNC43MjZTNy45NDIgMCA1LjAyNiAwWiIvPjwvc3ZnPg==" alt="LinkedIn Icon" width="30" height="30"> linkedin.com/in/emmilliarchi/
</a>
<br>
<a href="https://github.com/EmAchieng">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/github/github-original.svg" alt="GitHub Icon" width="30" height="30"> github.com/EmAchieng
</a>


---