package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1VideoToVideoCreateBodyStyle
type V1VideoToVideoCreateBodyStyle struct {
	ArtStyle V1VideoToVideoCreateBodyStyleArtStyleEnum `json:"art_style"`
	// * `Dreamshaper` - a good all-around model that works for both animations as well as realism.
	// * `Absolute Reality` - better at realism, but you'll often get similar results with Dreamshaper as well.
	// * `Flat 2D Anime` - best for a flat illustration style that's common in most anime.
	// * `default` - use the default recommended model for the selected art style.
	Model nullable.Nullable[V1VideoToVideoCreateBodyStyleModelEnum] `json:"model,omitempty"`
	// The prompt used for the video. Prompt is required if `prompt_type` is `custom` or `append_default`. If `prompt_type` is `default`, then the `prompt` value passed will be ignored.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
	// * `default` - Use the default recommended prompt for the art style.
	// * `custom` - Only use the prompt passed in the API. Note: for v1, lora prompt will still be auto added to apply the art style properly.
	// * `append_default` - Add the default recommended prompt to the end of the prompt passed in the API.
	PromptType nullable.Nullable[V1VideoToVideoCreateBodyStylePromptTypeEnum] `json:"prompt_type,omitempty"`
	// * `v1` - more detail, closer prompt adherence, and frame-by-frame previews.
	// * `v2` - faster, more consistent, and less noisy.
	// * `default` - use the selected art style's default version. Some art styles default to v3, which is unavailable through this API. For `Studio Ghibli`, explicitly set `version` to `v2` to use its supported preset.
	Version nullable.Nullable[V1VideoToVideoCreateBodyStyleVersionEnum] `json:"version,omitempty"`
}
