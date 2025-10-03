# Go-GenAI Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|-------------|
| WebSocket-based real-time communication | `examples/live/live_streaming_server.go` |
| Persistent conversation sessions | `examples/chats/chat.go` |
| Streaming text responses | `examples/chats/chat_stream.go` |
| File upload and processing | `examples/files/upload_file.go` |
| File enumeration and retrieval | `examples/files/list_download.go` |
| External API tool integration | `examples/mcptoolbox/mcp_toolbox.go` |
| Tool definitions with JSON schema | `examples/models/generate_content/function_declaration_json_schema.go` |
| Programmatic tool schema creation | `examples/models/generate_content/function_declaration_schema.go` |
| Basic text generation | `examples/models/generate_content/text.go` |
| Streaming text generation | `examples/models/generate_content/text_stream.go` |
| Image content analysis | `examples/models/generate_content/image_modality.go` |
| Streaming image processing | `examples/models/generate_content/image_modality_stream.go` |
| Audio content analysis | `examples/models/generate_content/audio_modality.go` |
| Token counting and cost estimation | `examples/models/count_tokens/text_tokens.go` |
| Token computation utilities | `examples/models/compute_tokens/text_tokens.go` |
| Image editing and manipulation | `examples/models/edit_image/image.go` |
| Recipe-based image modification | `examples/models/edit_image/ingredients.go` |
| Image segmentation | `examples/models/segment_image/segment_image.go` |
| Image enhancement and upscaling | `examples/models/upscale_image/image.go` |
| Image context modification | `examples/models/recontext_image/image.go` |
| Text-to-image generation | `examples/models/generate_images/images.go` |
| Text-to-video generation | `examples/models/generate_videos/videos.go` |
| Text embedding generation | `examples/models/embed_content/text.go` |
| Asynchronous batch processing | `examples/batches/create_get_cancel.go` |
| Content caching mechanisms | `examples/caches/create_get_delete.go` |
| Cache enumeration | `examples/caches/list.go` |
| Model fine-tuning operations | `examples/tunings/create_get.go` |
| Tuning job management | `examples/tunings/list.go` |
| Model resource retrieval | `examples/models/resource_methods/get.go` |
| Model resource listing | `examples/models/resource_methods/list.go` |

## Examples by Capability

### Real-Time Communication
**Use live streaming when you need:**
- WebSocket-based bidirectional communication
- Real-time message proxying between client and service
- Embedded HTML template serving
- Goroutine-based concurrent message handling

**File**: `examples/live/live_streaming_server.go`
**Key patterns**: WebSocket upgrade, goroutine message routing, GenAI session management, backend selection logic

### Conversational Interfaces
**Use chat sessions when you need:**
- Persistent conversation context across multiple messages
- Temperature-controlled response generation
- Sequential message-response interactions
- Chat history management

**File**: `examples/chats/chat.go`
**Key patterns**: Chat session creation, context preservation, backend detection

**Use streaming chat when you need:**
- Real-time response streaming as text generates
- Progressive content display
- Token-by-token output handling

**File**: `examples/chats/chat_stream.go`
**Key patterns**: Stream iteration, incremental text accumulation

### File Operations
**Use file upload when you need:**
- Multipart file transmission to GenAI service
- SHA256 checksum calculation
- File metadata handling
- Upload progress tracking

**File**: `examples/files/upload_file.go`
**Key patterns**: File reader creation, checksum computation, upload configuration

**Use file management when you need:**
- Enumeration of uploaded files
- File metadata retrieval
- Download operations
- File listing with filters

**File**: `examples/files/list_download.go`
**Key patterns**: File iteration, download streaming, metadata extraction

### External Tool Integration
**Use MCP toolbox integration when you need:**
- Conversion between external tool schemas and GenAI format
- Dynamic tool registration from external services
- Tool invocation with parameter mapping
- External API abstraction

**File**: `examples/mcptoolbox/mcp_toolbox.go`
**Key patterns**: Schema conversion, tool mapping, invocation proxying, response handling

### Function Calling and Tools
**Use JSON schema tools when you need:**
- Map-based parameter definition
- Required field specification
- Type validation through schema
- Complex nested parameter structures

**File**: `examples/models/generate_content/function_declaration_json_schema.go`
**Key patterns**: `ParametersJsonSchema` usage, map-based schema construction

**Use programmatic schema tools when you need:**
- Structured parameter definition with `genai.Schema`
- Type-safe schema construction
- Property definition with descriptions
- Required field arrays

**File**: `examples/models/generate_content/function_declaration_schema.go`
**Key patterns**: `genai.Schema` struct usage, property definition

### Content Generation
**Use text generation when you need:**
- Single-shot text completion
- Temperature-controlled output
- Non-streaming responses
- Synchronous generation

