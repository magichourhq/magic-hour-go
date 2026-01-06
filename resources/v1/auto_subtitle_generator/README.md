# v1.auto_subtitle_generator

## Module Functions

### Auto Subtitle Generator <a name="create"></a>

Automatically generate subtitles for your video in multiple languages.

**API Endpoint**: `POST /v1/auto-subtitle-generator`

#### Parameters

| Parameter          | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Example                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| ------------------ | :------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Assets`           |    ✓     | Provide the assets for auto subtitle generator                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | `V1AutoSubtitleGeneratorCreateBodyAssets {VideoFilePath: "api-assets/id/1234.mp4",}`                                                                                                                                                                                                                                                                                                                                                         |
| `└─ VideoFilePath` |    ✓     | This is the video used to add subtitles. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.                                                                                                                                                                                                                                                                                            | `"api-assets/id/1234.mp4"`                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `EndSeconds`       |    ✓     | End time of your clip (seconds). Must be greater than start_seconds.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | `15.0`                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `StartSeconds`     |    ✓     | Start time of your clip (seconds). Must be ≥ 0.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | `0.0`                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `Style`            |    ✓     | Style of the subtitle. At least one of `.style.template` or `.style.custom_config` must be provided. * If only `.style.template` is provided, default values for the template will be used. * If both are provided, the fields in `.style.custom_config` will be used to overwrite the fields in `.style.template`. * If only `.style.custom_config` is provided, then all fields in `.style.custom_config` will be used. To use custom config only, the following `custom_config` params are required: * `.style.custom_config.font` * `.style.custom_config.text_color` * `.style.custom_config.vertical_position` * `.style.custom_config.horizontal_position` | `V1AutoSubtitleGeneratorCreateBodyStyle {}`                                                                                                                                                                                                                                                                                                                                                                                                  |
| `└─ CustomConfig`  |    ✗     | Custom subtitle configuration.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | `V1AutoSubtitleGeneratorCreateBodyStyleCustomConfig {Font: nullable.NewValue("Noto Sans"),FontSize: nullable.NewValue(24.0),FontStyle: nullable.NewValue("normal"),HighlightedTextColor: nullable.NewValue("#FFD700"),HorizontalPosition: nullable.NewValue("center"),StrokeColor: nullable.NewValue("#000000"),StrokeWidth: nullable.NewValue(1.0),TextColor: nullable.NewValue("#FFFFFF"),VerticalPosition: nullable.NewValue("bottom"),}` |
| `└─ Template`      |    ✗     | Preset subtitle templates. Please visit https://magichour.ai/create/auto-subtitle-generator to see the style of the existing templates.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | `V1AutoSubtitleGeneratorCreateBodyStyleTemplateEnumCinematic`                                                                                                                                                                                                                                                                                                                                                                                |
| `Name`             |    ✗     | Give your video a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | `"My Auto Subtitle video"`                                                                                                                                                                                                                                                                                                                                                                                                                   |

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
		Name:         nullable.NewValue("My Auto Subtitle video"),
		StartSeconds: 0.0,
		Style:        types.V1AutoSubtitleGeneratorCreateBodyStyle{},
	})
}
```

#### Response

##### Type

[V1AutoSubtitleGeneratorCreateResponse](/types/v1_auto_subtitle_generator_create_response.go)

##### Example

```go
V1AutoSubtitleGeneratorCreateResponse {
CreditsCharged: 450,
EstimatedFrameCost: 450,
Id: "cuid-example",
}
```
