package types

// Optionally add an audio source if you'd like to incorporate audio into your video
type V1AnimationCreateBodyAssetsAudioSourceEnum string

const (
	V1AnimationCreateBodyAssetsAudioSourceEnumFile    V1AnimationCreateBodyAssetsAudioSourceEnum = "file"
	V1AnimationCreateBodyAssetsAudioSourceEnumNone    V1AnimationCreateBodyAssetsAudioSourceEnum = "none"
	V1AnimationCreateBodyAssetsAudioSourceEnumYoutube V1AnimationCreateBodyAssetsAudioSourceEnum = "youtube"
)
