package types

// The AI model to use for video generation.
//
// * `default`: uses our currently recommended model for general use. For paid tiers, defaults to `kling-3.0`. For free tiers, it defaults to `ltx-2.3`.
// * `ltx-2.3`: Fast iteration with lip-sync & end frame
// * `wan-2.2`: Fast, strong visuals with effects
// * `kling-2.5`: Motion, action, and camera control
// * `kling-3.0`: Cinematic, multi-scene storytelling
// * `veo3.1-lite`: Fast, affordable, high-quality
// * `veo3.1`: Realistic visuals and prompt adherence
// * `seedance`: Fast iteration
// * `seedance-2.0`: State-of-the-art quality and consistency
// * `sora-2`: Story-first concepts and creativity
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
	V1ImageToVideoCreateBodyModelEnumLtx23        V1ImageToVideoCreateBodyModelEnum = "ltx-2.3"
	V1ImageToVideoCreateBodyModelEnumSeedance     V1ImageToVideoCreateBodyModelEnum = "seedance"
	V1ImageToVideoCreateBodyModelEnumSeedance20   V1ImageToVideoCreateBodyModelEnum = "seedance-2.0"
	V1ImageToVideoCreateBodyModelEnumSora2        V1ImageToVideoCreateBodyModelEnum = "sora-2"
	V1ImageToVideoCreateBodyModelEnumVeo31        V1ImageToVideoCreateBodyModelEnum = "veo3.1"
	V1ImageToVideoCreateBodyModelEnumVeo31Audio   V1ImageToVideoCreateBodyModelEnum = "veo3.1-audio"
	V1ImageToVideoCreateBodyModelEnumVeo31Lite    V1ImageToVideoCreateBodyModelEnum = "veo3.1-lite"
	V1ImageToVideoCreateBodyModelEnumWan22        V1ImageToVideoCreateBodyModelEnum = "wan-2.2"
)
