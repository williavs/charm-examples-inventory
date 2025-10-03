# VHS Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|-------------|
| Basic VHS recording demonstration | `demo.tape` |
| Screenshot capture workflow | `screenshot.tape` |
| Chat interface recording | `bubbletea/chat.tape` |
| Package manager interface | `bubbletea/package-manager.tape` |
| Form input recording | `bubbletea/credit-card-form.tape`, `bubbletea/textinput.tape` |
| List interface recording | `bubbletea/list-default.tape`, `bubbletea/list-simple.tape`, `bubbletea/list-fancy.tape` |
| Table data display | `bubbletea/table.tape` |
| Progress indication | `bubbletea/progress-animated.tape`, `bubbletea/progress-static.tape` |
| Spinner animations | `bubbletea/spinner.tape`, `bubbletea/spinners.tape` |
| Split pane layouts | `bubbletea/split-editors.tape` |
| Tab navigation | `bubbletea/tabs.tape` |
| Help screen recording | `bubbletea/help.tape` |
| File navigation | `gum/file.tape` |
| Document viewer | `gum/pager.tape`, `glow/glow-simple.tape` |
| GitHub issue workflow | `gh-cli/gh-issue.tape` |
| GitHub PR workflow | `gh-cli/gh-pr.tape` |
| JSON query interface | `jqp/jqp.tape` |
| Text formatting output | `cli-ui/format.tape` |
| Interactive prompts | `cli-ui/interactive-prompt.tape`, `cli-ui/text-prompt.tape` |
| Status indicators | `cli-ui/status-widget.tape` |
| Keyboard input simulation | `commands/arrow.tape`, `commands/backspace.tape`, `commands/ctrl.tape` |
| Multi-format output | Any tape with multiple Output declarations |

## Examples by Capability

### Terminal Recording Basics
**Use demo.tape when you need:**
- Basic VHS command structure demonstration
- Simple echo command recording
- Minimal example for testing VHS installation

**File**: `demo.tape`
**Key patterns**: Output declaration, Require statement, Set commands, Type/Enter sequence

**Use screenshot.tape when you need:**
- Static terminal snapshot capture
- Non-animated terminal output recording

**File**: `screenshot.tape`
**Key patterns**: Single frame capture, minimal sleep timing

### Chat and Messaging Interfaces
**Use bubbletea/chat.tape when you need:**
- Real-time chat interface recording
- Message input and display patterns
- Interactive text entry demonstration

**File**: `bubbletea/chat.tape`
**Key patterns**: Build Go binary, execute TUI, text input with timing, cleanup

**Use bubbletea/send-msg.tape when you need:**
- Message sending functionality recording
- Button/action trigger demonstration

**File**: `bubbletea/send-msg.tape`
**Key patterns**: Space key for action trigger, automated message sending

### Package Management and Installation
**Use bubbletea/package-manager.tape when you need:**
- Software installation interface recording
- Package selection and management workflows
- Progress tracking for installations

**File**: `bubbletea/package-manager.tape`
**Key patterns**: Extended sleep for automated animation, minimal user interaction

### Form Input and Data Entry
**Use bubbletea/credit-card-form.tape when you need:**
- Multi-field form recording
- Structured data input patterns
- Input validation demonstration

**File**: `bubbletea/credit-card-form.tape`
**Key patterns**: Sequential field entry, Tab navigation between fields

**Use bubbletea/textinput.tape when you need:**
- Single text field recording
- Basic input component demonstration

**File**: `bubbletea/textinput.tape`
**Key patterns**: Simple text entry, Enter to submit

**Use bubbletea/textinputs.tape when you need:**
- Multiple text input fields
- Form with several text entries

**File**: `bubbletea/textinputs.tape`
**Key patterns**: Multiple field navigation, varied input types

**Use bubbletea/textarea.tape when you need:**
- Multi-line text input recording
- Long-form text entry demonstration

**File**: `bubbletea/textarea.tape`
**Key patterns**: Multiple Enter keys, paragraph entry

