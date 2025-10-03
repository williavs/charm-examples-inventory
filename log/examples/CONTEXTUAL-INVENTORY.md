# Log Examples - Contextual Reference

## Quick Reference

| Need | Example | File |
|------|---------|------|
| Basic output | log | `examples/log/log.go` |
| Multiple log levels | cookie | `examples/cookie/cookie.go` |
| Custom logger instance | new | `examples/new/new.go` |
| Error logging with context | error | `examples/error/error.go` |
| Helper functions and caller tracking | oven | `examples/oven/oven.go` |
| Logger configuration | options | `examples/options/options.go` |
| Persistent context values | batch2 | `examples/batch2/batch2.go` |
| Context logger creation | chocolate-chips | `examples/chocolate-chips/chocolate-chips.go` |
| Custom styling with lipgloss | styles | `examples/styles/styles.go` |
| Iterative logging patterns | format | `examples/format/format.go` |
| Incremental value logging | temperature | `examples/temperature/temperature.go` |
| Standard library slog integration | slog | `examples/slog/main.go` |
| Comprehensive usage patterns | app | `examples/app/app.go` |

## Examples by Capability

### Basic Logging
**Use log example when you need:**
- Minimal logging setup without configuration
- Single line output to default logger

**File**: `examples/log/log.go`
**Key patterns**: Direct Print() call without setup

---

**Use cookie example when you need:**
- Multiple log levels in simple form
- Debug and Info level differentiation

**File**: `examples/cookie/cookie.go`
**Key patterns**: Debug() and Info() level methods

---

**Use new example when you need:**
- Custom logger instances separate from global logger
- Output redirection to specific streams (os.Stderr)
- Structured key-value logging

**File**: `examples/new/new.go`
**Key patterns**: log.New() constructor, output stream specification, key-value pairs in Warn()

### Error Handling

**Use error example when you need:**
- Error logging with structured context
- Key-value pairs for error metadata
- Standard Go error type integration

**File**: `examples/error/error.go`
**Key patterns**: Error() method with "err" key, fmt.Errorf integration

---

**Use oven example when you need:**
- Helper function usage for accurate caller reporting
- Caller location tracking in logs
- Function-scoped logging with Helper() marking

**File**: `examples/oven/oven.go`
**Key patterns**: log.Helper() in functions, SetReportCaller(true)

### Configuration and Options

**Use options example when you need:**
- Logger configuration with prefix, timestamps, and caller info
- Custom time formatting
- Multiple configuration options set sequentially

**File**: `examples/options/options.go`
**Key patterns**: SetPrefix(), SetTimeFormat(), SetReportTimestamp(), SetReportCaller()

---

**Use batch2 example when you need:**
- Persistent key-value pairs across multiple log calls
- Logger context with pre-configured values
- Level configuration and output control

**File**: `examples/batch2/batch2.go`
**Key patterns**: With() method for context creation, SetLevel() for threshold control

---

**Use chocolate-chips example when you need:**
- Context logger creation from default logger
- Configuration chaining patterns
- Separate logger instances with different contexts

**File**: `examples/chocolate-chips/chocolate-chips.go`
**Key patterns**: Default().With() chaining, context inheritance from base logger

### Styling and Formatting

**Use styles example when you need:**
- Custom styling with lipgloss integration
- Level-specific visual customization
- Key and value styling for structured logs
- Background and foreground color control

**File**: `examples/styles/styles.go`
**Key patterns**: DefaultStyles() modification, lipgloss.NewStyle(), SetStyles(), per-level customization

---

**Use format example when you need:**
- Loop-based iterative logging
- Progress indication patterns
- Sequential log output with formatted strings

**File**: `examples/format/format.go`
**Key patterns**: Loop iteration with fmt.Sprintf(), time delays for visibility

---

**Use temperature example when you need:**
- Incremental value logging
- Formatted values in key-value pairs
- Progress tracking with changing numeric values

