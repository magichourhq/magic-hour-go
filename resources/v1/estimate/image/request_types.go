package image

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// ListRequest
type ListRequest struct {
	// The create endpoint to price, named by its path. For example `ai-image-editor` for `POST /v1/ai-image-editor`.
	Tool types.V1EstimateImageListToolEnum `json:"tool"`
	// The `image_count` of the create request. Used by `ai-image-generator`, `ai-image-editor`. Defaults to 1.
	ImageCount nullable.Nullable[int] `json:"image_count,omitempty"`
	// The `style.mode` of the create request. Used by `ai-image-upscaler`. One of `preserve`, `balanced`, `creative`. Defaults to `balanced`.
	Mode nullable.Nullable[string] `json:"mode,omitempty"`
	// The `model` of the create request. Used by `ai-image-generator`, `ai-image-editor`.
	Model nullable.Nullable[string] `json:"model,omitempty"`
	// The `resolution` of the create request. Used by `ai-image-generator`, `ai-image-editor`, `body-swap`. One of `640px`, `1k`, `2k`, `4k`.
	Resolution nullable.Nullable[string] `json:"resolution,omitempty"`
	// The `scale_factor` of the create request. Used by `ai-image-upscaler`. One of `2`, `4`.
	ScaleFactor nullable.Nullable[float64] `json:"scale_factor,omitempty"`
}
