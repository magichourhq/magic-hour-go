# v1.ai_headshot_generator

## Module Functions
### AI Headshots <a name="create"></a>

Create an AI headshot. Each headshot costs 50 credits.

**API Endpoint**: `POST /v1/ai-headshot-generator`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Assets` | ✓ | Provide the assets for headshot photo | `V1AiHeadshotGeneratorCreateBodyAssets {ImageFilePath: "api-assets/id/1234.png",}` |
| `└─ ImageFilePath` | ✓ | The image used to generate the headshot. This image must contain one detectable face. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).  Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.  | `"api-assets/id/1234.png"` |
| `Name` | ✗ | The name of image. This value is mainly used for your own identification of the image. | `"Ai Headshot image"` |
| `Style` | ✗ |  | `V1AiHeadshotGeneratorCreateBodyStyle {}` |
| `└─ Prompt` | ✗ | Prompt used to guide the style of your headshot. We recommend omitting the prompt unless you want to customize your headshot. You can visit [AI headshot generator](https://magichour.ai/create/ai-headshot-generator) to view an example of a good prompt used for our 'Professional' style. | `"string"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_headshot_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_headshot_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiHeadshotGenerator.Create(ai_headshot_generator.CreateRequest{
		Assets: types.V1AiHeadshotGeneratorCreateBodyAssets{
			ImageFilePath: "api-assets/id/1234.png",
		},
		Name: nullable.NewValue("Ai Headshot image"),
	})
}

```

#### Response

##### Type
[V1AiHeadshotGeneratorCreateResponse](/types/v1_ai_headshot_generator_create_response.go)

##### Example
`V1AiHeadshotGeneratorCreateResponse {
CreditsCharged: 50,
FrameCost: 50,
Id: "cuid-example",
}`
<!-- CUSTOM DOCS START -->

<!-- CUSTOM DOCS END -->

