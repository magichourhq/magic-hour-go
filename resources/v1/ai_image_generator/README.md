
### create <a name="create"></a>
Create AI Images

Create an AI image. Each image costs 5 frames.

**API Endpoint**: `POST /v1/ai-image-generator`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	ai_image_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_generator"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.AiImageGenerator.Create(ai_image_generator.CreateRequest { Data: types.PostV1AiImageGeneratorBody { ImageCount: 1, Name: nullable.NewValue("Ai Image image"), Orientation: types.PostV1AiImageGeneratorBodyOrientationEnumLandscape, Style: types.PostV1AiImageGeneratorBodyStyle { Prompt: "Cool image" } } })
```

**Upgrade to see all examples**