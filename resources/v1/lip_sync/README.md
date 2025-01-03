
### create <a name="create"></a>
Create Lip Sync video

Create a Lip Sync video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/lip-sync).
  

**API Endpoint**: `POST /v1/lip-sync`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	lip_sync "github.com/magichourhq/magic-hour-go/resources/v1/lip_sync"
	types "github.com/magichourhq/magic-hour-go/types"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.LipSync.Create(lip_sync.CreateRequest { Assets: types.PostV1LipSyncBodyAssets { AudioFilePath: "audio/id/1234.mp3", VideoSource: types.PostV1LipSyncBodyAssetsVideoSourceEnumFile }, EndSeconds: 15, Height: 960, StartSeconds: 0, Width: 512 })
```

**Upgrade to see all examples**
