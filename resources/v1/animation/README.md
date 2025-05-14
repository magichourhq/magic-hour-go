
### Animation <a name="create"></a>

Create a Animation video. The estimated frame cost is calculated based on the `fps` and `end_seconds` input.

**API Endpoint**: `POST /v1/animation`

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
			CameraEffect:    types.V1AnimationCreateBodyStyleCameraEffectEnumAccelerate,
			Prompt:          nullable.NewValue("Cyberpunk city"),
			PromptType:      types.V1AnimationCreateBodyStylePromptTypeEnumAiChoose,
			TransitionSpeed: 5,
		},
		Width: 512,
	})
}

```
