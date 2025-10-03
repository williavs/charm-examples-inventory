# Colorprofile Examples - Contextual Reference

## Quick Reference

| Need | Example File | Lines |
|------|-------------|-------|
| Terminal color capability detection | `examples/profile/main.go` | 76 |
| ANSI color stream filtering | `examples/writer/writer.go` | 24 |
| RGB to terminal color conversion | `examples/profile/main.go` | 76 |
| Color degradation pipeline | `examples/profile/main.go` | 76 |
| stdin/stdout color processing | `examples/writer/writer.go` | 24 |

## Examples by Capability

### Profile Example
**Use profile example when you need:**
- Terminal color capability detection (TrueColor, ANSI256, ANSI, ASCII, NoTTY)
- RGB color conversion to terminal-specific formats
- Color degradation demonstration across different terminal capabilities
- Interactive color profile information display

**File**: `examples/profile/main.go`
**Key patterns**: Environment-based detection, color.Color conversion, writer-based degradation

### Writer Example
**Use writer example when you need:**
- ANSI color sequence filtering through stdin/stdout pipeline
- Automatic color degradation based on terminal environment
- Stream processing with color adaptation
- Command-line color conversion utility

**File**: `examples/writer/writer.go`
**Key patterns**: io.Writer interface, stream filtering, environment-aware initialization

## Implementation Patterns

### Color Profile Detection Pattern
```go
p := colorprofile.Detect(os.Stdout, os.Environ())
```
- Inspects environment variables and stdout properties
- Returns profile enum: TrueColor, ANSI256, ANSI, Ascii, NoTTY
- No user configuration required

### Color Conversion Pattern
```go
targetColor := profile.Convert(sourceColor)
```
- Converts color.Color to profile-appropriate type
- Handles RGB to ANSI256, ANSI, or ASCII conversion
- Returns typed color value (colorprofile.ANSI256Color, colorprofile.ANSIColor, etc.)

### Writer-Based Degradation Pattern
```go
w := colorprofile.NewWriter(os.Stdout, os.Environ())
w.Profile = colorprofile.ANSI  // Optional: override detected profile
fmt.Fprintln(w, ansiColoredText)
```
- Creates writer that processes ANSI sequences
- Automatically degrades colors based on profile
- Supports profile override for testing or forced degradation

### Stream Filtering Pattern
```go
w := colorprofile.NewWriter(os.Stdout, os.Environ())
io.Copy(w, os.Stdin)
```
- Pipes input through color conversion writer
- Processes ANSI escape sequences in real-time
- Transparent to upstream/downstream processes

## Color Profile Types

| Profile | Description | Color Depth |
|---------|-------------|-------------|
| TrueColor | 24-bit RGB | 16.7M colors |
| ANSI256 | 8-bit color | 256 colors |
| ANSI | 4-bit color | 16 colors |
| Ascii | No color | Text only |
| NoTTY | Non-terminal | Stripped output |

## Dependencies

- `github.com/charmbracelet/colorprofile` - Core color profile detection and conversion
- `github.com/charmbracelet/x/ansi` - ANSI sequence handling
- `github.com/lucasb-eyer/go-colorful` - Color manipulation utilities
- `golang.org/x/sys` - System-level terminal interaction

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
