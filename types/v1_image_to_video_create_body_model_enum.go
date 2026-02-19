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
type V1ImageToVideoCreateBodyModelEnum string

const (
	V1ImageToVideoCreateBodyModelEnumDefault      V1ImageToVideoCreateBodyModelEnum = "default"
	V1ImageToVideoCreateBodyModelEnumKling16      V1ImageToVideoCreateBodyModelEnum = "kling-1.6"
	V1ImageToVideoCreateBodyModelEnumKling25      V1ImageToVideoCreateBodyModelEnum = "kling-2.5"
	V1ImageToVideoCreateBodyModelEnumKling25Audio V1ImageToVideoCreateBodyModelEnum = "kling-2.5-audio"
	V1ImageToVideoCreateBodyModelEnumKling30      V1ImageToVideoCreateBodyModelEnum = "kling-3.0"
	V1ImageToVideoCreateBodyModelEnumLtx2         V1ImageToVideoCreateBodyModelEnum = "ltx-2"
	V1ImageToVideoCreateBodyModelEnumSeedance     V1ImageToVideoCreateBodyModelEnum = "seedance"
	V1ImageToVideoCreateBodyModelEnumSora2        V1ImageToVideoCreateBodyModelEnum = "sora-2"
	V1ImageToVideoCreateBodyModelEnumVeo31        V1ImageToVideoCreateBodyModelEnum = "veo3.1"
	V1ImageToVideoCreateBodyModelEnumVeo31Audio   V1ImageToVideoCreateBodyModelEnum = "veo3.1-audio"
)
