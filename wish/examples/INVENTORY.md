# Wish Examples Inventory

## Overview
This directory contains 15 Go examples demonstrating SSH server implementations using the Wish framework. The examples progress from basic SSH servers to complex applications including TUI interfaces, file transfer capabilities, and Git repository hosting. All examples use the charmbracelet/wish library for SSH server functionality.

## Examples Catalog

### Basic SSH Servers
- **simple** - Minimal SSH server that prints "Hello, world!" to connecting clients. Uses ED25519 key authentication and basic middleware. Single file implementation.
- **graceful-shutdown** - SSH server with proper signal handling for clean shutdown. Demonstrates context-based shutdown with 30-second timeout.
- **banner** - SSH server that displays a custom banner from banner.txt file before allowing user interaction. Shows middleware for pre-session content.

### Authentication and Identity
- **identity** - SSH server that identifies users by their public keys. Maintains a hardcoded map of known users and greets them by name. Demonstrates key-based user identification.
- **multi-auth** - SSH server supporting multiple authentication methods. Shows implementation of both public key and password authentication mechanisms.

### Interactive Applications
- **bubbletea** - SSH server hosting a Bubble Tea TUI application. Displays terminal information (size, color profile, background) in real-time. Uses alt screen mode.
- **bubbleteaprogram** - SSH server that runs complete Bubble Tea programs over SSH connections. Demonstrates program lifecycle management.
- **bubbletea-exec** - Combines Bubble Tea interface with PTY allocation for running external programs. Shows interaction between TUI and shell execution.
- **multichat** - Multi-user chat application over SSH. Supports concurrent users in shared chat rooms with real-time message broadcasting. Uses textarea and viewport components.

### Command Line Integration
- **cobra** - SSH server that executes spf13/cobra CLI commands. Implements an echo command with reverse flag option. Shows integration of CLI frameworks over SSH.

### File Transfer and Repository Management
- **scp** - SSH server with SCP and SFTP support. Serves files from testdata directory with read-only filesystem handler. Implements secure file transfer protocols.
- **git** - SSH server for Git repository hosting. Provides repository listing interface and handles Git operations (clone, push, fetch). Stores repositories in .repos directory.

### System Integration
- **forward** - SSH server with reverse port forwarding capabilities. Enables network tunneling through SSH connections.
- **pty** - SSH server that allocates pseudo-terminals for interactive shell sessions. Demonstrates PTY handling for terminal applications.
- **exec** - SSH server that executes shell commands and scripts. Includes example.sh script for testing command execution capabilities.

## Implementation Patterns

### Common Architecture
- All examples use net.JoinHostPort for address binding
- Standard ED25519 key generation and management
- Middleware-based request processing pipeline
- Context-based graceful shutdown with signal handling
- Structured logging using charmbracelet/log

### Authentication Patterns
- Public key authentication with ED25519 keys
- User identification through key fingerprinting
- Access control through hardcoded user maps
- Session-based authentication state management

### TUI Integration
- Bubble Tea model-view-update architecture
- Session-specific renderer creation for color profiles
- Alt screen mode for fullscreen applications
- Real-time message passing between concurrent sessions

### File System Operations
- Filesystem abstraction for SCP/SFTP handlers
- Directory traversal with security boundaries
- Read-only and read-write access patterns
- Git repository management through filesystem

## Practical Applications

### NextJS Scaffolder with Auth/ShadCN
- **bubbletea** pattern can create interactive project scaffolding TUIs
- **identity** pattern provides SSH key-based authentication for secure access
- **multichat** pattern enables real-time collaboration during setup
- **exec** pattern allows running npm/yarn commands during scaffolding

### Inter-Agent Communication Systems
- **multichat** provides foundation for agent-to-agent messaging
- **forward** enables secure tunneling between agent environments
- **bubbletea** creates monitoring dashboards for agent status
- **git** pattern manages agent configuration repositories

### GitHub Search Tools
- **cobra** pattern integrates GitHub CLI commands over SSH
- **bubbletea** creates interactive search interfaces
- **scp** enables downloading search results and repositories
- **identity** provides secure access to organization repositories

### TUI for App Distribution/Package Management
- **bubbletea** creates package browser and installer interfaces
- **scp** handles package file transfers
- **exec** runs installation commands and scripts
- **git** manages package repository hosting
- **multichat** enables real-time installation status sharing

### Technical Specifications
- **File count**: 22 total files (15 Go source files, 3 configuration files, 4 data files)
- **Language**: Go 1.24+ with modern toolchain
- **Dependencies**: 67 external packages including Bubble Tea, SSH, and Git libraries
- **Architecture**: Event-driven SSH servers with middleware pipeline processing
- **Authentication**: SSH public key and password authentication support
- **Networking**: TCP listeners with configurable host/port binding