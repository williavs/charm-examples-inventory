# Glamour Examples - Contextual Reference

## Quick Reference

| Need | Example File | Style File |
|------|-------------|------------|
| Simple markdown rendering | `examples2/helloworld/main.go` | N/A (uses built-in "dark") |
| Stdin/stdout processing | `examples2/stdin-stdout/main.go` | N/A (auto-detect) |
| Custom word wrapping | `examples2/custom_renderer/main.go` | N/A (programmatic) |
| Environment-based config | `examples2/stdin-stdout-custom-styles/main.go` | Environment vars |
| Color profile detection | `examples2/artichokes/main.go` | Auto-downsampled |
| Block quotes | `examples/block_quote.md` | `examples/block_quote.style` |
| Code blocks with syntax highlighting | `examples/code_block.md` | `examples/code_block.style` |
| Inline code formatting | `examples/code.md` | `examples/code.style` |
| Emphasis (italic) | `examples/emph.md` | `examples/emph.style` |
| Strong emphasis (bold) | `examples/strong.md` | `examples/strong.style` |
| Strikethrough text | `examples/strikethrough.md` | `examples/strikethrough.style` |
| Hierarchical headings | `examples/heading.md` | `examples/heading.style` |
| Horizontal rules | `examples/hr.md` | `examples/hr.style` |
| Paragraphs | N/A | `examples/paragraph.style` |
| Numbered lists | `examples/enumeration.md` | `examples/enumeration.style` |
| Bullet lists | `examples/list.md` | `examples/list.style` |
| Ordered lists | `examples/ordered_list.md` | `examples/ordered_list.style` |
| Task lists (checkboxes) | `examples/task.md` | `examples/task.style` |
| Emoji rendering | `examples/emoji.md` | `examples/emoji.style` |
| Image placeholders | `examples/image.md` | `examples/image.style` |
| Hyperlinks | `examples/link.md` | `examples/link.style` |
| Basic tables | `examples/table.md` | `examples/table.style` |
| Table column alignment | `examples/table_align.md` | `examples/table_align.style` |
| Table content truncation | `examples/table_truncate.md` | `examples/table_truncate.style` |
| Table text wrapping | `examples/table_wrap.md` | `examples/table_wrap.style` |
| Tables with footer links | `examples/table_with_footer_links.md` | `examples/table_with_footer_links.style` |
| Tables with inline links | `examples/table_with_inline_links.md` | `examples/table_with_inline_links.style` |
| Tables with images | `examples/table_with_footer_images.md` | `examples/table_with_footer_images.style` |
| Auto-generated link footers | `examples/table_with_footer_auto_links.md` | `examples/table_with_footer_auto_links.style` |

## Examples by Capability

### Basic Markdown Rendering

**Use helloworld when you need:**
- Simple one-function markdown rendering
- Quick prototyping with default styles
- Direct string-to-terminal output

**File**: `examples2/helloworld/main.go`
**Key patterns**: `glamour.Render(markdown, "dark")` for immediate output

---

**Use stdin-stdout when you need:**
- Pipe-based markdown processing
- CLI tool integration with standard streams
- Auto-detected terminal styling

**File**: `examples2/stdin-stdout/main.go`
**Key patterns**: `io.ReadAll(os.Stdin)`, `glamour.WithAutoStyle()`, word wrapping configuration

---

**Use stdin-stdout-custom-styles when you need:**
- Environment variable style configuration
- System-wide style consistency
- Configuration file integration

**File**: `examples2/stdin-stdout-custom-styles/main.go`
**Key patterns**: `glamour.WithEnvironmentConfig()` for env-based theming

---

**Use custom_renderer when you need:**
- Programmatic width constraints
- Renderer configuration control
- Custom rendering pipelines

**File**: `examples2/custom_renderer/main.go`
**Key patterns**: `glamour.NewTermRenderer()`, `glamour.WithWordWrap(width)`, `glamour.WithStandardStyle()`

---

**Use artichokes when you need:**
- Terminal color profile detection
- Automatic color downsampling
- Cross-terminal compatibility
- Embedded file rendering

**File**: `examples2/artichokes/main.go`
**Key patterns**: `colorprofile.NewWriter()`, `embed.FS`, profile-aware rendering

### Text Formatting

**Use block_quote when you need:**
- Indented quotation display
- Highlighted citation blocks
- Nested quote rendering

**File**: `examples/block_quote.md`
**Style**: `examples/block_quote.style`
**Key properties**: Color, margin, indentation

---

**Use code when you need:**
- Inline code snippet formatting
- Monospace text highlighting within paragraphs
- Background color differentiation

**File**: `examples/code.md`
**Style**: `examples/code.style`
**Key properties**: Color, background_color

---

**Use code_block when you need:**
- Multi-line code display
- Syntax highlighting with themes
- Formatted code output rendering

