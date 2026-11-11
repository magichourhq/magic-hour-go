package types

// The AI model to use for video editing. Defaults to `ltx-2.5` on free tiers and `gemini-omni-1.2` on paid tiers.
//
// * `gemini-omni-1.2`: Best for precise clips up to 15s, references, and high-resolution output. Requires a paid tier.
// * `wan-3.0`: High-quality video with native audio, long clips, and end-frame control. Requires a paid tier.
// * `ltx-2.5`: Fastest for general scenes, long clips, audio, and rapid iteration.
// * `kling-4.0`: Create cinematic videos with image and video references and optional audio. Requires a paid tier.
type V1AiVideoEditorCreateBodyModelEnum string

const (
	V1AiVideoEditorCreateBodyModelEnumGeminiOmni   V1AiVideoEditorCreateBodyModelEnum = "gemini-omni"
	V1AiVideoEditorCreateBodyModelEnumGeminiOmni11 V1AiVideoEditorCreateBodyModelEnum = "gemini-omni-1.1"
	V1AiVideoEditorCreateBodyModelEnumGeminiOmni12 V1AiVideoEditorCreateBodyModelEnum = "gemini-omni-1.2"
	V1AiVideoEditorCreateBodyModelEnumKling40      V1AiVideoEditorCreateBodyModelEnum = "kling-4.0"
	V1AiVideoEditorCreateBodyModelEnumLtx23        V1AiVideoEditorCreateBodyModelEnum = "ltx-2.3"
	V1AiVideoEditorCreateBodyModelEnumLtx25        V1AiVideoEditorCreateBodyModelEnum = "ltx-2.5"
	V1AiVideoEditorCreateBodyModelEnumWan30        V1AiVideoEditorCreateBodyModelEnum = "wan-3.0"
)
