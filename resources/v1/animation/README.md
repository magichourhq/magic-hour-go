# v1.animation

## Module Functions

### Animation <a name="create"></a>

Create a Animation video. The estimated frame cost is calculated based on the `fps` and `end_seconds` input.

**API Endpoint**: `POST /v1/animation`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for animation. | `V1AnimationCreateBodyAssets {AudioFilePath: nullable.NewValue("api-assets/id/1234.mp3"),AudioSource: V1AnimationCreateBodyAssetsAudioSourceEnumFile,ImageFilePath: nullable.NewValue("api-assets/id/1234.png"),}` |
| `└─ AudioFilePath` | ✗ | The path of the input audio. This field is required if `audio_source` is `file`. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.mp3"` |
| `└─ AudioSource` | ✓ | Optionally add an audio source if you'd like to incorporate audio into your video | `V1AnimationCreateBodyAssetsAudioSourceEnumFile` |
| `└─ ImageFilePath` | ✗ | An initial image to use a the first frame of the video. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `└─ YoutubeUrl` | ✗ | Using a youtube video as the input source. This field is required if `audio_source` is `youtube` | `"http://www.example.com"` |
| `EndSeconds` | ✓ | This value determines the duration of the output video. | `15.0` |
| `Fps` | ✓ | The desire output video frame rate | `12.0` |
| `Height` | ✓ | The height of the final output video. The maximum height depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details | `960` |
| `Style` | ✓ | Defines the style of the output video | `V1AnimationCreateBodyStyle {ArtStyle: V1AnimationCreateBodyStyleArtStyleEnumPainterlyIllustration,CameraEffect: V1AnimationCreateBodyStyleCameraEffectEnumSimpleZoomIn,Prompt: nullable.NewValue("Cyberpunk city"),PromptType: V1AnimationCreateBodyStylePromptTypeEnumCustom,TransitionSpeed: 5,}` |
| `└─ ArtStyle` | ✓ | The art style used to create the output video | `V1AnimationCreateBodyStyleArtStyleEnumPainterlyIllustration` |
| `└─ ArtStyleCustom` | ✗ | Describe custom art style. This field is required if `art_style` is `Custom` | `"string"` |
| `└─ CameraEffect` | ✓ | The camera effect used to create the output video | `V1AnimationCreateBodyStyleCameraEffectEnumSimpleZoomIn` |
| `└─ Prompt` | ✗ | The prompt used for the video. Prompt is required if `prompt_type` is `custom`. Otherwise this value is ignored | `"Cyberpunk city"` |
| `└─ PromptType` | ✓ |  * `custom` - Use your own prompt for the video. * `use_lyrics` - Use the lyrics of the audio to create the prompt. If this option is selected, then `assets.audio_source` must be `file` or `youtube`. * `ai_choose` - Let AI write the prompt. If this option is selected, then `assets.audio_source` must be `file` or `youtube`. | `V1AnimationCreateBodyStylePromptTypeEnumCustom` |
| `└─ TransitionSpeed` | ✓ | Change determines how quickly the video's content changes across frames.  * Higher = more rapid transitions. * Lower = more stable visual experience. | `5` |
| `Width` | ✓ | The width of the final output video. The maximum width depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details | `512` |
| `Name` | ✗ | The name of video. This value is mainly used for your own identification of the video. | `"Animation video"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	animation "github.com/magichourhq/magic-hour-go/resources/v1/animation"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.Animation.Create(animation.CreateRequest{
		Assets: types.V1AnimationCreateBodyAssets{
			AudioFilePath: nullable.NewValue("api-assets/id/1234.mp3"),
			AudioSource:   types.V1AnimationCreateBodyAssetsAudioSourceEnumFile,
			ImageFilePath: nullable.NewValue("api-assets/id/1234.png"),
		},
		EndSeconds: 15.0,
		Fps:        12.0,
		Height:     960,
		Name:       nullable.NewValue("Animation video"),
		Style: types.V1AnimationCreateBodyStyle{
			ArtStyle:        types.V1AnimationCreateBodyStyleArtStyleEnumPainterlyIllustration,
			CameraEffect:    types.V1AnimationCreateBodyStyleCameraEffectEnumSimpleZoomIn,
			Prompt:          nullable.NewValue("Cyberpunk city"),
			PromptType:      types.V1AnimationCreateBodyStylePromptTypeEnumCustom,
			TransitionSpeed: 5,
		},
		Width: 512,
	})
}

```

#### Response

##### Type
[V1AnimationCreateResponse](/types/v1_animation_create_response.go)

##### Example
`V1AnimationCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}`


