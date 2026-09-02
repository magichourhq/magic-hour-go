package types

// Editing model. Defaults to `ltx-2.3` for free tier and `gemini-omni-1.1` for paid. `gemini-omni` is deprecated; use `gemini-omni-1.1` instead.
type V1AiVideoEditorCreateBodyModelEnum string

const (
	V1AiVideoEditorCreateBodyModelEnumGeminiOmni   V1AiVideoEditorCreateBodyModelEnum = "gemini-omni"
	V1AiVideoEditorCreateBodyModelEnumGeminiOmni11 V1AiVideoEditorCreateBodyModelEnum = "gemini-omni-1.1"
	V1AiVideoEditorCreateBodyModelEnumLtx23        V1AiVideoEditorCreateBodyModelEnum = "ltx-2.3"
)
