# v1.ai_image_generator

## Module Functions

### AI Image Generator <a name="create"></a>

Create an AI image with advanced model selection and quality controls.

**API Endpoint**: `POST /v1/ai-image-generator`

#### Parameters

| Parameter        | Required | Deprecated | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Example                                                                                                                                        |
| ---------------- | :------: | :--------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `ImageCount`     |    ✓     |     ✗      | Number of images to generate. Maximum varies by model.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | `1`                                                                                                                                            |
| `Style`          |    ✓     |     ✗      | The art style to use for image generation.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | `V1AiImageGeneratorCreateBodyStyle {Prompt: "Cool image",Tool: nullable.NewValue(V1AiImageGeneratorCreateBodyStyleToolEnumAiAnimeGenerator),}` |
| `└─ Prompt`      |    ✓     |     —      | The prompt used for the image(s).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | `"Cool image"`                                                                                                                                 |
| `└─ QualityMode` |    ✗     |     ✓      | DEPRECATED: Use `model` field instead for explicit model selection. Legacy quality mode mapping: - `standard` → `z-image-turbo` model - `pro` → `seedream` model If model is specified, it will take precedence over the legacy quality_mode field.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | `V1AiImageGeneratorCreateBodyStyleQualityModeEnumPro`                                                                                          |
| `└─ Tool`        |    ✗     |     —      | The art style to use for image generation. Defaults to 'general' if not provided.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | `V1AiImageGeneratorCreateBodyStyleToolEnumAiAnimeGenerator`                                                                                    |
| `AspectRatio`    |    ✗     |     ✗      | The aspect ratio of the output image(s). If not specified, defaults to `1:1` (square).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | `V1AiImageGeneratorCreateBodyAspectRatioEnum11`                                                                                                |
| `Model`          |    ✗     |     ✗      | The AI model to use for image generation. Each model has different capabilities and costs. **Models:** - `default` - Use the model we recommend, which will change over time. This is recommended unless you need a specific model. This is the default behavior. - `flux-schnell` - 5 credits/image - Supported resolutions: auto - Available for tiers: free, creator, pro, business - Image count allowed: 1, 2, 3, 4 - `z-image-turbo` - 5 credits/image - Supported resolutions: auto, 2k - Available for tiers: free, creator, pro, business - Image count allowed: 1, 2, 3, 4 - `seedream` - 30 credits/image - Supported resolutions: auto, 2k, 4k - Available for tiers: free, creator, pro, business - Image count allowed: 1, 2, 3, 4 - `nano-banana-pro` - 150 credits/image - Supported resolutions: auto - Available for tiers: creator, pro, business - Image count allowed: 1, 4, 9, 16 | `V1AiImageGeneratorCreateBodyModelEnumDefault`                                                                                                 |
| `Name`           |    ✗     |     ✗      | Give your image a custom name for easy identification.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | `"My Ai Image image"`                                                                                                                          |
| `Orientation`    |    ✗     |     ✓      | DEPRECATED: Use `aspect_ratio` instead. The orientation of the output image(s). `aspect_ratio` takes precedence when `orientation` if both are provided.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | `V1AiImageGeneratorCreateBodyOrientationEnumLandscape`                                                                                         |
| `Resolution`     |    ✗     |     ✗      | Maximum resolution for the generated image. **Options:** - `auto` - Automatic resolution (all tiers, default) - `2k` - Up to 2048px (requires Pro or Business tier) - `4k` - Up to 4096px (requires Business tier) Note: Resolution availability depends on the model and your subscription tier. See `model` field for which resolutions each model supports. Defaults to `auto` if not specified.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | `V1AiImageGeneratorCreateBodyResolutionEnumAuto`                                                                                               |

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
		AspectRatio: nullable.NewValue(types.V1AiImageGeneratorCreateBodyAspectRatioEnum11),
		ImageCount:  1,
		Model:       nullable.NewValue(types.V1AiImageGeneratorCreateBodyModelEnumDefault),
		Name:        nullable.NewValue("My Ai Image image"),
		Resolution:  nullable.NewValue(types.V1AiImageGeneratorCreateBodyResolutionEnumAuto),
		Style: types.V1AiImageGeneratorCreateBodyStyle{
			Prompt: "Cool image",
			Tool:   nullable.NewValue(types.V1AiImageGeneratorCreateBodyStyleToolEnumAiAnimeGenerator),
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
