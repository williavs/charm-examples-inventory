# Colorprofile Examples Inventory

## Overview
This directory contains 2 Go examples demonstrating color profile detection and ANSI color conversion functionality. The examples use the charmbracelet/colorprofile library to handle terminal color capabilities and automatic color degradation.

## Examples Catalog

### Profile Example (`profile/main.go`)
- **File count**: 1 Go source file (76 lines)
- **Language**: Go 1.23.1
- **Functionality**: Demonstrates color profile detection, color conversion between different terminal capabilities, and automatic ANSI color degradation
- **Key features**:
  - Detects terminal color capabilities (TrueColor, ANSI256, ANSI, ASCII, NoTTY)
  - Converts RGB colors to appropriate terminal color formats
  - Shows color degradation from true color to 4-bit ANSI to ASCII-only
  - Interactive demonstration with explanatory text output

### Writer Example (`writer/writer.go`)
- **File count**: 1 Go source file (24 lines)
- **Language**: Go 1.23.1
- **Functionality**: Provides a command-line utility that reads ANSI-colored input from stdin and outputs degraded colors based on terminal capabilities
- **Key features**:
  - Pipes stdin to stdout with automatic color conversion
  - Acts as a filter for ANSI color sequences
  - Handles color degradation transparently

## Implementation Patterns

### Color Profile Detection
- Uses environment variable inspection combined with stdout analysis
- Provides fallback mechanisms for unsupported terminals
- Implements automatic capability detection without user configuration

### Color Conversion Pipeline
- RGB to terminal-specific color format conversion
- Graceful degradation from high-color to low-color terminals
- ANSI escape sequence processing and modification

### Writer Interface Implementation
- Standard Go io.Writer pattern for color processing
- Environment-aware initialization
- Stream processing for real-time color conversion

## Practical Applications

### NextJS Scaffolder with Auth/Shadcn
- Color profile detection could ensure consistent UI color rendering across different development environments
- ANSI color degradation useful for build tool output in various CI/CD terminals
- Writer pattern applicable for streaming build logs with appropriate color support

### Inter-Agent Communication Systems
- Color-coded message severity levels that adapt to terminal capabilities
- Consistent color representation across different agent terminals
- Stream processing for real-time log colorization in agent communication

### GitHub Search Tools
- Terminal-appropriate syntax highlighting for search results
- Color-coded repository status indicators that work in any terminal
- Degraded color schemes for accessibility in limited terminal environments

### TUI for App Distribution/Package Management
- Adaptive color schemes based on terminal capabilities
- Progress indicators and status colors that work universally
- Color-coded package status (installed, available, deprecated) with fallbacks

## Dependencies
- `github.com/charmbracelet/colorprofile`: Core color profile detection and conversion
- `github.com/charmbracelet/x/ansi`: ANSI sequence handling
- `github.com/lucasb-eyer/go-colorful`: Color manipulation utilities
- `golang.org/x/sys`: System-level terminal interaction

## File Structure
```
examples/
├── go.mod (Go module configuration)
├── go.sum (Dependency checksums)
├── profile/
│   └── main.go (Interactive color profile demonstration)
└── writer/
    └── writer.go (CLI color conversion filter)
```