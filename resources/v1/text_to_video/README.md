
### create <a name="create"></a>
Create Text-to-Video

Create a Text To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/text-to-video).
  

**API Endpoint**: `POST /v1/text-to-video`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	text_to_video "github.com/magichourhq/magic-hour-go/resources/v1/text_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.TextToVideo.Create(text_to_video.CreateRequest { Data: types.PostV1TextToVideoBody { EndSeconds: 5, Name: nullable.NewValue("Text To Video video"), Orientation: types.PostV1TextToVideoBodyOrientationEnumLandscape, Style: types.PostV1TextToVideoBodyStyle { Prompt: "string" } } })
```

**Upgrade to see all examples**
