
### AI Face Editor <a name="create"></a>

Edit facial features of an image using AI. Each edit costs 1 frame. The height/width of the output image depends on your subscription. Please refer to our [pricing](/pricing) page for more details

**API Endpoint**: `POST /v1/ai-face-editor`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for face editor | `V1AiFaceEditorCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `style` | ✓ | Face editing parameters | `V1AiFaceEditorCreateBodyStyle {EnhanceFace: false,EyeGazeHorizontal: 0.0,EyeGazeVertical: 0.0,EyeOpenRatio: 0.0,EyebrowDirection: 0.0,HeadPitch: 0.0,HeadRoll: 0.0,HeadYaw: 0.0,LipOpenRatio: 0.0,MouthGrim: 0.0,MouthPositionHorizontal: 0.0,MouthPositionVertical: 0.0,MouthPout: 0.0,MouthPurse: 0.0,MouthSmile: 0.0,}` |
| `name` | ✗ | The name of image | `"Face Editor image"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_face_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_face_editor"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiFaceEditor.Create(ai_face_editor.CreateRequest{
		Assets: types.V1AiFaceEditorCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Face Editor image"),
		Style: types.V1AiFaceEditorCreateBodyStyle{
			EnhanceFace:             false,
			EyeGazeHorizontal:       0.0,
			EyeGazeVertical:         0.0,
			EyeOpenRatio:            0.0,
			EyebrowDirection:        0.0,
			HeadPitch:               0.0,
			HeadRoll:                0.0,
			HeadYaw:                 0.0,
			LipOpenRatio:            0.0,
			MouthGrim:               0.0,
			MouthPositionHorizontal: 0.0,
			MouthPositionVertical:   0.0,
			MouthPout:               0.0,
			MouthPurse:              0.0,
			MouthSmile:              0.0,
		},
	})
}

```

#### Response

##### Type
[V1AiFaceEditorCreateResponse](/types/v1_ai_face_editor_create_response.go)

##### Example
`V1AiFaceEditorCreateResponse {
CreditsCharged: 1,
FrameCost: 1,
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
