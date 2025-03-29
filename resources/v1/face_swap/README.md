
### create <a name="create"></a>
Face Swap video

Create a Face Swap video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/face-swap).
  

**API Endpoint**: `POST /v1/face-swap`

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	face_swap "github.com/magichourhq/magic-hour-go/resources/v1/face_swap"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.FaceSwap.Create(face_swap.CreateRequest{
		Assets: types.V1FaceSwapCreateBodyAssets{
			ImageFilePath: "image/id/1234.png",
			VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
			VideoSource:   types.V1FaceSwapCreateBodyAssetsVideoSourceEnumFile,
		},
		EndSeconds:   15.0,
		Height:       960,
		StartSeconds: 0.0,
		Width:        512,
	})
}

```
