# Anthropic SDK Go Examples Inventory

## Overview
This directory contains 12 example implementations demonstrating the Anthropic SDK for Go. The examples cover core messaging functionality, streaming responses, tool integration, multimodal inputs, file uploads, and platform-specific implementations for AWS Bedrock and Google Vertex AI.

## Examples Catalog

### Core Messaging
- **message** - Basic message API implementation with simple text input/output
- **message-streaming** - Streaming message responses with real-time token processing
- **message-mcp-streaming** - Model Context Protocol integration with external server communication

### Tool Integration
- **tools** - Function calling with three example tools (coordinates, weather, temperature units)
- **tools-streaming** - Streaming function calls with real-time tool execution feedback
- **tools-streaming-jsonschema** - Tool calling with automatic JSON schema generation from Go structs

### Multimodal & File Handling
- **multimodal** - Image analysis with base64 encoding for PNG input processing
- **file-upload** - Document upload via beta Files API with text analysis capabilities

### Platform Integrations
- **bedrock** - AWS Bedrock integration with standard message processing
- **bedrock-streaming** - AWS Bedrock with streaming response handling
- **vertex** - Google Vertex AI integration for cloud deployment
- **vertex-streaming** - Google Vertex AI with streaming capabilities

## Implementation Patterns

### Common Structures
- Context-based request handling using `context.TODO()`
- Anthropic client initialization with optional configuration
- Message parameter construction using builder patterns
- Error handling with panic-based failure management
- Color-coded terminal output for user/assistant distinction

### Tool Implementation
- Interface-based tool parameter definition
- JSON marshaling/unmarshaling for tool inputs/outputs
- Loop-based conversation flow with tool result injection
- Hardcoded mock responses for demonstration purposes

### Stream Processing
- Event-based streaming with type switching
- Message accumulation for complete response reconstruction
- Real-time output rendering during response generation
- Error propagation through stream error checking

### Platform Adaptations
- AWS SDK v2 integration for Bedrock access
- Model name variations for platform-specific endpoints
- Configuration loading for cloud service authentication
- Service-specific parameter formatting

## Practical Applications

### NextJS Scaffolder with Auth/ShadCN
- **tools** pattern can generate project scaffolding commands
- **file-upload** can analyze existing project structures
- **streaming** patterns provide real-time generation feedback
- Tool definitions can be adapted for file system operations

### Inter-Agent Communication Systems
- **tools-streaming** provides foundation for agent command execution
- **message-streaming** enables real-time agent communication
- **mcp-streaming** demonstrates external service integration
- JSON schema patterns support typed inter-agent messaging

### GitHub Search Tools
- **tools** framework can be adapted for GitHub API integration
- **file-upload** can process repository files for analysis
- **streaming** patterns provide progressive search results
- Tool parameter validation supports complex query structures

### TUI for App Distribution/Package Management
- **tools** can execute package management commands
- **streaming** provides real-time installation feedback
- **multimodal** can process package documentation images
- **file-upload** can analyze package manifests and dependencies

## Technical Specifications

### File Count
- 17 total files across 12 directories
- 12 Go source files (main.go in each example)
- 3 supporting files (go.mod, go.sum, nine_dogs.png)
- 2 data files (file.txt, .keep)

### Dependencies
- Go 1.23.0+ with toolchain 1.24.3
- anthropic-sdk-go (local development version)
- jsonschema package for automatic schema generation
- AWS SDK v2 for Bedrock integration
- Google Cloud libraries for Vertex AI support

### Code Characteristics
- Straightforward procedural implementations
- Minimal external dependencies beyond SDK
- Demonstration-focused with hardcoded responses
- Standard Go idioms and error handling patterns
- Terminal-based output with basic formatting