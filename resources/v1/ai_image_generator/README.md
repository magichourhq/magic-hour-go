
### AI Images <a name="create"></a>

Create an AI image. Each image costs 5 credits.

**API Endpoint**: `POST /v1/ai-image-generator`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `image_count` | ✓ | Number of images to generate. | `1` |
| `orientation` | ✓ | The orientation of the output image(s). | `V1AiImageGeneratorCreateBodyOrientationEnumLandscape` |
| `style` | ✓ | The art style to use for image generation. | `V1AiImageGeneratorCreateBodyStyle {Prompt: "Cool image",Tool: nullable.NewValue(V1AiImageGeneratorCreateBodyStyleToolEnumAiAnimeGenerator),}` |
| `name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Ai Image image"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_image_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiImageGenerator.Create(ai_image_generator.CreateRequest{
		ImageCount:  1,
		Name:        nullable.NewValue("Ai Image image"),
		Orientation: types.V1AiImageGeneratorCreateBodyOrientationEnumLandscape,
		Style: types.V1AiImageGeneratorCreateBodyStyle{
			Prompt: "Cool image",
			Tool:   nullable.NewValue(types.V1AiImageGeneratorCreateBodyStyleToolEnumAiAnimeGenerator),
		},
	})
}

```

#### Response

##### Type
[V1AiImageGeneratorCreateResponse](/types/v1_ai_image_generator_create_response.go)

##### Example
`V1AiImageGeneratorCreateResponse {
CreditsCharged: 5,
FrameCost: 5,
Id: "cuid-example",
}`