### List and Selection Interfaces
**Use bubbletea/list-default.tape when you need:**
- Standard list selection recording
- Navigation through items with descriptions

**File**: `bubbletea/list-default.tape`
**Key patterns**: Down/Up arrow navigation, Enter to select

**Use bubbletea/list-simple.tape when you need:**
- Minimal list interface recording
- Basic item selection without descriptions

**File**: `bubbletea/list-simple.tape`
**Key patterns**: Simple arrow navigation, compact display

**Use bubbletea/list-fancy.tape when you need:**
- Styled list interface recording
- Decorated item display

**File**: `bubbletea/list-fancy.tape`
**Key patterns**: Enhanced visual styling, custom item rendering

### Data Display and Presentation
**Use bubbletea/table.tape when you need:**
- Tabular data recording
- Column-based information display

**File**: `bubbletea/table.tape`
**Key patterns**: Row navigation, multi-column layout

**Use bubbletea/tabs.tape when you need:**
- Multi-view interface recording
- Tab switching demonstration

**File**: `bubbletea/tabs.tape`
**Key patterns**: Tab/Shift+Tab navigation, view switching

**Use bubbletea/pager.tape when you need:**
- Paginated content recording
- Large text display with scrolling

**File**: `bubbletea/pager.tape`
**Key patterns**: Page-based navigation, viewport scrolling

**Use bubbletea/paginator.tape when you need:**
- Page number navigation
- Discrete page switching

**File**: `bubbletea/paginator.tape`
**Key patterns**: Page indicator, numbered page selection

### Progress and Status Indicators
**Use bubbletea/progress-animated.tape when you need:**
- Animated progress bar recording
- Ongoing process visualization

**File**: `bubbletea/progress-animated.tape`
**Key patterns**: Auto-updating progress, percentage display

**Use bubbletea/progress-static.tape when you need:**
- Static progress display recording
- Manual progress updates

**File**: `bubbletea/progress-static.tape`
**Key patterns**: Fixed progress states, step-by-step updates

**Use bubbletea/spinner.tape when you need:**
- Loading animation recording
- Indeterminate progress indication

**File**: `bubbletea/spinner.tape`
**Key patterns**: Continuous rotation, loading states

**Use bubbletea/spinners.tape when you need:**
- Multiple spinner styles
- Variety of loading animations

**File**: `bubbletea/spinners.tape`
**Key patterns**: Different spinner characters, style comparison

**Use cli-ui/progress.tape when you need:**
- Ruby CLI progress bar recording
- Alternative progress implementation

**File**: `cli-ui/progress.tape`
**Key patterns**: CLI-UI gem progress patterns

**Use cli-ui/spinner.tape when you need:**
- Ruby CLI spinner recording
- CLI-UI gem loading indication

**File**: `cli-ui/spinner.tape`
**Key patterns**: CLI-UI gem spinner patterns

### Time-Based Components
**Use bubbletea/timer.tape when you need:**
- Countdown timer recording
- Time-limited action demonstration

**File**: `bubbletea/timer.tape`
**Key patterns**: Descending time display, timeout events

**Use bubbletea/stopwatch.tape when you need:**
- Elapsed time recording
- Duration tracking demonstration

**File**: `bubbletea/stopwatch.tape`
**Key patterns**: Ascending time display, start/stop actions

### Layout and Screen Management
**Use bubbletea/split-editors.tape when you need:**
- Multi-pane interface recording
- Split screen layout demonstration

**File**: `bubbletea/split-editors.tape`
**Key patterns**: Dual viewport management, pane navigation

**Use bubbletea/composable-views.tape when you need:**
- Modular component recording
- Nested view composition

**File**: `bubbletea/composable-views.tape`
**Key patterns**: Component hierarchy, view composition

**Use bubbletea/views.tape when you need:**
- Multiple view states
- View transition recording

**File**: `bubbletea/views.tape`
**Key patterns**: State-based view switching

