
### Text-to-Video <a name="create"></a>

Create a Text To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](https://magichour.ai/products/text-to-video).
  

**API Endpoint**: `POST /v1/text-to-video`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `end_seconds` | ✓ | The total duration of the output video in seconds.  The value must be greater than or equal to 5 seconds and less than or equal to 60 seconds.  Note: For 480p resolution, the value must be either 5 or 10. | `5.0` |
| `orientation` | ✓ | Determines the orientation of the output video | `V1TextToVideoCreateBodyOrientationEnumLandscape` |
| `style` | ✓ |  | `V1TextToVideoCreateBodyStyle {Prompt: "a dog running",}` |
| `name` | ✗ | The name of video. This value is mainly used for your own identification of the video. | `"Text To Video video"` |
| `resolution` | ✗ | Controls the output video resolution. Defaults to `720p` if not specified.  480p and 720p are available on Creator, Pro, or Business tiers. However, 1080p require Pro or Business tier.  **Options:** - `480p` - Supports only 5 or 10 second videos. Output: 24fps. Cost: 120 credits per 5 seconds. - `720p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 300 credits per 5 seconds. - `1080p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 600 credits per 5 seconds. | `V1TextToVideoCreateBodyResolutionEnum720p` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	text_to_video "github.com/magichourhq/magic-hour-go/resources/v1/text_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.TextToVideo.Create(text_to_video.CreateRequest{
		EndSeconds:  5.0,
		Name:        nullable.NewValue("Text To Video video"),
		Orientation: types.V1TextToVideoCreateBodyOrientationEnumLandscape,
		Resolution:  nullable.NewValue(types.V1TextToVideoCreateBodyResolutionEnum720p),
		Style: types.V1TextToVideoCreateBodyStyle{
			Prompt: "a dog running",
		},
	})
}

```

#### Response

##### Type
[V1TextToVideoCreateResponse](/types/v1_text_to_video_create_response.go)

##### Example
`V1TextToVideoCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}`
