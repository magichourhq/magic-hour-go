package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1AiQrCodeGeneratorBody
type PostV1AiQrCodeGeneratorBody struct {
	// The content of the QR code.
	Content string `json:"content"`
	// The name of image
	Name  nullable.Nullable[string]        `json:"name,omitempty"`
	Style PostV1AiQrCodeGeneratorBodyStyle `json:"style"`
}
