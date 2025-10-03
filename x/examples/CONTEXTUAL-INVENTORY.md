# X Examples - Contextual Reference

## Quick Reference

| Need | Example | File |
|------|---------|------|
| Terminal cell buffer manipulation | cellbuf | `examples/cellbuf/main.go` |
| Color palette generation | charmtone | `examples/charmtone/main.go` |
| PTY/fake TTY wrapper | faketty | `examples/faketty/main.go` |
| Image to Sixel graphics | img2term | `examples/img2term/main.go` |
| Lip Gloss layout patterns | layout | `examples/layout/main.go` |
| Image to ASCII conversion | mosaic | `examples/mosaic/main.go` |
| ANSI sequence parsing (high-level) | parserlog | `examples/parserlog/main.go` |
| ANSI sequence decoding (low-level) | parserlog2 | `examples/parserlog2/main.go` |
| Text wrapping with ANSI preservation | pen | `examples/pen/main.go` |
| Toner text processing | toner | `examples/toner/main.go` |

## Examples by Capability

### Terminal Buffer Management

**Use cellbuf when you need:**
- Direct terminal screen buffer control
- Raw terminal mode setup/cleanup
- Alternative screen buffer handling
- Mouse and keyboard event processing
- Real-time rendering with cell-level precision
- Window resize detection and response

**File**: `examples/cellbuf/main.go`
**Key patterns**: Event-driven input loop, screen writer with cursor positioning, cell-based rendering, hyperlink support in terminal output

---

### Color System Generation

**Use charmtone when you need:**
- CharmTone palette color generation
- CSS custom property output
- SCSS variable generation
- Vim colorscheme file creation
- Terminal color guide rendering with gradients
- Primary/secondary/tertiary color categorization

**Files**:
- `examples/charmtone/main.go` (CLI tool)
- `examples/charmtone/guide.go` (terminal rendering)

**Key patterns**: Cobra command structure with Fang execution, color blending algorithms, terminal background detection, formatted output generation

---

### PTY and Terminal Emulation

**Use faketty when you need:**
- Controlled terminal environment creation
- PTY (pseudo-terminal) management
- Specific terminal dimension enforcement
- Command execution with fake TTY
- File descriptor duplication for stdin/stdout redirection

**File**: `examples/faketty/main.go` (Unix-only, Windows excluded)
**Key patterns**: PTY winsize control, dual PTY setup for stdout/stderr, subprocess execution with controlled dimensions

---

### Graphics and Image Processing

**Use img2term when you need:**
- Sixel graphics protocol encoding
- JPEG/PNG to terminal conversion
- Terminal-displayable image output
- Image decoder integration (JPEG, PNG)

**File**: `examples/img2term/main.go`
**Key patterns**: Sixel encoder usage, image decoding from filesystem, ANSI graphics sequence generation

**Use mosaic when you need:**
- ASCII art from images
- Character-based image rendering
- JPEG image processing for terminal
- Customizable output dimensions

**File**: `examples/mosaic/main.go`
**Key patterns**: Mosaic renderer with width/height configuration, Lip Gloss layout integration, image loading utilities

---

### Layout and Styling

**Use layout when you need:**
- Comprehensive Lip Gloss styling reference
- Light/dark theme adaptation
- Column-based layouts
- Border and tab styling patterns
- Gradient text rendering
- Dialog box positioning
- Status bar composition
- Interactive UI with draggable elements

**File**: `examples/layout/main.go`
**Key patterns**: Background color detection, responsive width handling, cell buffer integration, mouse-driven element positioning, color blending for gradients

---

### ANSI Protocol Analysis

**Use parserlog when you need:**
- High-level ANSI sequence parsing
- Handler-based event processing
- CSI/ESC/DCS/OSC/APC/SOS/PM sequence categorization
- Print vs Execute differentiation
- Stream-based ANSI parsing

**File**: `examples/parserlog/main.go`
**Key patterns**: Parser handler registration, command extraction, parameter formatting, sequence type identification

