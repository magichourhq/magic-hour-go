
### create <a name="create"></a>
Generate asset upload urls

Create a list of urls used to upload the assets needed to generate a video. Each video type has their own requirements on what assets are required. Please refer to the specific mode API for more details. The response array will be in the same order as the request body.

Below is the list of valid extensions for each asset type:

- video: mp4, mov, webm
- audio: mp3, mpeg, wav, aac, aiff, flac
- image: png, jpg, jpeg, webp, avif, jp2, tiff, bmp

Note: `.gif` is supported for face swap API `video_file_path` field.

After receiving the upload url, you can upload the file by sending a PUT request with the header `'Content-Type: application/octet-stream'`.

For example using curl

```
curl -X PUT -H 'Content-Type: application/octet-stream' \
  --data '@/path/to/file/video.mp4' \
  https://videos.magichour.ai/api-assets/id/video.mp4?auth-value=1234567890
```


**API Endpoint**: `POST /v1/files/upload-urls`

#### Example Snippet

```go
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	upload_urls "github.com/magichourhq/magic-hour-go/resources/v1/files/upload_urls"
	types "github.com/magichourhq/magic-hour-go/types"
)

client := sdk.NewClient(sdk.WithBearerAuth(os.Getenv("API_TOKEN")))

res, err := client.V1.Files.UploadUrls.Create(upload_urls.CreateRequest { Items: []types.PostV1FilesUploadUrlsBodyItemsItem{types.PostV1FilesUploadUrlsBodyItemsItem { Extension: "mp4", Type: types.PostV1FilesUploadUrlsBodyItemsItemTypeEnumVideo }, types.PostV1FilesUploadUrlsBodyItemsItem { Extension: "mp3", Type: types.PostV1FilesUploadUrlsBodyItemsItemTypeEnumAudio }} })
```

**Upgrade to see all examples**
