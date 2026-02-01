# Configuration Manual

This guide covers how to configure and customize the CI/CD workflows for this repository.

## Go Version Management

*This repository only setup testing for one GO_VERSION to reduce complexity.*

This repository reads go version as a variable from github settings to provide some flexibility.

### Setting Up Go Version Variable

1. Go to **Repository Settings** → **Secrets and variables** → **Actions** → **Variables** tab
2. Create new variable: `GO_VERSION` with value `1.25.6` (or your desired version)
3. All workflows reference this via `${{ vars.GO_VERSION }}`

### Updating Go Version

To update the Go version across all workflows:
1. Change the `GO_VERSION` variable in repository settings
2. All workflows will automatically use the new version on next run

## Code Quality Configuration

The [`.golangci.yml`](../.golangci.yml) file configures which linters run and their behavior:

### Customizing Linters

Modify `.golangci.yml` to customize which checks run - no workflow changes needed!

**Configuration Reference:**
- [golangci-lint Configuration Guide](https://golangci-lint.run/usage/configuration/)
- [Available Linters](https://golangci-lint.run/usage/linters/)
- [Example configurations](https://github.com/golangci/golangci-lint/blob/master/.golangci.reference.yml)

### Example: Adding New Linter

```yaml
linters:
  enable:
    - goimports    # Handles formatting and import organization
    - govet        # Basic static analysis for common bugs
```

## Workflow Customization

### Adding New Workflow Triggers

To trigger workflows on additional events, modify the `on:` section:

```yaml
on:
  push:
    branches: [ main, develop ]
    tags: [ 'v*' ]      # Trigger on version tags like v1.0.0
  pull_request:
    branches: [ main ]
```

## Making Workflows Non-Blocking

Sometimes you would like to have certain jobs not blocking your smooth pushing flow, you might want change the "code burns" error into just "code smells" warning!

Add `continue-on-error: true` to any step in your workflow:

```yaml
- name: Install and run golangci-lint
  continue-on-error: true    # ← Makes this step non-blocking
  run: |
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.1
    echo "$(go env GOPATH)/bin" >> $GITHUB_PATH
    find . -name "go.mod" -execdir golangci-lint run ./... \;
```

**Result:** Step shows ⚠️ WARNING instead of ❌ FAILED, workflow continues successfully.

💡 However, sometimes you might have tools that can check for error but always exit with 0. You will want to manually throw exit 1 status to trigger failure!


## Local Testing Setup

### Prerequisites

1. Install Go (version matching `GO_VERSION` variable)
2. Install golangci-lint: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
3. Install govulncheck: `go install golang.org/x/vuln/cmd/govulncheck@latest`

### Running Checks Locally

It is always good to run tests locally first before you submit anything!

**Run tests:**
```bash
cd "Problem Directory"
go test -v ./...
go test -race -v ./...
```

**Run code quality checks:**
```bash
golangci-lint run ./...
```

**Run security scan:**
```bash
govulncheck ./...
```

**Build check:**
```bash
go build ./...
```

## Troubleshooting

### Common Issues

**Go version mismatch:**
- Check `GO_VERSION` variable in repository settings
- Ensure local Go version matches for consistent behavior

**Linter failures:**
- Run `golangci-lint run ./...` locally to see exact issues
- Check `.golangci.yml` configuration
- Consider adding exceptions for specific cases

**Module not found:**
- Ensure `go.mod` files are properly configured
- Run `go mod tidy` to clean up dependencies