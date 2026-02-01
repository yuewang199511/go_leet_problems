# CI/CD Concepts

This document explains key CI/CD concepts, particularly focusing on GitHub Actions compared to other systems like GitLab CI.

*Note: I only used GitLab before so this note will mention some reference to GitLab but still trying to offer the concept of how CI/CD is setup in GitHub Actions.*

Please also check [GitHub's Go tutorial](https://docs.github.com/en/actions/tutorials/build-and-test-code/go) where it has a good starting example for Go.

## GitHub Actions Basics

### Core Concepts

**Workflows vs Jobs vs Steps:**
- **Workflow**: A complete CI/CD process defined in `.github/workflows/*.yml`
- **Job**: A set of steps that run on the same runner machine
- **Step**: Individual commands or actions within a job

**Key Components:**
- **`uses`**: Runs pre-built actions instead of custom shell commands, good to check in [GitHub Marketplace](https://github.com/marketplace?query=checkout&type=actions)
- **`runs-on`**: Specifies what machine/VM to run the job on (like GitHub's free EC2)
- **`actions/checkout@v4`**: Downloads repository source code to the runner like running `git clone && cd repo && git checkout`
- **`.github/workflows/*.yml`**: GitHub's config location (vs GitLab's single `.gitlab-ci.yml`)

### Comparison with GitLab CI (check only if you know gitlab)

| GitHub Actions | GitLab CI | Purpose |
|----------------|-----------|---------|
| `uses: actions/checkout@v4` | `git clone` (automatic) | Get source code |
| `runs-on: ubuntu-latest` | `image: ubuntu:latest` | Specify runner |
| `.github/workflows/` | `.gitlab-ci.yml` | Config location |
| Multiple workflow files | Single config file | Organization |

## Caching (Speed Optimization)

Caching dramatically improves workflow performance by reusing dependencies between runs.

### How Caching Works

**Cache Lifecycle:**
1. **Restore**: Happens during `actions/cache@v4` step
2. **Use**: Cached files are available for the job
3. **Save**: Happens automatically at job end

### Cache Configuration

```yaml
- uses: actions/cache@v4
  with:
    path: ~/go/pkg/mod           # What to cache
    key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}  # Unique ID
    restore-keys: |              # Fallback keys
      ${{ runner.os }}-go-
```

**Key Components:**
- **`path`**: What directory/files to cache (e.g., `~/go/pkg/mod` for Go modules)
- **`key`**: Unique identifier for the cache (includes OS, Go version, dependency hash)
- **`restore-keys`**: Fallback cache keys to try if exact match not found

### Cache Strategy

**Exact Match**: Uses cached dependencies when `go.sum` hasn't changed
**Partial Match**: Uses fallback when some dependencies change
**Cache Miss**: Downloads all dependencies fresh (slowest)

**Note**: Go Actions has its own caching mechanism, this repo uses custom cache just to demonstrate one caching approach.

## Exit Codes and Failure Handling

### Understanding Exit Codes

**Success**: Exit code `0`
- Step shows green checkmark ✓
- Workflow continues

**Failure**: Exit code `1` (or any non-zero)
- Step shows red X ✗
- Workflow stops (unless configured otherwise)
- Triggers failure notifications

### Tool Behavior Examples

| Tool | Success (Exit 0) | Failure (Exit 1) |
|------|-----------------|-------------------|
| `go test` | All tests pass | Any test fails |
| `golangci-lint` | No issues found | Linting issues detected |
| `govulncheck` | No vulnerabilities | Vulnerabilities found |
| `go build` | Compilation successful | Compilation errors |

## Best Practices

### Workflow Organization

1. **Separate Concerns**: Different workflows for different purposes
2. **Clear Naming**: Use descriptive job and step names

### Performance Optimization

1. **Use Caching**: Cache dependencies and build artifacts
2. **Minimal Runners**: Use appropriate runner sizes (don't over-provision), they have all kinds of cost (money,  and save some carbon output if you care!)

### ⚠️ Security Considerations

1. **Secrets Management**: Use GitHub Secrets for sensitive data
2. **Dependency Scanning**: Regular vulnerability checks upon the packages you use, the vulnerability reports are maintained by trustworthy officials. For Go, this uses:
   - **Go Vulnerability Database**: https://vuln.go.dev/ - Official database maintained by Go security team
   - **Documentation fpr `govulncheck`**: https://go.dev/security/vuln/ - Go's official vulnerability management
3. **Supply Chain Security**: Pin action versions (`@v4` not `@main`) and all dependencies, use specific version if you know that is more stable and meets the requirement rather than `@main` or `latest`!
