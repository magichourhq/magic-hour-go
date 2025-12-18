package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Success
type V1AudioProjectsGetResponse struct {
	CreatedAt string `json:"created_at"`
	// The amount of credits deducted from your account to generate the audio. We charge credits right when the request is made.
	//
	// If an error occurred while generating the audio, credits will be refunded and this field will be updated to include the refund.
	CreditsCharged int                                       `json:"credits_charged"`
	Downloads      []V1AudioProjectsGetResponseDownloadsItem `json:"downloads"`
	// Indicates whether the resource is deleted
	Enabled bool `json:"enabled"`
	// In the case of an error, this object will contain the error encountered during video render
	Error nullable.Nullable[V1AudioProjectsGetResponseError] `json:"error,omitempty"`
	// Unique ID of the audio. This value can be used in the [get audio project API](https://docs.magichour.ai/api-reference/audio-projects/get-audio-details) to fetch additional details such as status
	Id string `json:"id"`
	// The name of the audio.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The status of the audio.
	Status V1AudioProjectsGetResponseStatusEnum `json:"status"`
	// The type of the audio project. Possible values are VOICE_GENERATOR, VOICE_CHANGER, VOICE_CLONER
	Type string `json:"type"`
}
