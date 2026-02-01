# Introduction

This repository contains notes and solutions for selected coding exercises, with examples of automated testing and CI/CD workflows.

For each problem, I focus on the thought process rather than providing a full, final solution. That includes analysis of time and space complexity, discussion of constraints, intuition, and invariants used to reason about correctness.

Solutions may not be optimal, but they serve as a starting point for further improvement and learning.

# CI/CD setup

This repository uses GitHub Actions for automated testing with a **modular workflow architecture**. Instead of putting everything in one large workflow file, we separate concerns into multiple focused workflows:

## Workflow Architecture

### **🧪 [test.yml](.github/workflows/test.yml)** - Testing
- Unit tests with `go test -v`
- Race condition detection with `-race` flag
- Coverage reporting

### **✨ [code-quality.yml](.github/workflows/code-quality.yml)** - Code Standards  
- Code formatting checks (`gofmt`)
- Static analysis (`go vet`)
- Comprehensive linting (`golangci-lint`)

### **🔒 [security.yml](.github/workflows/security.yml)** - Security
- Vulnerability scanning (`govulncheck`)
- Dependency security analysis

### **🏗️ [build.yml](.github/workflows/build.yml)** - Cross-compilation
- Linux x64, Windows x64, macOS x64/ARM64 builds
- Ensures code compiles across platforms

## Benefits of Separation

**🎯 Single Responsibility**: Each workflow has one clear purpose  
**⚡ Parallel Execution**: All workflows run simultaneously for faster feedback  
**🔧 Easy Maintenance**: Modify one aspect without affecting others  
**📊 Clear GitHub UI**: Each appears as separate status check  
**👥 Team Scalability**: Different team members can own different workflows  

## Triggers
All workflows trigger on:
- **Pushes** to `main`/`develop` branches  
- **Pull requests** targeting `main` branch

## Auto-discovery
Each workflow automatically finds and processes all directories containing `go.mod` files, making the setup scalable for multi-module repositories.

## Configuration

### Go Version Management
To maintain consistency across all workflows, set up a repository variable:

1. Go to **Repository Settings** → **Secrets and variables** → **Actions** → **Variables** tab
2. Create new variable: `GO_VERSION` with value `1.25.6` (or your desired version)
3. All workflows reference this via `${{ vars.GO_VERSION }}`

This ensures **single source of truth** for Go version - change once in repository settings, applies everywhere automatically.



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

## Best Practice: Workflow Separation

### ❌ **Anti-pattern: Monolithic Workflow**
```yaml
# DON'T: Everything in one job
jobs:
  everything:
    steps:
      - name: Test
      - name: Lint  
      - name: Security scan
      - name: Build
      - name: Deploy
```

### ✅ **Best Practice: Modular Workflows**
```yaml
# DO: Separate files for different concerns
.github/workflows/
├── test.yml          # Testing only
├── code-quality.yml  # Linting & formatting
├── security.yml      # Vulnerability scanning  
├── build.yml         # Cross-compilation
└── deploy.yml        # Deployment (if needed)
```

### **Why Separation Matters:**

1. **🚀 Faster Feedback**: Get quality issues immediately without waiting for all tests
2. **🔄 Independent Scaling**: Add more test scenarios without affecting code quality checks
3. **🎯 Focused Debugging**: When a workflow fails, you know exactly which aspect broke
4. **👥 Team Ownership**: Different team members can maintain different workflows
5. **📊 Better Metrics**: Track success rates for testing vs. code quality vs. security separately
6. **🔧 Selective Execution**: Disable specific checks (e.g., security scanning) without affecting core testing

This modular approach follows the **Single Responsibility Principle** in CI/CD design, making your development process more maintainable and scalable.

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