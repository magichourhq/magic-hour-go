package types

// Maximum resolution for the generated image.
//
// **Options:**
// - `auto` - Automatic resolution (all tiers, default)
// - `2k` - Up to 2048px (requires Pro or Business tier)
// - `4k` - Up to 4096px (requires Business tier)
//
// Note: Resolution availability depends on your subscription tier. Defaults to `auto` if not specified.
type V1AiImageEditorCreateBodyResolutionEnum string

const (
	V1AiImageEditorCreateBodyResolutionEnum2k   V1AiImageEditorCreateBodyResolutionEnum = "2k"
	V1AiImageEditorCreateBodyResolutionEnum4k   V1AiImageEditorCreateBodyResolutionEnum = "4k"
	V1AiImageEditorCreateBodyResolutionEnumAuto V1AiImageEditorCreateBodyResolutionEnum = "auto"
)
