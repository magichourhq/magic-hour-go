package ai_qr_code_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// The content of the QR code.
	Content string `json:"content"`
	// The name of image
	Name  nullable.Nullable[string]              `json:"name,omitempty"`
	Style types.PostV1AiQrCodeGeneratorBodyStyle `json:"style"`
}
