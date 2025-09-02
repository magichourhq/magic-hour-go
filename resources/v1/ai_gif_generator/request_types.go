package ai_gif_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// The name of gif. This value is mainly used for your own identification of the gif.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The output file format for the generated animation.
	OutputFormat nullable.Nullable[types.V1AiGifGeneratorCreateBodyOutputFormatEnum] `json:"output_format,omitempty"`
	Style        types.V1AiGifGeneratorCreateBodyStyle                               `json:"style"`
}
