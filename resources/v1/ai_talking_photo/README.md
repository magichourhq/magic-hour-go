# v1.ai_talking_photo

## Module Functions

### AI Talking Photo <a name="create"></a>

Create a talking photo from an image and audio or text input.

**API Endpoint**: `POST /v1/ai-talking-photo`

#### Parameters

| Parameter           | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Example                                                                                                                                                        |
| ------------------- | :------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Assets`            |    ✓     | Provide the assets for creating a talking photo                                                                                                                                                                                                                                                                                                                                                                                                                                           | `V1AiTalkingPhotoCreateBodyAssets {AudioFilePath: "api-assets/id/1234.mp3",ImageFilePath: "api-assets/id/1234.png",}`                                          |
| `└─ AudioFilePath`  |    ✓     | The audio file to sync with the image. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                                                                                      | `"api-assets/id/1234.mp3"`                                                                                                                                     |
| `└─ ImageFilePath`  |    ✓     | The source image to animate. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                                                                                                | `"api-assets/id/1234.png"`                                                                                                                                     |
| `EndSeconds`        |    ✓     | The end time of the input audio in seconds. The maximum duration allowed is 60 seconds.                                                                                                                                                                                                                                                                                                                                                                                                   | `15.0`                                                                                                                                                         |
| `StartSeconds`      |    ✓     | The start time of the input audio in seconds. The maximum duration allowed is 60 seconds.                                                                                                                                                                                                                                                                                                                                                                                                 | `0.0`                                                                                                                                                          |
| `MaxResolution`     |    ✗     | Constrains the larger dimension (height or width) of the output video. Allows you to set a lower resolution than your plan's maximum if desired. The value is capped by your plan's max resolution.                                                                                                                                                                                                                                                                                       | `1024`                                                                                                                                                         |
| `Name`              |    ✗     | Give your image a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                                                                    | `"My Talking Photo image"`                                                                                                                                     |
| `Style`             |    ✗     | Attributes used to dictate the style of the output                                                                                                                                                                                                                                                                                                                                                                                                                                        | `V1AiTalkingPhotoCreateBodyStyle {GenerationMode: nullable.NewValue(V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumPro),Intensity: nullable.NewValue(1.5),}` |
| `└─ GenerationMode` |    ✗     | Controls overall motion style. * `pro` - Higher fidelity, realistic detail, accurate lip sync, and faster generation. * `standard` - More expressive motion, but lower visual fidelity. * `expressive` - More motion and facial expressiveness; may introduce visual artifacts. (Deprecated: passing this value will be treated as `standard`) * `stable` - Reduced motion for cleaner output; may result in minimal animation. (Deprecated: passing this value will be treated as `pro`) | `V1AiTalkingPhotoCreateBodyStyleGenerationModeEnumPro`                                                                                                         |
| `└─ Intensity`      |    ✗     | Note: this value is only applicable when generation_mode is `expressive`. The value can include up to 2 decimal places. * Lower values yield more stability but can suppress mouth movement. * Higher values increase motion and expressiveness, with a higher risk of distortion.                                                                                                                                                                                                        | `1.5`                                                                                                                                                          |

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
		EndSeconds:    15.0,
		MaxResolution: nullable.NewValue(1024),
		Name:          nullable.NewValue("My Talking Photo image"),
		StartSeconds:  0.0,
	})
}
```

#### Response

##### Type

[V1AiTalkingPhotoCreateResponse](/types/v1_ai_talking_photo_create_response.go)

##### Example

```go
V1AiTalkingPhotoCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}
```
