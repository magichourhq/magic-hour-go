package types

// The content used to generate speech.
type V1AiVoiceGeneratorCreateBodyStyle struct {
	// Text used to generate speech. Starter tier users can use up to 200 characters, while Creator, Pro, or Business users can use up to 1000.
	Prompt string `json:"prompt"`
	// The voice to use for the speech. Available voices: Elon Musk, Mark Zuckerberg, Joe Rogan, Barack Obama, Morgan Freeman, Kanye West, Donald Trump, Joe Biden, Kim Kardashian, Taylor Swift
	VoiceName V1AiVoiceGeneratorCreateBodyStyleVoiceNameEnum `json:"voice_name"`
}
