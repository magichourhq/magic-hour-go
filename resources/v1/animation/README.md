
### create <a name="create"></a>
Animation

Create a Animation video. The estimated frame cost is calculated based on the `fps` and `end_seconds` input.

**API Endpoint**: `POST /v1/animation`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	animation "github.com/magichourhq/magic-hour-go/resources/v1/animation"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))
res, err := client.V1.Animation.Create(animation.CreateRequest { Assets: types.PostV1AnimationBodyAssets { AudioFilePath: nullable.NewValue("api-assets/id/1234.mp3"), AudioSource: types.PostV1AnimationBodyAssetsAudioSourceEnumFile, ImageFilePath: nullable.NewValue("api-assets/id/1234.png") }, EndSeconds: 15.0, Fps: 12.0, Height: 960, Style: types.PostV1AnimationBodyStyle { ArtStyle: types.PostV1AnimationBodyStyleArtStyleEnumPainterlyIllustration, CameraEffect: types.PostV1AnimationBodyStyleCameraEffectEnumAccelerate, Prompt: nullable.NewValue("Cyberpunk city"), PromptType: types.PostV1AnimationBodyStylePromptTypeEnumAiChoose, TransitionSpeed: 5 }, Width: 512 })
```
