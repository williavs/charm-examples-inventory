# Anthropic SDK Go Examples - Contextual Reference

## Quick Reference

| Need | Example File |
|------|-------------|
| Basic message API call | `examples/message/main.go` |
| Streaming message responses | `examples/message-streaming/main.go` |
| Function calling with tools | `examples/tools/main.go` |
| Streaming with tool use | `examples/tools-streaming/main.go` |
| JSON schema-based tool definitions | `examples/tools-streaming-jsonschema/main.go` |
| Image analysis (vision) | `examples/multimodal/main.go` |
| Document file upload and analysis | `examples/file-upload/main.go` |
| AWS Bedrock integration | `examples/bedrock/main.go` |
| AWS Bedrock streaming | `examples/bedrock-streaming/main.go` |
| Google Vertex AI integration | `examples/vertex/main.go` |
| Google Vertex AI streaming | `examples/vertex-streaming/main.go` |
| MCP server integration with streaming | `examples/message-mcp-streaming/main.go` |

## Examples by Capability

### Basic Message API
**Use message when you need:**
- Single-turn or multi-turn conversations
- Non-streaming request/response pattern
- Stop sequences for output control
- Model selection and token limits

**File**: `examples/message/main.go`
**Key patterns**: Client initialization, MessageNewParams configuration, response handling

### Streaming Responses
**Use message-streaming when you need:**
- Real-time token-by-token output delivery
- Event-based processing of AI responses
- Delta events for incremental text
- Reduced perceived latency

**File**: `examples/message-streaming/main.go`
**Key patterns**: NewStreaming method, event loop with stream.Next(), type switching on events

### Function Calling (Tools)
**Use tools when you need:**
- External function execution during conversation
- JSON schema validation for tool inputs
- Multi-turn tool use loops
- Structured data extraction from AI responses

**File**: `examples/tools/main.go`
**Key patterns**: ToolParam definition with InputSchema, tool result handling, message accumulation, agentic loop

### Streaming Tool Use
**Use tools-streaming when you need:**
- Real-time function calling with streaming responses
- Progressive display of tool invocations
- Event accumulation into complete messages
- Partial JSON streaming for tool arguments

**File**: `examples/tools-streaming/main.go`
**Key patterns**: Message.Accumulate() method, ContentBlockDeltaEvent processing, PartialJSON handling

### Type-Safe Tool Schemas
**Use tools-streaming-jsonschema when you need:**
- Go struct-based tool schema generation
- Type safety for tool inputs and outputs
- Automatic JSON schema reflection
- Reduced manual schema definition

**File**: `examples/tools-streaming-jsonschema/main.go`
**Key patterns**: jsonschema.Reflector usage, GenerateSchema helper, typed input structs with json tags

### Multimodal Input
**Use multimodal when you need:**
- Image analysis capabilities
- Vision model access
- Base64-encoded image inputs
- Combined text and image messages

**File**: `examples/multimodal/main.go`
**Key patterns**: NewImageBlockBase64 content blocks, image file encoding, multimodal message construction

### File Upload and Processing
**Use file-upload when you need:**
- Document analysis from uploaded files
- Beta Files API integration
- Server-side file storage
- Large document processing

**File**: `examples/file-upload/main.go`
**Key patterns**: Beta.Files.Upload, BetaDocumentBlock, Files API beta header

### AWS Bedrock Integration
**Use bedrock when you need:**
- Anthropic models via AWS Bedrock
- AWS credential-based authentication
- Region-specific model access
- AWS infrastructure integration

**File**: `examples/bedrock/main.go`
**Key patterns**: bedrock.WithLoadDefaultConfig, AWS model naming convention

### AWS Bedrock Streaming
**Use bedrock-streaming when you need:**
- Streaming responses through AWS Bedrock
- AWS-hosted model streaming
- AWS-authenticated streaming sessions

**File**: `examples/bedrock-streaming/main.go`
**Key patterns**: Bedrock client with NewStreaming, AWS context loading

### Google Vertex AI Integration
**Use vertex when you need:**
- Anthropic models via Google Vertex AI
- Google Cloud authentication
- Regional deployment in GCP
- GCP infrastructure integration

