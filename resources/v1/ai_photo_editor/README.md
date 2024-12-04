
### create <a name="create"></a>
AI Photo Editor

> **NOTE**: this API is still in early development stages, and should be avoided. Please reach out to us if you're interested in this API. 

Edit photo using AI. Each photo costs 10 frames.

**API Endpoint**: `POST /v1/ai-photo-editor`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	ai_photo_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_photo_editor"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.AiPhotoEditor.Create(ai_photo_editor.CreateRequest { Data: types.PostV1AiPhotoEditorBody { Assets: types.PostV1AiPhotoEditorBodyAssets { ImageFilePath: "image/id/1234.png" }, Name: nullable.NewValue("Photo Editor image"), Resolution: 768, Steps: 4, Style: types.PostV1AiPhotoEditorBodyStyle { ImageDescription: "A photo of a person", LikenessStrength: 5.2, NegativePrompt: nullable.NewValue("painting, cartoon, sketch"), Prompt: "A photo portrait of a person wearing a hat", PromptStrength: 3.75 } } })
```

**Upgrade to see all examples**