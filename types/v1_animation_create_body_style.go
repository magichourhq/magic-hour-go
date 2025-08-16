package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Defines the style of the output video
type V1AnimationCreateBodyStyle struct {
	// The art style used to create the output video
	ArtStyle V1AnimationCreateBodyStyleArtStyleEnum `json:"art_style"`
	// Describe custom art style. This field is required if `art_style` is `Custom`
	ArtStyleCustom nullable.Nullable[string] `json:"art_style_custom,omitempty"`
	// The camera effect used to create the output video
	CameraEffect V1AnimationCreateBodyStyleCameraEffectEnum `json:"camera_effect"`
	// The prompt used for the video. Prompt is required if `prompt_type` is `custom`. Otherwise this value is ignored
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
	//
	// * `custom` - Use your own prompt for the video.
	// * `use_lyrics` - Use the lyrics of the audio to create the prompt. If this option is selected, then `assets.audio_source` must be `file` or `youtube`.
	// * `ai_choose` - Let AI write the prompt. If this option is selected, then `assets.audio_source` must be `file` or `youtube`.
	PromptType V1AnimationCreateBodyStylePromptTypeEnum `json:"prompt_type"`
	// Change determines how quickly the video's content changes across frames.
	// * Higher = more rapid transitions.
	// * Lower = more stable visual experience.
	TransitionSpeed int `json:"transition_speed"`
}
