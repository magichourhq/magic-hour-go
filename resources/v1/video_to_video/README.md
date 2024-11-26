
### create <a name="create"></a>
Create Video-to-Video

Create a Video To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/video-to-video).
  

**API Endpoint**: `POST /v1/video-to-video`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	video_to_video "github.com/magichourhq/magic-hour-go/resources/v1/video_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.VideoToVideo.Create(video_to_video.CreateRequest { Data: types.PostV1VideoToVideoBody { Assets: types.PostV1VideoToVideoBodyAssets { VideoFilePath: nullable.NewValue("video/id/1234.mp4"), VideoSource: types.PostV1VideoToVideoBodyAssetsVideoSourceEnumFile, YoutubeUrl: nullable.NewValue("http://www.example.com") }, EndSeconds: 15, FpsResolution: nullable.NewValue(types.PostV1VideoToVideoBodyFpsResolutionEnumHalf), Height: 960, Name: nullable.NewValue("Video To Video video"), StartSeconds: 0, Style: types.PostV1VideoToVideoBodyStyle { ArtStyle: types.PostV1VideoToVideoBodyStyleArtStyleEnum3dRender, Model: types.PostV1VideoToVideoBodyStyleModelEnumAbsoluteReality, Prompt: nullable.NewValue("string"), PromptType: types.PostV1VideoToVideoBodyStylePromptTypeEnumAppendDefault, Version: types.PostV1VideoToVideoBodyStyleVersionEnumDefault }, Width: 512 } })
```

**Upgrade to see all examples**
