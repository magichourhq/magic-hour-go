
package text_to_video
import (
types "github.com/magichourhq/magic-hour-go/types"
nullable "github.com/magichourhq/magic-hour-go/nullable"
)
// CreateRequest
type CreateRequest struct {
    // The total duration of the output video in seconds.
    EndSeconds float64 `json:"end_seconds"`
    // The name of video
    Name nullable.Nullable[string] `json:"name,omitempty"`
    // Determines the orientation of the output video
    Orientation types.PostV1TextToVideoBodyOrientationEnum `json:"orientation"`
    Style types.PostV1TextToVideoBodyStyle `json:"style"`
}



