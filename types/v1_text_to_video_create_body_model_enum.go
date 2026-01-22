package types

// The AI model to use for video generation.
// * `default`: Our recommended model for general use (Kling 2.5 Audio). Note: For backward compatibility, if you use default and end_seconds > 10, we'll fall back to Kling 1.6.
// * `seedance`: Great for fast iteration and start/end frame
// * `kling-2.5-audio`: Great for motion, action, and camera control
// * `sora-2`: Great for story-telling, dialogue & creativity
// * `veo3.1-audio`: Great for dialogue + SFX generated natively
// * `veo3.1`: Great for realism, polish, & prompt adherence
// * `kling-1.6`: Great for dependable clips with smooth motion
type V1TextToVideoCreateBodyModelEnum string

const (
	V1TextToVideoCreateBodyModelEnumDefault      V1TextToVideoCreateBodyModelEnum = "default"
	V1TextToVideoCreateBodyModelEnumKling16      V1TextToVideoCreateBodyModelEnum = "kling-1.6"
	V1TextToVideoCreateBodyModelEnumKling25Audio V1TextToVideoCreateBodyModelEnum = "kling-2.5-audio"
	V1TextToVideoCreateBodyModelEnumSeedance     V1TextToVideoCreateBodyModelEnum = "seedance"
	V1TextToVideoCreateBodyModelEnumSora2        V1TextToVideoCreateBodyModelEnum = "sora-2"
	V1TextToVideoCreateBodyModelEnumVeo31        V1TextToVideoCreateBodyModelEnum = "veo3.1"
	V1TextToVideoCreateBodyModelEnumVeo31Audio   V1TextToVideoCreateBodyModelEnum = "veo3.1-audio"
)
