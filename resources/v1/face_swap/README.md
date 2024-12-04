
### create <a name="create"></a>
Create Face Swap video

Create a Face Swap video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/face-swap).
  

**API Endpoint**: `POST /v1/face-swap`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	face_swap "github.com/magichourhq/magic-hour-go/resources/v1/face_swap"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.FaceSwap.Create(face_swap.CreateRequest { Data: types.PostV1FaceSwapBody { Assets: types.PostV1FaceSwapBodyAssets { ImageFilePath: "image/id/1234.png", VideoFilePath: nullable.NewValue("video/id/1234.mp4"), VideoSource: types.PostV1FaceSwapBodyAssetsVideoSourceEnumFile, YoutubeUrl: nullable.NewValue("http://www.example.com") }, EndSeconds: 15, Height: 960, Name: nullable.NewValue("Face Swap video"), StartSeconds: 0, Width: 512 } })
```

**Upgrade to see all examples**