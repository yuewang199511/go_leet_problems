# Workflow Architecture

This repository uses GitHub Actions for automated testing with a **modular workflow architecture**. Instead of putting everything in one large workflow file, we separate concerns into multiple focused workflows.

## Overview

### **🧪 [test.yml](../.github/workflows/test.yml)** - Testing
- Unit tests with `go test -v`
- Race condition detection with `-race` flag
- Coverage reporting

### **✨ [code-quality.yml](../.github/workflows/code-quality.yml)** - Code Standards  
- Comprehensive linting via `golangci-lint` (includes formatting, static analysis, style checks)
- Configured via [`.golangci.yml`](../.golangci.yml) with 8 essential linters
- Single unified tool for all code quality checks

### **🔒 [security.yml](../.github/workflows/security.yml)** - Security
- Vulnerability scanning (`govulncheck`)
- Dependency security analysis

### **🏗️ [build.yml](../.github/workflows/build.yml)** - Cross-compilation
- Linux x64, Windows x64, macOS x64/ARM64 builds
- Ensures code compiles across platforms

## Benefits of Separation

**🎯 Single Responsibility**: Each workflow has one clear purpose  
**⚡ Fast Feedback**: Workflows provide quick status on code changes  
**🔧 Easy Maintenance**: Modify one aspect without affecting others  
**📊 Clear GitHub UI**: Each appears as separate status check  
**👥 Team Scalability**: Different team members can keep and modify theier desired parts

## Workflow Triggers

All workflows trigger on:
- **Pushes** to `main`/`develop` branches  
- **Pull requests** targeting `main` branch

## Auto-discovery

Each workflow automatically finds and processes all directories containing `go.mod` files, making the setup scalable for multi-module repositories.

## Monitoring Workflow Performance

- **Actions tab**: View all workflow runs and their duration
- **Status checks**: See which workflows passed/failed on PRs
- **Workflow badges**: Add status badges to README for public visibility