**Use parserlog2 when you need:**
- Low-level ANSI sequence decoding
- State machine parsing
- Parameter extraction with subparameter support
- Intermediate/prefix/final command component access
- Byte-by-byte sequence analysis

**File**: `examples/parserlog2/main.go`
**Key patterns**: DecodeSequence state management, command component inspection, parameter iteration with HasMore tracking

---

### Text Processing

**Use pen when you need:**
- Automatic text wrapping
- ANSI sequence preservation during wrapping
- PenWriter usage patterns
- Terminal-aware text formatting

**File**: `examples/pen/main.go`
**Key patterns**: PenWriter initialization, ANSI.Wrap integration, stdin to stdout processing

**Use toner when you need:**
- Experimental toner package text processing
- Basic stdin/stdout pipeline
- Toner Writer usage

**File**: `examples/toner/main.go`
**Key patterns**: Toner Writer wrapper, stream processing

---

## Implementation Patterns

### Terminal State Management
- Raw mode setup with restoration cleanup pattern
- Alternative screen buffer toggling
- Terminal size detection via `term.GetSize`
- Cross-platform compatibility checks (Windows exclusions)
- SIGWINCH signal handling for resize events

### Input Handling
- Input driver creation with terminal type detection
- Mouse mode activation (ButtonEvent, SgrExt)
- Event loop processing with ReadEvents
- Keyboard/mouse/resize event discrimination
- Deferred cleanup for terminal state restoration

### Rendering Pipelines
- Screen initialization with ScreenOptions (profile, alt screen, relative cursor)
- Cell buffer filling and selective content placement
- Screen writer for positioned text output
- Render/Flush separation for buffer management
- Logging to file for debugging without terminal interference

### Color and Style
- Light/dark background detection via `lipgloss.HasDarkBackground`
- LightDarkFunc creation for adaptive theming
- Color blending using colorful library (Luv color space)
- Gradient generation across multiple color stops
- Style composition with borders, padding, alignment

### CLI Architecture
- Cobra command structure with subcommands
- Fang enhanced CLI execution wrapper
- Flag-based configuration (IntVar, Parse)
- Context-aware command execution
- Formatted output generation (CSS, SCSS, Vim)

### ANSI Protocol Handling
- Parser initialization with handler registration
- State machine for sequence decoding
- Parameter extraction and iteration
- Command component inspection (prefix, intermediate, final)
- Sequence type detection (HasCsiPrefix, HasOscPrefix, etc.)

### Image Processing
- Image decoder registration (import side effects: `_ "image/jpeg"`)
- Sixel encoder integration
- Mosaic renderer configuration (width, height)
- Lip Gloss layout integration for image placement

### Resource Management
- Deferred cleanup with error suppression (`//nolint:errcheck`)
- File opening/closing patterns
- PTY resource handling
- Buffer pooling for sequence processing

---

## Technology Dependencies

**Core Libraries:**
- `github.com/charmbracelet/x/ansi` - ANSI sequence handling
- `github.com/charmbracelet/x/cellbuf` - Terminal cell buffer
- `github.com/charmbracelet/x/input` - Input event processing
- `github.com/charmbracelet/x/term` - Terminal control
- `github.com/charmbracelet/lipgloss/v2` - Styling and layout
- `github.com/charmbracelet/colorprofile` - Color capability detection
- `github.com/charmbracelet/x/exp/charmtone` - CharmTone color system
- `github.com/charmbracelet/x/exp/toner` - Experimental text processing
- `github.com/charmbracelet/x/mosaic` - ASCII art generation

**CLI and Utilities:**
- `github.com/spf13/cobra` - Command structure
- `github.com/charmbracelet/fang` - Enhanced CLI execution
- `github.com/creack/pty` - PTY management
- `github.com/lucasb-eyer/go-colorful` - Color manipulation
- `github.com/rivo/uniseg` - Unicode grapheme segmentation