**Use bubbletea/fullscreen.tape when you need:**
- Full terminal usage recording
- Maximum viewport utilization

**File**: `bubbletea/fullscreen.tape`
**Key patterns**: Entire screen rendering

**Use bubbletea/altscreen-toggle.tape when you need:**
- Alternate screen buffer recording
- Screen mode switching demonstration

**File**: `bubbletea/altscreen-toggle.tape`
**Key patterns**: Screen buffer toggle, mode switching

**Use split/split.tape when you need:**
- Multiple terminal split recording
- Concurrent session demonstration

**File**: `split/split.tape`
**Key patterns**: Split terminal layouts

### External Integration
**Use bubbletea/http.tape when you need:**
- HTTP request interface recording
- API interaction demonstration

**File**: `bubbletea/http.tape`
**Key patterns**: Network request display, response handling

**Use bubbletea/exec.tape when you need:**
- External command execution recording
- Subprocess integration demonstration

**File**: `bubbletea/exec.tape`
**Key patterns**: Command execution within TUI, output capture

**Use bubbletea/tui-daemon-combo.tape when you need:**
- Background process coordination
- Daemon interaction recording

**File**: `bubbletea/tui-daemon-combo.tape`
**Key patterns**: TUI with background service

**Use bubbletea/pipe.tape when you need:**
- Stdin/stdout handling recording
- Pipeline integration demonstration

**File**: `bubbletea/pipe.tape`
**Key patterns**: Data piping, stream processing

### Real-Time and Reactive Interfaces
**Use bubbletea/realtime.tape when you need:**
- Live data update recording
- Real-time information display

**File**: `bubbletea/realtime.tape`
**Key patterns**: Auto-refreshing data, live updates

**Use bubbletea/debounce.tape when you need:**
- Input debouncing recording
- Delayed action demonstration

**File**: `bubbletea/debounce.tape`
**Key patterns**: Debounced input handling

**Use bubbletea/sequence.tape when you need:**
- Sequential animation recording
- Step-by-step process demonstration

**File**: `bubbletea/sequence.tape`
**Key patterns**: Ordered action sequence

### Help and Documentation
**Use bubbletea/help.tape when you need:**
- Help screen recording
- Keyboard shortcut documentation

**File**: `bubbletea/help.tape`
**Key patterns**: Help menu display, key binding lists

**Use bubbletea/glamour.tape when you need:**
- Markdown rendering recording
- Styled text display

**File**: `bubbletea/glamour.tape`
**Key patterns**: Rich text formatting, markdown display

### File and Document Management
**Use gum/file.tape when you need:**
- File browser recording
- Directory navigation with selection

**File**: `gum/file.tape`
**Key patterns**: Gum file picker, directory traversal

**Use gum/pager.tape when you need:**
- Document viewer recording with Gum
- Content scrolling with navigation

**File**: `gum/pager.tape`
**Key patterns**: Gum pager component, scrolling content

**Use glow/glow-simple.tape when you need:**
- Markdown viewer recording
- Document browsing interface

**File**: `glow/glow-simple.tape`
**Key patterns**: Glow markdown rendering, navigation

**Use glow/vhs-glow.tape when you need:**
- Advanced markdown workflow recording
- Search and edit operations

**File**: `glow/vhs-glow.tape`
**Key patterns**: Search functionality, content editing, Hide/Show commands

### GitHub Integration
**Use gh-cli/gh-issue.tape when you need:**
- GitHub issue browsing recording
- Issue viewing workflow

**File**: `gh-cli/gh-issue.tape`
**Key patterns**: gh issue commands, list and view operations

**Use gh-cli/gh-pr.tape when you need:**
- Pull request listing recording
- PR management workflow

**File**: `gh-cli/gh-pr.tape`
**Key patterns**: gh pr commands, state filtering

### Data Query and Processing
**Use jqp/jqp.tape when you need:**
- JSON query interface recording
- Interactive jq expression building

**File**: `jqp/jqp.tape`
**Key patterns**: JSON navigation, query building

