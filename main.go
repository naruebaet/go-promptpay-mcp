package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/naruebaet/go-promptpay/pp"
)

type GeneratePromptPayRequest struct {
	AccountType   pp.AccountType `json:"accountType"`
	AccountNumber string         `json:"accountNumber"`
	Amount        *float64       `json:"amount,omitempty"`
}

func main() {
	// Initialize MCP server
	s := server.NewMCPServer(
		"PromptPay MCP Server",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	tool := mcp.NewTool(
		"generate_promptpay_code",
		mcp.WithDescription("Generate PromptPay QR Code no amount"),
		mcp.WithString("accountType",
			mcp.Required(),
			mcp.Description("Type of account for PromptPay (phone or ID)"),
		),
		mcp.WithString("accountNumber",
			mcp.Required(),
			mcp.Description("Account number for PromptPay (phone number or Thai ID)"),
		),
		mcp.WithNumber("amount",
			mcp.Description("Amount for PromptPay QR Code (optional)"),
		),
	)

	s.AddTool(tool, generatePromptPayCode)

	// Start the MCP server
	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Failed to start MCP server: %v", err)
	}
}

func generatePromptPayCode(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	var req GeneratePromptPayRequest

	err := request.BindArguments(&req)
	if err != nil {
		return nil, err
	}

	// Generate PromptPay QR code payload
	var qrCode string
	if req.Amount != nil {
		qrCode, err = pp.GenPromptpayWithAmount(req.AccountType, req.AccountNumber, *req.Amount)
	} else {
		qrCode, err = pp.GenPromptpay(req.AccountType, req.AccountNumber)
	}

	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(fmt.Sprintf("Here is qr code payload : %s", qrCode)), nil
}
