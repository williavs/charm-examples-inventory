# Ultraviolet Examples Inventory

## Overview

This directory contains 9 Go examples demonstrating the Ultraviolet terminal UI library. The examples include 12 total files across 9 subdirectories, with a shared Go module configuration. Examples range from basic terminal setup to advanced features like image rendering and mouse interaction.

## Examples Catalog

### altscreen
- **Files**: 1 Go file (1,873 bytes)
- **Function**: Demonstrates alternate screen buffer switching with inline mode toggle
- **Key features**: Screen mode switching, cursor visibility control, viewport management
- **Input handling**: Space to toggle modes, Ctrl+C/q to exit
- **Application**: Foundation for applications requiring different display modes

### draw
- **Files**: 1 Go file (4,682 bytes)
- **Function**: Interactive drawing application with mouse and keyboard input
- **Key features**: Mouse event handling, character drawing, altscreen buffer, focus events
- **Input handling**: Mouse drawing, keyboard input, character rendering
- **Application**: Base for interactive drawing tools or text editors with mouse support

### helloworld
- **Files**: 1 Go file (2,354 bytes)
- **Function**: Basic terminal application displaying styled text in a fixed area
- **Key features**: Terminal initialization, event loop, styled string rendering, suspend/resume
- **Input handling**: q/Ctrl+C to quit, Ctrl+Z for suspend
- **Application**: Template for basic TUI applications with event handling

### image
- **Files**: 1 Go file (11,095 bytes) + 1 JPEG file (6,300 bytes)
- **Function**: Image rendering with multiple encoding formats
- **Key features**: Supports blocks, sixel, iTerm2, and Kitty image protocols
- **Input handling**: Command-line flags for encoding selection
- **Application**: Image viewers, media applications, terminal graphics

### layout
- **Files**: 1 Go file (13,636 bytes)
- **Function**: Complex layout system with tabs, columns, and styled components
- **Key features**: Lipgloss styling, responsive design, color theme detection, tabbed interface
- **Input handling**: Not specified in examined portion
- **Application**: Dashboard layouts, multi-panel applications, styled interfaces

### panic
- **Files**: 1 Go file (1,491 bytes)
- **Function**: Demonstrates panic recovery and graceful terminal restoration
- **Key features**: Panic handling, countdown timer, terminal cleanup
- **Input handling**: q/Ctrl+C to exit before panic
- **Application**: Error handling patterns, graceful shutdown mechanisms

### prependline
- **Files**: 1 Go file (3,881 bytes)
- **Function**: Line prepending functionality with color cycling
- **Key features**: Dynamic line addition, color backgrounds, frame height management
- **Input handling**: Event-driven line addition
- **Application**: Log viewers, streaming text applications, status displays

### space
- **Files**: 1 Go file (3,161 bytes)
- **Function**: Animated space-like visualization with randomized grayscale patterns
- **Key features**: Color gradients, animation loops, mathematical pattern generation
- **Input handling**: Standard quit commands
- **Application**: Screensavers, visual effects, background animations

### tv
- **Files**: 1 Go file (3,680 bytes)
- **Function**: TV test pattern generator with color bars
- **Key features**: Color bar rendering, predefined color palette, test pattern display
- **Input handling**: Standard terminal events
- **Application**: Display testing, color calibration, retro aesthetics

## Implementation Patterns

### Common Patterns
- **Terminal initialization**: All examples use `uv.DefaultTerminal()` or `uv.NewTerminal()`
- **Event handling**: Context-based cancellation with event channels
- **Error handling**: Consistent error checking with log.Fatalf()
- **Cleanup**: Deferred terminal restoration and shutdown procedures

### Input Management
- **Standard quit**: q, Ctrl+C patterns across examples
- **Context cancellation**: Proper goroutine management
- **Event streaming**: Channel-based event distribution

### Display Techniques
- **Alternate screen**: Most examples use altscreen buffer for full-screen applications
- **Styled rendering**: Integration with Lipgloss for advanced styling
- **Dynamic sizing**: Responsive to terminal size changes

### Architecture
- **Single-file applications**: Each example is self-contained
- **Module structure**: Shared dependencies via go.mod
- **Logging**: Debug logging to separate files

## Practical Applications

### NextJS Scaffolder with Auth/Shadcn
- **layout example**: Provides tabbed interface patterns for project configuration screens
- **helloworld**: Basic template for CLI tool interfaces
- **altscreen**: Mode switching for different scaffolding phases

### Inter-Agent Communication Systems
- **draw example**: Input handling patterns for interactive agent interfaces
- **prependline**: Real-time message display systems
- **panic**: Error recovery for agent communication failures

### GitHub Search Tools
- **layout**: Multi-column display for search results and repository information
- **image**: Display repository assets or user avatars
- **space**: Loading animations during search operations

### TUI for App Distribution/Package Management
- **layout**: Package listing interfaces with categorized tabs
- **draw**: Interactive package selection and marking
- **altscreen**: Switch between package lists and detailed views
- **tv**: System status displays and health monitoring
- **panic**: Graceful handling of package installation failures

## File Structure Summary
- **Total files**: 12 (9 Go files, 1 JPEG, 1 go.mod, 1 go.sum)
- **Languages**: Go (primary), module configuration
- **Dependencies**: Ultraviolet TUI library, Lipgloss styling, supporting utilities
- **Size range**: 1,491 - 13,636 bytes per Go file
- **Complexity**: Basic examples (helloworld, panic) to advanced (layout, image, draw)