package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiHeadshotGeneratorCreateBodyStyle
type V1AiHeadshotGeneratorCreateBodyStyle struct {
	// Prompt used to guide the style of your headshot. We recommend omitting the prompt unless you want to customize your headshot. You can visit [AI headshot generator](https://magichour.ai/create/ai-headshot-generator) to view an example of a good prompt used for our 'Professional' style.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
