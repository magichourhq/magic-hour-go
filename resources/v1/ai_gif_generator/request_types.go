package ai_gif_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Give your gif a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The output file format for the generated animation.
	OutputFormat nullable.Nullable[types.V1AiGifGeneratorCreateBodyOutputFormatEnum] `json:"output_format,omitempty"`
	Style        types.V1AiGifGeneratorCreateBodyStyle                               `json:"style"`
}
