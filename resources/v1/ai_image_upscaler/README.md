
### create <a name="create"></a>
AI Image Upscaler

Upscale your image using AI. Each 2x upscale costs 50 frames, and 4x upscale costs 200 frames.

**API Endpoint**: `POST /v1/ai-image-upscaler`

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	ai_image_upscaler "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_upscaler"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiImageUpscaler.Create(ai_image_upscaler.CreateRequest{
		Assets: types.PostV1AiImageUpscalerBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		ScaleFactor: 2.0,
		Style: types.PostV1AiImageUpscalerBodyStyle{
			Enhancement: types.PostV1AiImageUpscalerBodyStyleEnhancementEnumBalanced,
		},
	})
}

```
