
### AI Talking Photo <a name="create"></a>

Create a talking photo from an image and audio or text input.

**API Endpoint**: `POST /v1/ai-talking-photo`

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
