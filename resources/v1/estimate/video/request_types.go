package video

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// ListRequest
type ListRequest struct {
	// The `end_seconds` of the create request. Used by every tool.
	EndSeconds float64 `json:"end_seconds"`
	// The create endpoint to price, named by its path. For example `image-to-video` for `POST /v1/image-to-video`.
	Tool types.V1EstimateVideoListToolEnum `json:"tool"`
	// The `audio` of the create request. Used by `image-to-video`, `text-to-video`. Defaults to `false`.
	Audio nullable.Nullable[bool] `json:"audio,omitempty"`
	// Whether any of the create request's `references` is a video. Used by `text-to-video`. Defaults to `false`.
	HasVideoReference nullable.Nullable[bool] `json:"has_video_reference,omitempty"`
	// The `model` of the create request. Used by `image-to-video`, `text-to-video`, `ai-ugc-ad`, `ai-video-editor`, `character-replace`.
	Model nullable.Nullable[string] `json:"model,omitempty"`
	// The `resolution` of the create request. Used by `image-to-video`, `text-to-video`, `ai-ugc-ad`, `ai-video-editor`, `character-replace`, `audio-to-video`, `ai-video-translator`, `ai-video-upscaler`. `ai-video-translator` defaults to `480p`.
	Resolution nullable.Nullable[string] `json:"resolution,omitempty"`
	// The `start_seconds` of the create request, for tools whose create request has it. Defaults to 0.
	StartSeconds nullable.Nullable[float64] `json:"start_seconds,omitempty"`
}
