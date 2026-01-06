package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiQrCodeGeneratorCreateBody
type V1AiQrCodeGeneratorCreateBody struct {
	// The content of the QR code.
	Content string `json:"content"`
	// Give your image a custom name for easy identification.
	Name  nullable.Nullable[string]          `json:"name,omitempty"`
	Style V1AiQrCodeGeneratorCreateBodyStyle `json:"style"`
}
