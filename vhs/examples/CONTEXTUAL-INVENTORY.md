# VHS Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|-------------|
| Basic animation recording | `examples/demo.tape` |
| Single screenshot capture | `examples/screenshot.tape` |
| Publishing workflow | `examples/publish/publish.tape` |
| Environment variables | `examples/env/env.tape` |
| Split pane workflows | `examples/split/split.tape` |
| Visual styling showcase | `examples/decorations/decorations.tape` |
| Input simulation | `examples/commands/type.tape` |
| Theme configuration | `examples/settings/set-theme.tape` |
| Spinner component | `examples/bubbletea/spinner.tape` |
| Progress indicators | `examples/bubbletea/progress-animated.tape` |
| List navigation | `examples/bubbletea/list-default.tape` |
| Text input forms | `examples/bubbletea/textinput.tape` |
| Table rendering | `examples/bubbletea/table.tape` |
| Modal prompts | `examples/cli-ui/interactive-prompt.tape` |

## Examples by Capability

### Core VHS Commands

**Use examples/commands/ when you need:**
- Keyboard input simulation patterns
- Arrow navigation patterns
- Text entry with timing controls
- Clipboard operations
- Comment syntax
- Control key sequences

**Files:**
- `examples/commands/type.tape` - Text input with variable speed
- `examples/commands/arrow.tape` - Directional navigation
- `examples/commands/alt.tape` - Alt key combinations
- `examples/commands/ctrl.tape` - Control key combinations
- `examples/commands/backspace.tape` - Deletion operations
- `examples/commands/clipboard.tape` - Copy/paste workflows
- `examples/commands/comment.tape` - Comment syntax usage
- `examples/commands/enter.tape` - Return key execution
- `examples/commands/hide.tape` - Hide terminal output
- `examples/commands/show.tape` - Show terminal output
- `examples/commands/space.tape` - Space key input
- `examples/commands/tab.tape` - Tab key completion

**Key patterns:** Timing control with `@` suffix, multi-step input sequences

### VHS Settings & Configuration

**Use examples/settings/ when you need:**
- Output formatting configuration
- Terminal appearance customization
- Shell environment configuration
- Dimension control patterns
- Typography settings

**Files:**
- `examples/settings/set-theme.tape` - Custom color scheme application
- `examples/settings/set-theme-name.tape` - Named theme selection
- `examples/settings/set-font-family.tape` - Font selection
- `examples/settings/set-font-size-10.tape` - Small text rendering
- `examples/settings/set-font-size-20.tape` - Medium text rendering
- `examples/settings/set-font-size-40.tape` - Large text rendering
- `examples/settings/height.tape` - Vertical dimension control
- `examples/settings/width.tape` - Horizontal dimension control
- `examples/settings/set-padding.tape` - Internal spacing
- `examples/settings/set-margin.tape` - External spacing
- `examples/settings/set-border-radius.tape` - Corner rounding
- `examples/settings/set-letter-spacing.tape` - Character spacing
- `examples/settings/set-line-height.tape` - Vertical text spacing
- `examples/settings/set-cursor-blink.tape` - Cursor animation
- `examples/settings/set-typing-speed.tape` - Global input timing
- `examples/settings/set-loop-offset.tape` - Loop timing control
- `examples/settings/set-window-bar.tape` - Window decoration
- `examples/settings/set-shell-bash.tape` - Bash shell configuration
- `examples/settings/set-shell-zsh.tape` - Zsh shell configuration
- `examples/settings/set-shell-fish.tape` - Fish shell configuration
- `examples/settings/set-shell-nu.tape` - Nushell configuration
- `examples/settings/set-shell-pwsh.tape` - PowerShell configuration
- `examples/settings/set-shell-cmd.tape` - Windows CMD configuration
- `examples/settings/set-shell-osh.tape` - Oil shell configuration
- `examples/settings/set-shell-xonsh.tape` - Xonsh configuration
- `examples/settings/set-shell-custom.tape` - Custom shell setup

**Key patterns:** `Set` command syntax, dimension units, JSON theme objects

### Bubbletea TUI Components

**Use examples/bubbletea/ when you need:**
- Interactive UI component recording
- State management workflows
- Complex navigation patterns
- Form input sequences
- Component lifecycle demonstrations

#### Input Components
- `examples/bubbletea/textinput.tape` - Single-line text entry
- `examples/bubbletea/textinputs.tape` - Multi-field forms
- `examples/bubbletea/textarea.tape` - Multi-line text editing
- `examples/bubbletea/credit-card-form.tape` - Structured form validation

#### Selection & Navigation
- `examples/bubbletea/list-simple.tape` - Basic list selection
- `examples/bubbletea/list-default.tape` - Standard list with search
- `examples/bubbletea/list-fancy.tape` - Styled list rendering
- `examples/bubbletea/table.tape` - Tabular data navigation
- `examples/bubbletea/tabs.tape` - Tab-based interface switching

