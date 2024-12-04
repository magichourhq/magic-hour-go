
### create <a name="create"></a>
Create Animation

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

res, err := client.V1.Animation.Create(animation.CreateRequest { Data: types.PostV1AnimationBody { Assets: types.PostV1AnimationBodyAssets { AudioFilePath: nullable.NewValue("api-assets/id/1234.mp3"), AudioSource: types.PostV1AnimationBodyAssetsAudioSourceEnumFile, ImageFilePath: nullable.NewValue("api-assets/id/1234.png"), YoutubeUrl: nullable.NewValue("http://www.example.com") }, EndSeconds: 15, Fps: 12, Height: 960, Name: nullable.NewValue("Animation video"), Style: types.PostV1AnimationBodyStyle { ArtStyle: types.PostV1AnimationBodyStyleArtStyleEnumPainterlyIllustration, ArtStyleCustom: nullable.NewValue("string"), CameraEffect: types.PostV1AnimationBodyStyleCameraEffectEnumAccelerate, Prompt: nullable.NewValue("Cyberpunk city"), PromptType: types.PostV1AnimationBodyStylePromptTypeEnumAiChoose, TransitionSpeed: 5 }, Width: 512 } })
```

**Upgrade to see all examples**
