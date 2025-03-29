
### create <a name="create"></a>
AI Photo Editor

> **NOTE**: this API is still in early development stages, and should be avoided. Please reach out to us if you're interested in this API. 

Edit photo using AI. Each photo costs 10 frames.

**API Endpoint**: `POST /v1/ai-photo-editor`

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
		Resolution: 768,
		Style: types.V1AiPhotoEditorCreateBodyStyle{
			ImageDescription: "A photo of a person",
			LikenessStrength: 5.2,
			NegativePrompt:   nullable.NewValue("painting, cartoon, sketch"),
			Prompt:           "A photo portrait of a person wearing a hat",
			PromptStrength:   3.75,
			Steps:            nullable.NewValue(4),
		},
	})
}

```