**File**: `examples/models/generate_content/text.go`
**Key patterns**: `GenerateContent` call, temperature configuration, result extraction

**Use streaming generation when you need:**
- Progressive text delivery
- Real-time output display
- Token-by-token processing
- Reduced perceived latency

**File**: `examples/models/generate_content/text_stream.go`
**Key patterns**: `GenerateContentStream` call, iterator usage, incremental processing

**Use image analysis when you need:**
- Visual content understanding
- Image-to-text conversion
- Visual question answering
- Multimodal input processing

**File**: `examples/models/generate_content/image_modality.go`
**Key patterns**: Image blob creation, multimodal part construction, inline data handling

**Use audio processing when you need:**
- Audio content transcription
- Audio-to-text conversion
- Audio analysis and understanding
- Multimodal audio input

**File**: `examples/models/generate_content/audio_modality.go`
**Key patterns**: Audio blob creation, base64 encoding, MIME type handling

### Token Management
**Use token counting when you need:**
- Cost estimation before generation
- Input size validation
- Rate limiting calculations
- Usage tracking

**File**: `examples/models/count_tokens/text_tokens.go`
**Key patterns**: `CountTokens` call, token count extraction

**Use token computation when you need:**
- Detailed token analysis
- Token-to-cost mapping
- Budget enforcement
- Usage analytics

**File**: `examples/models/compute_tokens/text_tokens.go`
**Key patterns**: `ComputeTokens` call, detailed metrics extraction

### Image Processing
**Use image editing when you need:**
- Programmatic image modification
- Edit instruction processing
- Mask-based editing
- Transformation operations

**File**: `examples/models/edit_image/image.go`
**Key patterns**: Edit request construction, mask specification, output handling

**Use image segmentation when you need:**
- Object isolation in images
- Region identification
- Mask generation
- Spatial analysis

**File**: `examples/models/segment_image/segment_image.go`
**Key patterns**: Segmentation request creation, region extraction

**Use image upscaling when you need:**
- Resolution enhancement
- Quality improvement
- Detail preservation
- Enlargement operations

**File**: `examples/models/upscale_image/image.go`
**Key patterns**: Upscale configuration, quality parameters

**Use image recontextualization when you need:**
- Background modification
- Context replacement
- Scene transformation
- Environmental changes

**File**: `examples/models/recontext_image/image.go`
**Key patterns**: Context specification, transformation parameters

### Media Generation
**Use image generation when you need:**
- Text-to-image synthesis
- Visual content creation from descriptions
- Concept visualization
- Asset generation

**File**: `examples/models/generate_images/images.go`
**Key patterns**: Image generation request, prompt engineering, output retrieval

**Use video generation when you need:**
- Text-to-video synthesis
- Motion content creation
- Sequential frame generation
- Animated content production

**File**: `examples/models/generate_videos/videos.go`
**Key patterns**: Video generation configuration, long-form content handling

### Embeddings
**Use text embeddings when you need:**
- Semantic similarity calculations
- Vector database population
- Document clustering
- Search relevance scoring

**File**: `examples/models/embed_content/text.go`
**Key patterns**: Embedding request creation, vector extraction

### Batch Processing
**Use batch operations when you need:**
- Asynchronous bulk processing
- Long-running job management
- Cost-optimized batch inference
- GCS-based input/output

**File**: `examples/batches/create_get_cancel.go`
**Key patterns**: Batch job creation, status polling, cancellation handling, GCS URI configuration

### Caching
**Use content caching when you need:**
- Response memoization
- Repeated query optimization
- Latency reduction
- Cost savings on duplicate requests

**File**: `examples/caches/create_get_delete.go`
**Key patterns**: Cache creation, retrieval, deletion, TTL configuration

**Use cache management when you need:**
- Cache inventory
- Resource cleanup
- Cache enumeration
- Usage monitoring

**File**: `examples/caches/list.go`
**Key patterns**: Cache listing, metadata inspection

### Model Tuning
**Use fine-tuning when you need:**
- Custom model training
- Domain-specific adaptation
- Behavior modification
- Performance optimization

**File**: `examples/tunings/create_get.go`
**Key patterns**: Tuning job creation, training configuration, checkpoint retrieval

**Use tuning management when you need:**
- Job status monitoring
- Training progress tracking
- Model version control
- Job enumeration

**File**: `examples/tunings/list.go`
**Key patterns**: Job listing, status inspection

### Resource Management
**Use model retrieval when you need:**
- Model metadata inspection
- Capability discovery
- Version information
- Configuration details

**File**: `examples/models/resource_methods/get.go`
**Key patterns**: Model get operation, metadata extraction

**Use resource listing when you need:**
- Available model enumeration
- Capability survey
- Version discovery
- Resource inventory

