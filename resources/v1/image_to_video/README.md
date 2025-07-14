
### Image-to-Video <a name="create"></a>

Create a Image To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/image-to-video).
  

**API Endpoint**: `POST /v1/image-to-video`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for image-to-video. | `V1ImageToVideoCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `end_seconds` | ✓ | The total duration of the output video in seconds. | `5.0` |
| `style` | ✓ | Attributed used to dictate the style of the output | `V1ImageToVideoCreateBodyStyle {Prompt: nullable.NewValue("a dog running"),}` |
| `height` | ✗ | This field does not affect the output video's resolution. The video's orientation will match that of the input image.  It is retained solely for backward compatibility and will be deprecated in the future. | `960` |
| `name` | ✗ | The name of video | `"Image To Video video"` |
| `resolution` | ✗ | Controls the output video resolution. Defaults to `720p` if not specified.  **Options:** - `480p` - Supports only 5 or 10 second videos. Output: 24fps. Cost: 120 credits per 5 seconds. - `720p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 300 credits per 5 seconds. - `1080p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 600 credits per 5 seconds. **Requires** `pro` or `business` tier. | `V1ImageToVideoCreateBodyResolutionEnum1080p` |
| `width` | ✗ | This field does not affect the output video's resolution. The video's orientation will match that of the input image.  It is retained solely for backward compatibility and will be deprecated in the future. | `512` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	image_to_video "github.com/magichourhq/magic-hour-go/resources/v1/image_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.ImageToVideo.Create(image_to_video.CreateRequest{
		Assets: types.V1ImageToVideoCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		EndSeconds: 5.0,
		Height:     nullable.NewValue(960),
		Name:       nullable.NewValue("Image To Video video"),
		Style: types.V1ImageToVideoCreateBodyStyle{
			Prompt: nullable.NewValue("a dog running"),
		},
		Width: nullable.NewValue(512),
	})
}

```

#### Response

##### Type
[V1ImageToVideoCreateResponse](/types/v1_image_to_video_create_response.go)

##### Example
`V1ImageToVideoCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
