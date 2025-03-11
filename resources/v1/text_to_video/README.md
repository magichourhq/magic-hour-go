
### create <a name="create"></a>
Text-to-Video

Create a Text To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/text-to-video).
  

**API Endpoint**: `POST /v1/text-to-video`

#### Example Snippet

```go
package main
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	text_to_video "github.com/magichourhq/magic-hour-go/resources/v1/text_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
)
func main(){
client := sdk.NewClient(
sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
)
res, err := client.V1.TextToVideo.Create(text_to_video.CreateRequest {
EndSeconds: 5.0,
Orientation: types.PostV1TextToVideoBodyOrientationEnumLandscape,
Style: types.PostV1TextToVideoBodyStyle {
Prompt: "string",
},
})
}
```
