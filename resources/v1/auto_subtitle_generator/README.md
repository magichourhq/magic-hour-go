
### Auto Subtitle Generator <a name="create"></a>

Automatically generate subtitles for your video in multiple languages.

**API Endpoint**: `POST /v1/auto-subtitle-generator`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for auto subtitle generator | `V1AutoSubtitleGeneratorCreateBodyAssets {VideoFilePath: "api-assets/id/1234.mp4",}` |
| `end_seconds` | ✓ | The end time of the input video in seconds | `15.0` |
| `start_seconds` | ✓ | The start time of the input video in seconds | `0.0` |
| `style` | ✓ | Style of the subtitle. At least one of `.style.template` or `.style.custom_config` must be provided.  * If only `.style.template` is provided, default values for the template will be used. * If both are provided, the fields in `.style.custom_config` will be used to overwrite the fields in `.style.template`. * If only `.style.custom_config` is provided, then all fields in `.style.custom_config` will be used.  To use custom config only, the following `custom_config` params are required: * `.style.custom_config.font` * `.style.custom_config.text_color` * `.style.custom_config.vertical_position` * `.style.custom_config.horizontal_position`  | `V1AutoSubtitleGeneratorCreateBodyStyle {}` |
| `name` | ✗ | The name of video | `"Auto Subtitle video"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	auto_subtitle_generator "github.com/magichourhq/magic-hour-go/resources/v1/auto_subtitle_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AutoSubtitleGenerator.Create(auto_subtitle_generator.CreateRequest{
		Assets: types.V1AutoSubtitleGeneratorCreateBodyAssets{
			VideoFilePath: "api-assets/id/1234.mp4",
		},
		EndSeconds:   15.0,
		Name:         nullable.NewValue("Auto Subtitle video"),
		StartSeconds: 0.0,
		Style:        types.V1AutoSubtitleGeneratorCreateBodyStyle{},
	})
}

```

#### Response

##### Type
[V1AutoSubtitleGeneratorCreateResponse](/types/v1_auto_subtitle_generator_create_response.go)

##### Example
`V1AutoSubtitleGeneratorCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "clx7uu86w0a5qp55yxz315r6r",
}`
