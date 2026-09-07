# v1.account

## Module Functions

### Get account details <a name="list"></a>

Get the current credit balance and subscription details of the account that owns the API key.

**API Endpoint**: `GET /v1/account`

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.Account.List()
}
```

#### Response

##### Type

[V1AccountListResponse](/types/v1_account_list_response.go)

##### Example

```go
V1AccountListResponse {
Credits: 12500,
Email: nullable.NewValue("user@example.com"),
Id: "cuid-example",
Subscription: nullable.NewValue(V1AccountListResponseSubscription {
BillingInterval: nullable.NewValue(V1AccountListResponseSubscriptionBillingIntervalEnumMonth),
CancelAtPeriodEnd: false,
CurrentPeriodEnd: nullable.NewValue("2026-10-01T00:00:00.000Z"),
Discount: nullable.NewValue(V1AccountListResponseSubscriptionDiscount {
AmountOff: nullable.NewValue(123),
PercentOff: nullable.NewValue(20.0),
}),
Name: nullable.NewValue("Pro"),
Price: V1AccountListResponseSubscriptionPrice {
Amount: 4900,
Currency: "usd",
},
Status: V1AccountListResponseSubscriptionStatusEnumActive,
}),
Tier: V1AccountListResponseTierEnumPro,
}
```
