
### AI Talking Photo <a name="create"></a>

Create a talking photo from an image and audio or text input.

**API Endpoint**: `POST /v1/ai-talking-photo`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for creating a talking photo | `V1AiTalkingPhotoCreateBodyAssets {AudioFilePath: "api-assets/id/1234.mp3",ImageFilePath: "api-assets/id/1234.png",}` |
| `end_seconds` | ✓ | The end time of the input audio in seconds. The maximum duration allowed is 60 seconds. | `15.0` |
| `start_seconds` | ✓ | The start time of the input audio in seconds. The maximum duration allowed is 60 seconds. | `0.0` |
| `name` | ✗ | The name of image | `"Talking Photo image"` |
| `style` | ✗ | Attributes used to dictate the style of the output | `V1AiTalkingPhotoCreateBodyStyle {GenerationMode: nullable.NewValue(V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumExpressive),Intensity: nullable.NewValue(1.5),}` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_talking_photo "github.com/magichourhq/magic-hour-go/resources/v1/ai_talking_photo"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiTalkingPhoto.Create(ai_talking_photo.CreateRequest{
		Assets: types.V1AiTalkingPhotoCreateBodyAssets{
			AudioFilePath: "api-assets/id/1234.mp3",
			ImageFilePath: "api-assets/id/1234.png",
		},
		EndSeconds:   15.0,
		Name:         nullable.NewValue("Talking Photo image"),
		StartSeconds: 0.0,
	})
}

```

#### Response

##### Type
[V1AiTalkingPhotoCreateResponse](/types/v1_ai_talking_photo_create_response.go)

##### Example
`V1AiTalkingPhotoCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
