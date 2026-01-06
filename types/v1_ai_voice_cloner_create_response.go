package types

// Success
type V1AiVoiceClonerCreateResponse struct {
	// The amount of credits deducted from your account to generate the audio. We charge credits right when the request is made.
	//
	// If an error occurred while generating the audio, credits will be refunded and this field will be updated to include the refund.
	CreditsCharged int `json:"credits_charged"`
	// Unique ID of the audio. Use it with the [Get audio Project API](https://docs.magichour.ai/api-reference/audio-projects/get-audio-details) to fetch status and downloads.
	Id string `json:"id"`
}
