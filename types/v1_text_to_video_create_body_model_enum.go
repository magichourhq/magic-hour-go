package types

// The AI model to use for video generation.
//
// * `default`: uses our currently recommended model for general use. For paid tiers, defaults to `kling-3.0`. For free tiers, it defaults to `ltx-2.3`.
// * `ltx-2.3`: Fast iteration with audio, lip-sync, and end frame
// * `wan-2.2`: Fast, strong visuals with effects
// * `kling-2.5`: Motion, action, and camera control
// * `kling-3.0`: Cinematic, multi-scene storytelling
// * `veo3.1-lite`: Fast, affordable, high-quality
// * `veo3.1`: Realistic visuals and prompt adherence
// * `seedance`: Fast iteration and start/end frames
// * `seedance-2.0`: State-of-the-art quality and consistency
// * `sora-2`: Story-first concepts and creativity
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
	V1TextToVideoCreateBodyModelEnumLtx23        V1TextToVideoCreateBodyModelEnum = "ltx-2.3"
	V1TextToVideoCreateBodyModelEnumSeedance     V1TextToVideoCreateBodyModelEnum = "seedance"
	V1TextToVideoCreateBodyModelEnumSeedance20   V1TextToVideoCreateBodyModelEnum = "seedance-2.0"
	V1TextToVideoCreateBodyModelEnumSora2        V1TextToVideoCreateBodyModelEnum = "sora-2"
	V1TextToVideoCreateBodyModelEnumVeo31        V1TextToVideoCreateBodyModelEnum = "veo3.1"
	V1TextToVideoCreateBodyModelEnumVeo31Audio   V1TextToVideoCreateBodyModelEnum = "veo3.1-audio"
	V1TextToVideoCreateBodyModelEnumVeo31Lite    V1TextToVideoCreateBodyModelEnum = "veo3.1-lite"
	V1TextToVideoCreateBodyModelEnumWan22        V1TextToVideoCreateBodyModelEnum = "wan-2.2"
)
