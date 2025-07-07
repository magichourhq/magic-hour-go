
### AI Headshots <a name="create"></a>

Create an AI headshot. Each headshot costs 50 credits.

**API Endpoint**: `POST /v1/ai-headshot-generator`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for headshot photo | `V1AiHeadshotGeneratorCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `name` | ✗ | The name of image | `"Ai Headshot image"` |
| `style` | ✗ |  | `V1AiHeadshotGeneratorCreateBodyStyle {Prompt: nullable.NewValue("professional passport photo, business attire, smiling, good posture, light blue background, centered, plain background"),}` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_headshot_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_headshot_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiHeadshotGenerator.Create(ai_headshot_generator.CreateRequest{
		Assets: types.V1AiHeadshotGeneratorCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Ai Headshot image"),
	})
}

```

#### Response

##### Type
[V1AiHeadshotGeneratorCreateResponse](/types/v1_ai_headshot_generator_create_response.go)

##### Example
`V1AiHeadshotGeneratorCreateResponse {
CreditsCharged: 50,
FrameCost: 50,
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
