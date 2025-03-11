package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Success
type GetV1ImageProjectsIdResponse struct {
	CreatedAt string                                      `json:"created_at"`
	Downloads []GetV1ImageProjectsIdResponseDownloadsItem `json:"downloads"`
	// Indicates whether the resource is deleted
	Enabled bool `json:"enabled"`
	// In the case of an error, this object will contain the error encountered during video render
	Error nullable.Nullable[GetV1ImageProjectsIdResponseError] `json:"error,omitempty"`
	// Unique ID of the image. This value can be used in the [get image project API](https://docs.magichour.ai/api-reference/image-projects/get-image-details) to fetch additional details such as status
	Id string `json:"id"`
	// Number of images generated
	ImageCount int `json:"image_count"`
	// The name of the image.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The status of the image.
	Status GetV1ImageProjectsIdResponseStatusEnum `json:"status"`
	// The amount of frames used to generate the image.
	TotalFrameCost int                                  `json:"total_frame_cost"`
	Type           GetV1ImageProjectsIdResponseTypeEnum `json:"type"`
}
