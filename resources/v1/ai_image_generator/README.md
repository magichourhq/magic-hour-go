# v1.ai_image_generator

## Module Functions

### AI Image Generator <a name="create"></a>

Create an AI image. Each standard image costs 5 credits. Pro quality images cost 30 credits.

**API Endpoint**: `POST /v1/ai-image-generator`

#### Parameters

| Parameter        | Required | Description                                                                                                                                                                                                                                                                                                                                        | Example                                                                                                                                                                                                                                 |
| ---------------- | :------: | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ImageCount`     |    ✓     | Number of images to generate.                                                                                                                                                                                                                                                                                                                      | `1`                                                                                                                                                                                                                                     |
| `Orientation`    |    ✓     | The orientation of the output image(s).                                                                                                                                                                                                                                                                                                            | `V1AiImageGeneratorCreateBodyOrientationEnumLandscape`                                                                                                                                                                                  |
| `Style`          |    ✓     | The art style to use for image generation.                                                                                                                                                                                                                                                                                                         | `V1AiImageGeneratorCreateBodyStyle {Prompt: "Cool image",QualityMode: nullable.NewValue(V1AiImageGeneratorCreateBodyStyleQualityModeEnumStandard),Tool: nullable.NewValue(V1AiImageGeneratorCreateBodyStyleToolEnumAiAnimeGenerator),}` |
| `└─ Prompt`      |    ✓     | The prompt used for the image(s).                                                                                                                                                                                                                                                                                                                  | `"Cool image"`                                                                                                                                                                                                                          |
| `└─ QualityMode` |    ✗     | Controls the quality of the generated image. Defaults to 'standard' if not specified. **Options:** - `standard` - Standard quality generation. Cost: 5 credits per image. - `pro` - Pro quality generation with enhanced details and quality. Cost: 30 credits per image. Note: Pro mode is available for users on Creator, Pro, or Business tier. | `V1AiImageGeneratorCreateBodyStyleQualityModeEnumStandard`                                                                                                                                                                              |
| `└─ Tool`        |    ✗     | The art style to use for image generation. Defaults to 'general' if not provided.                                                                                                                                                                                                                                                                  | `V1AiImageGeneratorCreateBodyStyleToolEnumAiAnimeGenerator`                                                                                                                                                                             |
| `Name`           |    ✗     | The name of image. This value is mainly used for your own identification of the image.                                                                                                                                                                                                                                                             | `"Ai Image image"`                                                                                                                                                                                                                      |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_image_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_image_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiImageGenerator.Create(ai_image_generator.CreateRequest{
		ImageCount:  1,
		Name:        nullable.NewValue("Ai Image image"),
		Orientation: types.V1AiImageGeneratorCreateBodyOrientationEnumLandscape,
		Style: types.V1AiImageGeneratorCreateBodyStyle{
			Prompt:      "Cool image",
			QualityMode: nullable.NewValue(types.V1AiImageGeneratorCreateBodyStyleQualityModeEnumStandard),
			Tool:        nullable.NewValue(types.V1AiImageGeneratorCreateBodyStyleToolEnumAiAnimeGenerator),
		},
	})
}
```

#### Response

##### Type

[V1AiImageGeneratorCreateResponse](/types/v1_ai_image_generator_create_response.go)

##### Example

```go
V1AiImageGeneratorCreateResponse {
CreditsCharged: 5,
FrameCost: 5,
Id: "cuid-example",
}
```
