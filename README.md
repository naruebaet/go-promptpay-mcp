# Go PromptPay MCP

A Model Context Protocol (MCP) server implementation for generating PromptPay QR codes payload in Go.

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/naruebaet/go-promptpay-mcp.svg)](https://pkg.go.dev/github.com/naruebaet/go-promptpay-mcp)

## Features

- 🔄 Generate PromptPay QR codes with EMVCo standard
- 📱 Support for phone numbers (with proper formatting)
- 🆔 Support for Thai national ID numbers
- 💰 Optional amount specification with decimal support
- 🚀 MCP compliant server implementation
- ✅ Input validation and error handling

## Prerequisites

- Go 1.16 or higher
- Basic understanding of MCP (Model Context Protocol)

## Installation

```bash
go install github.com/naruebaet/go-promptpay-mcp
```

## Use in vscode workspace
- Setup in your workspace
    ```shell
    mkdir .vscode
    touch mcp.json
    ```
- Insert this value to mcp.json and start server from here.
    ```json
    {
        "servers": {
            "go-promptpay-mcp": {
                "type": "stdio",
                "command": "go-promptpay-mcp",
                "args": []
            }
        }
    }
    ```

## Usage

The server implements one tool:

### generate_promptpay_code

Generates a PromptPay QR code payload.

Parameters:
- `accountType` (required): Type of account ("phone" or "id")
- `accountNumber` (required): Phone number or Thai ID
- `amount` (optional): Payment amount in THB

Example request:
```json
{
    "tool": "generate_promptpay_code",
    "arguments": {
        "accountType": "phone",
        "accountNumber": "0812345678",
        "amount": 100.50
    }
}
```

Example success response:
```json
{
    "success": true,
    "data": {
        "payload": "00020101021229370016A000000677010111011300668123456785802TH53037645406100.506304A14F"
    }
}
```

Example error response:
```json
{
    "success": false,
    "error": {
        "message": "Invalid account number format"
    }
}
```

## Development

1. Clone the repository
2. Install dependencies: `go mod download`
3. Run the server: `go run main.go`

## Dependencies

- [github.com/mark3labs/mcp-go](https://github.com/mark3labs/mcp-go)
- [github.com/naruebaet/go-promptpay](https://github.com/naruebaet/go-promptpay)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -am 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Create a new Pull Request

## License

MIT License
