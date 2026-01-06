# v1.ai_voice_cloner

## Module Functions

### AI Voice Cloner <a name="create"></a>

Clone a voice from an audio sample and generate speech.

- Each character costs 0.05 credits.
- The cost is rounded up to the nearest whole number

**API Endpoint**: `POST /v1/ai-voice-cloner`

#### Parameters

| Parameter          | Required | Description                                                                                                                                                                                                                                                                                                                                                      | Example                                                                       |
| ------------------ | :------: | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `Assets`           |    ✓     | Provide the assets for voice cloning.                                                                                                                                                                                                                                                                                                                            | `V1AiVoiceClonerCreateBodyAssets {AudioFilePath: "api-assets/id/1234.mp3",}`  |
| `└─ AudioFilePath` |    ✓     | The audio used to clone the voice. This value is either - a direct URL to the video file - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls). See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details. | `"api-assets/id/1234.mp3"`                                                    |
| `Style`            |    ✓     |                                                                                                                                                                                                                                                                                                                                                                  | `V1AiVoiceClonerCreateBodyStyle {Prompt: "Hello, this is my cloned voice.",}` |
| `└─ Prompt`        |    ✓     | Text used to generate speech from the cloned voice. The character limit is 1000 characters.                                                                                                                                                                                                                                                                      | `"Hello, this is my cloned voice."`                                           |
| `Name`             |    ✗     | Give your audio a custom name for easy identification.                                                                                                                                                                                                                                                                                                           | `"My Voice Cloner audio"`                                                     |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_voice_cloner "github.com/magichourhq/magic-hour-go/resources/v1/ai_voice_cloner"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiVoiceCloner.Create(ai_voice_cloner.CreateRequest{
		Assets: types.V1AiVoiceClonerCreateBodyAssets{
			AudioFilePath: "api-assets/id/1234.mp3",
		},
		Name: nullable.NewValue("My Voice Cloner audio"),
		Style: types.V1AiVoiceClonerCreateBodyStyle{
			Prompt: "Hello, this is my cloned voice.",
		},
	})
}
```

#### Response

##### Type

[V1AiVoiceClonerCreateResponse](/types/v1_ai_voice_cloner_create_response.go)

##### Example

```go
V1AiVoiceClonerCreateResponse {
CreditsCharged: 1,
Id: "cuid-example",
}
```
