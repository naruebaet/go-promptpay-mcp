# Go PromptPay MCP

A Model Context Protocol (MCP) server implementation for generating PromptPay QR codes payload in Go.

## Features

- Generate PromptPay QR codes
- Support for both phone numbers and Thai ID numbers
- Optional amount specification
- MCP compliant server implementation

## Installation

```bash
go get github.com/naruebaet/go-promptpay-mcp
```

## Usage

The server implements one tool:

### generate_promptpay_code

Generates a PromptPay QR code payload.

Parameters:
- `accountType` (required): Type of account ("phone" or "id")
- `accountNumber` (required): Phone number or Thai ID
- `amount` (optional): Payment amount

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

## Development

1. Clone the repository
2. Install dependencies: `go mod download`
3. Run the server: `go run main.go`

## Dependencies

- [github.com/mark3labs/mcp-go](https://github.com/mark3labs/mcp-go)
- [github.com/naruebaet/go-promptpay](https://github.com/naruebaet/go-promptpay)

## License

MIT License
