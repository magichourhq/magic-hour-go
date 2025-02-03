package test_video_projects_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	video_projects "github.com/magichourhq/magic-hour-go/resources/v1/video_projects"
)

func TestDelete204GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.8.0"))
	err := client.V1.VideoProjects.Delete(video_projects.DeleteRequest{Id: "string"})

	if err != nil {
		t.Fatalf("TestDelete204GeneratedSuccess - failed making request with error: %#v", err)
	}

}

func TestGet200GeneratedSuccess(t *testing.T) {
	// Success case test
	client := sdk.NewClient(sdk.WithBearerAuth("API_TOKEN"), sdk.WithBaseURL("https://api.sideko.dev/v1/mock/magichour/magic-hour/0.8.0"))
	res, err := client.V1.VideoProjects.Get(video_projects.GetRequest{Id: "string"})

	if err != nil {
		t.Fatalf("TestGet200GeneratedSuccess - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
