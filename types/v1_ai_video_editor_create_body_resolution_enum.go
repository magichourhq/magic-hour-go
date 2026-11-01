package types

// Output resolution. Defaults to `480p` for free tier and `720p` for paid. Supported per model: `gemini-omni-1.1`: 720p, 1080p; `wan-3.0`: 480p, 720p, 1080p; `ltx-2.5`: 480p, 720p, 1080p.
type V1AiVideoEditorCreateBodyResolutionEnum string

const (
	V1AiVideoEditorCreateBodyResolutionEnum1080p V1AiVideoEditorCreateBodyResolutionEnum = "1080p"
	V1AiVideoEditorCreateBodyResolutionEnum480p  V1AiVideoEditorCreateBodyResolutionEnum = "480p"
	V1AiVideoEditorCreateBodyResolutionEnum720p  V1AiVideoEditorCreateBodyResolutionEnum = "720p"
)