**File**: `examples/code_block.md`
**Style**: `examples/code_block.style`
**Key properties**: Color, theme (e.g., "solarized-dark")

---

**Use emph when you need:**
- Italic text emphasis
- Secondary text highlighting
- Stylistic differentiation

**File**: `examples/emph.md`
**Style**: `examples/emph.style`
**Key properties**: Color, italic styling

---

**Use strong when you need:**
- Bold text emphasis
- Primary text highlighting
- Weight-based differentiation

**File**: `examples/strong.md`
**Style**: `examples/strong.style`
**Key properties**: Color, bold attribute

---

**Use strikethrough when you need:**
- Deprecated text marking
- Completed item indication
- Text correction display

**File**: `examples/strikethrough.md`
**Style**: `examples/strikethrough.style`
**Key properties**: Color, strikethrough rendering

### Document Structure

**Use heading when you need:**
- Hierarchical section organization
- Multi-level heading display (h1-h6)
- Customizable heading prefixes/suffixes

**File**: `examples/heading.md`
**Style**: `examples/heading.style`
**Key properties**: Per-level color/background, prefix/suffix, margins, bold

---

**Use hr when you need:**
- Section separation
- Visual content dividers
- Document segmentation

**File**: `examples/hr.md`
**Style**: `examples/hr.style`
**Key properties**: Color, character styling

---

**Use paragraph when you need:**
- Basic text block styling
- Baseline paragraph rendering
- Text flow configuration

**File**: N/A (style-only)
**Style**: `examples/paragraph.style`
**Key properties**: Margin, indent

### Lists and Tasks

**Use enumeration when you need:**
- Numbered list rendering
- Sequential item display
- Custom numbering styles

**File**: `examples/enumeration.md`
**Style**: `examples/enumeration.style`
**Key properties**: Color, numbering format, indent

---

**Use list when you need:**
- Unordered list rendering
- Bullet point display
- Custom bullet characters

**File**: `examples/list.md`
**Style**: `examples/list.style`
**Key properties**: Color, bullet style, indent

---

**Use ordered_list when you need:**
- Sequential numbered items
- Step-by-step procedures
- Hierarchical numbering

**File**: `examples/ordered_list.md`
**Style**: `examples/ordered_list.style`
**Key properties**: Numbering format, indent levels

---

**Use task when you need:**
- Checkbox list rendering
- TODO item display
- Completion tracking UI
- Custom tick/untick symbols

**File**: `examples/task.md`
**Style**: `examples/task.style`
**Key properties**: ticked/unticked characters, list styling

### Links and Media

**Use emoji when you need:**
- Unicode emoji rendering
- Symbol-based status indicators
- Visual markers in text

**File**: `examples/emoji.md`
**Style**: `examples/emoji.style`
**Key properties**: Rendering behavior

---

**Use image when you need:**
- Image placeholder display
- Alt-text rendering
- Media reference handling

**File**: `examples/image.md`
**Style**: `examples/image.style`
**Key properties**: Color, prefix/suffix, alt-text display

---

**Use link when you need:**
- Hyperlink formatting
- URL highlighting
- Link text differentiation
- Custom URL brackets

**File**: `examples/link.md`
**Style**: `examples/link.style`
**Key properties**: Color, underline, block_prefix/suffix, link_text styling

### Table Rendering

**Use table when you need:**
- Basic row/column display
- Header/body separation
- Structured data rendering

**File**: `examples/table.md`
**Style**: `examples/table.style`
**Key properties**: Margin, cell styling

---

**Use table_align when you need:**
- Column alignment control
- Left/center/right cell alignment
- Numeric data display

**File**: `examples/table_align.md`
**Style**: `examples/table_align.style`
**Key properties**: Alignment directives

---

**Use table_truncate when you need:**
- Width-constrained table display
- Content overflow handling
- Ellipsis truncation

**File**: `examples/table_truncate.md`
**Style**: `examples/table_truncate.style`
**Key properties**: Truncation behavior

---

**Use table_wrap when you need:**
- Multi-line cell content
- Text wrapping within cells
- Flexible width tables

**File**: `examples/table_wrap.md`
**Style**: `examples/table_wrap.style`
**Key properties**: Wrap behavior

---

**Use table_with_footer_links when you need:**
- Reference-style links in tables
- Footer link collection
- Deduplication of repeated URLs

**File**: `examples/table_with_footer_links.md`
**Style**: `examples/table_with_footer_links.style`
**Key properties**: Link footer rendering

---

**Use table_with_footer_links_complex when you need:**
- Multi-element table cells
- Combined text, links, and images
- Rich data table display

**File**: `examples/table_with_footer_links_complex.md`
**Style**: `examples/table_with_footer_links_complex.style`
**Key properties**: Complex cell rendering

---

