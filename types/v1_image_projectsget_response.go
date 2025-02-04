package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Success
type V1ImageProjectsgetResponse struct {
	CreatedAt string                                    `json:"created_at"`
	Downloads []V1ImageProjectsgetResponseDownloadsItem `json:"downloads"`
	// Indicates whether the resource is deleted
	Enabled bool `json:"enabled"`
	// In the case of an error, this object will contain the error encountered during video render
	Error nullable.Nullable[V1ImageProjectsgetResponseError] `json:"error,omitempty"`
	// Unique ID of the image. This value can be used in the [get image project API](/api/tag/image-projects/get/v1/image-projects/{id}) to fetch additional details such as status
	Id string `json:"id"`
	// Number of images generated
	ImageCount int `json:"image_count"`
	// The name of the image.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The status of the image.
	Status V1ImageProjectsgetResponseStatusEnum `json:"status"`
	// The amount of frames used to generate the image.
	TotalFrameCost int                                `json:"total_frame_cost"`
	Type           V1ImageProjectsgetResponseTypeEnum `json:"type"`
}
