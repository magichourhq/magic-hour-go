# v1.ai_image_editor

## Module Functions
### AI Image Editor <a name="create"></a>

Edit images with AI. Each edit costs 50 credits.

**API Endpoint**: `POST /v1/ai-image-editor`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for image edit | `V1AiImageEditorCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `└─ ImageFilePath` | ✓ | The image used in the edit. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `Style` | ✓ |  | `V1AiImageEditorCreateBodyStyle {Prompt: "Give me sunglasses",}` |
| `└─ Prompt` | ✓ | The prompt used to edit the image. | `"Give me sunglasses"` |
| `Name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Ai Image Editor image"` |

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
Id: "cuid-example",
}`
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

