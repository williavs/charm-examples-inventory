# Go-GenAI Examples Inventory

## Overview
Collection of 52 files (30 Go source files) demonstrating Google GenAI library functionality across 9 categories. Examples cover chat interfaces, file operations, streaming, tool integration, and multimodal content generation. All examples support both Vertex AI and Gemini API backends through environment variable configuration.

## Examples Catalog

### Chat Systems (`chats/`)
- **chat.go**: Basic chat session implementation with temperature configuration
- **chat_stream.go**: Streaming chat with real-time response processing
- **Files**: 2 Go files, 1 README
- **Function**: Creates persistent chat sessions, handles streaming responses

### Live Streaming (`live/`)
- **live_streaming_server.go**: WebSocket server for real-time AI interactions (6,881 lines)
- **live_streaming.html**: Frontend interface with WebSocket client implementation
- **Files**: 1 Go file, 1 HTML file, 1 README
- **Function**: Real-time bidirectional communication between browser and AI model via WebSockets

### File Management (`files/`)
- **upload_file.go**: File upload and processing pipeline
- **list_download.go**: File enumeration and retrieval operations
- **Files**: 2 Go files, 1 README
- **Function**: Handles file uploads to AI service, downloads processed results

### Tool Integration (`mcptoolbox/`)
- **mcp_toolbox.go**: MCP Toolbox SDK integration for database operations
- **go.mod/go.sum**: Dependencies for MCP integration
- **Files**: 1 Go file, 2 module files, 1 README
- **Function**: Converts MCP tools to GenAI-compatible format, handles database tool invocation

### Content Generation (`models/generate_content/`)
- **function_declaration_json_schema.go**: JSON schema-based tool definitions
- **function_declaration_schema.go**: Programmatic tool schema creation
- **text.go**: Basic text generation
- **text_stream.go**: Streaming text generation
- **image_modality.go**: Image analysis and processing
- **image_modality_stream.go**: Streaming image processing
- **audio_modality.go**: Audio content analysis
- **Files**: 7 Go files, 1 README
- **Function**: Demonstrates multimodal content generation with tool calling

### Token Operations (`models/count_tokens/`, `models/compute_tokens/`)
- **text_tokens.go**: Token counting and computation utilities
- **Files**: 1 Go file per directory, 2 READMEs
- **Function**: Token usage analysis and cost estimation

### Image Processing (`models/edit_image/`, `models/recontext_image/`, `models/segment_image/`, `models/upscale_image/`)
- **image.go**: Core image manipulation functions
- **ingredients.go**: Recipe-based image editing
- **segment_image.go**: Image segmentation operations
- **Files**: 4-6 Go files across directories
- **Function**: Image editing, enhancement, segmentation, and upscaling

### Media Generation (`models/generate_images/`, `models/generate_videos/`)
- **images.go**: Image generation from text prompts
- **videos.go**: Video generation from text descriptions
- **Files**: 1 Go file per directory, 2 READMEs
- **Function**: Creates visual media from text descriptions

### Content Processing (`models/embed_content/`)
- **text.go**: Text embedding generation for semantic analysis
- **Files**: 1 Go file, 1 README
- **Function**: Converts text to vector embeddings

### Batch Operations (`batches/`)
- **create_get_cancel.go**: Batch job management operations
- **Files**: 1 Go file, 1 README
- **Function**: Handles asynchronous batch processing tasks

### Caching System (`caches/`)
- **create_get_delete.go**: Cache management operations
- **list.go**: Cache enumeration functionality
- **Files**: 2 Go files, 1 README
- **Function**: Content caching for performance optimization

### Model Tuning (`tunings/`)
- **create_get.go**: Model fine-tuning operations
- **list.go**: Tuning job listing and management
- **Files**: 2 Go files, 1 README
- **Function**: Custom model training and management

### Resource Management (`models/resource_methods/`)
- **get.go**: Resource retrieval operations
- **list.go**: Resource enumeration
- **Files**: 2 Go files, 1 README
- **Function**: API resource management and inspection

## Implementation Patterns

### Authentication Configuration
- Environment-based backend selection (Vertex AI vs Gemini API)
- API key and project configuration through environment variables
- Consistent authentication pattern across all examples

### Client Initialization
- Standard `genai.NewClient(ctx, config)` pattern
- Backend detection and appropriate initialization
- Error handling with fatal logging

### Content Generation
- Temperature-based response control
- Streaming vs non-streaming response handling
- Multimodal input processing (text, image, audio)

### Tool Integration
- JSON schema-based tool definitions
- Function declaration conversion utilities
- Tool execution and response handling

### WebSocket Communication
- Gorilla WebSocket for real-time communication
- Bidirectional message proxying
- Embedded HTML templates

### Error Handling
- Consistent error checking with `log.Fatal()`
- Context-based cancellation support
- Resource cleanup patterns

## Practical Applications

### NextJS Scaffolder with Auth/Shadcn
- **Relevant**: `live/` for real-time scaffolding feedback
- **Relevant**: `models/generate_content/function_declaration_*` for component generation tools
- **Adaptation**: Use streaming responses for build progress, tool definitions for template generation

### Inter-Agent Communication Systems
- **Relevant**: `live/` WebSocket implementation for real-time agent communication
- **Relevant**: `chats/` for persistent conversation management
- **Adaptation**: Replace AI model with agent message routing, maintain WebSocket infrastructure

### GitHub Search Tools
- **Relevant**: `mcptoolbox/` for external API integration patterns
- **Relevant**: `models/generate_content/function_declaration_*` for search tool definitions
- **Adaptation**: Replace database tools with GitHub API tools, use schema patterns for search parameters

### TUI for App Distribution/Package Management
- **Relevant**: `files/` for package upload/download operations
- **Relevant**: `batches/` for asynchronous package processing
- **Relevant**: `caches/` for package caching mechanisms
- **Adaptation**: Replace AI file operations with package management, use batch patterns for installations

### Cross-Pattern Reuse
- WebSocket server architecture from `live/` applicable to any real-time interface
- Tool conversion utilities from `mcptoolbox/` adaptable to any external API
- Streaming patterns from multiple examples useful for progress indication
- File management patterns applicable to any content distribution system