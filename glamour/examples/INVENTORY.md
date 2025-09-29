# Glamour Examples Inventory

## Overview
This directory contains 77 files demonstrating Glamour markdown rendering capabilities for terminal interfaces. The examples consist of 30 markdown files, 33 style files, and 14 PNG screenshots showing rendered output. Each example demonstrates specific markdown element rendering with customizable styling through JSON configuration files.

## Examples Catalog

### Text Formatting
- **block_quote**: Demonstrates quote block rendering with indentation and styling
- **code**: Inline code snippet formatting with background highlighting
- **code_block**: Multi-line code block rendering with syntax highlighting support
- **emph**: Italic text emphasis rendering
- **strong**: Bold text emphasis rendering
- **strikethrough**: Text strikethrough formatting

### Document Structure
- **heading**: Hierarchical heading rendering (h1, h2, h3) with prefix/suffix customization
- **hr**: Horizontal rule rendering for document separation
- **paragraph**: Basic paragraph text formatting

### Lists and Tasks
- **enumeration**: Numbered list rendering with customizable numbering styles
- **list**: Unordered list rendering with bullet point customization
- **ordered_list**: Sequential numbered list implementation
- **task**: Checkbox task list rendering for TODO items

### Links and Media
- **emoji**: Emoji character rendering in terminal output
- **image**: Image placeholder rendering with alt-text display
- **link**: Hyperlink rendering with URL highlighting

### Table Rendering (13 variations)
- **table**: Basic table structure with headers and rows
- **table_align**: Table column alignment demonstration
- **table_truncate**: Table content truncation for width management
- **table_wrap**: Table text wrapping within cells
- **table_with_footer_auto_links**: Auto-generated footer links
- **table_with_footer_auto_links_short**: Shortened footer link format
- **table_with_footer_images**: Image references in table footers
- **table_with_footer_images_no_alt**: Image handling without alt text
- **table_with_footer_images_same_alt**: Duplicate alt text handling
- **table_with_footer_links**: Reference links in table footers
- **table_with_footer_links_complex**: Multi-developer profile table with links and images
- **table_with_footer_links_huge**: Large reference link collections
- **table_with_footer_links_no_color**: Colorless link rendering
- **table_with_footer_links_repeated**: Duplicate link reference handling
- **table_with_inline_links**: Direct inline link embedding

## Implementation Patterns

### Style Configuration Pattern
Each example follows a consistent structure:
- `.md` file: Contains markdown content to render
- `.style` file: JSON configuration defining colors, margins, prefixes, and formatting
- `.png` file: Screenshot showing expected terminal output (when available)

### Style Properties
Style files use JSON objects with element-specific configuration:
- Color values (numeric color codes)
- Background colors
- Margins and indentation
- Text prefixes and suffixes
- Font styling (bold, italic)

### Table Rendering Complexity
Tables demonstrate the most complex rendering scenarios with:
- Footer reference management
- Image placeholder handling
- Link processing and deduplication
- Content truncation and wrapping
- Multi-column alignment

## Practical Applications

### NextJS Scaffolder with Auth/Shadcn
- **task** examples: Progress tracking for scaffolding steps
- **table** examples: Configuration option displays
- **heading** patterns: Section organization for setup wizards
- **code_block** styling: Command output and configuration display

### Inter-Agent Communication Systems
- **emoji** rendering: Status indicators and alert symbols
- **strong/emph** formatting: Message priority highlighting
- **list** patterns: Agent status displays and command menus
- **table** layouts: Agent registry and status dashboards

### GitHub Search Tools
- **table_with_footer_links_complex**: Repository and user profile displays
- **link** formatting: Repository and issue linking
- **code** styling: Repository path and branch displays
- **image** handling: Avatar and organization logo display

### TUI App Distribution/Package Management
- **enumeration** patterns: Package listing and version displays
- **task** rendering: Installation progress tracking
- **table** variations: Package dependency matrices
- **hr** separation: Section dividers in package information
- **block_quote** styling: Package description and changelog display

The table rendering examples provide the most relevant patterns for data-heavy applications, while text formatting examples offer foundation elements for status displays and user interface text.