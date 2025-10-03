# Wish Examples - Contextual Reference

## Quick Reference

| Need | Example | File |
|------|---------|------|
| Basic SSH server | simple | `examples/simple/main.go` |
| Server with graceful shutdown | graceful-shutdown | `examples/graceful-shutdown/main.go` |
| Custom banner and password auth | banner | `examples/banner/main.go` |
| Public key authentication | identity | `examples/identity/main.go` |
| Multiple authentication methods | multi-auth | `examples/multi-auth/main.go` |
| CLI with flags over SSH | cobra | `examples/cobra/main.go` |
| Bubbletea TUI over SSH | bubbletea | `examples/bubbletea/main.go` |
| Custom Bubbletea middleware | bubbleteaprogram | `examples/bubbleteaprogram/main.go` |
| Reverse port forwarding | forward | `examples/forward/main.go` |
| Multi-user chat server | multichat | `examples/multichat/main.go` |
| Git repository hosting | git | `examples/git/main.go` |
| SCP and SFTP file transfer | scp | `examples/scp/main.go` |
| PTY allocation | pty | `examples/pty/main.go` |
| Bubbletea + external program execution | bubbletea-exec | `examples/bubbletea-exec/main.go` |
| External command execution | exec | `examples/exec/main.go` |

## Examples by Capability

### Server Basics

**Use simple when you need:**
- Minimal SSH server setup
- Basic middleware chain structure
- "Hello, world" response to connections

**File**: `examples/simple/main.go`
**Key patterns**: NewServer, WithAddress, WithHostKeyPath, basic middleware

**Use graceful-shutdown when you need:**
- Signal handling for clean server shutdown
- Context-based timeout for active connections
- Server lifecycle management

**File**: `examples/graceful-shutdown/main.go`
**Key patterns**: signal.Notify, context.WithTimeout, srv.Shutdown

### Authentication

**Use banner when you need:**
- Custom banner text before authentication
- Password authentication
- Session duration tracking middleware

**File**: `examples/banner/main.go`
**Key patterns**: WithBannerHandler, WithPasswordAuth, elapsed.Middleware

**Use identity when you need:**
- Public key authentication by key type
- User identification from public key
- Known user mapping

**File**: `examples/identity/main.go`
**Key patterns**: WithPublicKeyAuth, ssh.KeysEqual, ssh.ParseAuthorizedKey

**Use multi-auth when you need:**
- Multiple authentication method support
- Keyboard-interactive challenges
- Authentication fallback chain

**File**: `examples/multi-auth/main.go`
**Key patterns**: WithPublicKeyAuth, WithPasswordAuth, WithKeyboardInteractiveAuth

### CLI and TUI Applications

**Use cobra when you need:**
- CLI command parsing over SSH
- Flag-based argument handling
- Command execution with session I/O wiring

**File**: `examples/cobra/main.go`
**Key patterns**: cobra.Command, SetArgs, SetIn/SetOut/SetErr

**Use bubbletea when you need:**
- Full-screen TUI over SSH
- Session-specific styling and rendering
- Terminal capability detection

**File**: `examples/bubbletea/main.go`
**Key patterns**: bubbletea.Middleware, MakeRenderer, activeterm.Middleware

**Use bubbleteaprogram when you need:**
- Custom Bubbletea program creation
- Program-level message injection
- Advanced program options configuration

**File**: `examples/bubbleteaprogram/main.go`
**Key patterns**: MiddlewareWithProgramHandler, tea.NewProgram, custom middleware

### Multi-User Applications

**Use multichat when you need:**
- Multi-session state coordination
- Broadcasting messages to all sessions
- Shared program list management

**File**: `examples/multichat/main.go`
**Key patterns**: app struct with progs slice, p.Send for broadcasting, session tracking

### File Operations

**Use git when you need:**
- Git protocol over SSH
- Repository access control
- Push/fetch hooks

**File**: `examples/git/main.go`
**Key patterns**: git.Middleware, git.AccessLevel, GitHooks interface

**Use scp when you need:**
- SCP file transfer
- SFTP subsystem
- Filesystem access control

**File**: `examples/scp/main.go`
**Key patterns**: scp.Middleware, WithSubsystem, NewFileSystemHandler

### Advanced Session Management

**Use forward when you need:**
- Reverse TCP port forwarding
- Port binding callbacks
- Request handler registration

**File**: `examples/forward/main.go`
**Key patterns**: ForwardedTCPHandler, ReversePortForwardingCallback, RequestHandlers

**Use pty when you need:**
- Pseudoterminal allocation
- Terminal capability detection
- PTY metadata access

**File**: `examples/pty/main.go`
**Key patterns**: ssh.AllocatePty, sess.Pty, terminal background detection

**Use bubbletea-exec when you need:**
- External program execution from TUI
- Editor integration over SSH
- Shell spawning within session

**File**: `examples/bubbletea-exec/main.go`
**Key patterns**: tea.Exec, wish.Command, editor.Cmd

**Use exec when you need:**
- Direct command execution on connection
- Script running with session I/O
- Non-interactive command handling

**File**: `examples/exec/main.go`
**Key patterns**: wish.Command, cmd.Run, activeterm.Middleware

## Implementation Patterns

### Server Initialization
All examples follow the pattern:
```go
srv, err := wish.NewServer(
    wish.WithAddress(net.JoinHostPort(host, port)),
    wish.WithHostKeyPath(".ssh/id_ed25519"),
    wish.WithMiddleware(...),
)
```

### Middleware Chain
Middleware executes in reverse order (last defined runs first):
```go
wish.WithMiddleware(
    customHandler,        // runs third
    loggingMiddleware,    // runs second
    authMiddleware,       // runs first
)
```

### Session Handling
Sessions provide access to:
- `sess.User()` - authenticated username
- `sess.Command()` - SSH command arguments
- `sess.PublicKey()` - user's public key
- `sess.Pty()` - PTY information
- Session I/O via Reader/Writer interface

### Graceful Shutdown Pattern
Standard approach across examples:
```go
done := make(chan os.Signal, 1)
signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
go func() { srv.ListenAndServe() }()
<-done
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
srv.Shutdown(ctx)
```

### Bubbletea Integration
Handler functions create models per session:
```go
func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
    renderer := bubbletea.MakeRenderer(s)
    // Use renderer for session-specific styling
    return model, []tea.ProgramOption{tea.WithAltScreen()}
}
```

### Authentication Callbacks
Return boolean for auth success:
```go
wish.WithPublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
    return validateKey(key)
})
```

### Command Execution
Wire session I/O to external commands:
```go
cmd := wish.Command(sess, "program", "arg1", "arg2")
err := cmd.Run()
```

## Common Dependencies

- `github.com/charmbracelet/wish` - Core SSH server
- `github.com/charmbracelet/ssh` - SSH implementation (forked)
- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/log` - Structured logging
- `github.com/charmbracelet/lipgloss` - Terminal styling
- `github.com/spf13/cobra` - CLI commands (cobra example)
- `github.com/pkg/sftp` - SFTP protocol (scp example)

## Notes

- All examples use port 23234 except: git (23233), scp (23235), cobra (23235)
- Host key path defaults to `.ssh/id_ed25519`
- PTY allocation required for: bubbletea, pty, bubbletea-exec, exec
- activeterm.Middleware ensures PTY availability for TUI applications
- Bubbletea renderers must be created per-session for correct color profiles
- Windows has limited PTY support (pseudoconsole)
