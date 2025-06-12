
### AI Image Upscaler <a name="create"></a>

Upscale your image using AI. Each 2x upscale costs 50 credits, and 4x upscale costs 200 credits.

**API Endpoint**: `POST /v1/ai-image-upscaler`

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_image_upscaler "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_upscaler"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiImageUpscaler.Create(ai_image_upscaler.CreateRequest{
		Assets: types.V1AiImageUpscalerCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name:        nullable.NewValue("Image Upscaler image"),
		ScaleFactor: 2.0,
		Style: types.V1AiImageUpscalerCreateBodyStyle{
			Enhancement: types.V1AiImageUpscalerCreateBodyStyleEnhancementEnumBalanced,
		},
	})
}

```

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for upscaling | `V1AiImageUpscalerCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `scale_factor` | ✓ | How much to scale the image. Must be either 2 or 4 | `2.0` |
| `style` | ✓ |  | `V1AiImageUpscalerCreateBodyStyle {Enhancement: V1AiImageUpscalerCreateBodyStyleEnhancementEnumBalanced,}` |
| `name` | ✗ | The name of image | `"Image Upscaler image"` |
