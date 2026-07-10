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
type V1TextToVideoCreateBodyModelEnum string

const (
	V1TextToVideoCreateBodyModelEnumDefault        V1TextToVideoCreateBodyModelEnum = "default"
	V1TextToVideoCreateBodyModelEnumKling16        V1TextToVideoCreateBodyModelEnum = "kling-1.6"
	V1TextToVideoCreateBodyModelEnumKling25        V1TextToVideoCreateBodyModelEnum = "kling-2.5"
	V1TextToVideoCreateBodyModelEnumKling25Audio   V1TextToVideoCreateBodyModelEnum = "kling-2.5-audio"
	V1TextToVideoCreateBodyModelEnumKling30        V1TextToVideoCreateBodyModelEnum = "kling-3.0"
	V1TextToVideoCreateBodyModelEnumLtx2           V1TextToVideoCreateBodyModelEnum = "ltx-2"
	V1TextToVideoCreateBodyModelEnumLtx23          V1TextToVideoCreateBodyModelEnum = "ltx-2.3"
	V1TextToVideoCreateBodyModelEnumSeedance       V1TextToVideoCreateBodyModelEnum = "seedance"
	V1TextToVideoCreateBodyModelEnumSeedance15     V1TextToVideoCreateBodyModelEnum = "seedance-1.5"
	V1TextToVideoCreateBodyModelEnumSeedance20     V1TextToVideoCreateBodyModelEnum = "seedance-2.0"
	V1TextToVideoCreateBodyModelEnumSeedance20Mini V1TextToVideoCreateBodyModelEnum = "seedance-2.0-mini"
	V1TextToVideoCreateBodyModelEnumSora2          V1TextToVideoCreateBodyModelEnum = "sora-2"
	V1TextToVideoCreateBodyModelEnumVeo31          V1TextToVideoCreateBodyModelEnum = "veo3.1"
	V1TextToVideoCreateBodyModelEnumVeo31Audio     V1TextToVideoCreateBodyModelEnum = "veo3.1-audio"
	V1TextToVideoCreateBodyModelEnumVeo31Lite      V1TextToVideoCreateBodyModelEnum = "veo3.1-lite"
	V1TextToVideoCreateBodyModelEnumWan22          V1TextToVideoCreateBodyModelEnum = "wan-2.2"
)
