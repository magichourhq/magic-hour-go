
### create <a name="create"></a>
AI Images

Create an AI image. Each image costs 5 frames.

**API Endpoint**: `POST /v1/ai-image-generator`

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_image_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiImageGenerator.Create(ai_image_generator.CreateRequest{
		ImageCount:  1,
		Name:        nullable.NewValue("Ai Image image"),
		Orientation: types.V1AiImageGeneratorCreateBodyOrientationEnumLandscape,
		Style: types.V1AiImageGeneratorCreateBodyStyle{
			Prompt: "Cool image",
		},
	})
}

```