**File**: `examples/temperature/temperature.go`
**Key patterns**: Loop-based logging with formatted key-value pairs, fmt.Sprintf() for value formatting

### Integration

**Use slog example when you need:**
- Standard library slog handler integration
- UTC time functions and RFC3339 formatting
- Compatibility layer between charmbracelet/log and slog
- Timestamp configuration with custom time functions

**File**: `examples/slog/main.go`
**Key patterns**: NewWithOptions() with Options struct, slog.New() wrapping, TimeFunction configuration

---

**Use app example when you need:**
- Comprehensive logging patterns in single example
- Custom type String() method implementation for logging
- Helper functions with log.Helper()
- Conditional logging based on values
- Multiple log levels in workflow context
- Fatal() termination on critical errors

**File**: `examples/app/app.go`
**Key patterns**: Custom type formatting, conditional log levels, time.Sleep() for sequential operations, multiline structured values

## Implementation Patterns

### Logger Initialization
- **Global logger**: Use package-level functions (log.Print, log.Info, etc.)
- **Custom instances**: Create with log.New(writer) for isolated loggers
- **Context loggers**: Derive with With() method for persistent key-value pairs
- **Options-based**: Use NewWithOptions() for complex initial configuration

### Configuration Methods
- **SetPrefix()**: Add identifier prefix to all log messages
- **SetTimeFormat()**: Customize timestamp format (time.Kitchen, time.RFC3339, etc.)
- **SetReportTimestamp()**: Toggle timestamp display
- **SetReportCaller()**: Enable/disable file:line caller information
- **SetLevel()**: Set minimum log level threshold
- **SetStyles()**: Apply custom lipgloss styling

### Structured Logging Patterns
- **Key-value pairs**: Pass alternating key (string) and value (any) arguments
- **Error context**: Use "err" key for error values
- **Persistent context**: Use With() to create logger with fixed key-value pairs
- **Custom formatters**: Implement String() method on types for automatic formatting

### Helper Function Pattern
```go
func operation() {
    log.Helper()  // Mark this function as helper
    log.Info("message")  // Caller will be operation's caller, not this line
}
```

### Styling Pattern
```go
styles := log.DefaultStyles()
styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
    SetString("CUSTOM").
    Background(lipgloss.Color("204")).
    Foreground(lipgloss.Color("0"))
styles.Keys["keyname"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
logger.SetStyles(styles)
```

### Integration Pattern
```go
// charmbracelet/log as slog handler
handler := log.NewWithOptions(os.Stdout, log.Options{
    ReportTimestamp: true,
    TimeFunction:    log.NowUTC,
})
slogLogger := slog.New(handler)
```

## Output Stream Control
- **Default**: Logs to stderr by default
- **Custom stream**: Pass io.Writer to log.New()
- **Multiple loggers**: Create separate instances for different outputs
- **Stream redirection**: Use os.Stderr, os.Stdout, or custom writers

## Level Hierarchy
Levels from most to least verbose:
1. DebugLevel - Detailed debugging information
2. InfoLevel - General informational messages (default)
3. WarnLevel - Warning messages
4. ErrorLevel - Error messages
5. FatalLevel - Fatal errors (calls os.Exit(1))

## Time Formatting Options
- **time.Kitchen**: "3:04PM"
- **time.RFC3339**: "2006-01-02T15:04:05Z07:00"
- **Custom format**: Any valid Go time format string

## Common Use Cases

### Progress Tracking
Use format or temperature patterns with loop-based logging and formatted values.

### Error Reporting
Use error pattern with structured key-value pairs including "err" key.

### Configuration Logging
Use options pattern to track setup steps with timestamps and caller info.

### Context-Based Logging
Use batch2 or chocolate-chips patterns to create loggers with persistent context values.

### Visual Differentiation
Use styles pattern to apply custom colors and formatting for different log levels or keys.

### Standard Library Integration
Use slog pattern to bridge between charmbracelet/log and Go's standard slog package.

### Complex Workflows
Use app pattern as reference for multi-level logging, conditional output, and helper functions.