**Use gum/table.tape when you need:**
- CSV data display recording
- Tabular data presentation with Gum

**File**: `gum/table.tape`
**Key patterns**: Gum table component, column formatting

### Text Formatting and Output
**Use cli-ui/format.tape when you need:**
- Colored text output recording
- Text styling demonstration

**File**: `cli-ui/format.tape`
**Key patterns**: ANSI color codes, text decoration

**Use cli-ui/symbols.tape when you need:**
- Symbol and icon output recording
- Unicode character display

**File**: `cli-ui/symbols.tape`
**Key patterns**: Special characters, status symbols

**Use cli-ui/nested-frames.tape when you need:**
- Hierarchical layout recording
- Nested frame structures

**File**: `cli-ui/nested-frames.tape`
**Key patterns**: Frame nesting, border composition

### User Interaction Patterns
**Use cli-ui/interactive-prompt.tape when you need:**
- Interactive question recording
- User confirmation dialog

**File**: `cli-ui/interactive-prompt.tape`
**Key patterns**: Yes/no prompts, user confirmation

**Use cli-ui/text-prompt.tape when you need:**
- Text prompt recording
- User input request

**File**: `cli-ui/text-prompt.tape`
**Key patterns**: Input prompt, response capture

**Use cli-ui/status-widget.tape when you need:**
- Status indicator recording
- Real-time status display

**File**: `cli-ui/status-widget.tape`
**Key patterns**: Status symbols, state indication

