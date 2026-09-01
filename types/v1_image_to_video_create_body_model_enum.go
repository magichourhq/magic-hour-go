package types

// The AI model to use for video generation.
//
// * `default`: uses our currently recommended model for general use. For paid tiers, defaults to `kling-3.0`. For free tiers, it defaults to `ltx-2.3`.
// * `kling-2.6`: Best for action, motion blur, and controlled camera moves.
// * `kling-3.0`: Best for cinematic stories, references, and optional audio.
// * `ltx-2.3`: Fastest for general scenes, long clips, audio, and rapid iteration.
// * `minimax-h3`: Great for reference-driven clips with native audio and longer durations.
// * `seedance-1.5`: Best for smooth, consistent motion with an end frame.
// * `seedance-2.0`: Best for reference-led clips with precise subject control.
// * `seedance-2.0-mini`: Faster reference-led clips with consistent motion and audio.
// * `seedance-2.5`: Best for premium realism, detail, and natural motion.
// * `sora-2`: Best for creative concepts and longer clips with audio.
// * `veo3.1`: Best for romantic interactions and expressive action, with realistic detail.
// * `veo3.1-lite`: Balanced realism and audio at a lower cost than Veo 3.1.
// * `wan-2.2`: Best for physical motion, action, and camera movement.
//
// If you specify the deprecated model value that includes the `-audio` suffix, this will be the same as included `audio` as `true`.
type V1ImageToVideoCreateBodyModelEnum string

const (
	V1ImageToVideoCreateBodyModelEnumDefault        V1ImageToVideoCreateBodyModelEnum = "default"
	V1ImageToVideoCreateBodyModelEnumGoogleOmni11   V1ImageToVideoCreateBodyModelEnum = "google-omni-1.1"
	V1ImageToVideoCreateBodyModelEnumKling16        V1ImageToVideoCreateBodyModelEnum = "kling-1.6"
	V1ImageToVideoCreateBodyModelEnumKling25        V1ImageToVideoCreateBodyModelEnum = "kling-2.5"
	V1ImageToVideoCreateBodyModelEnumKling25Audio   V1ImageToVideoCreateBodyModelEnum = "kling-2.5-audio"
	V1ImageToVideoCreateBodyModelEnumKling26        V1ImageToVideoCreateBodyModelEnum = "kling-2.6"
	V1ImageToVideoCreateBodyModelEnumKling30        V1ImageToVideoCreateBodyModelEnum = "kling-3.0"
	V1ImageToVideoCreateBodyModelEnumLtx2           V1ImageToVideoCreateBodyModelEnum = "ltx-2"
	V1ImageToVideoCreateBodyModelEnumLtx23          V1ImageToVideoCreateBodyModelEnum = "ltx-2.3"
	V1ImageToVideoCreateBodyModelEnumLtx25          V1ImageToVideoCreateBodyModelEnum = "ltx-2.5"
	V1ImageToVideoCreateBodyModelEnumMinimaxH3      V1ImageToVideoCreateBodyModelEnum = "minimax-h3"
	V1ImageToVideoCreateBodyModelEnumSeedance       V1ImageToVideoCreateBodyModelEnum = "seedance"
	V1ImageToVideoCreateBodyModelEnumSeedance15     V1ImageToVideoCreateBodyModelEnum = "seedance-1.5"
	V1ImageToVideoCreateBodyModelEnumSeedance20     V1ImageToVideoCreateBodyModelEnum = "seedance-2.0"
	V1ImageToVideoCreateBodyModelEnumSeedance20Mini V1ImageToVideoCreateBodyModelEnum = "seedance-2.0-mini"
	V1ImageToVideoCreateBodyModelEnumSeedance25     V1ImageToVideoCreateBodyModelEnum = "seedance-2.5"
	V1ImageToVideoCreateBodyModelEnumSora2          V1ImageToVideoCreateBodyModelEnum = "sora-2"
	V1ImageToVideoCreateBodyModelEnumVeo31          V1ImageToVideoCreateBodyModelEnum = "veo3.1"
	V1ImageToVideoCreateBodyModelEnumVeo31Audio     V1ImageToVideoCreateBodyModelEnum = "veo3.1-audio"
	V1ImageToVideoCreateBodyModelEnumVeo31Lite      V1ImageToVideoCreateBodyModelEnum = "veo3.1-lite"
	V1ImageToVideoCreateBodyModelEnumWan22          V1ImageToVideoCreateBodyModelEnum = "wan-2.2"
)
