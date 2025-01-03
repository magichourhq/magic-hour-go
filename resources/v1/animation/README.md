
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

res, err := client.V1.Animation.Create(animation.CreateRequest { Assets: types.PostV1AnimationBodyAssets { AudioSource: types.PostV1AnimationBodyAssetsAudioSourceEnumFile }, EndSeconds: 15, Fps: 12, Height: 960, Style: types.PostV1AnimationBodyStyle { ArtStyle: types.PostV1AnimationBodyStyleArtStyleEnumPainterlyIllustration, CameraEffect: types.PostV1AnimationBodyStyleCameraEffectEnumAccelerate, Prompt: nullable.NewValue("Cyberpunk city"), PromptType: types.PostV1AnimationBodyStylePromptTypeEnumAiChoose, TransitionSpeed: 5 }, Width: 512 })
```

**Upgrade to see all examples**
