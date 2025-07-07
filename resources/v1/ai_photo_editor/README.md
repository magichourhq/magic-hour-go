
### AI Photo Editor <a name="create"></a>

> **NOTE**: this API is still in early development stages, and should be avoided. Please reach out to us if you're interested in this API. 

Edit photo using AI. Each photo costs 10 credits.

**API Endpoint**: `POST /v1/ai-photo-editor`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for photo editor | `V1AiPhotoEditorCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `resolution` | ✓ | The resolution of the final output image. The allowed value is based on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details | `768` |
| `style` | ✓ |  | `V1AiPhotoEditorCreateBodyStyle {ImageDescription: "A photo of a person",LikenessStrength: 5.2,NegativePrompt: nullable.NewValue("painting, cartoon, sketch"),Prompt: "A photo portrait of a person wearing a hat",PromptStrength: 3.75,Steps: nullable.NewValue(4),UpscaleFactor: nullable.NewValue(2),UpscaleFidelity: nullable.NewValue(0.5),}` |
| `name` | ✗ | The name of image | `"Photo Editor image"` |
| `steps` | ✗ | Deprecated: Please use `.style.steps` instead. Number of iterations used to generate the output. Higher values improve quality and increase the strength of the prompt but increase processing time. | `123` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_photo_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_photo_editor"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiPhotoEditor.Create(ai_photo_editor.CreateRequest{
		Assets: types.V1AiPhotoEditorCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name:       nullable.NewValue("Photo Editor image"),
		Resolution: 768,
		Style: types.V1AiPhotoEditorCreateBodyStyle{
			ImageDescription: "A photo of a person",
			LikenessStrength: 5.2,
			NegativePrompt:   nullable.NewValue("painting, cartoon, sketch"),
			Prompt:           "A photo portrait of a person wearing a hat",
			PromptStrength:   3.75,
			Steps:            nullable.NewValue(4),
			UpscaleFactor:    nullable.NewValue(2),
			UpscaleFidelity:  nullable.NewValue(0.5),
		},
	})
}

```

#### Response

##### Type
[V1AiPhotoEditorCreateResponse](/types/v1_ai_photo_editor_create_response.go)

##### Example
`V1AiPhotoEditorCreateResponse {
CreditsCharged: 10,
FrameCost: 10,
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
