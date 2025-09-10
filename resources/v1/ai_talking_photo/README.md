# v1.ai_talking_photo

## Module Functions
### AI Talking Photo <a name="create"></a>

Create a talking photo from an image and audio or text input.

**API Endpoint**: `POST /v1/ai-talking-photo`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for creating a talking photo | `V1AiTalkingPhotoCreateBodyAssets {AudioFilePath: "api-assets/id/1234.mp3",ImageFilePath: "api-assets/id/1234.png",}` |
| `└─ AudioFilePath` | ✓ | The audio file to sync with the image. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.mp3"` |
| `└─ ImageFilePath` | ✓ | The source image to animate. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `EndSeconds` | ✓ | The end time of the input audio in seconds. The maximum duration allowed is 60 seconds. | `15.0` |
| `StartSeconds` | ✓ | The start time of the input audio in seconds. The maximum duration allowed is 60 seconds. | `0.0` |
| `Name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Talking Photo image"` |
| `Style` | ✗ | Attributes used to dictate the style of the output | `V1AiTalkingPhotoCreateBodyStyle {GenerationMode: nullable.NewValue(V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumExpressive),Intensity: nullable.NewValue(1.5),}` |
| `└─ GenerationMode` | ✗ | Controls overall motion style. * `pro` -  Realistic, high fidelity, accurate lip sync, slower. * `expressive` - More motion and facial expressiveness; may introduce visual artifacts. * `stable` -  Reduced motion for cleaner output; may result in minimal animation. (Deprecated: passing this value will be treated as `pro`) | `V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumExpressive` |
| `└─ Intensity` | ✗ | Note: this value is only applicable when generation_mode is `expressive`. The value can include up to 2 decimal places. * Lower values yield more stability but can suppress mouth movement. * Higher values increase motion and expressiveness, with a higher risk of distortion. | `1.5` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_talking_photo "github.com/magichourhq/magic-hour-go/resources/v1/ai_talking_photo"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiTalkingPhoto.Create(ai_talking_photo.CreateRequest{
		Assets: types.V1AiTalkingPhotoCreateBodyAssets{
			AudioFilePath: "api-assets/id/1234.mp3",
			ImageFilePath: "api-assets/id/1234.png",
		},
		EndSeconds:   15.0,
		Name:         nullable.NewValue("Talking Photo image"),
		StartSeconds: 0.0,
	})
}

```

#### Response

##### Type
[V1AiTalkingPhotoCreateResponse](/types/v1_ai_talking_photo_create_response.go)

##### Example
`V1AiTalkingPhotoCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}`
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