**File**: `examples/vertex/main.go`
**Key patterns**: vertex.WithGoogleAuth, GCP region and project ID configuration, Vertex model naming

### Google Vertex AI Streaming
**Use vertex-streaming when you need:**
- Streaming responses through Vertex AI
- GCP-hosted model streaming
- Google-authenticated streaming sessions

**File**: `examples/vertex-streaming/main.go`
**Key patterns**: Vertex client with NewStreaming, GCP auth context

### MCP Server Integration
**Use message-mcp-streaming when you need:**
- Model Context Protocol server connections
- Remote tool execution via MCP
- SSE-based MCP communication
- Tool filtering and configuration for MCP servers

**File**: `examples/message-mcp-streaming/main.go`
**Key patterns**: BetaRequestMCPServerURLDefinitionParam, MCP beta header, ToolConfiguration, event accumulation

## Implementation Patterns

### Client Initialization
```go
// Standard API
client := anthropic.NewClient()

// AWS Bedrock
client := anthropic.NewClient(bedrock.WithLoadDefaultConfig(ctx))

// Google Vertex AI
client := anthropic.NewClient(vertex.WithGoogleAuth(ctx, region, projectID))
```

### Message Construction
```go
// Text-only
messages := []anthropic.MessageParam{
    anthropic.NewUserMessage(anthropic.NewTextBlock(content)),
}

// With images
messages := []anthropic.MessageParam{
    anthropic.NewUserMessage(
        anthropic.NewTextBlock(content),
        anthropic.NewImageBlockBase64(mediaType, encodedData),
    ),
}

// With tools
messages = append(messages, anthropic.NewUserMessage(toolResults...))
```

### Streaming Event Handling
```go
stream := client.Messages.NewStreaming(ctx, params)

for stream.Next() {
    event := stream.Current()
    switch eventVariant := event.AsAny().(type) {
    case anthropic.ContentBlockDeltaEvent:
        // Handle text delta
        print(eventVariant.Delta.Text)
    case anthropic.MessageDeltaEvent:
        // Handle message-level events
    }
}

if stream.Err() != nil {
    // Handle errors
}
```

### Tool Definition
```go
// Manual schema
tool := anthropic.ToolParam{
    Name: "function_name",
    Description: anthropic.String("description"),
    InputSchema: anthropic.ToolInputSchemaParam{
        Properties: map[string]interface{}{
            "param": map[string]interface{}{
                "type": "string",
                "description": "param description",
            },
        },
    },
}

// JSON schema reflection
type ToolInput struct {
    Param string `json:"param" jsonschema_description:"param description"`
}
schema := GenerateSchema[ToolInput]()
```

### Tool Result Processing
```go
for _, block := range message.Content {
    switch variant := block.AsAny().(type) {
    case anthropic.ToolUseBlock:
        // Parse input
        var input ToolInput
        json.Unmarshal([]byte(variant.JSON.Input.Raw()), &input)

        // Execute function
        result := ExecuteFunction(input)

        // Return result
        toolResults = append(toolResults,
            anthropic.NewToolResultBlock(block.ID, jsonResult, false))
    }
}
```

### Event Accumulation
```go
message := anthropic.Message{}
for stream.Next() {
    event := stream.Current()
    err := message.Accumulate(event)
    if err != nil {
        // Handle accumulation errors
    }
}
// message now contains complete accumulated response
```

## Platform-Specific Patterns

### AWS Bedrock
- Use `bedrock.WithLoadDefaultConfig(context.Background())` for authentication
- Model names follow format: `us.anthropic.claude-sonnet-4-20250514-v1:0`
- Supports both streaming and non-streaming

### Google Vertex AI
- Use `vertex.WithGoogleAuth(ctx, region, projectID)` for authentication
- Model names follow format: `claude-sonnet-4-v1@20250514`
- Requires project ID and region specification

### Model Context Protocol (MCP)
- Requires beta header: `anthropic.AnthropicBetaMCPClient2025_04_04`
- Server configuration includes URL, auth token, and tool filtering
- Use Beta client methods: `client.Beta.Messages.NewStreaming()`
