package types

// Editing model. Defaults to `ltx-2.3` for free tier and `gemini-omni` for paid. Use `ltx-2.3` for LTX video edit.
type V1AiVideoEditorCreateBodyModelEnum string

const (
	V1AiVideoEditorCreateBodyModelEnumGeminiOmni V1AiVideoEditorCreateBodyModelEnum = "gemini-omni"
	V1AiVideoEditorCreateBodyModelEnumLtx23      V1AiVideoEditorCreateBodyModelEnum = "ltx-2.3"
)
