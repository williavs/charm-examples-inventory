# Glamour Examples2 - Contextual Reference

## Quick Reference

| Need | Example File |
|------|-------------|
| Basic markdown-to-terminal rendering | `examples2/helloworld/main.go` |
| Configured renderer with word wrap | `examples2/custom_renderer/main.go` |
| Pipeline-based stdin/stdout processing | `examples2/stdin-stdout/main.go` |
| Environment-based style detection | `examples2/stdin-stdout-custom-styles/main.go` |
| Embedded content with color downsampling | `examples2/artichokes/main.go` |

## Examples by Capability

### Basic Rendering
**Use helloworld when you need:**
- Simple string-to-terminal markdown rendering
- Direct `glamour.Render()` function usage
- Hardcoded markdown content output
- Fixed style theme application

**File**: `examples2/helloworld/main.go`
**Key patterns**: Direct function call with string input and style parameter

### Renderer Configuration
**Use custom_renderer when you need:**
- Custom renderer instance with specific settings
- Word wrap at defined character limits
- Reusable renderer configuration
- Multiple render operations with same settings

**File**: `examples2/custom_renderer/main.go`
**Key patterns**: `glamour.NewTermRenderer()` with option functions, controlled text wrapping

### Pipeline Processing
**Use stdin-stdout when you need:**
- Unix pipeline integration (cat, echo, redirection)
- Stream-based markdown processing
- Auto-detection of terminal color capabilities
- Byte-based input handling

**File**: `examples2/stdin-stdout/main.go`
**Key patterns**: `io.ReadAll(os.Stdin)`, `RenderBytes()`, stderr error handling

**Use stdin-stdout-custom-styles when you need:**
- Environment variable-based style configuration
- Pipeline processing with explicit environment detection
- Alternative to auto-style detection

**File**: `examples2/stdin-stdout-custom-styles/main.go`
**Key patterns**: `glamour.WithEnvironmentConfig()` instead of `WithAutoStyle()`

### Embedded Content and Color Handling
**Use artichokes when you need:**
- Embedded file resources using `//go:embed`
- Color profile detection and downsampling
- Adaptive rendering based on terminal capabilities
- Manual color profile writer usage

**File**: `examples2/artichokes/main.go`
**Key patterns**: `embed.FS`, `colorprofile.NewWriter()`, terminal capability detection

## Implementation Patterns

### Rendering Methods
```go
// Direct rendering
glamour.Render(string, "style")

// Configured renderer
r, _ := glamour.NewTermRenderer(options...)
r.Render(string)
r.RenderBytes([]byte)
```

### Configuration Options
- **Style detection**: `WithAutoStyle()` (auto-detect), `WithEnvironmentConfig()` (env vars), `WithStandardStyle("dark")` (fixed)
- **Word wrapping**: `WithWordWrap(40)`, `WithWordWrap(80)`
- **Color profiles**: Manual detection via `colorprofile.NewWriter()`

### Input Handling
- Hardcoded strings in source code
- Standard input streams (`os.Stdin`)
- Embedded files (`//go:embed`, `embed.FS`)
- Byte buffers (`bytes.Buffer`)

### Error Management
- Stderr output for processing failures
- Early exit with `os.Exit(1)` on critical errors
- Deferred file closures
- Separate error handling for read, render, write stages

### Output Patterns
- Direct stdout (`fmt.Print`, `fmt.Fprintf`)
- Color-aware output (`colorprofile.NewWriter`)
- Pipeline-compatible stdout writing

## Code Structure

All examples follow shared patterns:
- Single `main.go` per example
- Shared `go.mod` at examples2 root
- Dependencies: `github.com/charmbracelet/glamour`, `github.com/charmbracelet/colorprofile`
- Line count: 20-57 lines per example
- Total implementation: 191 lines across 5 examples

## Usage Contexts

### Stream Processing
Use stdin-stdout examples when you need to integrate with existing command-line workflows, process files through pipes, or build automation that transforms markdown content.

### Embedded Documentation
Use artichokes pattern when you need to bundle markdown content within compiled binaries, distribute help files without external dependencies, or package documentation.

### Configuration Control
Use custom_renderer when you need explicit control over rendering parameters, reusable renderer instances, or consistent styling across multiple operations.

### Rapid Prototyping
Use helloworld when you need quick markdown rendering tests, simple demonstrations, or minimal implementation examples.

### Terminal Adaptation
Use color profile patterns when you need consistent output across terminal types, graceful degradation on limited color support, or explicit color capability handling.
