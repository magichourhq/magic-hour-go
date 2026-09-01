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
type V1TextToVideoCreateBodyModelEnum string

const (
	V1TextToVideoCreateBodyModelEnumDefault        V1TextToVideoCreateBodyModelEnum = "default"
	V1TextToVideoCreateBodyModelEnumGoogleOmni11   V1TextToVideoCreateBodyModelEnum = "google-omni-1.1"
	V1TextToVideoCreateBodyModelEnumKling16        V1TextToVideoCreateBodyModelEnum = "kling-1.6"
	V1TextToVideoCreateBodyModelEnumKling25        V1TextToVideoCreateBodyModelEnum = "kling-2.5"
	V1TextToVideoCreateBodyModelEnumKling25Audio   V1TextToVideoCreateBodyModelEnum = "kling-2.5-audio"
	V1TextToVideoCreateBodyModelEnumKling26        V1TextToVideoCreateBodyModelEnum = "kling-2.6"
	V1TextToVideoCreateBodyModelEnumKling30        V1TextToVideoCreateBodyModelEnum = "kling-3.0"
	V1TextToVideoCreateBodyModelEnumLtx2           V1TextToVideoCreateBodyModelEnum = "ltx-2"
	V1TextToVideoCreateBodyModelEnumLtx23          V1TextToVideoCreateBodyModelEnum = "ltx-2.3"
	V1TextToVideoCreateBodyModelEnumLtx25          V1TextToVideoCreateBodyModelEnum = "ltx-2.5"
	V1TextToVideoCreateBodyModelEnumMinimaxH3      V1TextToVideoCreateBodyModelEnum = "minimax-h3"
	V1TextToVideoCreateBodyModelEnumSeedance       V1TextToVideoCreateBodyModelEnum = "seedance"
	V1TextToVideoCreateBodyModelEnumSeedance15     V1TextToVideoCreateBodyModelEnum = "seedance-1.5"
	V1TextToVideoCreateBodyModelEnumSeedance20     V1TextToVideoCreateBodyModelEnum = "seedance-2.0"
	V1TextToVideoCreateBodyModelEnumSeedance20Mini V1TextToVideoCreateBodyModelEnum = "seedance-2.0-mini"
	V1TextToVideoCreateBodyModelEnumSeedance25     V1TextToVideoCreateBodyModelEnum = "seedance-2.5"
	V1TextToVideoCreateBodyModelEnumSora2          V1TextToVideoCreateBodyModelEnum = "sora-2"
	V1TextToVideoCreateBodyModelEnumVeo31          V1TextToVideoCreateBodyModelEnum = "veo3.1"
	V1TextToVideoCreateBodyModelEnumVeo31Audio     V1TextToVideoCreateBodyModelEnum = "veo3.1-audio"
	V1TextToVideoCreateBodyModelEnumVeo31Lite      V1TextToVideoCreateBodyModelEnum = "veo3.1-lite"
	V1TextToVideoCreateBodyModelEnumWan22          V1TextToVideoCreateBodyModelEnum = "wan-2.2"
)
