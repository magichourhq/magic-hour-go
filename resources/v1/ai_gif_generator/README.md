
### create <a name="create"></a>
AI GIFs

Create an AI GIF. Each GIF costs 5 frames.

**API Endpoint**: `POST /v1/ai-gif-generator`

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_gif_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_gif_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiGifGenerator.Create(ai_gif_generator.CreateRequest{
		Name: nullable.NewValue("Ai Gif gif"),
		Style: types.V1AiGifGeneratorCreateBodyStyle{
			Prompt: "Cute dancing cat, pixel art",
		},
	})
}

```
