# Introduction

This repository contains notes and solutions for selected coding exercises, with examples of automated testing and CI/CD workflows.

For each problem, I focus on the thought process rather than providing a full, final solution. That includes analysis of time and space complexity, discussion of constraints, intuition, and invariants used to reason about correctness.

Solutions may not be optimal, but they serve as a starting point for further improvement and learning.

# CI/CD Setup

This repository uses GitHub Actions with a **modular workflow architecture** for automated testing, code quality checks, security scanning, and builds.

## Quick Overview

- **🧪 [test.yml](.github/workflows/test.yml)** - Unit tests and race detection
- **✨ [code-quality.yml](.github/workflows/code-quality.yml)** - Linting and formatting
- **🔒 [security.yml](.github/workflows/security.yml)** - Vulnerability scanning
- **🏗️ [build.yml](.github/workflows/build.yml)** - Cross-platform compilation

TODO: releasing workflow needed to be built

All workflows trigger on pushes to `main`/`develop` and pull requests to `main`.

## Documentation

📖 **Detailed documentation available in [`docs/`](docs/):**

**🚀 START HERE FOR CICD USAGE:** [**Configuration Manual**](docs/configuration-manual.md) - Setup, customization, and troubleshooting

**Additional docs:**
- [**Workflow Architecture**](docs/workflow-architecture.md) - How workflows are organized and why
- [**CI/CD Concepts**](docs/cicd-concepts.md) - GitHub Actions fundamentals and best practices

## Local Testing

For running tests and quality checks locally, see [**Running Checks Locally**](docs/configuration-manual.md#running-checks-locally) in the Configuration Manual.