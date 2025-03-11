
package types
import (
nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1TextToVideoBody
type PostV1TextToVideoBody struct {
    // The total duration of the output video in seconds.
    EndSeconds float64 `json:"end_seconds"`
    // The name of video
    Name nullable.Nullable[string] `json:"name,omitempty"`
    // Determines the orientation of the output video
    Orientation PostV1TextToVideoBodyOrientationEnum `json:"orientation"`
    Style PostV1TextToVideoBodyStyle `json:"style"`
}



