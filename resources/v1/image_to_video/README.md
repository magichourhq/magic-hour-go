
### create <a name="create"></a>
Image-to-Video

Create a Image To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/image-to-video).
  

**API Endpoint**: `POST /v1/image-to-video`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	image_to_video "github.com/magichourhq/magic-hour-go/resources/v1/image_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))
res, err := client.V1.ImageToVideo.Create(image_to_video.CreateRequest { Assets: types.PostV1ImageToVideoBodyAssets { ImageFilePath: "api-assets/id/1234.png" }, EndSeconds: 5.0, Height: 960, Style: types.PostV1ImageToVideoBodyStyle { Prompt: nullable.NewNull[string]() }, Width: 512 })
```
