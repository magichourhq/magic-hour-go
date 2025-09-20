# v1.ai_qr_code_generator

## Module Functions

### AI QR Code <a name="create"></a>

Create an AI QR code. Each QR code costs 20 credits.

**API Endpoint**: `POST /v1/ai-qr-code-generator`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Content` | ✓ | The content of the QR code. | `"https://magichour.ai"` |
| `Style` | ✓ |  | `V1AiQrCodeGeneratorCreateBodyStyle {ArtStyle: "Watercolor",}` |
| `└─ ArtStyle` | ✓ | To use our templates, pass in one of Watercolor, Cyberpunk City, Ink Landscape, Interior Painting, Japanese Street, Mech, Minecraft, Picasso Painting, Game Map, Spaceship, Chinese Painting, Winter Village, or pass any custom art style. | `"Watercolor"` |
| `Name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Qr Code image"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_qr_code_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_qr_code_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiQrCodeGenerator.Create(ai_qr_code_generator.CreateRequest{
		Content: "https://magichour.ai",
		Name:    nullable.NewValue("Qr Code image"),
		Style: types.V1AiQrCodeGeneratorCreateBodyStyle{
			ArtStyle: "Watercolor",
		},
	})
}

```

#### Response

##### Type
[V1AiQrCodeGeneratorCreateResponse](/types/v1_ai_qr_code_generator_create_response.go)

##### Example
`V1AiQrCodeGeneratorCreateResponse {
CreditsCharged: 20,
FrameCost: 20,
Id: "cuid-example",
}`


