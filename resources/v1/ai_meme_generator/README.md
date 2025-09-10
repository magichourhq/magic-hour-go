# v1.ai_meme_generator

## Module Functions
### AI Meme Generator <a name="create"></a>

Create an AI generated meme. Each meme costs 10 credits.

**API Endpoint**: `POST /v1/ai-meme-generator`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Style` | ✓ |  | `V1AiMemeGeneratorCreateBodyStyle {SearchWeb: nullable.NewValue(false),Template: V1AiMemeGeneratorCreateBodyStyleTemplateEnumDrakeHotlineBling,Topic: "When the code finally works",}` |
| `└─ SearchWeb` | ✗ | Whether to search the web for meme content. | `false` |
| `└─ Template` | ✓ | To use our templates, pass in one of the enum values. | `V1AiMemeGeneratorCreateBodyStyleTemplateEnumDrakeHotlineBling` |
| `└─ Topic` | ✓ | The topic of the meme. | `"When the code finally works"` |
| `Name` | ✗ | The name of the meme. | `"My Funny Meme"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_meme_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_meme_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiMemeGenerator.Create(ai_meme_generator.CreateRequest{
		Name: nullable.NewValue("My Funny Meme"),
		Style: types.V1AiMemeGeneratorCreateBodyStyle{
			SearchWeb: nullable.NewValue(false),
			Template:  types.V1AiMemeGeneratorCreateBodyStyleTemplateEnumDrakeHotlineBling,
			Topic:     "When the code finally works",
		},
	})
}

```

#### Response

##### Type
[V1AiMemeGeneratorCreateResponse](/types/v1_ai_meme_generator_create_response.go)

##### Example
`V1AiMemeGeneratorCreateResponse {
CreditsCharged: 10,
FrameCost: 10,
Id: "cuid-example",
}`
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

