# v1.text_to_video

## Module Functions

### Text-to-Video <a name="create"></a>

**What this API does**

Create the same Text To Video you can make in the browser, but programmatically, so you can automate it, run it at scale, or connect it to your own app or workflow.

**Good for**

- Automation and batch processing
- Adding text to video into apps, pipelines, or tools

**How it works (3 steps)**

1. Upload your inputs (video, image, or audio) with [Generate Upload URLs](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls) and copy the `file_path`.
2. Send a request to create a text to video job with the basic fields.
3. Check the job status until it's `complete`, then download the result from `downloads`.

**Key options**

- Inputs: usually a file, sometimes a YouTube link, depending on project type
- Resolution: free users are limited to 512px; higher plans unlock HD and larger sizes
- Extra fields: e.g. `face_swap_mode`, `start_seconds`/`end_seconds`, or a text prompt

**Cost**\
Credits are only charged for the frames that actually render. You'll see an estimate when the job is queued, and the final total after it's done.

For detailed examples, see the [product page](https://magichour.ai/products/text-to-video).

**API Endpoint**: `POST /v1/text-to-video`

#### Parameters

| Parameter        | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Example                                                   |
| ---------------- | :------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| `EndSeconds`     |    ✓     | The total duration of the output video in seconds. The value must be greater than or equal to 5 seconds and less than or equal to 60 seconds. Note: For 480p resolution, the value must be either 5 or 10.                                                                                                                                                                                                                                                                                              | `5.0`                                                     |
| `Orientation`    |    ✓     | Determines the orientation of the output video                                                                                                                                                                                                                                                                                                                                                                                                                                                          | `V1TextToVideoCreateBodyOrientationEnumLandscape`         |
| `Style`          |    ✓     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | `V1TextToVideoCreateBodyStyle {Prompt: "a dog running",}` |
| `└─ Prompt`      |    ✓     | The prompt used for the video.                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | `"a dog running"`                                         |
| `└─ QualityMode` |    ✗     | DEPRECATED: Please use `resolution` field instead. For backward compatibility: * `quick` maps to 720p resolution * `studio` maps to 1080p resolution This field will be removed in a future version. Use the `resolution` field to directly to specify the resolution.                                                                                                                                                                                                                                  | `V1TextToVideoCreateBodyStyleQualityModeEnumQuick`        |
| `Name`           |    ✗     | Give your video a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                                                                                  | `"My Text To Video video"`                                |
| `Resolution`     |    ✗     | Controls the output video resolution. Defaults to `720p` if not specified. 480p and 720p are available on Creator, Pro, or Business tiers. However, 1080p require Pro or Business tier. **Options:** - `480p` - Supports only 5 or 10 second videos. Output: 24fps. Cost: 120 credits per 5 seconds. - `720p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 300 credits per 5 seconds. - `1080p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 600 credits per 5 seconds. | `V1TextToVideoCreateBodyResolutionEnum720p`               |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	text_to_video "github.com/magichourhq/magic-hour-go/resources/v1/text_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.TextToVideo.Create(text_to_video.CreateRequest{
		EndSeconds:  5.0,
		Name:        nullable.NewValue("My Text To Video video"),
		Orientation: types.V1TextToVideoCreateBodyOrientationEnumLandscape,
		Resolution:  nullable.NewValue(types.V1TextToVideoCreateBodyResolutionEnum720p),
		Style: types.V1TextToVideoCreateBodyStyle{
			Prompt: "a dog running",
		},
	})
}
```

#### Response

##### Type

[V1TextToVideoCreateResponse](/types/v1_text_to_video_create_response.go)

##### Example

```go
V1TextToVideoCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}
```