**File**: `examples/models/resource_methods/list.go`
**Key patterns**: List operation, pagination handling

## Implementation Patterns

### Client Initialization
All examples follow a consistent initialization pattern:
```go
ctx := context.Background()
client, err := genai.NewClient(ctx, nil)
// Backend auto-detected from environment:
// - Vertex AI if GOOGLE_CLOUD_PROJECT set
// - Gemini API if GOOGLE_API_KEY set
```

### Backend Detection
Most examples include backend-aware logic:
```go
if client.ClientConfig().Backend == genai.BackendVertexAI {
    // Vertex AI specific configuration
} else {
    // Gemini API specific configuration
}
```

### Error Handling
Consistent error handling across all examples:
```go
if err != nil {
    log.Fatal(err)  // Fatal logging for critical errors
}
```

### Configuration Objects
Examples use configuration objects for fine-tuning:
```go
config := &genai.GenerateContentConfig{
    Temperature: genai.Ptr[float32](0.5),
    Tools: tools,
    // Additional parameters...
}
```

### Streaming Patterns
Streaming examples use iterator-based consumption:
```go
stream, err := client.Models.GenerateContentStream(ctx, model, prompt, config)
for {
    response, err := stream.Recv()
    if err == io.EOF {
        break
    }
    // Process response incrementally
}
```

### Goroutine Concurrency
Real-time examples use goroutines for bidirectional communication:
```go
go func() {
    for {
        message, err := session.Receive()
        // Process incoming messages
    }
}()
// Main thread handles outgoing messages
```

### Tool Invocation Flow
Function calling examples demonstrate the tool execution cycle:
1. Define tool schema with function declarations
2. Include tools in generation config
3. Parse function call from response
4. Execute function locally
5. Return result in next request

### Multimodal Input Construction
Examples handling images/audio use part-based construction:
```go
parts := []genai.Part{
    genai.Text("Analyze this content"),
    genai.Blob{MimeType: "image/jpeg", Data: imageData},
}
```

### Resource Lifecycle Management
Examples with stateful resources demonstrate proper cleanup:
```go
defer client.Close()
defer session.Close()
defer file.Close()
```

## Environment Configuration

### Required Variables
- `GOOGLE_API_KEY`: Gemini API authentication
- `GOOGLE_CLOUD_PROJECT`: Vertex AI project ID
- `GOOGLE_APPLICATION_CREDENTIALS`: Service account key path (Vertex AI)

### Model Selection
Examples use flag-based or hardcoded model selection:
```go
var model = flag.String("model", "gemini-2.0-flash", "model name")
```

Common models referenced:
- `gemini-2.0-flash`: Fast, general-purpose
- `gemini-2.0-flash-live-preview-04-09`: Live streaming (Vertex AI)
- `gemini-live-2.5-flash-preview`: Live streaming (Gemini API)
- `gemini-1.5-flash-002`: Batch operations
- `gemini-1.5-pro-002`: Advanced capabilities

## Common Patterns Across Examples

### File Path Resolution
Examples include helper functions for repository-relative paths:
```go
func moduleRootDir() string {
    // Walks up directories to find go.mod
}
```

### JSON Marshaling
Structured output through JSON formatting:
```go
jsonData, err := json.MarshalIndent(result, "", "  ")
fmt.Println(string(jsonData))
```

### Embedded Resources
Static resources embedded at compile time:
```go
//go:embed live_streaming.html
var homeTemplate string
```

### Context Propagation
All API calls accept context for cancellation:
```go
ctx := context.Background()
// Pass ctx to all API calls
```

## Reusable Code Segments

### WebSocket Server Setup
`examples/live/live_streaming_server.go` provides production-ready WebSocket server with:
- HTTP to WebSocket upgrade handling
- Concurrent bidirectional message routing
- Session lifecycle management
- Error propagation and logging

### Tool Schema Conversion
`examples/mcptoolbox/mcp_toolbox.go` demonstrates external tool integration with:
- Schema format translation
- Tool registration and lookup
- Parameter marshaling
- Invocation proxying

### Streaming Response Handling
Multiple examples show streaming consumption patterns applicable to:
- Real-time text generation
- Progressive image processing
- Live audio transcription
- Incremental result delivery

### Multimodal Content Assembly
Image and audio examples demonstrate part-based content construction for:
- Mixed text and media inputs
- Inline data encoding
- MIME type specification
- Blob data handling

### Batch Job Management
`examples/batches/create_get_cancel.go` shows asynchronous job patterns for:
- Long-running operations
- GCS-based I/O
- Status polling
- Cancellation handling

### File Upload Pipeline
`examples/files/upload_file.go` demonstrates complete upload workflow with:
- Checksum calculation
- Metadata construction
- Progress tracking
- Error recovery
