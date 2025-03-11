
### create <a name="create"></a>
AI Clothes Changer

Change outfits in photos in seconds with just a photo reference. Each photo costs 25 frames.

**API Endpoint**: `POST /v1/ai-clothes-changer`

#### Example Snippet

```go
package main
import (
	sdk "github.com/magichourhq/magic-hour-go/client"
	os "os"
	ai_clothes_changer "github.com/magichourhq/magic-hour-go/resources/v1/ai_clothes_changer"
	types "github.com/magichourhq/magic-hour-go/types"
)
func main(){
client := sdk.NewClient(
sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
)
res, err := client.V1.AiClothesChanger.Create(ai_clothes_changer.CreateRequest {
Assets: types.PostV1AiClothesChangerBodyAssets {
GarmentFilePath: "api-assets/id/outfit.png",
GarmentType: types.PostV1AiClothesChangerBodyAssetsGarmentTypeEnumDresses,
PersonFilePath: "api-assets/id/model.png",
},
})
}
```
