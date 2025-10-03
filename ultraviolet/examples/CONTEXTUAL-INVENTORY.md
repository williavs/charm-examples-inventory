# Ultraviolet Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|-------------|
| Basic terminal setup and event handling | `examples/helloworld/main.go` |
| Screen mode switching (altscreen/inline) | `examples/altscreen/main.go` |
| Mouse input and interactive drawing | `examples/draw/main.go` |
| Image rendering with multiple protocols | `examples/image/main.go` |
| Complex layouts with styling | `examples/layout/main.go` |
| Panic recovery and cleanup | `examples/panic/main.go` |
| Dynamic line prepending | `examples/prependline/main.go` |
| Animated grayscale patterns | `examples/space/main.go` |
| Color bar rendering | `examples/tv/main.go` |

## Examples by Capability

### Terminal Initialization & Event Loop
**Use helloworld when you need:**
- Basic terminal setup with `uv.NewTerminal()` or `uv.DefaultTerminal()`
- Event loop with context-based cancellation
- Suspend/resume handling (Ctrl+Z)
- Styled text rendering in fixed areas
- Standard quit patterns (q, Ctrl+C)

**File**: `examples/helloworld/main.go`
**Key patterns**: Terminal lifecycle, event streaming, context management, viewport rendering

### Screen Mode Management
**Use altscreen when you need:**
- Runtime switching between alternate screen and inline mode
- Dynamic viewport resizing based on screen mode
- Cursor visibility control
- Frame height management

**File**: `examples/altscreen/main.go`
**Key patterns**: Screen buffer switching, cursor control, mode toggling

### Mouse Interaction
**Use draw when you need:**
- Mouse event handling (clicks, motion, drag)
- Interactive character placement
- Wide character handling and cell fitting
- Temporary overlay displays (help dialogs)
- Buffer cloning and restoration
- Focus event handling

**File**: `examples/draw/main.go`
**Key patterns**: Mouse input, character drawing, buffer manipulation, modal overlays

### Image Rendering
**Use image when you need:**
- Terminal image display with protocol detection
- Support for blocks, sixel, iTerm2, or Kitty graphics
- Device capability querying
- Image positioning and movement
- Pixel-to-cell coordinate conversion
- Base64 encoding for inline images

**File**: `examples/image/main.go` (includes `charm.jpg`)
**Key patterns**: Multi-protocol graphics, capability detection, image placement

### Complex Styling & Layout
**Use layout when you need:**
- Multi-column layouts
- Tabbed interfaces
- Lipgloss integration for advanced styling
- Color gradients and blending
- Background color detection (light/dark mode)
- Responsive width handling
- Dialog boxes and buttons
- List components with checkmarks
- Status bars

**File**: `examples/layout/main.go`
**Key patterns**: Grid layouts, styling systems, component composition, grapheme handling

### Error Handling
**Use panic when you need:**
- Graceful panic recovery
- Terminal restoration after crashes
- Deferred cleanup handlers
- Timer-based events
- Custom event types

**File**: `examples/panic/main.go`
**Key patterns**: Panic recovery, defer mechanisms, stack traces, custom events

### Dynamic Text Display
**Use prependline when you need:**
- Line prepending functionality
- Event-driven text addition
- Background color cycling
- Inline mode text streaming
- Frame height adjustment

**File**: `examples/prependline/main.go`
**Key patterns**: Text prepending, color cycling, inline display

### Animation & Visual Effects
**Use space when you need:**
- Continuous animation loops
- FPS calculation and display
- Mathematical pattern generation
- Grayscale gradients
- Double-height rendering with half blocks
- Per-frame display updates
- Custom tick events

**File**: `examples/space/main.go`
**Key patterns**: Animation loops, FPS tracking, procedural graphics, event-driven rendering

### Test Patterns & Calibration
**Use tv when you need:**
- Color bar test patterns
- Predefined color palettes
- Percentage-based layout divisions
- RGB color specification
- Area filling with solid colors
- Display calibration patterns

**File**: `examples/tv/main.go`
**Key patterns**: Color bars, geometric divisions, test patterns

## Implementation Patterns

### Terminal Lifecycle
All examples follow this pattern:
```
1. Create terminal: uv.DefaultTerminal() or uv.NewTerminal()
2. Start: t.Start()
3. Optional: t.EnterAltScreen()
4. Event loop with context
5. Shutdown: t.Shutdown(ctx)
```

### Event Handling
Standard event patterns across examples:
- Context-based cancellation for event loops
- Event channel streaming via `t.StreamEvents(ctx, evch)`
- WindowSizeEvent triggers resize and redraw
- KeyPressEvent for user input
- Custom events via `t.SendEvent(ctx, event)`

### Display Updates
Common display approach:
- Modify terminal buffer state
- Call `t.Display()` to render changes
- Use `t.Erase()` or `screen.Clear(t)` to reset
- Call display function after each event

### Cell Manipulation
- `t.SetCell(x, y, &uv.Cell{...})` for individual cells
- `t.FillArea(&cell, area)` for regions
- `screen.Fill(t, &cell)` for entire screen
- StyledString components for text rendering

### Context Management
- Always use `context.WithCancel()` for event loops
- Defer cancel() to ensure cleanup
- Pass context to Shutdown() for graceful termination
- Optional timeout context for shutdown operations

### Input Patterns
- KeyPressEvent matching: `ev.MatchStrings("q", "ctrl+c")`
- Mouse events: MouseClickEvent, MouseMotionEvent
- Window events: WindowSizeEvent, WindowPixelSizeEvent
- Device queries: PrimaryDeviceAttributesEvent, TerminalVersionEvent

### Error Handling
- Check errors from Start(), Display(), Shutdown()
- Use log.Fatalf() for critical failures in examples
- Defer cleanup operations (ExitAltScreen, DisableMouse, Restore)
- Panic recovery with defer + recover()

### Styling Integration
- Direct color specification via uv.Style{Fg, Bg}
- Lipgloss integration for advanced styling
- ANSI color support (BasicColor, IndexedColor)
- True color via color.RGBA
- Background detection for theme adaptation

### Resource Management
- Close event channels in goroutines
- Use defer for cleanup operations
- Timeout contexts for shutdown operations
- Debug logging to separate files

## File Structure Notes

- All examples are self-contained single Go files
- Shared dependencies via `go.mod` in examples directory
- Debug logging typically to `*_debug.log` or `*.log` files
- File sizes range from 1.5KB (panic) to 13.6KB (layout)
- Most examples use alternate screen buffer
- Standard imports: ultraviolet, ultraviolet/screen, x/ansi
