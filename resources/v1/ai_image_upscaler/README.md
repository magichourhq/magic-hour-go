
### create <a name="create"></a>
Create Upscaled Image

Upscale your image using AI. Each 2x upscale costs 50 frames, and 4x upscale costs 200 frames.

**API Endpoint**: `POST /v1/ai-image-upscaler`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	ai_image_upscaler "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_upscaler"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.AiImageUpscaler.Create(ai_image_upscaler.CreateRequest { Data: types.PostV1AiImageUpscalerBody { Assets: types.PostV1AiImageUpscalerBodyAssets { ImageFilePath: "image/id/1234.png" }, Name: nullable.NewValue("Image Upscaler image"), ScaleFactor: 123.45, Style: types.PostV1AiImageUpscalerBodyStyle { Enhancement: types.PostV1AiImageUpscalerBodyStyleEnhancementEnumBalanced, Prompt: nullable.NewValue("string") } } })
```

**Upgrade to see all examples**
