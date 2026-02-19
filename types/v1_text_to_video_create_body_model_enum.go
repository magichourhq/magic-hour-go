package types

// The AI model to use for video generation.
//
// * `default`: uses our currently recommended model for general use. For paid tiers, defaults to `kling-2.5`. For free tiers, it defaults to `ltx-2`.
// * `ltx-2`: Great for fast iteration with audio, lip-sync, and expressive faces
// * `seedance`: Great for fast iteration and start/end frame
// * `kling-2.5`: Great for motion, action, and camera control
// * `kling-3.0`: Great for cinematic, multi-scene storytelling with control
// * `sora-2`: Great for story-telling, dialogue & creativity
// * `veo3.1`: Great for realism, polish, & prompt adherence
//
// Legacy models:
// * `kling-1.6`: Great for dependable clips with smooth motion
//
// If you specify the deprecated model value that includes the `-audio` suffix, this will be the same as included `audio` as `true`.
type V1TextToVideoCreateBodyModelEnum string

const (
	V1TextToVideoCreateBodyModelEnumDefault      V1TextToVideoCreateBodyModelEnum = "default"
	V1TextToVideoCreateBodyModelEnumKling16      V1TextToVideoCreateBodyModelEnum = "kling-1.6"
	V1TextToVideoCreateBodyModelEnumKling25      V1TextToVideoCreateBodyModelEnum = "kling-2.5"
	V1TextToVideoCreateBodyModelEnumKling25Audio V1TextToVideoCreateBodyModelEnum = "kling-2.5-audio"
	V1TextToVideoCreateBodyModelEnumKling30      V1TextToVideoCreateBodyModelEnum = "kling-3.0"
	V1TextToVideoCreateBodyModelEnumLtx2         V1TextToVideoCreateBodyModelEnum = "ltx-2"
	V1TextToVideoCreateBodyModelEnumSeedance     V1TextToVideoCreateBodyModelEnum = "seedance"
	V1TextToVideoCreateBodyModelEnumSora2        V1TextToVideoCreateBodyModelEnum = "sora-2"
	V1TextToVideoCreateBodyModelEnumVeo31        V1TextToVideoCreateBodyModelEnum = "veo3.1"
	V1TextToVideoCreateBodyModelEnumVeo31Audio   V1TextToVideoCreateBodyModelEnum = "veo3.1-audio"
)
