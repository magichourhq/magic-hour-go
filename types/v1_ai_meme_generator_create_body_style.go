package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiMemeGeneratorCreateBodyStyle
type V1AiMemeGeneratorCreateBodyStyle struct {
	// Whether to search the web for meme content.
	SearchWeb nullable.Nullable[bool] `json:"searchWeb,omitempty"`
	// To use our templates, pass in one of the enum values.
	Template V1AiMemeGeneratorCreateBodyStyleTemplateEnum `json:"template"`
	// The topic of the meme.
	Topic string `json:"topic"`
}
