
### Get face detection details <a name="get"></a>

Get the details of a face detection task.

**API Endpoint**: `GET /v1/face-detection/{id}`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `id` | ✓ | The id of the task | `"string"` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	face_detection "github.com/magichourhq/magic-hour-go/resources/v1/face_detection"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.FaceDetection.Get(face_detection.GetRequest{
		Id: "string",
	})
}

```

#### Response

##### Type
[V1FaceDetectionGetResponse](/types/v1_face_detection_get_response.go)

##### Example
`V1FaceDetectionGetResponse {
CreditsCharged: 123,
Faces: []V1FaceDetectionGetResponseFacesItem{
V1FaceDetectionGetResponseFacesItem {
Path: "api-assets/id/0-0.png",
Url: "https://videos.magichour.ai/api-assets/id/0-0.png",
},
},
Id: "string",
Status: V1FaceDetectionGetResponseStatusEnumComplete,
}`

### Face Detection <a name="create"></a>

Detect faces in an image or video. 

Note: Face detection is free to use for the near future. Pricing may change in the future.

**API Endpoint**: `POST /v1/face-detection`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `assets` | ✓ | Provide the assets for face detection | `V1FaceDetectionCreateBodyAssets {TargetFilePath: "api-assets/id/1234.png",}` |
| `confidence_score` | ✗ | Confidence threshold for filtering detected faces.  * Higher values (e.g., 0.9) include only faces detected with high certainty, reducing false positives.  * Lower values (e.g., 0.3) include more faces, but may increase the chance of incorrect detections. | `0.5` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	face_detection "github.com/magichourhq/magic-hour-go/resources/v1/face_detection"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.FaceDetection.Create(face_detection.CreateRequest{
		Assets: types.V1FaceDetectionCreateBodyAssets{
			TargetFilePath: "api-assets/id/1234.png",
		},
		ConfidenceScore: nullable.NewValue(0.5),
	})
}

```

#### Response

##### Type
[V1FaceDetectionCreateResponse](/types/v1_face_detection_create_response.go)

##### Example
`V1FaceDetectionCreateResponse {
CreditsCharged: 123,
Id: "string",
}`
