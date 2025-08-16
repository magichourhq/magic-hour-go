package ai_qr_code_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// The content of the QR code.
	Content string `json:"content"`
	// The name of image. This value is mainly used for your own identification of the image.
	Name  nullable.Nullable[string]                `json:"name,omitempty"`
	Style types.V1AiQrCodeGeneratorCreateBodyStyle `json:"style"`
}
