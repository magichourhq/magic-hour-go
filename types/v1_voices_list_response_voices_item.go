package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1VoicesListResponseVoicesItem
type V1VoicesListResponseVoicesItem struct {
	// URL of the voice's avatar image, or null when the voice has no avatar.
	AvatarUrl nullable.Nullable[string] `json:"avatar_url,omitempty"`
	// Unique ID of the voice. Pass it as `style.voice_name` when generating speech. Names are not unique, so use the ID.
	Id string `json:"id"`
	// Display name of the voice.
	Name string `json:"name"`
	// URL of an audio sample of the voice.
	PreviewUrl string `json:"preview_url"`
}
