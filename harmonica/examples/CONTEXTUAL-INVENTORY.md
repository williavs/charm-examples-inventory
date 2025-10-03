# Harmonica Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|--------------|
| Projectile physics with gravity | `examples/particle/main.go` |
| Spring physics in terminal | `examples/spring/tui/main.go` |
| Interactive spring physics with graphics | `examples/spring/opengl/main.go` |

## Examples by Capability

### Projectile Physics Simulation
**Use particle when you need:**
- Gravity-based motion calculations in terminal space
- Frame-based physics updates at 60 FPS
- ASCII coordinate visualization of moving objects
- Automatic termination based on position thresholds

**File**: `examples/particle/main.go`
**Key patterns**: BubbleTea framework integration, harmonica.Projectile with TerminalGravity, position-based rendering with spaces/newlines

### Terminal Spring Physics
**Use spring/tui when you need:**
- Damped spring oscillation in text interfaces
- Styled terminal output with colored blocks
- Spring-based horizontal motion toward targets
- Automatic settling detection and termination

**File**: `examples/spring/tui/main.go`
**Key patterns**: harmonica.Spring with configurable frequency/damping, Lipgloss styling, position-to-target convergence

### Interactive Graphics Spring Physics
**Use spring/opengl when you need:**
- Real-time mouse-driven spring animations
- Adjustable physics parameters during runtime
- Hardware-accelerated circular sprite rendering
- Keyboard-controlled frequency and damping modification

**File**: `examples/spring/opengl/main.go`
**Key patterns**: Pixel graphics library, gg rendering context, user input parameter adjustment, state-based sprite behavior

## Implementation Patterns

### Physics Integration
- **Frame timing**: 60 FPS tick-based updates using `tea.Tick(time.Second/fps, ...)`
- **Spring mechanics**: `harmonica.NewSpring(fps, frequency, damping)` for oscillation
- **Projectile motion**: `harmonica.NewProjectile(fps, position, velocity, acceleration)` for ballistic paths
- **State updates**: Immutable model updates returning `(updatedModel, nextCommand)`

### Rendering Techniques
- **ASCII positioning**: Space/newline insertion for coordinate-based text placement
- **Styled blocks**: Lipgloss for colored rectangular sprites in terminal
- **OpenGL graphics**: Pixel library with gg context for circle rendering and text overlay
- **Font handling**: TrueType font loading for text rendering in graphics mode

### Input Handling
- **Keyboard exit**: Universal `tea.KeyMsg` → `tea.Quit` pattern
- **Mouse targeting**: Click position capture for spring target setting
- **Parameter adjustment**: Arrow keys with shift modifier for fine/coarse control
- **Real-time feedback**: Immediate visual response to input changes

### Architecture Patterns
- **BubbleTea Model-View-Update**: Init() → Update() → View() cycle for TUI examples
- **Pixel game loop**: Time-based delta updates with VSync rendering
- **State machines**: Enum-based sprite states (moving, stopped, hiding, gone)
- **Spring coefficient recalculation**: Dynamic physics updates on parameter changes

## Physics Configuration Examples

### Spring Parameters
```
Frequency: 7.0 (medium bounce)
Damping: 0.15 (light resistance)

Frequency: 10.0 (faster oscillation)
Damping: 0.2 (moderate settling)

Frequency: 32.0 (rapid response)
Damping: 1.0 (critical damping, no overshoot)
```

### Projectile Setup
```
Initial Position: (0, 0)
Initial Velocity: (5, 0) - horizontal motion
Acceleration: TerminalGravity - downward pull
Max Height: 100 - termination threshold
```

## Reusable Code Patterns

### Frame Animation Command
```go
func animate() tea.Cmd {
    return tea.Tick(time.Second/fps, func(t time.Time) tea.Msg {
        return frameMsg(t)
    })
}
```

### Spring Update Pattern
```go
// Update position and velocity toward target
newPos, newVel := spring.Update(currentPos, currentVel, targetPos)
```

### Settling Detection
```go
// Check if spring has settled near target
if math.Abs(current - target) < 0.01 {
    // Trigger completion action
}
```

### State-Based Rendering
```go
// Only render when not in terminal state
if sprite.State() != gone {
    // Perform rendering
}
```

## Technical Requirements

### Dependencies
- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/harmonica` - Physics library
- `github.com/charmbracelet/lipgloss` - Terminal styling
- `github.com/faiface/pixel` - OpenGL graphics
- `github.com/fogleman/gg` - 2D rendering context
- `golang.org/x/image/font` - Font handling

### External Assets
- `JetBrainsMono-Regular.ttf` - Required for OpenGL text rendering (18 TrueType tables)

### Performance Characteristics
- **Frame rate**: 60 FPS for all examples
- **TUI overhead**: Minimal, suitable for terminal applications
- **OpenGL performance**: Hardware-accelerated, suitable for real-time interaction
- **Memory usage**: Low for TUI, moderate for OpenGL sprite rendering
