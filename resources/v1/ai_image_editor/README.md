# v1.ai_image_editor

## Module Functions

### AI Image Editor <a name="create"></a>

Edit images with AI. Each edit costs 10 credits.

**API Endpoint**: `POST /v1/ai-image-editor`

#### Parameters

| Parameter           | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Example                                                                                                                                                                                         |
| ------------------- | :------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Assets`            |    ✓     | Provide the assets for image edit                                                                                                                                                                                                                                                                                                                                                                                                                            | `V1AiImageEditorCreateBodyAssets {ImageFilePath: nullable.NewValue("api-assets/id/1234.png"),ImageFilePaths: nullable.NewValue([]string{"api-assets/id/1234.png","api-assets/id/1235.png",}),}` |
| `└─ ImageFilePath`  |    ✗     | Deprecated: Please use `image_file_paths` instead as edits with multiple images are now supported. The image used in the edit. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details. | `"api-assets/id/1234.png"`                                                                                                                                                                      |
| `└─ ImageFilePaths` |    ✗     | The image(s) used in the edit, maximum of 10 images. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                                           | `[]string{"api-assets/id/1234.png","api-assets/id/1235.png",}`                                                                                                                                  |
| `Style`             |    ✓     |                                                                                                                                                                                                                                                                                                                                                                                                                                                              | `V1AiImageEditorCreateBodyStyle {Model: nullable.NewValue(V1AiImageEditorCreateBodyStyleModelEnumNanoBanana),Prompt: "Give me sunglasses",}`                                                    |
| `└─ Model`          |    ✗     | The AI model to use for image editing. * `Nano Banana` - Precise, realistic edits with consistent results * `Seedream` - Creative, imaginative images with artistic freedom * `default` - Use the model we recommend, which will change over time. This is recommended unless you need a specific model. This is the default behavior.                                                                                                                       | `V1AiImageEditorCreateBodyStyleModelEnumNanoBanana`                                                                                                                                             |
| `└─ Prompt`         |    ✓     | The prompt used to edit the image.                                                                                                                                                                                                                                                                                                                                                                                                                           | `"Give me sunglasses"`                                                                                                                                                                          |
| `Name`              |    ✗     | Give your image a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                                       | `"My Ai Image Editor image"`                                                                                                                                                                    |

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
			ImageFilePath: nullable.NewValue("api-assets/id/1234.png"),
			ImageFilePaths: nullable.NewValue([]string{
				"api-assets/id/1234.png",
				"api-assets/id/1235.png",
			}),
		},
		Name: nullable.NewValue("My Ai Image Editor image"),
		Style: types.V1AiImageEditorCreateBodyStyle{
			Model:  nullable.NewValue(types.V1AiImageEditorCreateBodyStyleModelEnumNanoBanana),
			Prompt: "Give me sunglasses",
		},
	})
}
```

#### Response

##### Type

[V1AiImageEditorCreateResponse](/types/v1_ai_image_editor_create_response.go)

##### Example

```go
V1AiImageEditorCreateResponse {
CreditsCharged: 10,
FrameCost: 10,
Id: "cuid-example",
}
```
