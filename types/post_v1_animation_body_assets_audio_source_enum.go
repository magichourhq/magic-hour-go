package types

// Optionally add an audio source if you'd like to incorporate audio into your video
type PostV1AnimationBodyAssetsAudioSourceEnum string

const (
	PostV1AnimationBodyAssetsAudioSourceEnumFile    PostV1AnimationBodyAssetsAudioSourceEnum = "file"
	PostV1AnimationBodyAssetsAudioSourceEnumNone    PostV1AnimationBodyAssetsAudioSourceEnum = "none"
	PostV1AnimationBodyAssetsAudioSourceEnumYoutube PostV1AnimationBodyAssetsAudioSourceEnum = "youtube"
)
