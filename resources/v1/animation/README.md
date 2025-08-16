
### Animation <a name="create"></a>

Create a Animation video. The estimated frame cost is calculated based on the `fps` and `end_seconds` input.

**API Endpoint**: `POST /v1/animation`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for animation. | `V1AnimationCreateBodyAssets {AudioFilePath: nullable.NewValue("api-assets/id/1234.mp3"),AudioSource: V1AnimationCreateBodyAssetsAudioSourceEnumFile,ImageFilePath: nullable.NewValue("api-assets/id/1234.png"),}` |
| `end_seconds` | ✓ | This value determines the duration of the output video. | `15.0` |
| `fps` | ✓ | The desire output video frame rate | `12.0` |
| `height` | ✓ | The height of the final output video. The maximum height depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details | `960` |
| `style` | ✓ | Defines the style of the output video | `V1AnimationCreateBodyStyle {ArtStyle: V1AnimationCreateBodyStyleArtStyleEnumPainterlyIllustration,CameraEffect: V1AnimationCreateBodyStyleCameraEffectEnumSimpleZoomIn,Prompt: nullable.NewValue("Cyberpunk city"),PromptType: V1AnimationCreateBodyStylePromptTypeEnumCustom,TransitionSpeed: 5,}` |
| `width` | ✓ | The width of the final output video. The maximum width depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details | `512` |
| `name` | ✗ | The name of video. This value is mainly used for your own identification of the video. | `"Animation video"` |

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
