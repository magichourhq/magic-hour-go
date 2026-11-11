package types

// Editing model. Defaults to LTX 2.5 for free tier and `gemini-omni-1.2` for paid. `wan-3.0` requires a paid tier. `gemini-omni` and `gemini-omni-1.1` are deprecated and run on `gemini-omni-1.2`.
type V1AiVideoEditorCreateBodyModelEnum string

const (
	V1AiVideoEditorCreateBodyModelEnumGeminiOmni   V1AiVideoEditorCreateBodyModelEnum = "gemini-omni"
	V1AiVideoEditorCreateBodyModelEnumGeminiOmni11 V1AiVideoEditorCreateBodyModelEnum = "gemini-omni-1.1"
	V1AiVideoEditorCreateBodyModelEnumGeminiOmni12 V1AiVideoEditorCreateBodyModelEnum = "gemini-omni-1.2"
	V1AiVideoEditorCreateBodyModelEnumLtx23        V1AiVideoEditorCreateBodyModelEnum = "ltx-2.3"
	V1AiVideoEditorCreateBodyModelEnumLtx25        V1AiVideoEditorCreateBodyModelEnum = "ltx-2.5"
	V1AiVideoEditorCreateBodyModelEnumWan30        V1AiVideoEditorCreateBodyModelEnum = "wan-3.0"
)
