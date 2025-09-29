# Log Examples Inventory

## Overview
This directory contains 13 Go examples demonstrating usage patterns for the charmbracelet/log library. The examples total 247 lines of Go code across 13 files, with 12 examples including corresponding .tape files for terminal output recording. All examples use a baking/cooking theme for demonstration purposes.

## Examples Catalog

### Basic Logging
- **log/log.go** (7 lines): Minimal usage example showing basic Print function
- **cookie/cookie.go** (8 lines): Simple Debug and Info level logging demonstration
- **new/new.go** (12 lines): Creating custom logger instances with output redirection

### Error Handling
- **error/error.go** (12 lines): Error logging with structured key-value pairs
- **oven/oven.go** (13 lines): Helper function usage and caller reporting

### Configuration and Options
- **options/options.go** (19 lines): Logger configuration including prefix, time format, timestamps, and caller reporting
- **batch2/batch2.go** (15 lines): Logger context creation with persistent key-value pairs
- **chocolate-chips/chocolate-chips.go** (18 lines): Context logger creation and configuration chaining

### Styling and Formatting
- **styles/styles.go** (25 lines): Custom styling with lipgloss integration, error level customization, and key-value styling
- **format/format.go** (15 lines): Loop-based logging for progress tracking scenarios
- **temperature/temperature.go** (15 lines): Incremental value logging with formatting

### Integration
- **slog/main.go** (26 lines): Standard library slog integration with timestamp configuration and UTC time functions
- **app/app.go** (62 lines): Comprehensive example with custom types, helper functions, multiple log levels, and conditional logging

### Recording Files
12 .tape files (7-10 lines each) for generating terminal output recordings using VHS or similar tools.

## Implementation Patterns

### Logger Configuration
- Custom output destinations (os.Stderr redirection)
- Time format customization (time.Kitchen, time.RFC3339)
- Log level management (DebugLevel, InfoLevel, etc.)
- Caller reporting toggle
- Timestamp reporting control

### Structured Logging
- Key-value pair logging pattern throughout examples
- Context logger creation with With() method
- Persistent context values across multiple log calls
- Custom type String() method implementation for logging

### Styling Integration
- Lipgloss styling library integration
- Custom level styling (background, foreground, padding)
- Key and value styling customization
- Error highlighting patterns

### Helper Functions
- log.Helper() usage for accurate caller reporting
- Function-scoped logging contexts
- Error condition logging patterns

## Practical Applications

### NextJS Scaffolder with Auth/ShadCN
- **Configuration logging**: Use options.go pattern for setup process logging with timestamps
- **Error reporting**: Adapt error.go pattern for authentication failures and setup errors
- **Progress tracking**: Apply format.go and temperature.go patterns for scaffolding progress
- **Styled output**: Implement styles.go pattern for user-friendly CLI feedback

### Inter-Agent Communication Systems
- **Structured messaging**: Use batch2.go and chocolate-chips.go patterns for agent context logging
- **Error propagation**: Apply error.go pattern for communication failures
- **Status reporting**: Adapt app.go comprehensive logging for agent state management
- **Integration logging**: Use slog.go pattern for system log integration

### GitHub Search Tools
- **Query logging**: Apply options.go pattern with timestamps for search operations
- **Result formatting**: Use format.go pattern for search result iteration
- **Error handling**: Implement error.go pattern for API failures
- **Progress indication**: Adapt temperature.go pattern for search progress

### TUI for App Distribution/Package Management
- **Operation logging**: Use app.go comprehensive pattern for package operations
- **Status updates**: Apply format.go pattern for installation progress
- **Error reporting**: Implement styles.go pattern for user-visible error highlighting
- **Configuration tracking**: Use options.go pattern for package manager settings

The examples provide foundational patterns for structured logging, error handling, progress reporting, and user interface feedback suitable for terminal-based applications and development tools.