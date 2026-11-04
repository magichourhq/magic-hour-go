package types

// V1AiVoiceClonerCreateBodyStyle
type V1AiVoiceClonerCreateBodyStyle struct {
	// Text used to generate speech from the cloned voice. The character limit is 5000 characters.
	Prompt string `json:"prompt"`
}
