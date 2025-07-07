package test_image_projects_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	image_projects "github.com/magichourhq/magic-hour-go/resources/v1/image_projects"
)

func TestDelete204SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	err := client.V1.ImageProjects.Delete(image_projects.DeleteRequest{
		Id: "cm6pvghix03bvyz0zwash6noj",
	})

	if err != nil {
		t.Fatalf("TestDelete204SuccessAllParams - failed making request with error: %#v", err)
	}

}

func TestGet200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.ImageProjects.Get(image_projects.GetRequest{
		Id: "cm6pvghix03bvyz0zwash6noj",
	})

	if err != nil {
		t.Fatalf("TestGet200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
