
### create <a name="create"></a>
Create AI QR Code

Create an AI QR code. Each QR code costs 20 frames.

**API Endpoint**: `POST /v1/ai-qr-code-generator`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	ai_qr_code_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_qr_code_generator"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.AiQrCodeGenerator.Create(ai_qr_code_generator.CreateRequest { Data: types.PostV1AiQrCodeGeneratorBody { Content: "https://magichour.ai", Name: nullable.NewValue("Qr Code image"), Style: types.PostV1AiQrCodeGeneratorBodyStyle { ArtStyle: "Watercolor" } } })
```

**Upgrade to see all examples**