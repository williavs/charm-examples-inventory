# Pop Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|--------------|
| Command-line email with attachments | `examples/cli.tape` |
| Interactive TUI email composition | `examples/demo.tape` |
| Dynamic recipient/sender selection | `examples/gum-example.tape` |
| File generation and attachment workflow | `examples/invoice-example.tape` |
| AI-generated content in emails | `examples/mods-example.tape` |

## Examples by Capability

### CLI Email Sending
**Use cli.tape when you need:**
- Command-line email sending with file attachments
- Piped content from files as email body
- Non-interactive scripted email workflows
- API key authentication pattern (RESEND_API_KEY)

**File**: `examples/cli.tape`
**Key patterns**:
- Stdin redirection (`pop < message.md`)
- Multiple flag usage (--from, --to, --subject, --attach)
- Markdown file as email body

### Interactive TUI Composition
**Use demo.tape when you need:**
- Full terminal UI email composition
- Step-by-step form-based input
- Markdown editing within TUI
- Visual email preview before sending

**File**: `examples/demo.tape`
**Key patterns**:
- No-argument invocation launches TUI
- Tab navigation between fields
- Inline markdown composition
- Multi-line message formatting

### Dynamic Input Selection
**Use gum-example.tape when you need:**
- Interactive sender selection from predefined list
- Recipient filtering from contact file
- Command substitution for email parameters
- Integration with selection tools (gum choose/filter)

**File**: `examples/gum-example.tape`
**Key patterns**:
- Command substitution with `$(gum choose)`
- Contact file filtering with `$(gum filter < contacts.txt)`
- Dynamic flag population from user selection

### File Attachment Workflows
**Use invoice-example.tape when you need:**
- Generated file attachment patterns
- Command chaining for document creation and sending
- PDF attachment workflows
- Inline body text with attachments

**File**: `examples/invoice-example.tape`
**Key patterns**:
- Sequential command execution (generate then send)
- --attach flag for PDF files
- --body flag for inline message text
- Integration with document generation tools

### AI Content Generation
**Use mods-example.tape when you need:**
- AI-generated email content
- LLM integration in email workflows
- Preview mode for generated content
- Here-string input from command output

**File**: `examples/mods-example.tape`
**Key patterns**:
- Here-string syntax (`pop <<< "$(mods ...)"`)
- AI API key authentication (OPENAI_API_KEY)
- --preview flag for content review
- Command substitution for dynamic body content

## Implementation Patterns

### Authentication
All examples use environment variable pattern for API keys:
```bash
export RESEND_API_KEY=$(pass RESEND_CHARM_API_KEY)
```

### Input Methods
- **Stdin redirection**: `pop < file.md`
- **Here-string**: `pop <<< "content"`
- **Flag-based**: `--body "content"`
- **TUI entry**: No arguments, interactive mode

### Attachment Handling
- Single file: `--attach filename.ext`
- Multiple files: Chain multiple `--attach` flags
- Generated files: Create first, then attach in second command

### Recipient Patterns
- Static: `--to 'email@domain.com'`
- Dynamic selection: `--to "$(gum filter < contacts.txt)"`
- Multiple recipients: Comma-separated or multiple flags

## Support Files

- `contacts.txt`: Sample contact list for filtering examples
- `message.md`: Sample markdown email content
- `README.md`: Minimal project identifier

## Tape File Structure

All `.tape` files follow VHS format:
- Output configuration (GIF settings)
- Hide/Show blocks for API key setup
- Command demonstrations
- Timing controls (Sleep, Enter)

Use these as templates for creating similar demonstration recordings.
