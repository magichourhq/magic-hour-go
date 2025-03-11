
### create <a name="create"></a>
Lip Sync

Create a Lip Sync video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/lip-sync).
  

**API Endpoint**: `POST /v1/lip-sync`

#### Example Snippet

```go
package main
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	lip_sync "github.com/magichourhq/magic-hour-go/resources/v1/lip_sync"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)
func main(){
client := sdk.NewClient(
sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
)
res, err := client.V1.LipSync.Create(lip_sync.CreateRequest {
Assets: types.PostV1LipSyncBodyAssets {
AudioFilePath: "api-assets/id/1234.mp3",
VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
VideoSource: types.PostV1LipSyncBodyAssetsVideoSourceEnumFile,
},
EndSeconds: 15.0,
Height: 960,
StartSeconds: 0.0,
Width: 512,
})
}
```
