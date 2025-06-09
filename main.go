package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
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

// makeGeneratePromptPayHandler returns a ToolHandlerFunc that handles the generation of PromptPay QR codes.
// It parses the incoming request to extract PromptPay account details and amount, generates a QR code string,
// encodes it as a PNG image, and returns the image as a base64-encoded string in the response.
// The handler uses a temporary file to store the generated QR code image, which is cleaned up after use.
// If any error occurs during request binding, QR code generation, or image encoding, an error is returned.
func makeGeneratePromptPayHandler(service *promptpay.Service) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var req types.GeneratePromptPayRequest
		if err := request.BindArguments(&req); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("invalid request: %w", err)), err
		}

		qrCode, err := service.GenerateQRCode(req.AccountType, req.AccountNumber, req.Amount)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("failed to generate QR code: %w", err)), err
		}

		// save file to disk if needed
		// Create the barcode
		qrCodeImg, _ := qr.Encode(qrCode, qr.M, qr.Auto)

		// Scale the barcode to 200x200 pixels
		qrCodeImg, _ = barcode.Scale(qrCodeImg, 200, 200)

		// create the output file
		file, _ := os.CreateTemp("", "qrcode_*.png")
		defer os.Remove(file.Name()) // Clean up the file after use
		defer file.Close()

		// encode the barcode as png
		png.Encode(file, qrCodeImg)

		// read the file from temp
		qrCodeByte, _ := os.ReadFile(file.Name())

		b64 := base64.StdEncoding.EncodeToString(qrCodeByte)

		// Return the QR code as an image result
		return mcp.NewToolResultImage(qrCode, b64, "image/png"), nil
	}
}
