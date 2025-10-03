# VHS Examples

Contextual reference examples for VHS terminal recorder. Each `.tape` file demonstrates specific capabilities and patterns.

## Quick Start

Browse examples by need in [CONTEXTUAL-INVENTORY.md](./CONTEXTUAL-INVENTORY.md).

## Categories

### Commands
Input simulation patterns (`commands/`) - keyboard input, navigation, control sequences.

### Settings
Configuration examples (`settings/`) - dimensions, fonts, themes, shell environments.

### Bubbletea
Interactive TUI components (`bubbletea/`) - forms, lists, progress, navigation, layout.

### CLI-UI
Non-interactive UI patterns (`cli-ui/`) - formatted output, status displays, progress bars.

### Tool Examples
Real-world tool recordings (`gum/`, `gh-cli/`, `glow/`, `jqp/`, `neofetch/`, `slides/`) - practical workflow demonstrations.

### Utilities
Special-purpose examples (`env/`, `errors/`, `fixtures/`, `publish/`, `split/`, `decorations/`) - environment setup, error handling, publishing, styling.

## Running Examples

```bash
# Generate GIF from any tape file
vhs path/to/example.tape

# Example
vhs bubbletea/spinner.tape
```

## File Locations

All tape files use relative paths from examples directory root:
- Bubbletea: `examples/bubbletea/[name].tape`
- Commands: `examples/commands/[name].tape`
- Settings: `examples/settings/[name].tape`
- Etc.

See [CONTEXTUAL-INVENTORY.md](./CONTEXTUAL-INVENTORY.md) for complete file listing and usage patterns.
