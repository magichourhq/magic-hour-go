package types

// The AI model to use for video generation.
//
// * `default`: uses our currently recommended model for general use. For paid tiers, defaults to `kling-3.0`. For free tiers, it defaults to `ltx-2.3`.
// * `ltx-2.3`: Fastest output. Best for rapid iteration.
// * `wan-2.2`: Strong physics, camera moves, and motion.
// * `kling-2.5`: Great for action, motion blur, and camera moves.
// * `kling-3.0`: Best overall quality for cinematic storytelling.
// * `veo3.1-lite`: Veo quality at a more accessible cost.
// * `veo3.1`: Google's model. Highest realism and detail.
// * `seedance-1.5`: Smooth, consistent motion with precision.
// * `seedance-2.0-mini`: Fast, consistent video with strong motion quality
// * `seedance-2.0`: Top quality with reference-to-video control.
// * `sora-2`: Open AI's model. Great for creativity and viral clips.
//
// If you specify the deprecated model value that includes the `-audio` suffix, this will be the same as included `audio` as `true`.
type V1ImageToVideoCreateBodyModelEnum string

const (
	V1ImageToVideoCreateBodyModelEnumDefault        V1ImageToVideoCreateBodyModelEnum = "default"
	V1ImageToVideoCreateBodyModelEnumKling16        V1ImageToVideoCreateBodyModelEnum = "kling-1.6"
	V1ImageToVideoCreateBodyModelEnumKling25        V1ImageToVideoCreateBodyModelEnum = "kling-2.5"
	V1ImageToVideoCreateBodyModelEnumKling25Audio   V1ImageToVideoCreateBodyModelEnum = "kling-2.5-audio"
	V1ImageToVideoCreateBodyModelEnumKling30        V1ImageToVideoCreateBodyModelEnum = "kling-3.0"
	V1ImageToVideoCreateBodyModelEnumLtx2           V1ImageToVideoCreateBodyModelEnum = "ltx-2"
	V1ImageToVideoCreateBodyModelEnumLtx23          V1ImageToVideoCreateBodyModelEnum = "ltx-2.3"
	V1ImageToVideoCreateBodyModelEnumSeedance       V1ImageToVideoCreateBodyModelEnum = "seedance"
	V1ImageToVideoCreateBodyModelEnumSeedance15     V1ImageToVideoCreateBodyModelEnum = "seedance-1.5"
	V1ImageToVideoCreateBodyModelEnumSeedance20     V1ImageToVideoCreateBodyModelEnum = "seedance-2.0"
	V1ImageToVideoCreateBodyModelEnumSeedance20Mini V1ImageToVideoCreateBodyModelEnum = "seedance-2.0-mini"
	V1ImageToVideoCreateBodyModelEnumSora2          V1ImageToVideoCreateBodyModelEnum = "sora-2"
	V1ImageToVideoCreateBodyModelEnumVeo31          V1ImageToVideoCreateBodyModelEnum = "veo3.1"
	V1ImageToVideoCreateBodyModelEnumVeo31Audio     V1ImageToVideoCreateBodyModelEnum = "veo3.1-audio"
	V1ImageToVideoCreateBodyModelEnumVeo31Lite      V1ImageToVideoCreateBodyModelEnum = "veo3.1-lite"
	V1ImageToVideoCreateBodyModelEnumWan22          V1ImageToVideoCreateBodyModelEnum = "wan-2.2"
)
