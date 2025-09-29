# X Examples Inventory

## Overview
Collection of 10 Go examples demonstrating terminal UI components and utilities using the Charm ecosystem (Lip Gloss, cellbuf, ANSI handling). Contains 17 files across terminal manipulation, color theming, image processing, and ANSI sequence parsing.

## Examples Catalog

### cellbuf (3 files)
Interactive terminal screen buffer demonstration with mouse and keyboard input handling. Shows raw terminal mode, alternative screen buffer usage, and real-time rendering with resize support. Useful for building terminal applications requiring direct screen manipulation.

### charmtone (2 files)
Color palette utility for the CharmTone design system. Generates CSS, SCSS, and Vim colorscheme configurations. Includes terminal color guide rendering with gradient blending. Applicable for design system tooling and color palette management.

### faketty (1 file)
PTY (pseudo-terminal) wrapper for running commands with controlled terminal dimensions. Creates fake TTY environment for testing terminal applications with specific row/column configurations. Relevant for terminal emulation and command execution testing.

### img2term (1 file)
Image to terminal converter using Sixel graphics protocol. Reads JPEG/PNG images and outputs terminal-displayable graphics. Useful for terminal-based image viewers and media applications.

### layout (1 file)
Comprehensive Lip Gloss styling demonstration showing advanced layout features, color detection, and responsive terminal UI design. Contains patterns for column layouts, styling, and dark/light theme adaptation.

### mosaic (2 files)
Image-to-ASCII art converter that renders JPEG images as terminal output using character-based graphics. Demonstrates image processing for terminal display with customizable dimensions.

### parserlog (1 file)
ANSI escape sequence parser with detailed logging output. Processes stdin and categorizes ANSI sequences (CSI, ESC, DCS, OSC). Useful for debugging terminal output and understanding ANSI protocol implementation.

### parserlog2 (1 file)
Alternative ANSI sequence decoder using lower-level parsing approach. Provides detailed sequence analysis with parameter extraction and state tracking. Applicable for ANSI protocol analysis and terminal emulator development.

### pen (1 file)
Text wrapping utility using cellbuf PenWriter. Demonstrates automatic text wrapping with ANSI sequence preservation. Relevant for terminal text formatting and content display systems.

### toner (1 file)
Simple stdin-to-stdout processor using the experimental toner package. Minimal implementation for text processing pipelines with toner formatting capabilities.

## Implementation Patterns

### Terminal Management
- Raw terminal mode setup and cleanup
- Alternative screen buffer handling
- Window resize detection and response
- Cross-platform compatibility (Windows exclusions)

### Input/Output Processing
- Event-driven input handling (keyboard, mouse)
- ANSI sequence parsing and generation
- PTY creation and management
- Stream processing with proper error handling

### Styling and Rendering
- Light/dark theme detection and adaptation
- Color palette management and generation
- Layout composition with Lip Gloss
- Image processing for terminal display

### CLI Framework Integration
- Cobra command structure
- Fang enhanced CLI execution
- Flag-based configuration
- Subcommand organization

## Practical Applications

### NextJS Scaffolder with Auth/Shadcn
- **charmtone**: Color system integration for consistent theming across web and terminal interfaces
- **layout**: Terminal UI patterns adaptable to web component layout systems
- **cellbuf**: Interactive terminal menus for project configuration during scaffolding

### Inter-Agent Communication Systems
- **parserlog/parserlog2**: ANSI protocol understanding for terminal-based agent interfaces
- **faketty**: Controlled terminal environments for agent testing and simulation
- **pen**: Text formatting for agent message display and logging

### GitHub Search Tools
- **layout**: Terminal UI framework for search result display and navigation
- **cellbuf**: Interactive search interfaces with real-time filtering
- **mosaic**: Image preview capabilities for repository avatars and media

### TUI Package Management/App Distribution
- **cellbuf**: Interactive package browsers with mouse support
- **layout**: Multi-column package information display
- **img2term**: Logo and icon display for packages
- **charmtone**: Consistent color theming for package status indicators
- **toner**: Text processing for package descriptions and metadata

## Technology Stack
- **Language**: Go 1.24+
- **Dependencies**: Charm ecosystem (Lip Gloss v2, cellbuf, ANSI utilities)
- **Capabilities**: Terminal manipulation, image processing, CLI frameworks, PTY handling
- **Platform Support**: Unix-like systems (Windows excluded for PTY features)