#### Feedback Components
- `examples/bubbletea/spinner.tape` - Loading indicator
- `examples/bubbletea/spinners.tape` - Multiple spinner styles
- `examples/bubbletea/progress-static.tape` - Fixed progress display
- `examples/bubbletea/progress-animated.tape` - Dynamic progress updates
- `examples/bubbletea/timer.tape` - Countdown functionality
- `examples/bubbletea/stopwatch.tape` - Elapsed time tracking

#### Content Display
- `examples/bubbletea/pager.tape` - Scrollable content viewer
- `examples/bubbletea/paginator.tape` - Page-based navigation
- `examples/bubbletea/help.tape` - Help text rendering
- `examples/bubbletea/glamour.tape` - Markdown rendering

#### Layout & Composition
- `examples/bubbletea/views.tape` - Multi-view switching
- `examples/bubbletea/composable-views.tape` - Nested view composition
- `examples/bubbletea/split-editors.tape` - Split pane editing
- `examples/bubbletea/fullscreen.tape` - Fullscreen mode handling
- `examples/bubbletea/altscreen-toggle.tape` - Alternate screen buffer

#### Workflow Demonstrations
- `examples/bubbletea/package-manager.tape` - Installation workflow recording
- `examples/bubbletea/chat.tape` - Message exchange interface
- `examples/bubbletea/result.tape` - Operation result display
- `examples/bubbletea/sequence.tape` - Multi-step process flow

#### Advanced Patterns
- `examples/bubbletea/exec.tape` - External command execution
- `examples/bubbletea/http.tape` - Network request handling
- `examples/bubbletea/pipe.tape` - Pipeline data handling
- `examples/bubbletea/realtime.tape` - Live data updates
- `examples/bubbletea/send-msg.tape` - Message passing patterns
- `examples/bubbletea/debounce.tape` - Input debouncing
- `examples/bubbletea/tui-daemon-combo.tape` - Background service coordination
- `examples/bubbletea/simple.tape` - Minimal implementation

**Key patterns:** Build/execute/cleanup workflow, interaction timing, state transitions

### CLI-UI Components

**Use examples/cli-ui/ when you need:**
- Non-interactive UI patterns
- Formatted output rendering
- Status display patterns
- Progress feedback without user input

**Files:**
- `examples/cli-ui/format.tape` - Text formatting patterns
- `examples/cli-ui/interactive-prompt.tape` - Selection prompts
- `examples/cli-ui/nested-frames.tape` - Bordered content nesting
- `examples/cli-ui/progress.tape` - Progress bar rendering
- `examples/cli-ui/spinner.tape` - Loading animation
- `examples/cli-ui/status-widget.tape` - Status indicator display
- `examples/cli-ui/symbols.tape` - Icon and symbol usage
- `examples/cli-ui/text-prompt.tape` - Simple text input

**Key patterns:** Ruby IRB setup, library loading, API usage patterns

### Real-World Tool Examples

**Use examples/gum/ when you need:**
- File selection workflows
- Scrollable content patterns
- Table display with navigation

**Files:**
- `examples/gum/file.tape` - File browser navigation
- `examples/gum/pager.tape` - Long content scrolling
- `examples/gum/table.tape` - Data table interaction

**Use examples/gh-cli/ when you need:**
- List filtering workflows
- Detail view patterns

**Files:**
- `examples/gh-cli/gh-issue.tape` - Issue listing and viewing
- `examples/gh-cli/gh-pr.tape` - Pull request listing

**Use examples/glow/ when you need:**
- Document viewer workflows
- Search and navigation patterns
- Content editing sequences

**Files:**
- `examples/glow/glow-simple.tape` - Basic document viewing
- `examples/glow/glow.tape` - Full navigation workflow
- `examples/glow/glow-edit.tape` - Document modification

**Use examples/jqp/ when you need:**
- JSON query interface recording

**File:** `examples/jqp/jqp.tape`

**Use examples/neofetch/ when you need:**
- System information display

**File:** `examples/neofetch/neofetch.tape`

**Use examples/slides/ when you need:**
- Presentation navigation

**File:** `examples/slides/slides.tape`

**Key patterns:** Tool-specific command patterns, realistic interaction timing

### Utility & Testing Examples

**Use examples/env/ when you need:**
- Environment variable configuration patterns

**File:** `examples/env/env.tape`

**Use examples/errors/ when you need:**
- VHS error handling patterns
- Parser error demonstrations
- Dimension validation
- Dependency checking

**Files:**
- `examples/errors/parser.tape` - Invalid syntax examples
- `examples/errors/dimensions.tape` - Size constraint errors
- `examples/errors/require.tape` - Missing dependency handling

**Use examples/fixtures/ when you need:**
- Comprehensive VHS feature testing

**File:** `examples/fixtures/all.tape`

**Use examples/split/ when you need:**
- Tmux integration patterns
- Multi-pane workflows

