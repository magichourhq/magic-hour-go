package types

// The create endpoint to price, named by its path. For example `ai-voice-generator` for `POST /v1/ai-voice-generator`.
type V1EstimateAudioListToolEnum string

const (
	V1EstimateAudioListToolEnumAiMusicGenerator V1EstimateAudioListToolEnum = "ai-music-generator"
	V1EstimateAudioListToolEnumAiVoiceCloner    V1EstimateAudioListToolEnum = "ai-voice-cloner"
	V1EstimateAudioListToolEnumAiVoiceGenerator V1EstimateAudioListToolEnum = "ai-voice-generator"
)
