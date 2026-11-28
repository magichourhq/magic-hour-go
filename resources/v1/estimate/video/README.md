# v1.estimate.video

## Module Functions

### Estimate video credits <a name="list"></a>

Get the credits a video create request charges when it is submitted. Pass `tool`, then the fields below that the tool's create request uses, with the same values. Tools billed per frame charge an estimate at submission, assuming 30 fps for an uploaded video, and settle to the frames that render. Does not create a project or spend credits. It prices the fields passed here. It does not check your plan, your balance or the rest of the create request.

**API Endpoint**: `GET /v1/estimate/video`

#### Parameters

| Parameter           | Required | Description                                                                                                                                                                                                                                     | Example                                   |
| ------------------- | :------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------- |
| `EndSeconds`        |    ✓     | The `end_seconds` of the create request. Used by every tool.                                                                                                                                                                                    | `5.0`                                     |
| `Tool`              |    ✓     | The create endpoint to price, named by its path. For example `image-to-video` for `POST /v1/image-to-video`.                                                                                                                                    | `V1EstimateVideoListToolEnumImageToVideo` |
| `Audio`             |    ✗     | The `audio` of the create request. Used by `image-to-video`, `text-to-video`. Defaults to `false`.                                                                                                                                              | `true`                                    |
| `HasVideoReference` |    ✗     | Whether any of the create request's `references` is a video. Used by `text-to-video`. Defaults to `false`.                                                                                                                                      | `true`                                    |
| `Model`             |    ✗     | The `model` of the create request. Used by `image-to-video`, `text-to-video`, `ai-ugc-ad`, `ai-video-editor`, `character-replace`.                                                                                                              | `"kling-3.0"`                             |
| `Resolution`        |    ✗     | The `resolution` of the create request. Used by `image-to-video`, `text-to-video`, `ai-ugc-ad`, `ai-video-editor`, `character-replace`, `audio-to-video`, `ai-video-translator`, `ai-video-upscaler`. `ai-video-translator` defaults to `480p`. | `"720p"`                                  |
| `StartSeconds`      |    ✗     | The `start_seconds` of the create request, for tools whose create request has it. Defaults to 0.                                                                                                                                                | `123.0`                                   |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	video "github.com/magichourhq/magic-hour-go/resources/v1/estimate/video"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.Estimate.Video.List(video.ListRequest{
		EndSeconds: 5.0,
		Tool:       types.V1EstimateVideoListToolEnumImageToVideo,
	})
}
```

#### Response

##### Type

[V1EstimateVideoListResponse](/types/v1_estimate_video_list_response.go)

##### Example

```go
V1EstimateVideoListResponse {
EstimatedCredits: 150,
}
```
