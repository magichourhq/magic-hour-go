# v1.account.checkout

## Module Functions

### Create checkout <a name="create"></a>

Create a Stripe Checkout link for the account that owns the API key. Nothing is charged until the checkout is completed in a browser.

**API Endpoint**: `POST /v1/account/checkout`

#### Parameters

| Parameter         | Required | Description                                                                                                                                                 | Example                                         |
| ----------------- | :------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------- |
| `Type`            |    ✓     | What to buy. `credit_pack` is a one-time credit purchase.                                                                                                   | `V1AccountCheckoutCreateBodyTypeEnumCreditPack` |
| `CreditShortfall` |    ✗     | Credits a blocked request is missing, e.g. `credit_shortfall` from a 402 response. The checkout starts with enough packs to cover it. Defaults to one pack. | `150.0`                                         |
| `PackCount`       |    ✗     | Number of credit packs the checkout starts with. Takes priority over `credit_shortfall`.                                                                    | `3`                                             |

#### Example Snippet

```go
package main

import (
	os "os"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	checkout "github.com/magichourhq/magic-hour-go/resources/v1/account/checkout"
	types "github.com/magichourhq/magic-hour-go/types"
)

func main() {
	client := sdk.NewClient(
		sdk.WithBearerAuth(os.Getenv("API_TOKEN")),
	)
	res, err := client.V1.Account.Checkout.Create(checkout.CreateRequest{
		CreditShortfall: nullable.NewValue(150.0),
		PackCount:       nullable.NewValue(3),
		Type:            types.V1AccountCheckoutCreateBodyTypeEnumCreditPack,
	})
}
```

#### Response

##### Type

[V1AccountCheckoutCreateResponse](/types/v1_account_checkout_create_response.go)

##### Example

```go
V1AccountCheckoutCreateResponse {
CheckoutUrl: "https://pay.magichour.ai/c/pay/cs_live_...",
}
```
