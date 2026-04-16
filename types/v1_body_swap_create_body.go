package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1BodySwapCreateBody
type V1BodySwapCreateBody struct {
	// Person image and scene image for body swap
	Assets V1BodySwapCreateBodyAssets `json:"assets"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output resolution. Determines credits charged for the run.
	Resolution V1BodySwapCreateBodyResolutionEnum `json:"resolution"`
}
