# v1.ai_gif_generator

## Module Functions
### AI GIFs <a name="create"></a>

Create an AI GIF. Each GIF costs 50 credits.

**API Endpoint**: `POST /v1/ai-gif-generator`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Style` | ✓ |  | `V1AiGifGeneratorCreateBodyStyle {Prompt: "Cute dancing cat, pixel art",}` |
| `└─ Prompt` | ✓ | The prompt used for the GIF. | `"Cute dancing cat, pixel art"` |
| `Name` | ✗ | The name of gif. This value is mainly used for your own identification of the gif. | `"Ai Gif gif"` |
| `OutputFormat` | ✗ | The output file format for the generated animation. | `V1AiGifGeneratorCreateBodyOutputFormatEnumGif` |

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
		Name:         nullable.NewValue("Ai Gif gif"),
		OutputFormat: nullable.NewValue(types.V1AiGifGeneratorCreateBodyOutputFormatEnumGif),
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
CreditsCharged: 50,
FrameCost: 50,
Id: "cuid-example",
}`
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

