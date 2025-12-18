# v1.ai_photo_editor

## Module Functions

### AI Photo Editor <a name="create"></a>

> **NOTE**: this API is still in early development stages, and should be avoided. Please reach out to us if you're interested in this API. 

Edit photo using AI. Each photo costs 10 credits.

**API Endpoint**: `POST /v1/ai-photo-editor`

#### Parameters

| Parameter | Required | Deprecated | Description | Example |
|-----------|:--------:|:----------:|-------------|--------|
| `Assets` | ✓ | ✗ | Provide the assets for photo editor | `V1AiPhotoEditorCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `└─ ImageFilePath` | ✓ | — | The image used to generate the output. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `Resolution` | ✓ | ✗ | The resolution of the final output image. The allowed value is based on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details | `768` |
| `Style` | ✓ | ✗ |  | `V1AiPhotoEditorCreateBodyStyle {ImageDescription: "A photo of a person",LikenessStrength: 5.2,NegativePrompt: nullable.NewValue("painting, cartoon, sketch"),Prompt: "A photo portrait of a person wearing a hat",PromptStrength: 3.75,Steps: nullable.NewValue(4),UpscaleFactor: nullable.NewValue(2),UpscaleFidelity: nullable.NewValue(0.5),}` |
| `└─ ImageDescription` | ✓ | — | Use this to describe what your input image is. This helps maintain aspects of the image you don't want to change. | `"A photo of a person"` |
| `└─ LikenessStrength` | ✓ | — | Determines the input image's influence. Higher values align the output more with the initial image. | `5.2` |
| `└─ NegativePrompt` | ✗ | — | What you want to avoid seeing in the final output; has a minor effect. | `"painting, cartoon, sketch"` |
| `└─ Prompt` | ✓ | — | What you want your final output to look like. We recommend starting with the image description and making minor edits for best results. | `"A photo portrait of a person wearing a hat"` |
| `└─ PromptStrength` | ✓ | — | Determines the prompt's influence. Higher values align the output more with the prompt. | `3.75` |
| `└─ Steps` | ✗ | — | Number of iterations used to generate the output. Higher values improve quality and increase the strength of the prompt but increase processing time. | `4` |
| `└─ UpscaleFactor` | ✗ | — | The multiplier applied to an image's original dimensions during the upscaling process. For example, a scale of 2 doubles the width and height (e.g., from 512x512 to 1024x1024). | `2` |
| `└─ UpscaleFidelity` | ✗ | — | Upscale fidelity refers to the level of quality desired in the generated image. Fidelity value of 1 means more details. | `0.5` |
| `Name` | ✗ | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Photo Editor image"` |
| `Steps` | ✗ | ✓ | Deprecated: Please use `.style.steps` instead. Number of iterations used to generate the output. Higher values improve quality and increase the strength of the prompt but increase processing time. | `123` |

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
Id: "cuid-example",
}`