**File:** `examples/split/split.tape` (requires tmux)

**Key patterns:** `Require` directive usage, error demonstration patterns

### Publishing & Deployment

**Use examples/publish/ when you need:**
- VHS publishing workflow recording
- Output upload patterns

**Files:**
- `examples/publish/publish.tape` - Publishing command demo
- `examples/publish/cassette.tape` - Source tape for publishing

**Key patterns:** `--publish` flag usage, extended sleep for upload completion

### Visual Styling Showcase

**Use examples/decorations/ when you need:**
- Comprehensive styling demonstration
- Multiple setting combinations

**File:** `examples/decorations/decorations.tape`

**Key patterns:** Margin fills, window bars, border radius, extensive padding

### Meta Examples

**Use when you need:**
- VHS recording VHS workflow
- Self-referential demonstrations

**Files:**
- `examples/meta.tape` - VHS recording itself
- `examples/welcome.tape` - Introductory animation

**Key patterns:** Recursive recording patterns

## Implementation Patterns

### Standard Build/Execute/Cleanup
Pattern used in most Bubbletea examples:
```
Hide
Type "go build -o [binary] ."
Enter
Type "clear"
Enter
Show

Type "./[binary]"
Enter
[interactions]

Hide
Ctrl+C
Type "rm [binary]"
Enter
```

### Timing Control
- `Sleep [duration]` - Pause before next action
- `Type@[speed] "text"` - Variable typing speed
- `Down@[interval] [count]` - Repeated actions with timing

### Output Management
- `Output [path].gif` - GIF animation output
- `Output [path].mp4` - Video output
- `Output [path].webm` - WebM video output
- `Output [path].png` - Static screenshot
- `Output [path].ascii` - ASCII art output

### Visibility Control
- `Hide` - Suppress output during setup
- `Show` - Resume visible recording
- Pattern: Hide during compilation/cleanup, show during interaction

### Environment Setup
- `Set [Setting] [Value]` - Configure VHS behavior
- `Env [VAR] "value"` - Set environment variables
- `Require [tool]` - Enforce dependency availability

### Navigation Patterns
- `Down@[speed] [count]` - Repeated down movement
- `Up@[speed] [count]` - Repeated up movement
- `Right`/`Left` - Horizontal navigation
- `Enter` - Confirm/execute
- `Escape` - Cancel/exit

### Cleanup Patterns
- `Ctrl+C` - Interrupt running process
- `Ctrl+L` - Clear screen
- `Ctrl+U` - Clear line
- `Ctrl+A` - Move to line start

## Cross-Reference Notes

### For Package Installation Workflows
See `examples/bubbletea/package-manager.tape` - demonstrates timed installation process recording with automated progression.

### For Form Input Patterns
See:
- `examples/bubbletea/textinput.tape` - Single field
- `examples/bubbletea/textinputs.tape` - Multi-field
- `examples/bubbletea/credit-card-form.tape` - Validated structured input

### For Progress Feedback
See:
- `examples/bubbletea/spinner.tape` - Indeterminate progress
- `examples/bubbletea/progress-animated.tape` - Determinate progress
- `examples/cli-ui/progress.tape` - Simple progress bar

### For Search/Filter Workflows
See:
- `examples/bubbletea/list-default.tape` - List with search (`/` key)
- `examples/glow/glow.tape` - Document search and navigation

### For Scrollable Content
See:
- `examples/bubbletea/pager.tape` - Integrated pager component
- `examples/gum/pager.tape` - Standalone pager tool

### For Split View Interfaces
See:
- `examples/bubbletea/split-editors.tape` - Application-level splits
- `examples/split/split.tape` - Terminal multiplexer integration

### For Configuration Examples
All VHS settings have dedicated files in `examples/settings/` directory. Each demonstrates single setting in isolation for clear reference.

## Directory Structure

```
examples/
├── bubbletea/          # Interactive TUI components (36 files)
├── cli-ui/             # Non-interactive UI patterns (8 files)
├── commands/           # VHS command demonstrations (12 files)
├── decorations/        # Visual styling showcase (1 file)
├── env/                # Environment configuration (1 file)
├── errors/             # Error handling patterns (3 files)
├── fixtures/           # Testing fixtures (1 file)
├── gh-cli/             # GitHub CLI examples (2 files)
├── glow/               # Glow markdown viewer (3 files)
├── gum/                # Gum tool examples (3 files)
├── jqp/                # JSON query processor (1 file)
├── neofetch/           # System info display (1 file)
├── publish/            # Publishing workflow (2 files)
├── settings/           # VHS configuration (28 files)
├── slides/             # Presentation tool (1 file)
├── split/              # Terminal multiplexing (1 file)
├── demo.tape           # Main demonstration
├── meta.tape           # Self-referential recording
├── screenshot.tape     # Screenshot capture
└── welcome.tape        # Introductory animation
```

Total: 109 tape files demonstrating VHS capabilities and usage patterns.
