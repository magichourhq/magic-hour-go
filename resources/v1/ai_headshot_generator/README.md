
### create <a name="create"></a>
Create AI Headshots

Create an AI headshot. Each headshot costs 50 frames.

**API Endpoint**: `POST /v1/ai-headshot-generator`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	ai_headshot_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_headshot_generator"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.AiHeadshotGenerator.Create(ai_headshot_generator.CreateRequest { Data: types.PostV1AiHeadshotGeneratorBody { Assets: types.PostV1AiHeadshotGeneratorBodyAssets { ImageFilePath: "image/id/1234.png" }, Name: nullable.NewValue("Ai Headshot image") } })
```

**Upgrade to see all examples**
