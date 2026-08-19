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
	// Whether this resource is active. If false, it is deleted.
	Enabled bool `json:"enabled"`
	// In the case of an error, this object will contain the error encountered during video render
	Error nullable.Nullable[V1AudioProjectsGetResponseError] `json:"error,omitempty"`
	// Unique ID of the audio. Use it with the [Get audio Project API](https://docs.magichour.ai/api-reference/audio-projects/get-audio-details) to fetch status and downloads.
	Id string `json:"id"`
	// The name of the audio.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The status of the audio.
	//
	// - `draft` - the project was created but has not been submitted for rendering
	// - `queued` - the job is waiting for an available server
	// - `rendering` - the job is being processed; the `audio.started` webhook event fires when rendering begins
	// - `complete` - the job finished successfully; fires `audio.completed`
	// - `error` - the job failed during processing; fires `audio.errored`
	// - `canceled` - the job was manually canceled (for example from the Magic Hour web app)
	//
	// **Note:** `rendering`, `complete`, and `error` have matching webhook events; `canceled` does not - a canceled job emits no webhook event, so poll this endpoint to detect cancellation.
	Status V1AudioProjectsGetResponseStatusEnum `json:"status"`
	// The type of the audio project. Possible values are VOICE_GENERATOR, VOICE_CHANGER, VOICE_CLONER, VIDEO_TO_AUDIO, MUSIC_GENERATOR
	Type string `json:"type"`
}
