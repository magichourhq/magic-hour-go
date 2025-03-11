
package types



// 
// * `custom` - use your own prompt for the video.
// * `use_lyrics` - Use the lyrics of the audio to create the prompt. If this option is selected, then `assets.audio_source` must be `file` or `youtube`.
// * `ai_choose` - Let AI write the prompt. If this option is selected, then `assets.audio_source` must be `file` or `youtube`.
type PostV1AnimationBodyStylePromptTypeEnum string
const (
    PostV1AnimationBodyStylePromptTypeEnumAiChoose PostV1AnimationBodyStylePromptTypeEnum = "ai_choose"
    PostV1AnimationBodyStylePromptTypeEnumCustom PostV1AnimationBodyStylePromptTypeEnum = "custom"
    PostV1AnimationBodyStylePromptTypeEnumUseLyrics PostV1AnimationBodyStylePromptTypeEnum = "use_lyrics"
)


