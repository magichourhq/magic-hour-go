### create <a name="create"></a>

AI Headshots

Create an AI headshot. Each headshot costs 50 frames.

**API Endpoint**: `POST /v1/ai-headshot-generator`

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	ai_headshot_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_headshot_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiHeadshotGenerator.Create(ai_headshot_generator.CreateRequest{
		Assets: types.PostV1AiHeadshotGeneratorBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
	})
}
```
