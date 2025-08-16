package types

// * `custom` - Use your own prompt for the video.
// * `use_lyrics` - Use the lyrics of the audio to create the prompt. If this option is selected, then `assets.audio_source` must be `file` or `youtube`.
// * `ai_choose` - Let AI write the prompt. If this option is selected, then `assets.audio_source` must be `file` or `youtube`.
type V1AnimationCreateBodyStylePromptTypeEnum string

const (
	V1AnimationCreateBodyStylePromptTypeEnumAiChoose  V1AnimationCreateBodyStylePromptTypeEnum = "ai_choose"
	V1AnimationCreateBodyStylePromptTypeEnumCustom    V1AnimationCreateBodyStylePromptTypeEnum = "custom"
	V1AnimationCreateBodyStylePromptTypeEnumUseLyrics V1AnimationCreateBodyStylePromptTypeEnum = "use_lyrics"
)
