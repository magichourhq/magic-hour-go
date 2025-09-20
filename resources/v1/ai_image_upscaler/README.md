# v1.ai_image_upscaler

## Module Functions

### AI Image Upscaler <a name="create"></a>

Upscale your image using AI. Each 2x upscale costs 50 credits, and 4x upscale costs 200 credits.

**API Endpoint**: `POST /v1/ai-image-upscaler`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for upscaling | `V1AiImageUpscalerCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `└─ ImageFilePath` | ✓ | The image to upscale. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `ScaleFactor` | ✓ | How much to scale the image. Must be either 2 or 4.              Note: 4x upscale is only available on Creator, Pro, or Business tier. | `2.0` |
| `Style` | ✓ |  | `V1AiImageUpscalerCreateBodyStyle {Enhancement: V1AiImageUpscalerCreateBodyStyleEnhancementEnumBalanced,}` |
| `└─ Enhancement` | ✓ |  | `V1AiImageUpscalerCreateBodyStyleEnhancementEnumBalanced` |
| `└─ Prompt` | ✗ | A prompt to guide the final image. This value is ignored if `enhancement` is not Creative | `"string"` |
| `Name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Image Upscaler image"` |

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

#### Response

##### Type
[V1AiImageUpscalerCreateResponse](/types/v1_ai_image_upscaler_create_response.go)

##### Example
`V1AiImageUpscalerCreateResponse {
CreditsCharged: 50,
FrameCost: 50,
Id: "cuid-example",
}`


