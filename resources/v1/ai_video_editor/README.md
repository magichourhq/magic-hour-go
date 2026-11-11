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
- Resolution: output size and plan requirements vary by workflow and model. Some workflows derive output size from the input video.
- Extra fields: e.g. `face_swap_mode`, `start_seconds`/`end_seconds`, or a text prompt

**Cost**\
Credits are charged when the job is queued (`credits_charged` in the response). Tools billed per frame charge an estimate and settle to the frames that actually render once the job finishes; tools billed per second charge the final amount upfront. Failed jobs are refunded.

For detailed examples, see the [product page](https://magichour.ai/products/ai-video-editor).

**API Endpoint**: `POST /v1/ai-video-editor`

#### Parameters

| Parameter          | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Example                                                                                                                                  |
| ------------------ | :------: | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `Assets`           |    ✓     | Provide the assets for video editing.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | `V1AiVideoEditorCreateBodyAssets {VideoFilePath: "api-assets/id/1234.mp4",}`                                                             |
| `└─ VideoFilePath` |    ✓     | The video to edit. This value is either - a direct URL to the file. The URL path must end in a supported file extension. - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                                                                                                                   | `"api-assets/id/1234.mp4"`                                                                                                               |
| `EndSeconds`       |    ✓     | End time of your clip in seconds. Must be greater than `start_seconds`. Allowed duration depends on model: `gemini-omni-1.2`: 3-15s, `wan-3.0`: 2-15s, `ltx-2.5`: 0.5-60s (30s at 1080p), `kling-4.0`: 3-30s.                                                                                                                                                                                                                                                                                                                                      | `5.0`                                                                                                                                    |
| `Style`            |    ✓     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | `V1AiVideoEditorCreateBodyStyle {Prompt: "Change the car color to blue",}`                                                               |
| `└─ Prompt`        |    ✓     | The prompt used to edit the video.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | `"Change the car color to blue"`                                                                                                         |
| `Model`            |    ✗     | The AI model to use for video editing. Defaults to `ltx-2.5` on free tiers and `gemini-omni-1.2` on paid tiers. * `gemini-omni-1.2`: Best for precise clips up to 15s, references, and high-resolution output. Requires a paid tier. * `wan-3.0`: High-quality video with native audio, long clips, and end-frame control. Requires a paid tier. * `ltx-2.5`: Fastest for general scenes, long clips, audio, and rapid iteration. * `kling-4.0`: Create cinematic videos with image and video references and optional audio. Requires a paid tier. | `V1AiVideoEditorCreateBodyModelEnumLtx25`                                                                                                |
| `Name`             |    ✗     | Give your video a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | `"My Video Editor video"`                                                                                                                |
| `References`       |    ✗     | Name references without @ and refer to them as @name in style.prompt. Names must be unique, ignoring case. Maximum 10 references total. Omit for no references. Reference images for gemini-omni-1.2, wan-3.0, kling-4.0.                                                                                                                                                                                                                                                                                                                          | `[]V1AiVideoEditorCreateBodyReferencesItem{V1AiVideoEditorCreateBodyReferencesItem {FilePath: "api-assets/id/1234.png",Name: "hero",},}` |
| `Resolution`       |    ✗     | Output resolution. Defaults to `480p` for free tier and `720p` for paid. Supported per model: `gemini-omni-1.2`: 360p, 720p, 1080p, 4k; `wan-3.0`: 480p, 720p, 1080p; `ltx-2.5`: 480p, 720p, 1080p; `kling-4.0`: 720p, 1080p.                                                                                                                                                                                                                                                                                                                      | `V1AiVideoEditorCreateBodyResolutionEnum480p`                                                                                            |
| `StartSeconds`     |    ✗     | Start time of your clip (seconds). Must be ≥ 0.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | `0.0`                                                                                                                                    |

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
		Model:        nullable.NewValue(types.V1AiVideoEditorCreateBodyModelEnumLtx25),
		Name:         nullable.NewValue("My Video Editor video"),
		Resolution:   nullable.NewValue(types.V1AiVideoEditorCreateBodyResolutionEnum480p),
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
EstimatedFrameCost: 123,
Id: "cuid-example",
}
```
