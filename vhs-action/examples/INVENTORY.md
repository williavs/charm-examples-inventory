# VHS-Action Examples Inventory

## Overview
This examples directory contains 2 GitHub Actions workflow files that demonstrate integration patterns for the VHS terminal recorder with automated Git operations and pull request workflows.

## Examples Catalog

### auto-commit.yml
- **File Count**: 1 YAML workflow file
- **Language**: GitHub Actions YAML
- **Functionality**: Triggers on push events when `vhs.tape` files change, executes VHS recording, commits generated GIF files back to repository
- **Dependencies**: charmbracelet/vhs-action@v1, stefanzweifel/git-auto-commit-action@v4
- **Trigger Pattern**: Push to main branch with changes to vhs.tape files
- **Output**: Automatically commits generated *.gif files to repository

### comment-pr.yml
- **File Count**: 1 YAML workflow file
- **Language**: GitHub Actions YAML
- **Functionality**: Triggers on pull requests with vhs.tape changes, generates GIF, uploads to Imgur, posts GIF link as PR comment
- **Dependencies**: charmbracelet/vhs-action@v1, devicons/public-upload-to-imgur@v2.2.2, github-actions-up-and-running/pr-comment@v1.0.1
- **Trigger Pattern**: Pull request events with vhs.tape file modifications
- **Output**: Imgur-hosted GIF link posted as pull request comment

## Implementation Patterns

### Conditional Workflow Triggers
Both examples use path-based triggering that activates only when specific files (vhs.tape) are modified, preventing unnecessary workflow runs.

### Multi-Step Action Chains
Sequential action execution pattern: checkout → VHS generation → external service integration (Git commit or image hosting + commenting).

### Secret Management
Both workflows require GitHub secrets (GITHUB_TOKEN, IMGUR_CLIENT_ID) for external service authentication.

### File Pattern Matching
Auto-commit example uses glob pattern `*.gif` to target generated files, while comment example targets specific `./vhs.gif` output.

## Practical Applications

### NextJS Scaffolder with Auth/ShadCN
- Could adapt auto-commit pattern to regenerate demo GIFs when scaffolding templates change
- Pull request workflow could show visual previews of UI component changes
- VHS could record authentication flows or component interactions for documentation

### Inter-Agent Communication Systems
- Auto-commit pattern could document agent interaction flows with terminal recordings
- PR comments could provide visual verification of agent communication protocols
- VHS recordings could serve as integration test artifacts

### GitHub Search Tools
- Could record search interface usage patterns and commit as documentation
- PR workflow could validate search functionality changes with recorded demos
- Generated GIFs could serve as visual API documentation

### TUI for App Distribution/Package Management
- Auto-commit could regenerate installation/usage demos when TUI changes
- PR comments could show before/after TUI behavior comparisons
- VHS recordings could document package installation flows and error handling

Both workflows provide templates for automated visual documentation generation that could be adapted for any terminal-based application development workflow.