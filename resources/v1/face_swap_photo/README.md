
### create <a name="create"></a>
Create Face Swap Photo

Create a face swap photo. Each photo costs 5 frames. The height/width of the output image depends on your subscription. Please refer to our [pricing](/pricing) page for more details

**API Endpoint**: `POST /v1/face-swap-photo`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	face_swap_photo "github.com/magichourhq/magic-hour-go/resources/v1/face_swap_photo"
	types "github.com/magichourhq/magic-hour-go/types"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.FaceSwapPhoto.Create(face_swap_photo.CreateRequest { Data: types.PostV1FaceSwapPhotoBody { Assets: types.PostV1FaceSwapPhotoBodyAssets { SourceFilePath: "image/id/1234.png", TargetFilePath: "image/id/1234.png" }, Name: nullable.NewValue("Face Swap image") } })
```

**Upgrade to see all examples**
