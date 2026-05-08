package types

// Maximum resolution (longest edge) for the output image.
//
// **Options:**
// - `640px` — up to 640px
// - `1k` — up to 1024px
// - `2k` — up to 2048px
// - `4k` — up to 4096px
// - `auto` — **Deprecated.** Mapped server-side from your subscription tier to the best matching resolution the model supports
//
// **Per-model support:**
// - `qwen-edit` - 640px, 1k, 2k
// - `nano-banana` - 640px, 1k
// - `nano-banana-2` - 640px, 1k, 2k, 4k
// - `seedream-v4` - 640px, 1k, 2k, 4k
// - `nano-banana-pro` - 1k, 2k, 4k
// - `seedream-v4.5` - 640px, 1k, 2k, 4k
// - `gpt-image-2` - 640px, 1k, 2k, 4k
//
// Note: Resolution availability depends on the model and your subscription tier.
type V1AiImageEditorCreateBodyResolutionEnum string

const (
	V1AiImageEditorCreateBodyResolutionEnum1k    V1AiImageEditorCreateBodyResolutionEnum = "1k"
	V1AiImageEditorCreateBodyResolutionEnum2k    V1AiImageEditorCreateBodyResolutionEnum = "2k"
	V1AiImageEditorCreateBodyResolutionEnum4k    V1AiImageEditorCreateBodyResolutionEnum = "4k"
	V1AiImageEditorCreateBodyResolutionEnum640px V1AiImageEditorCreateBodyResolutionEnum = "640px"
	V1AiImageEditorCreateBodyResolutionEnumAuto  V1AiImageEditorCreateBodyResolutionEnum = "auto"
)
