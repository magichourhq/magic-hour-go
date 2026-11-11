package ai_video_editor

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for video editing.
	Assets types.V1AiVideoEditorCreateBodyAssets `json:"assets"`
	// End time of your clip in seconds. Must be greater than `start_seconds`. Allowed duration depends on model: `gemini-omni-1.2`: 3-15s, `wan-3.0`: 2-15s, `ltx-2.5`: 0.5-60s (30s at 1080p), `kling-4.0`: 3-30s.
	EndSeconds float64 `json:"end_seconds"`
	// The AI model to use for video editing. Defaults to `ltx-2.5` on free tiers and `gemini-omni-1.2` on paid tiers.
	//
	// * `gemini-omni-1.2`: Best for precise clips up to 15s, references, and high-resolution output. Requires a paid tier.
	// * `wan-3.0`: High-quality video with native audio, long clips, and end-frame control. Requires a paid tier.
	// * `ltx-2.5`: Fastest for general scenes, long clips, audio, and rapid iteration.
	// * `kling-4.0`: Create cinematic videos with image and video references and optional audio. Requires a paid tier.
	Model nullable.Nullable[types.V1AiVideoEditorCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Name references without @ and refer to them as @name in style.prompt. Names must be unique, ignoring case. Maximum 10 references total. Omit for no references.
	//
	// Reference images for gemini-omni-1.2, wan-3.0, kling-4.0.
	References nullable.Nullable[[]types.V1AiVideoEditorCreateBodyReferencesItem] `json:"references,omitempty"`
	// Output resolution. Defaults to `480p` for free tier and `720p` for paid. Supported per model: `gemini-omni-1.2`: 360p, 720p, 1080p, 4k; `wan-3.0`: 480p, 720p, 1080p; `ltx-2.5`: 480p, 720p, 1080p; `kling-4.0`: 720p, 1080p.
	Resolution nullable.Nullable[types.V1AiVideoEditorCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64]           `json:"start_seconds,omitempty"`
	Style        types.V1AiVideoEditorCreateBodyStyle `json:"style"`
}