### System Information Display
**Use neofetch/*.tape when you need:**
- System information display recording
- OS/hardware specs with ASCII art

**Files**: `neofetch/neofetch.tape`, `neofetch/colors-*.tape` (6 variants)
**Key patterns**: System info display, ASCII art, color schemes

### Keyboard Input Simulation
**Use commands/arrow.tape when you need:**
- Arrow key navigation recording
- Directional input demonstration

**File**: `commands/arrow.tape`
**Key patterns**: Up/Down/Left/Right key sequences

**Use commands/backspace.tape when you need:**
- Backspace key recording
- Text deletion demonstration

**File**: `commands/backspace.tape`
**Key patterns**: Backspace key timing

**Use commands/ctrl.tape when you need:**
- Control key combination recording
- Keyboard shortcuts demonstration

**File**: `commands/ctrl.tape`
**Key patterns**: Ctrl+[key] sequences

**Use commands/enter.tape when you need:**
- Enter key recording
- Command submission demonstration

**File**: `commands/enter.tape`
**Key patterns**: Enter key timing

**Use commands/space.tape when you need:**
- Space key recording
- Whitespace input demonstration

**File**: `commands/space.tape`
**Key patterns**: Space key timing

**Use commands/tab.tape when you need:**
- Tab key recording
- Field navigation or completion demonstration

**File**: `commands/tab.tape`
**Key patterns**: Tab key timing

**Use commands/type.tape when you need:**
- Text typing recording
- Character-by-character input

**File**: `commands/type.tape`
**Key patterns**: Type command with custom timing

**Use commands/clipboard.tape when you need:**
- Paste operation recording
- Clipboard interaction demonstration

**File**: `commands/clipboard.tape`
**Key patterns**: Clipboard paste simulation

**Use commands/alt.tape when you need:**
- Alt key combination recording
- Alternative keyboard shortcuts

**File**: `commands/alt.tape`
**Key patterns**: Alt+[key] sequences

### Recording Control and Visibility
**Use commands/hide.tape when you need:**
- Hidden command execution
- Setup steps concealment

**File**: `commands/hide.tape`
**Key patterns**: Hide/Show command pairs, invisible preparation

**Use commands/show.tape when you need:**
- Visible command demonstration
- Display toggle patterns

**File**: `commands/show.tape`
**Key patterns**: Show command, visibility control

**Use commands/comment.tape when you need:**
- Inline documentation
- Tape file commenting

**File**: `commands/comment.tape`
**Key patterns**: Comment syntax, documentation

### VHS Configuration Examples
**Use settings/set-font-size-*.tape when you need:**
- Font size variation recording
- Text scale demonstration

**Files**: `settings/set-font-size-10.tape`, `settings/set-font-size-20.tape`, `settings/set-font-size-40.tape`
**Key patterns**: Set FontSize command, size comparison

**Use settings/set-font-family.tape when you need:**
- Font family configuration recording
- Typeface customization

**File**: `settings/set-font-family.tape`
**Key patterns**: Set FontFamily command, font selection

**Use settings/set-letter-spacing.tape when you need:**
- Character spacing recording
- Typography adjustment

**File**: `settings/set-letter-spacing.tape`
**Key patterns**: Set LetterSpacing command, tracking adjustment

**Use settings/set-line-height.tape when you need:**
- Line spacing recording
- Vertical spacing adjustment

**File**: `settings/set-line-height.tape`
**Key patterns**: Set LineHeight command, leading adjustment

**Use settings/height.tape when you need:**
- Terminal height configuration
- Viewport size recording

**File**: `settings/height.tape`
**Key patterns**: Set Height command

**Use settings/width.tape when you need:**
- Terminal width configuration
- Horizontal dimension setting

**File**: `settings/width.tape`
**Key patterns**: Set Width command

**Use settings/set-padding.tape when you need:**
- Interior spacing recording
- Terminal padding configuration

**File**: `settings/set-padding.tape`
**Key patterns**: Set Padding command

**Use settings/set-margin.tape when you need:**
- Exterior spacing recording
- Terminal margin configuration

**File**: `settings/set-margin.tape`
**Key patterns**: Set Margin command, MarginFill usage

**Use settings/set-border-radius.tape when you need:**
- Rounded corner recording
- Border styling demonstration

**File**: `settings/set-border-radius.tape`
**Key patterns**: Set BorderRadius command

**Use settings/set-window-bar.tape when you need:**
- Title bar style recording
- Window decoration configuration

**File**: `settings/set-window-bar.tape`
**Key patterns**: Set WindowBar command, bar style options

**Use settings/set-theme.tape when you need:**
- Color scheme recording with JSON
- Custom theme configuration

**File**: `settings/set-theme.tape`
**Key patterns**: Set Theme with JSON object

**Use settings/set-theme-name.tape when you need:**
- Named theme recording
- Predefined color scheme selection

**File**: `settings/set-theme-name.tape`
**Key patterns**: Set Theme with string name

**Use settings/set-typing-speed.tape when you need:**
- Typing speed customization
- Input timing adjustment

**File**: `settings/set-typing-speed.tape`
**Key patterns**: Set TypingSpeed command

**Use settings/set-cursor-blink.tape when you need:**
- Cursor animation recording
- Cursor behavior configuration

**File**: `settings/set-cursor-blink.tape`
**Key patterns**: Set CursorBlink command

**Use settings/set-loop-offset.tape when you need:**
- GIF loop timing recording
- Animation loop customization

**File**: `settings/set-loop-offset.tape`
**Key patterns**: Set LoopOffset command

**Use settings/set-shell-*.tape when you need:**
- Shell environment recording
- Cross-platform shell demonstration

**Files**: `settings/set-shell-bash.tape`, `settings/set-shell-zsh.tape`, `settings/set-shell-fish.tape`, `settings/set-shell-nu.tape`, `settings/set-shell-cmd.tape`, `settings/set-shell-pwsh.tape`, `settings/set-shell-osh.tape`, `settings/set-shell-xonsh.tape`, `settings/set-shell-custom.tape`
**Key patterns**: Set Shell command, shell-specific syntax

### Error Handling and Validation
**Use errors/dimensions.tape when you need:**
- Dimension validation error demonstration
- Invalid size configuration example

**File**: `errors/dimensions.tape`
**Key patterns**: Invalid dimension values, error messages

**Use errors/parser.tape when you need:**
- Syntax error demonstration
- Tape file parsing failure

**File**: `errors/parser.tape`
**Key patterns**: Invalid syntax, parser errors

**Use errors/require.tape when you need:**
- Missing dependency error demonstration
- Require statement validation

**File**: `errors/require.tape`
**Key patterns**: Require command, missing program error

### Publishing and Output Formats
**Use publish/*.tape when you need:**
- Multi-format output generation
- Distribution preparation

**Files**: `publish/animated-gif.tape`, `publish/ascii.tape`, `publish/static.tape`
**Key patterns**: Multiple Output commands, format variants

### Environment Configuration
**Use env/env.tape when you need:**
- Environment variable usage
- Dynamic configuration

**File**: `env/env.tape`
**Key patterns**: Environment variable references

### Visual Enhancements
**Use decorations/decorations.tape when you need:**
- Visual styling recording
- Decoration techniques

**File**: `decorations/decorations.tape`
**Key patterns**: Visual enhancement commands

### Presentation Mode
**Use slides/slides.tape when you need:**
- Slide presentation recording
- Sequential content display

**File**: `slides/slides.tape`
**Key patterns**: Slide transitions, presentation flow

### Simple Examples
**Use bubbletea/simple.tape when you need:**
- Minimal Bubble Tea program recording
- Basic TUI skeleton

**File**: `bubbletea/simple.tape`
**Key patterns**: Minimal TUI structure

**Use bubbletea/result.tape when you need:**
- Result display recording
- Output presentation

**File**: `bubbletea/result.tape`
**Key patterns**: Result rendering

## Implementation Patterns

### VHS Tape File Structure
All VHS recordings follow this general pattern:
1. Output declaration (GIF, MP4, WebM, ASCII)
2. Optional Require statements for dependencies
3. Set commands for visual configuration
4. Hide command for invisible setup
5. Build/preparation commands
6. Show command to reveal recording
7. Visible demonstration commands
8. Hide command for cleanup
9. Cleanup commands

### Bubble Tea Application Recording
Standard pattern for Go TUI applications:
1. Hide initial commands
2. Build Go binary: `go build -o [name] .`
3. Clear screen: `clear`
4. Show visible portion
5. Execute binary: `./[name]`
6. Interact with TUI
7. Hide cleanup
8. Exit application (Ctrl+C)
9. Remove binary: `rm [name]`

### Command Timing Patterns
- `Sleep [duration]` - Pause recording
- `Type@[speed] "[text]"` - Custom typing speed
- `[Key]@[interval] [count]` - Repeated key with timing
- `Sleep` after `Enter` - Wait for command completion
- `Sleep` before user action - Display time

### Visibility Control Pattern
```
Hide
[setup commands]
Show
[visible demonstration]
Hide
[cleanup commands]
```

### Multi-Format Output Pattern
```
Output example.gif
Output example.mp4
Output example.webm
Output example.ascii
```

### Navigation Patterns
- **List navigation**: Down@[speed] [count], Up@[speed] [count]
- **Search**: Type "/", Type "[query]", Enter
- **Mode switch**: Escape, Type "[command]"
- **Selection**: Enter (after navigation)

### Form Input Patterns
- **Field entry**: Type "[value]"
- **Field navigation**: Tab (next), Shift+Tab (previous)
- **Submission**: Enter (final field or dedicated button)

### Configuration Patterns
All Set commands apply globally and persist:
- **Visual**: FontSize, FontFamily, LetterSpacing, LineHeight
- **Layout**: Width, Height, Padding, Margin, BorderRadius
- **Behavior**: Shell, TypingSpeed, PlaybackSpeed, Framerate
- **Styling**: Theme, WindowBar, CursorBlink, LoopOffset

### Shell Compatibility Pattern
Use `Set Shell "[shell]"` for shell-specific syntax:
- bash, zsh, fish - POSIX-like shells
- cmd, pwsh - Windows shells
- nu, osh, xonsh - Alternative shells

### Error Demonstration Pattern
Intentionally invalid configurations for documentation:
- Invalid dimension values
- Missing required programs
- Syntax errors in tape files
