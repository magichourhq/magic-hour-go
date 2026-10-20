package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for creating a talking photo
type V1AiTalkingPhotoCreateBody struct {
	// Provide the assets for creating a talking photo
	Assets V1AiTalkingPhotoCreateBodyAssets `json:"assets"`
	// The end time of the input audio in seconds. Maximum clip length depends on style.generation_mode: realistic 300s, prompted 45s.
	EndSeconds float64 `json:"end_seconds"`
	// Deprecated: use `resolution`. Still accepted as a larger-dimension cap, limited by your plan. Ignored when `resolution` is provided.
	//
	// After plan capping:
	// * Below 1280 maps to `480p` (intermediate values round down).
	// * 1280 or above maps to `720p`.
	MaxResolution nullable.Nullable[int] `json:"max_resolution,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output preset. Overrides `max_resolution`. Defaults to 480p in prompted mode; realistic mode uses the mapped plan maximum.
	Resolution nullable.Nullable[V1AiTalkingPhotoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// The start time of the input audio in seconds. Maximum clip length depends on style.generation_mode: realistic 300s, prompted 45s.
	StartSeconds float64 `json:"start_seconds"`
	// Attributes used to dictate the style of the output
	Style nullable.Nullable[V1AiTalkingPhotoCreateBodyStyle] `json:"style,omitempty"`
}
