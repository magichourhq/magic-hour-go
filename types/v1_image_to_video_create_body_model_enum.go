package types

// The AI model to use for video generation.
// * `default`: Our recommended model for general use (Kling 2.5 Audio). Note: For backward compatibility, if you use default and end_seconds > 10, we'll fall back to Kling 1.6.
// * `seedance`: Great for fast iteration and start/end frame
// * `kling-2.5`: Great for motion, action, and camera control
// * `kling-3.0`: Great for cinematic, multi-scene storytelling with control
// * `sora-2`: Great for story-telling, dialogue & creativity
// * `veo3.1`: Great for realism, polish, & prompt adherence
// * `kling-1.6`: Great for dependable clips with smooth motion
type V1ImageToVideoCreateBodyModelEnum string

const (
	V1ImageToVideoCreateBodyModelEnumDefault      V1ImageToVideoCreateBodyModelEnum = "default"
	V1ImageToVideoCreateBodyModelEnumKling16      V1ImageToVideoCreateBodyModelEnum = "kling-1.6"
	V1ImageToVideoCreateBodyModelEnumKling25      V1ImageToVideoCreateBodyModelEnum = "kling-2.5"
	V1ImageToVideoCreateBodyModelEnumKling25Audio V1ImageToVideoCreateBodyModelEnum = "kling-2.5-audio"
	V1ImageToVideoCreateBodyModelEnumKling30      V1ImageToVideoCreateBodyModelEnum = "kling-3.0"
	V1ImageToVideoCreateBodyModelEnumSeedance     V1ImageToVideoCreateBodyModelEnum = "seedance"
	V1ImageToVideoCreateBodyModelEnumSora2        V1ImageToVideoCreateBodyModelEnum = "sora-2"
	V1ImageToVideoCreateBodyModelEnumVeo31        V1ImageToVideoCreateBodyModelEnum = "veo3.1"
	V1ImageToVideoCreateBodyModelEnumVeo31Audio   V1ImageToVideoCreateBodyModelEnum = "veo3.1-audio"
)
