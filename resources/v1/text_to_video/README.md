
### Text-to-Video <a name="create"></a>

Create a Text To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/text-to-video).
  

**API Endpoint**: `POST /v1/text-to-video`

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
		Style: types.V1TextToVideoCreateBodyStyle{
			Prompt: "a dog running",
		},
	})
}

```

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `end_seconds` | ✓ | The total duration of the output video in seconds. | `5.0` |
| `orientation` | ✓ | Determines the orientation of the output video | `V1TextToVideoCreateBodyOrientationEnumLandscape` |
| `style` | ✓ |  | `V1TextToVideoCreateBodyStyle {Prompt: "a dog running",}` |
| `name` | ✗ | The name of video | `"Text To Video video"` |