**Platform Support:**
- Primary: Unix-like systems (Linux, macOS)
- Limited: Windows (no PTY support in faketty)
- Go Version: 1.24+ (uses range over integer patterns)

---

## File Organization

```
examples/
├── cellbuf/
│   ├── main.go              # Primary implementation
│   ├── winsize_other.go     # Unix resize handling
│   └── winsize_windows.go   # Windows resize handling
├── charmtone/
│   ├── main.go              # CLI commands
│   └── guide.go             # Terminal color guide
├── faketty/
│   └── main.go              # PTY wrapper (Unix only)
├── img2term/
│   └── main.go              # Sixel graphics
├── layout/
│   └── main.go              # Lip Gloss showcase
├── mosaic/
│   └── main.go              # ASCII art converter
├── parserlog/
│   └── main.go              # High-level ANSI parser
├── parserlog2/
│   └── main.go              # Low-level ANSI decoder
├── pen/
│   └── main.go              # Text wrapping
├── toner/
│   └── main.go              # Toner processing
├── go.mod                   # Module dependencies
├── go.sum                   # Dependency checksums
├── JetBrainsMono-Regular.ttf
├── INVENTORY.md             # Prescriptive inventory (existing)
└── CONTEXTUAL-INVENTORY.md  # Generalized reference (this file)
```

---

## Usage Patterns by Context

### Building Interactive TUIs
- Start with `layout` for styling patterns
- Use `cellbuf` for event-driven interaction
- Reference `input` handling from cellbuf/layout examples
- Apply color themes from `charmtone` patterns

### Processing Terminal Output
- Use `parserlog` for sequence identification
- Use `parserlog2` for detailed parameter extraction
- Apply `pen` patterns for text wrapping needs
- Reference `toner` for experimental processing

### Image Integration
- Use `img2term` for Sixel-capable terminals
- Use `mosaic` for ASCII art fallback
- Check color profile support via layout patterns

### CLI Tool Development
- Reference `charmtone` for Cobra + Fang structure
- Apply flag parsing from `faketty`
- Use formatted output patterns from charmtone subcommands

### Testing and Debugging
- Use `faketty` for controlled terminal environments
- Reference logging patterns from cellbuf/layout (file output)
- Apply ANSI debugging from parserlog examples

---

## Common Integration Points

**Cellbuf + Lip Gloss**: Combine cell buffer control with Lip Gloss styling (see layout example)

**ANSI Parser + Pen**: Parse sequences then wrap with preservation (complementary tools)

**CharmTone + Layout**: Color system integration with adaptive theming

**Image Tools + Cellbuf**: Position images in cell buffer layouts

**Input + Screen**: Event processing with screen rendering coordination

---

## Build and Execution

All examples are standalone Go programs. Build and run:

```bash
cd examples/<example-name>
go run main.go [args]
```

Some examples require specific inputs:
- `img2term`: Requires image file path argument
- `mosaic`: Requires `pekinas.jpg` in directory
- `parserlog`, `parserlog2`, `pen`, `toner`: Process stdin

Example invocations:
```bash
# Interactive demos
go run examples/cellbuf/main.go
go run examples/layout/main.go

# CLI tools
go run examples/charmtone/main.go          # Terminal guide
go run examples/charmtone/main.go css      # CSS output
go run examples/charmtone/main.go scss     # SCSS output

# Image processing
go run examples/img2term/main.go image.png
go run examples/mosaic/main.go

# Stream processing
echo "test" | go run examples/pen/main.go
cat file.txt | go run examples/toner/main.go

# ANSI debugging
printf "\033[31mRed\033[0m" | go run examples/parserlog/main.go
```

---

## Notes

- Examples demonstrate terminal UI patterns without prescribing specific applications
- Code is reference material for understanding X package capabilities
- Platform-specific code uses build tags (`//go:build !windows`)
- Error handling uses `//nolint` directives for example brevity
- Logging to files prevents terminal output interference during debugging
- Mouse support requires explicit mode activation in terminal
