package types

// Output resolution. Defaults to `480p` for free tier and `720p` for paid. Google Omni supports 720p only; LTX-2.3 supports 480p, 720p, and 1080p.
type V1AiVideoEditorCreateBodyResolutionEnum string

const (
	V1AiVideoEditorCreateBodyResolutionEnum1080p V1AiVideoEditorCreateBodyResolutionEnum = "1080p"
	V1AiVideoEditorCreateBodyResolutionEnum480p  V1AiVideoEditorCreateBodyResolutionEnum = "480p"
	V1AiVideoEditorCreateBodyResolutionEnum720p  V1AiVideoEditorCreateBodyResolutionEnum = "720p"
)
