# Glamour Examples2 Inventory

## Overview
Contains 5 Go examples demonstrating markdown rendering capabilities using the Glamour library. Total of 191 lines of Go code across examples, with shared dependencies managed through a single go.mod file that includes Charm Bracelet ecosystem packages for terminal styling and color profile detection.

## Examples Catalog

### helloworld (20 lines)
Basic markdown rendering demonstration. Takes hardcoded markdown string and renders to terminal using dark theme. Shows fundamental Glamour usage pattern of `glamour.Render(input, style)` for simple markdown-to-terminal output.

### custom_renderer (22 lines)
Terminal renderer configuration example. Demonstrates custom renderer creation with word wrapping at 40 characters and dark style theme. Uses `glamour.NewTermRenderer()` with configuration options for controlled output formatting.

### stdin-stdout (46 lines)
Command-line markdown processor. Reads markdown from stdin, renders through Glamour with auto-style detection and 80-character word wrap, outputs to stdout. Includes error handling and can be used as unix pipeline component.

### stdin-stdout-custom-styles (46 lines)
Alternative stdin processor using environment-based configuration. Similar to stdin-stdout but uses `glamour.WithEnvironmentConfig()` instead of `glamour.WithAutoStyle()` for theme detection from terminal environment.

### artichokes (57 lines)
Color profile demonstration with embedded content. Uses Go embed directive to include artichokes.md file, demonstrates color profile detection and downsampling through colorprofile package. Shows terminal color capability detection and adaptive rendering.

## Implementation Patterns

### Rendering Approaches
- Simple string rendering: `glamour.Render(string, style)`
- Configured renderer: `glamour.NewTermRenderer(options...)`
- Byte-based processing: `renderer.RenderBytes([]byte)`

### Configuration Options
- Style themes: "dark", auto-detection, environment-based
- Word wrapping: Fixed width (40, 80 characters)
- Color handling: Auto-detection, manual profile specification

### Input Sources
- Hardcoded strings in source
- Standard input streams (stdin)
- Embedded files using `//go:embed`
- File system reads with `io.ReadAll()`

### Error Handling
- Stderr output for processing errors
- Exit codes for failure conditions
- Graceful degradation for unsupported features

## Practical Applications

### NextJS Scaffolder with Auth/ShadCN
- **stdin-stdout examples**: Can process README.md files and documentation during scaffolding
- **custom_renderer**: Format configuration summaries and setup instructions
- **color profile detection**: Ensure consistent output across different terminal environments during setup

### Inter-Agent Communication Systems
- **stdin-stdout pattern**: Process markdown messages between agents in pipeline fashion
- **embedded content approach**: Package help documentation and communication templates
- **error handling patterns**: Robust message processing with failure recovery

### GitHub Search Tools
- **markdown rendering**: Format search results, README previews, and issue content
- **color adaptation**: Ensure search results display properly across terminal types
- **word wrapping**: Format long descriptions and code snippets in search output

### TUI for App Distribution/Package Management
- **custom renderer configuration**: Format package descriptions and changelogs
- **embedded documentation**: Package help files and installation guides
- **pipeline processing**: Transform package metadata and documentation files
- **color profile support**: Consistent rendering across different terminal capabilities

### Implementation Notes
All examples use the Charm Bracelet ecosystem (lipgloss, colorprofile) for terminal styling. The modular approach allows selective adoption of features. The stdin-stdout pattern enables integration with existing command-line workflows and automation scripts.