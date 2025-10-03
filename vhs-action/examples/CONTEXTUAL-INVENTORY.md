# VHS-Action Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|-------------|
| Automated GIF generation on file changes | `examples/auto-commit.yml` |
| Visual demos posted to pull requests | `examples/comment-pr.yml` |
| Terminal recordings committed to repository | `examples/auto-commit.yml` |
| Externally-hosted GIF links in PR comments | `examples/comment-pr.yml` |

## Examples by Capability

### Automated Recording and Commit
**Use auto-commit when you need:**
- Automated GIF generation triggered by specific file changes
- Generated terminal recordings committed back to the repository
- Path-based workflow triggers that prevent unnecessary runs
- Direct repository integration without external hosting

**File**: `examples/auto-commit.yml`
**Key patterns**:
- Path-based trigger configuration (`paths: - vhs.tape`)
- Glob pattern file matching for commits (`file_pattern: '*.gif'`)
- Sequential action chaining (checkout → generate → commit)
- Custom commit metadata (user name, email, author)

### Pull Request Visual Feedback
**Use comment-pr when you need:**
- Visual terminal demos in pull request discussions
- External image hosting integration (Imgur)
- Automated PR comment posting with generated content
- Path-triggered workflows on pull request events

**File**: `examples/comment-pr.yml`
**Key patterns**:
- Pull request event triggering with path filters
- Multi-service integration (VHS → Imgur → GitHub Comments)
- Environment variable interpolation for dynamic content
- Output chaining between workflow steps (`steps.imgur_step.outputs.imgur_urls`)
- Secret management for external services (`IMGUR_CLIENT_ID`)

## Implementation Patterns

### Path-Based Workflow Triggering
Both examples use conditional execution based on file path changes:
```yaml
on:
  [event]:
    paths:
      - vhs.tape
```
This prevents workflows from running on unrelated repository changes.

### Action Composition
Sequential step execution pattern:
1. Checkout repository state
2. Execute VHS recording
3. Integrate with external service (commit or upload + comment)

### Secret Management Requirements
- `GITHUB_TOKEN`: Required for both repository commits and PR comment posting
- `IMGUR_CLIENT_ID`: Required for external image hosting (comment-pr only)

### File Pattern Targeting
- Glob patterns (`*.gif`): Matches all generated GIF files
- Specific paths (`./vhs.gif`): Targets single known output file

### Output Reference Syntax
Access previous step outputs using:
```yaml
${{ steps.[step_id].outputs.[output_name] }}
```

### Dynamic Content Formatting
Use `format()` function with environment variables for templated messages:
```yaml
message: ${{ format(env.MESSAGE, env.IMG_URL) }}
```

## Technical Characteristics

### auto-commit.yml
- **Trigger**: `push` events with `vhs.tape` path changes
- **Dependencies**: 2 external actions (checkout, vhs-action, git-auto-commit-action)
- **Authentication**: GITHUB_TOKEN environment variable
- **Output location**: Repository root (committed files)
- **File targeting**: Glob pattern matching

### comment-pr.yml
- **Trigger**: `pull_request` events with `vhs.tape` path changes
- **Dependencies**: 3 external actions (checkout, vhs-action, imgur upload, pr-comment)
- **Authentication**: GITHUB_TOKEN + IMGUR_CLIENT_ID secrets
- **Output location**: Imgur (external hosting) + PR comment thread
- **File targeting**: Specific file path reference

## Use Cases by Pattern Type

### Documentation Automation
Path-triggered recording generation captures terminal interactions when source files change.

### Visual Change Verification
Pull request workflows provide visual confirmation of behavior modifications.

### Artifact Generation
Automated creation of visual assets from tape definitions without manual execution.

### Integration Testing
Terminal recordings serve as visual regression test artifacts.
