package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/naruebaet/go-promptpay-mcp/promptpay"
	"github.com/naruebaet/go-promptpay-mcp/types"
)

func main() {
	// Initialize MCP server
	s := server.NewMCPServer(
		"PromptPay MCP Server",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	tool := mcp.NewTool(
		"generate_promptpay_code",
		mcp.WithDescription("Generate PromptPay QR Code for payment via phone number or Thai ID, with optional amount"),
		mcp.WithString("accountType",
			mcp.Required(),
			mcp.Description("Account type: 'phone' for phone numbers, 'id' for Thai ID"),
		),
		mcp.WithString("accountNumber",
			mcp.Required(),
			mcp.Description("Phone number (e.g., 0812345678) or Thai ID (13 digits)"),
		),
		mcp.WithNumber("amount",
			mcp.Description("Payment amount in THB (optional)"),
		),
	)

	ppService := promptpay.NewService()
	s.AddTool(tool, makeGeneratePromptPayHandler(ppService))

	// Start the MCP server
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Failed to start MCP server: %v", err)
	}
}

func makeGeneratePromptPayHandler(service *promptpay.Service) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req types.GeneratePromptPayRequest
		if err := request.BindArguments(&req); err != nil {
			return nil, fmt.Errorf("invalid request: %w", err)
		}

		qrCode, err := service.GenerateQRCode(req.AccountType, req.AccountNumber, req.Amount)
		if err != nil {
			return nil, fmt.Errorf("failed to generate QR code: %w", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("Here is qr code payload : %s", qrCode)), nil
	}
}
