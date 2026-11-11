package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiVideoEditorCreateBody
type V1AiVideoEditorCreateBody struct {
	// Provide the assets for video editing.
	Assets V1AiVideoEditorCreateBodyAssets `json:"assets"`
	// End time of your clip in seconds. Must be greater than `start_seconds`. Allowed duration depends on model: `gemini-omni-1.2`: 3-15s, `wan-3.0`: 2-15s, `ltx-2.5`: 0.5-60s (30s at 1080p).
	EndSeconds float64 `json:"end_seconds"`
	// Editing model. Defaults to LTX 2.5 for free tier and `gemini-omni-1.2` for paid. `wan-3.0` requires a paid tier. `gemini-omni` and `gemini-omni-1.1` are deprecated and run on `gemini-omni-1.2`.
	Model nullable.Nullable[V1AiVideoEditorCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Name references without @ and refer to them as @name in style.prompt. Names must be unique, ignoring case. Maximum 10 references total. Omit for no references.
	//
	// Reference images for gemini-omni-1.2, wan-3.0.
	References nullable.Nullable[[]V1AiVideoEditorCreateBodyReferencesItem] `json:"references,omitempty"`
	// Output resolution. Defaults to `480p` for free tier and `720p` for paid. Supported per model: `gemini-omni-1.2`: 360p, 720p, 1080p, 4k; `wan-3.0`: 480p, 720p, 1080p; `ltx-2.5`: 480p, 720p, 1080p.
	Resolution nullable.Nullable[V1AiVideoEditorCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64]     `json:"start_seconds,omitempty"`
	Style        V1AiVideoEditorCreateBodyStyle `json:"style"`
}
