package test_saved_items_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	saved_items "github.com/magichourhq/magic-hour-go/resources/v1/saved_items"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestList200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.SavedItems.List(saved_items.ListRequest{
		Cursor: nullable.NewValue("string"),
		Limit:  nullable.NewValue(20),
		Type:   nullable.NewValue(types.V1SavedItemsListTypeEnumCharacter),
	})

	if err != nil {
		t.Fatalf("TestList200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestList200SuccessRequiredOnly(t *testing.T) {
	// Success test using only required fields
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.SavedItems.List(saved_items.ListRequest{
		Limit: nullable.NewValue(20),
		Type:  nullable.NewValue(types.V1SavedItemsListTypeEnumCharacter),
	})

	if err != nil {
		t.Fatalf("TestList200SuccessRequiredOnly - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
