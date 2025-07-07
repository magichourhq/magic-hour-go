
### AI GIFs <a name="create"></a>

Create an AI GIF. Each GIF costs 25 credits.

**API Endpoint**: `POST /v1/ai-gif-generator`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `style` | ✓ |  | `V1AiGifGeneratorCreateBodyStyle {Prompt: "Cute dancing cat, pixel art",}` |
| `name` | ✗ | The name of gif | `"Ai Gif gif"` |

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

#### Response

##### Type
[V1AiGifGeneratorCreateResponse](/types/v1_ai_gif_generator_create_response.go)

##### Example
`V1AiGifGeneratorCreateResponse {
CreditsCharged: 25,
FrameCost: 25,
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
