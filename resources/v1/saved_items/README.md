# v1.saved_items

## Module Functions

### List saved items <a name="list"></a>

Returns active saved items owned by the authenticated account, newest first. Each item includes every saved asset with a durable file_path for reuse in compatible generation APIs and a temporary signed URL for previewing or downloading. Filter by type to find characters, references, voices, moodboards, or brand kits. To fetch the next page, pass the response's next_cursor as cursor.

**API Endpoint**: `GET /v1/saved-items`

#### Parameters

| Parameter | Required | Description                                                        | Example                             |
| --------- | :------: | ------------------------------------------------------------------ | ----------------------------------- |
| `Cursor`  |    ✗     | Opaque pagination cursor from the previous response's next_cursor. | `"string"`                          |
| `Limit`   |    ✗     | Maximum number of saved items to return. Defaults to 20.           | `20`                                |
| `Type`    |    ✗     | Only return saved items of this type.                              | `V1SavedItemsListTypeEnumCharacter` |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	saved_items "github.com/magichourhq/magic-hour-go/resources/v1/saved_items"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.SavedItems.List(saved_items.ListRequest{
		Limit: nullable.NewValue(20),
		Type:  nullable.NewValue(types.V1SavedItemsListTypeEnumCharacter),
	})
}
```

#### Response

##### Type

[V1SavedItemsListResponse](/types/v1_saved_items_list_response.go)

##### Example

```go
V1SavedItemsListResponse {
Items: []V1SavedItemsListResponseItemsItem{
V1SavedItemsListResponseItemsItem {
Assets: []V1SavedItemsListResponseItemsItemAssetsItem{
V1SavedItemsListResponseItemsItemAssetsItem {
FilePath: "saved-items/user-id/item-id/image.png",
IsPrimary: true,
MediaKind: V1SavedItemsListResponseItemsItemAssetsItemMediaKindEnumImage,
Url: "http://www.example.com",
UrlExpiresAt: "2026-09-17T00:00:00.000Z",
},
},
Id: "cuid-example",
Name: nullable.NewValue("Alex"),
Type: V1SavedItemsListResponseItemsItemTypeEnumCharacter,
},
},
NextCursor: nullable.NewValue("string"),
}
```
