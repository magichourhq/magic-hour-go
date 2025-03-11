
### create <a name="create"></a>
Video-to-Video

Create a Video To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
  
Get more information about this mode at our [product page](/products/video-to-video).
  

**API Endpoint**: `POST /v1/video-to-video`

#### Example Snippet

```go
package main
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	video_to_video "github.com/magichourhq/magic-hour-go/resources/v1/video_to_video"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)
func main(){
client := sdk.NewClient(
sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
)
res, err := client.V1.VideoToVideo.Create(video_to_video.CreateRequest {
Assets: types.PostV1VideoToVideoBodyAssets {
VideoFilePath: nullable.NewValue("api-assets/id/1234.mp4"),
VideoSource: types.PostV1VideoToVideoBodyAssetsVideoSourceEnumFile,
},
EndSeconds: 15.0,
Height: 960,
StartSeconds: 0.0,
Style: types.PostV1VideoToVideoBodyStyle {
ArtStyle: types.PostV1VideoToVideoBodyStyleArtStyleEnum3dRender,
Model: types.PostV1VideoToVideoBodyStyleModelEnumAbsoluteReality,
Prompt: nullable.NewNull[string](),
PromptType: types.PostV1VideoToVideoBodyStylePromptTypeEnumAppendDefault,
Version: types.PostV1VideoToVideoBodyStyleVersionEnumDefault,
},
Width: 512,
})
}
```
