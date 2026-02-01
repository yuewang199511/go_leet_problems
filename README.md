# Introduction

This repository contains notes and solutions for selected coding exercises, with examples of automated testing and CI/CD workflows.

For each problem, I focus on the thought process rather than providing a full, final solution. That includes analysis of time and space complexity, discussion of constraints, intuition, and invariants used to reason about correctness.

Solutions may not be optimal, but they serve as a starting point for further improvement and learning.

# CI/CD setup

This repository uses GitHub Actions for automated testing. The workflow:

- **Triggers**: Runs on pushes to `main`/`develop` branches and pull requests to `main`
- **Go versions**: Tests against Go 1.21, 1.22, and 1.23 for compatibility
- **Test steps**:
  1. Downloads dependencies for all Go modules in the repository
  2. Runs tests with `go test -v ./...`
  3. Runs tests with race detector enabled
  4. Verifies code builds successfully


Each problem directory with a `go.mod` file is automatically discovered and tested independently.



# Notes on CI/CD

I only used gitlab before so this note will mention some reference to gitlab but still trying to offer the concept of how CI/CD is setup in github acitons.

Please also check [text](https://docs.github.com/en/actions/tutorials/build-and-test-code/go) where it has a good starting example for go.

## Key Concepts

**GitHub Actions Basics:**
- **`uses`**: Runs pre-built actions instead of custom shell commands, good to check in [text](https://github.com/marketplace?query=checkout&type=actions)
- **`runs-on`**: Specifies what machine/VM to run the job on (like GitHub's free EC2)
- **`actions/checkout@v4`**: Downloads repository source code to the runner like running clone - cd - checkout
- **`.github/workflows/*.yml`**: GitHub's config location (vs GitLab's single `.gitlab-ci.yml`)

**Matrix Strategy (Parallel Jobs):**
- **`strategy` + `matrix`**: Creates parallel jobs with different configurations (like GitLab's parallel:matrix)
- **Cross-platform testing**: Simply use multiple OS runners (ubuntu/windows/macos)

**Caching (Speed Optimization):**
- **`actions/cache@v4`**: Speeds up workflows by saving/restoring files between CI/CD runs
- **`path`**: What directory/files to cache (~/go/pkg/mod for Go modules)
- **`key`**: Unique identifier for the cache (includes OS, Go version, dependency hash)
- **`restore-keys`**: Fallback cache keys to try if exact match not found
- **Cache timing**: Restore happens on the step of actions/cache, save happens automatically at job end. go actions has its own caching mechanism, this repo use the custom cache just to show one of the machanism for caching.
- **Cache fallback**: Uses partial matches when dependencies change incrementally, but still try to get the previous installed packages

**Go Commands:**
- **`go test -v`**: Runs tests with verbose output showing individual test names
- **`go build`**: Compiles Go source code into executables (compilation verification)
- **`find . -name "go.mod" -execdir`**: Finds all Go modules and runs commands in their directories

## Local testing

To run tests locally for a specific problem:

```bash
cd "Problem Directory"
go test -v ./...
```

To run tests with race detection:

```bash
go test -race -v ./...
```