**Use table_with_footer_links_huge when you need:**
- Large link collections
- Bulk reference management
- High-density link tables

**File**: `examples/table_with_footer_links_huge.md`
**Style**: `examples/table_with_footer_links_huge.style`
**Key properties**: Scalable link handling

---

**Use table_with_footer_links_no_color when you need:**
- Monochrome table output
- High-contrast displays
- Color-blind accessibility

**File**: `examples/table_with_footer_links_no_color.md`
**Style**: `examples/table_with_footer_links_no_color.style`
**Key properties**: Colorless rendering

---

**Use table_with_footer_links_repeated when you need:**
- Duplicate link reference handling
- Shared URL management
- Reference consolidation

**File**: `examples/table_with_footer_links_repeated.md`
**Style**: `examples/table_with_footer_links_repeated.style`
**Key properties**: Link deduplication

---

**Use table_with_inline_links when you need:**
- Direct URL embedding in cells
- No footer separation
- Compact link display

**File**: `examples/table_with_inline_links.md`
**Style**: `examples/table_with_inline_links.style`
**Key properties**: Inline link rendering

---

**Use table_with_footer_images when you need:**
- Image reference management
- Footer-based image links
- Alt-text handling

**File**: `examples/table_with_footer_images.md`
**Style**: `examples/table_with_footer_images.style`
**Key properties**: Image placeholder rendering

---

**Use table_with_footer_images_no_alt when you need:**
- Image handling without alt text
- URL-only image references
- Minimal image metadata

**File**: `examples/table_with_footer_images_no_alt.md`
**Style**: `examples/table_with_footer_images_no_alt.style`
**Key properties**: Alt-less image rendering

---

**Use table_with_footer_images_same_alt when you need:**
- Duplicate alt-text handling
- Multiple images with shared descriptions
- Alt-text deduplication

**File**: `examples/table_with_footer_images_same_alt.md`
**Style**: `examples/table_with_footer_images_same_alt.style`
**Key properties**: Shared alt-text management

---

**Use table_with_footer_auto_links when you need:**
- Automatic footer generation
- Link detection and extraction
- Auto-numbered references

**File**: `examples/table_with_footer_auto_links.md`
**Style**: `examples/table_with_footer_auto_links.style`
**Key properties**: Automatic link footers

---

**Use table_with_footer_auto_links_short when you need:**
- Compact auto-generated footers
- Shortened reference format
- Space-efficient link lists

**File**: `examples/table_with_footer_auto_links_short.md`
**Style**: `examples/table_with_footer_auto_links_short.style`
**Key properties**: Compact footer format

## Implementation Patterns

### Style Configuration Structure

Style files use JSON objects with element-specific configuration:

```json
{
  "element_name": {
    "color": "numeric_color_code",
    "background_color": "numeric_color_code",
    "margin": integer,
    "indent": integer,
    "prefix": "string",
    "suffix": "string",
    "bold": boolean,
    "italic": boolean,
    "underline": boolean
  }
}
```

### Renderer Initialization Patterns

**Simple rendering:**
```go
out, _ := glamour.Render(markdown, "dark")
```

**Custom renderer:**
```go
r, _ := glamour.NewTermRenderer(
    glamour.WithStandardStyle("dark"),
    glamour.WithWordWrap(80),
)
out, _ := r.Render(markdown)
```

**Environment-based:**
```go
r, _ := glamour.NewTermRenderer(
    glamour.WithEnvironmentConfig(),
    glamour.WithWordWrap(defaultWidth),
)
```

**Auto-style detection:**
```go
r, _ := glamour.NewTermRenderer(
    glamour.WithAutoStyle(),
)
```

**Color profile aware:**
```go
w := colorprofile.NewWriter(os.Stdout, os.Environ())
r, _ := glamour.NewTermRenderer(glamour.WithEnvironmentConfig())
md, _ := r.RenderBytes(input)
fmt.Fprintf(w, "%s\n", md)
```

### File Organization Pattern

Each markdown example follows:
- `{name}.md` - Markdown content to render
- `{name}.style` - JSON style configuration
- `{name}.png` - Screenshot of expected output (when available)

### Common Style Properties

- **Color codes**: Terminal 256-color numeric codes
- **Margins**: Integer spacing (vertical whitespace)
- **Indents**: Integer spacing (horizontal offset)
- **Prefixes/suffixes**: String decorations
- **Boolean flags**: bold, italic, underline, strikethrough

### Table Footer Reference System

Tables support reference-style links and images:
- Automatic detection and extraction of URLs
- Footer consolidation and deduplication
- Numbered or labeled references
- Combined text, link, and image cells

### Terminal Compatibility

- Auto-detection via `WithAutoStyle()`
- Environment variable configuration via `WithEnvironmentConfig()`
- Color profile detection via `colorprofile` package
- Automatic color downsampling for limited terminals
