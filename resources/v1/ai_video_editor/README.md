# v1.ai_video_editor

## Module Functions

### AI Video Editor <a name="create"></a>

**What this API does**

Create the same Video Editor you can make in the browser, but programmatically, so you can automate it, run it at scale, or connect it to your own app or workflow.

**Good for**

- Automation and batch processing
- Adding video editor into apps, pipelines, or tools

**How it works (3 steps)**

1. Upload your inputs (video, image, or audio) with [Generate Upload URLs](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls) and copy the `file_path`.
2. Send a request to create a video editor job with the basic fields.
3. Check the job status until it's `complete`, then download the result from `downloads`.

**Key options**

- Inputs: usually a file, sometimes a YouTube link, depending on project type
- Resolution: free users are limited to 576px; higher plans unlock HD and larger sizes
- Extra fields: e.g. `face_swap_mode`, `start_seconds`/`end_seconds`, or a text prompt

**Cost**\
Credits are only charged for the frames that actually render. You'll see an estimate when the job is queued, and the final total after it's done.

For detailed examples, see the [product page](https://magichour.ai/products/video-editor).

**API Endpoint**: `POST /v1/ai-video-editor`

#### Parameters

| Parameter          | Required | Description                                                                                                                                                                                                                                                                                                                                      | Example                                                                      |
| ------------------ | :------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------- |
| `Assets`           |    ✓     | Provide the assets for video editing.                                                                                                                                                                                                                                                                                                            | `V1AiVideoEditorCreateBodyAssets {VideoFilePath: "api-assets/id/1234.mp4",}` |
| `└─ VideoFilePath` |    ✓     | The video to edit. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details. | `"api-assets/id/1234.mp4"`                                                   |
| `EndSeconds`       |    ✓     | End time of your clip in seconds. Must be greater than `start_seconds`. Duration must be between 3 and 10 seconds.                                                                                                                                                                                                                               | `5.0`                                                                        |
| `Style`            |    ✓     |                                                                                                                                                                                                                                                                                                                                                  | `V1AiVideoEditorCreateBodyStyle {Prompt: "Change the car color to blue",}`   |
| `└─ Prompt`        |    ✓     | The prompt used to edit the video.                                                                                                                                                                                                                                                                                                               | `"Change the car color to blue"`                                             |
| `Model`            |    ✗     | Editing model. Defaults to `ltx-2.3` for free tier and `gemini-omni` for paid. Use `ltx-2.3` for LTX video edit.                                                                                                                                                                                                                                 | `V1AiVideoEditorCreateBodyModelEnumGeminiOmni`                               |
| `Name`             |    ✗     | Give your video a custom name for easy identification.                                                                                                                                                                                                                                                                                           | `"My Video Editor video"`                                                    |
| `Resolution`       |    ✗     | Output resolution. Defaults to `480p` for free tier and `720p` for paid. Google Omni supports 720p only; LTX-2.3 supports 480p, 720p, and 1080p.                                                                                                                                                                                                 | `V1AiVideoEditorCreateBodyResolutionEnum720p`                                |
| `StartSeconds`     |    ✗     | Start time of your clip (seconds). Must be ≥ 0.                                                                                                                                                                                                                                                                                                  | `0.0`                                                                        |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_video_editor "github.com/magichourhq/magic-hour-go/resources/v1/ai_video_editor"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiVideoEditor.Create(ai_video_editor.CreateRequest{
		Assets: types.V1AiVideoEditorCreateBodyAssets{
			VideoFilePath: "api-assets/id/1234.mp4",
		},
		EndSeconds:   5.0,
		Model:        nullable.NewValue(types.V1AiVideoEditorCreateBodyModelEnumGeminiOmni),
		Name:         nullable.NewValue("My Video Editor video"),
		Resolution:   nullable.NewValue(types.V1AiVideoEditorCreateBodyResolutionEnum720p),
		StartSeconds: nullable.NewValue(0.0),
		Style: types.V1AiVideoEditorCreateBodyStyle{
			Prompt: "Change the car color to blue",
		},
	})
}
```

#### Response

##### Type

[V1AiVideoEditorCreateResponse](/types/v1_ai_video_editor_create_response.go)

##### Example

```go
V1AiVideoEditorCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}
```
