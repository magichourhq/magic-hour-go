package body_swap

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Person image and scene image for body swap
	Assets types.V1BodySwapCreateBodyAssets `json:"assets"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output resolution. Determines credits charged for the run.
	Resolution types.V1BodySwapCreateBodyResolutionEnum `json:"resolution"`
}
