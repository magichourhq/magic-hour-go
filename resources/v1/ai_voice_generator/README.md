# v1.ai_voice_generator

## Module Functions

### AI Voice Generator <a name="create"></a>

Generate speech from text. Each character costs 0.1 credits. The cost is rounded up to the nearest whole number.

**API Endpoint**: `POST /v1/ai-voice-generator`

#### Parameters

| Parameter      | Required | Description                                                                                                               | Example                                                                                                                                                                                                                                                 |
| -------------- | :------: | ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Style`        |    ✓     | The content used to generate speech.                                                                                      | `V1AiVoiceGeneratorCreateBodyStyle {Language: nullable.NewValue(V1AiVoiceGeneratorCreateBodyStyleLanguageEnumSpanish),Prompt: "Hello, how are you?",Speed: nullable.NewValue(1.15),VoiceName: V1AiVoiceGeneratorCreateBodyStyleVoiceNameEnumElonMusk,}` |
| `└─ Language`  |    ✗     | Spoken language. Use a canonical lowercase language name, or auto to detect the language from the text. Defaults to auto. | `V1AiVoiceGeneratorCreateBodyStyleLanguageEnumSpanish`                                                                                                                                                                                                  |
| `└─ Prompt`    |    ✓     | Text used to generate speech. The character limit is 5000 characters.                                                     | `"Hello, how are you?"`                                                                                                                                                                                                                                 |
| `└─ Speed`     |    ✗     | Speech speed multiplier, from 0.75 to 1.30 inclusive. Defaults to 1 (normal speed).                                       | `1.15`                                                                                                                                                                                                                                                  |
| `└─ VoiceName` |    ✓     | The voice to use for the speech. Use the `id` of a voice returned by `GET /v1/voices`.                                    | `V1AiVoiceGeneratorCreateBodyStyleVoiceNameEnumElonMusk`                                                                                                                                                                                                |
| `Name`         |    ✗     | Give your audio a custom name for easy identification.                                                                    | `"My Voice Generator audio"`                                                                                                                                                                                                                            |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	ai_voice_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_voice_generator"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.AiVoiceGenerator.Create(ai_voice_generator.CreateRequest{
		Name: nullable.NewValue("My Voice Generator audio"),
		Style: types.V1AiVoiceGeneratorCreateBodyStyle{
			Language:  nullable.NewValue(types.V1AiVoiceGeneratorCreateBodyStyleLanguageEnumSpanish),
			Prompt:    "Hello, how are you?",
			Speed:     nullable.NewValue(1.15),
			VoiceName: types.V1AiVoiceGeneratorCreateBodyStyleVoiceNameEnumElonMusk,
		},
	})
}
```

#### Response

##### Type

[V1AiVoiceGeneratorCreateResponse](/types/v1_ai_voice_generator_create_response.go)

##### Example

```go
V1AiVoiceGeneratorCreateResponse {
CreditsCharged: 1,
Id: "cuid-example",
}
```
