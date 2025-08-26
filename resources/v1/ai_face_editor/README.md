# v1_aifaceeditor

## Module Functions
### AI Face Editor <a name="create"></a>

Edit facial features of an image using AI. Each edit costs 1 frame. The height/width of the output image depends on your subscription. Please refer to our [pricing](/pricing) page for more details

**API Endpoint**: `POST /v1/ai-face-editor`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for face editor | `V1AiFaceEditorCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `└─ ImageFilePath` | ✓ | This is the image whose face will be edited. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `Style` | ✓ | Face editing parameters | `V1AiFaceEditorCreateBodyStyle {EnhanceFace: nullable.NewValue(false),EyeGazeHorizontal: nullable.NewValue(0.0),EyeGazeVertical: nullable.NewValue(0.0),EyeOpenRatio: nullable.NewValue(0.0),EyebrowDirection: nullable.NewValue(0.0),HeadPitch: nullable.NewValue(0.0),HeadRoll: nullable.NewValue(0.0),HeadYaw: nullable.NewValue(0.0),LipOpenRatio: nullable.NewValue(0.0),MouthGrim: nullable.NewValue(0.0),MouthPositionHorizontal: nullable.NewValue(0.0),MouthPositionVertical: nullable.NewValue(0.0),MouthPout: nullable.NewValue(0.0),MouthPurse: nullable.NewValue(0.0),MouthSmile: nullable.NewValue(0.0),}` |
| `└─ EnhanceFace` | ✗ | Enhance face features | `false` |
| `└─ EyeGazeHorizontal` | ✗ | Horizontal eye gaze (-100 to 100), in increments of 5 | `0.0` |
| `└─ EyeGazeVertical` | ✗ | Vertical eye gaze (-100 to 100), in increments of 5 | `0.0` |
| `└─ EyeOpenRatio` | ✗ | Eye open ratio (-100 to 100), in increments of 5 | `0.0` |
| `└─ EyebrowDirection` | ✗ | Eyebrow direction (-100 to 100), in increments of 5 | `0.0` |
| `└─ HeadPitch` | ✗ | Head pitch (-100 to 100), in increments of 5 | `0.0` |
| `└─ HeadRoll` | ✗ | Head roll (-100 to 100), in increments of 5 | `0.0` |
| `└─ HeadYaw` | ✗ | Head yaw (-100 to 100), in increments of 5 | `0.0` |
| `└─ LipOpenRatio` | ✗ | Lip open ratio (-100 to 100), in increments of 5 | `0.0` |
| `└─ MouthGrim` | ✗ | Mouth grim (-100 to 100), in increments of 5 | `0.0` |
| `└─ MouthPositionHorizontal` | ✗ | Horizontal mouth position (-100 to 100), in increments of 5 | `0.0` |
| `└─ MouthPositionVertical` | ✗ | Vertical mouth position (-100 to 100), in increments of 5 | `0.0` |
| `└─ MouthPout` | ✗ | Mouth pout (-100 to 100), in increments of 5 | `0.0` |
| `└─ MouthPurse` | ✗ | Mouth purse (-100 to 100), in increments of 5 | `0.0` |
| `└─ MouthSmile` | ✗ | Mouth smile (-100 to 100), in increments of 5 | `0.0` |
| `Name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Face Editor image"` |

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
			EnhanceFace:             nullable.NewValue(false),
			EyeGazeHorizontal:       nullable.NewValue(0.0),
			EyeGazeVertical:         nullable.NewValue(0.0),
			EyeOpenRatio:            nullable.NewValue(0.0),
			EyebrowDirection:        nullable.NewValue(0.0),
			HeadPitch:               nullable.NewValue(0.0),
			HeadRoll:                nullable.NewValue(0.0),
			HeadYaw:                 nullable.NewValue(0.0),
			LipOpenRatio:            nullable.NewValue(0.0),
			MouthGrim:               nullable.NewValue(0.0),
			MouthPositionHorizontal: nullable.NewValue(0.0),
			MouthPositionVertical:   nullable.NewValue(0.0),
			MouthPout:               nullable.NewValue(0.0),
			MouthPurse:              nullable.NewValue(0.0),
			MouthSmile:              nullable.NewValue(0.0),
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
Id: "cuid-example",
}`
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

