
### AI Image Editor <a name="create"></a>

Edit images with AI. Each edit costs 50 credits.

**API Endpoint**: `POST /v1/ai-image-editor`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for image edit | `V1AiImageEditorCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `style` | ✓ |  | `V1AiImageEditorCreateBodyStyle {Prompt: "Give me sunglasses",}` |
| `name` | ✗ | The name of image | `"Ai Image Editor image"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_image_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_editor"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiImageEditor.Create(ai_image_editor.CreateRequest{
		Assets: types.V1AiImageEditorCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Ai Image Editor image"),
		Style: types.V1AiImageEditorCreateBodyStyle{
			Prompt: "Give me sunglasses",
		},
	})
}

```

#### Response

##### Type
[V1AiImageEditorCreateResponse](/types/v1_ai_image_editor_create_response.go)

##### Example
`V1AiImageEditorCreateResponse {
CreditsCharged: 50,
FrameCost: 50,
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
