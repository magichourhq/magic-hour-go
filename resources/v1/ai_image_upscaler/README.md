
### create <a name="create"></a>
AI Image Upscaler

Upscale your image using AI. Each 2x upscale costs 50 frames, and 4x upscale costs 200 frames.

**API Endpoint**: `POST /v1/ai-image-upscaler`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	ai_image_upscaler "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_upscaler"
	types "github.com/magichourhq/magic-hour-go/types"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))
res, err := client.V1.AiImageUpscaler.Create(ai_image_upscaler.CreateRequest { Assets: types.PostV1AiImageUpscalerBodyAssets { ImageFilePath: "image/id/1234.png" }, ScaleFactor: 123.45, Style: types.PostV1AiImageUpscalerBodyStyle { Enhancement: types.PostV1AiImageUpscalerBodyStyleEnhancementEnumBalanced } })
```
