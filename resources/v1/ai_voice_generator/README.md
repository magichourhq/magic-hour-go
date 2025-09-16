# v1.ai_voice_generator

## Module Functions

### AI Voice Generator <a name="create"></a>

Generate speech from text. Each character costs 0.05 credits. The cost is rounded up to the nearest whole number.

**API Endpoint**: `POST /v1/ai-voice-generator`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `Style` | ✓ | The content used to generate speech. | `V1AiVoiceGeneratorCreateBodyStyle {Prompt: "Hello, how are you?",VoiceName: V1AiVoiceGeneratorCreateBodyStyleVoiceNameEnumElonMusk,}` |
| `└─ Prompt` | ✓ | Text used to generate speech. Starter tier users can use up to 200 characters, while Creator, Pro, or Business users can use up to 1000. | `"Hello, how are you?"` |
| `└─ VoiceName` | ✓ | The voice to use for the speech. Available voices: Elon Musk, Mark Zuckerberg, Joe Rogan, Barack Obama, Morgan Freeman, Kanye West, Donald Trump, Joe Biden, Kim Kardashian, Taylor Swift | `V1AiVoiceGeneratorCreateBodyStyleVoiceNameEnumElonMusk` |
| `Name` | ✗ | The name of audio. This value is mainly used for your own identification of the audio. | `"Voice Generator audio"` |

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
		Name: nullable.NewValue("Voice Generator audio"),
		Style: types.V1AiVoiceGeneratorCreateBodyStyle{
			Prompt:    "Hello, how are you?",
			VoiceName: types.V1AiVoiceGeneratorCreateBodyStyleVoiceNameEnumElonMusk,
		},
	})
}

```

#### Response

##### Type
[V1AiVoiceGeneratorCreateResponse](/types/v1_ai_voice_generator_create_response.go)

##### Example
`V1AiVoiceGeneratorCreateResponse {
CreditsCharged: 1,
Id: "cuid-example",
}`


