# Pop Examples Inventory

## Overview
This directory contains 8 files demonstrating the "pop" command-line email client functionality. The examples focus on email composition, sending, and integration with other CLI tools. All demonstration scripts are written in VHS tape format for terminal recording.

## Examples Catalog

### Demonstration Scripts (.tape files)
- **cli.tape** (387 bytes): Command-line email sending with file attachment. Shows piping markdown content to pop with explicit from/to/subject parameters and README.md attachment.
- **demo.tape** (596 bytes): Interactive TUI workflow demonstration. Records step-by-step email composition using pop's form interface for sender, recipient, subject, and message body.
- **gum-example.tape** (404 bytes): Integration with gum selection tools. Demonstrates dynamic recipient selection from contact list and sender choice using gum's choose and filter commands.
- **invoice-example.tape** (366 bytes): Business workflow integration. Shows PDF invoice generation followed by email attachment using pop with custom body text.
- **mods-example.tape** (469 bytes): AI integration demonstration. Uses mods (AI text generation) to create email content and sends via pop with preview mode enabled.

### Supporting Files
- **contacts.txt** (148 bytes): Sample email address list containing 10 @charm.sh addresses for use with contact selection examples.
- **message.md** (63 bytes): Template markdown email content for demonstrations. Contains basic formatting with bold and italic text.
- **README.md** (6 bytes): Minimal documentation containing only "# Pop" header.

## Implementation Patterns

### Email Workflow Patterns
- **Pipeline Integration**: Input redirection and command chaining for automated email sending
- **Interactive Forms**: TUI-based step-through email composition with tab navigation
- **External Tool Integration**: Combining pop with gum for contact selection and mods for content generation
- **File Attachment**: Direct file attachment capability with various document types

### Demonstration Patterns
- **VHS Recording Configuration**: Consistent framerate (28-30 fps), padding (40-70px), and resolution settings across all demos
- **Environment Setup**: API key management using pass command for secure credential handling
- **User Interaction Simulation**: Keyboard input timing, navigation delays, and visual presentation sequences

## Practical Applications

### For NextJS Scaffolder with Auth/Shadcn
- **Email Notification System**: Pop's command-line integration could send deployment status emails after scaffolding completion
- **User Onboarding**: Automated welcome email sending for new project setups with attachment capability for documentation
- **Template Distribution**: Email delivery mechanism for sharing scaffolded project templates

### For Inter-Agent Communication Systems
- **Agent Status Reporting**: Email-based notification system for cross-agent communication when direct messaging fails
- **Error Notification**: Automated email alerts for critical agent failures or system issues
- **Command Coordination**: Email-based task assignment between distributed agents

### For GitHub Search Tools
- **Search Result Distribution**: Email delivery of search results with attachment capability for exporting findings
- **Repository Monitoring**: Automated notifications for repository changes or issue updates
- **Team Coordination**: Email-based sharing of GitHub search results with stakeholders

### For TUI App Distribution/Package Management
- **Update Notifications**: Email alerts for available package updates or new releases
- **Installation Reports**: Automated installation status emails with logs attached
- **User Communication**: Email-based user support and notification system for package management operations

## Technical Integration Points
- **API Integration**: Requires RESEND_API_KEY environment variable for email service backend
- **CLI Chaining**: Compatible with Unix pipeline operations and command substitution
- **File System**: Direct file attachment support for documentation and report distribution
- **External Dependencies**: Integrates with gum (UI components), mods (AI text), and invoice (document generation) tools