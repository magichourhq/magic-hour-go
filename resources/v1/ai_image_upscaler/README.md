# v1.ai_image_upscaler

## Module Functions

### AI Image Upscaler <a name="create"></a>

Upscale your image using AI. Each 2x upscale costs 50 credits for balanced/creative modes, and 25 credits for preserve. 4x upscale costs 200 and 100 credits respectively.

**API Endpoint**: `POST /v1/ai-image-upscaler`

#### Parameters

| Parameter          | Required | Description                                                                                                                                                                                                                                                                                                                                                                                        | Example                                                                                                         |
| ------------------ | :------: | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `Assets`           |    ✓     | Provide the assets for upscaling                                                                                                                                                                                                                                                                                                                                                                   | `V1AiImageUpscalerCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}`                                  |
| `└─ ImageFilePath` |    ✓     | The image to upscale. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details. . The maximum input image size is 4096x4096px. | `"api-assets/id/1234.png"`                                                                                      |
| `ScaleFactor`      |    ✓     | How much to scale the image. Must be either 2 or 4. Note: 4x upscale is only available on Creator, Pro, or Business tier.                                                                                                                                                                                                                                                                          | `2.0`                                                                                                           |
| `Name`             |    ✗     | Give your image a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                             | `"My Image Upscaler image"`                                                                                     |
| `Style`            |    ✗     | Style settings for the upscale. Use `mode` (`"preserve"`, `"balanced"`, or `"creative"`). Defaults to `"balanced"`.                                                                                                                                                                                                                                                                                | `V1AiImageUpscalerCreateBodyStyle {Mode: nullable.NewValue(V1AiImageUpscalerCreateBodyStyleModeEnumBalanced),}` |
| `└─ Enhancement`   |    ✗     | Deprecated: use `mode` instead. `"Resemblance"` maps to `"preserve"`. `"Balanced"` and `"Creative"` map to the same-named modes.                                                                                                                                                                                                                                                                   | `V1AiImageUpscalerCreateBodyStyleEnhancementEnumBalanced`                                                       |
| `└─ Mode`          |    ✗     | The upscaling mode. `"preserve"` uses the fast pro pipeline (1× credit multiplier). `"balanced"` and `"creative"` use the creative pipeline (2× credit multiplier). `"pro"` is deprecated and maps to `"preserve"`. Defaults to `"balanced"`.                                                                                                                                                      | `V1AiImageUpscalerCreateBodyStyleModeEnumBalanced`                                                              |
| `└─ Prompt`        |    ✗     | A prompt to guide the final image. Only used when mode is `creative`.                                                                                                                                                                                                                                                                                                                              | `"string"`                                                                                                      |

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
		Name:        nullable.NewValue("My Image Upscaler image"),
		ScaleFactor: 2.0,
	})
}
```

#### Response

##### Type

[V1AiImageUpscalerCreateResponse](/types/v1_ai_image_upscaler_create_response.go)

##### Example

```go
V1AiImageUpscalerCreateResponse {
CreditsCharged: 50,
FrameCost: 50,
Id: "cuid-example",
}
```
