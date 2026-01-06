package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiGifGeneratorCreateBody
type V1AiGifGeneratorCreateBody struct {
	// Give your gif a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The output file format for the generated animation.
	OutputFormat nullable.Nullable[V1AiGifGeneratorCreateBodyOutputFormatEnum] `json:"output_format,omitempty"`
	Style        V1AiGifGeneratorCreateBodyStyle                               `json:"style"`
}
