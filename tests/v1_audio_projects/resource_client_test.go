package test_audio_projects_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	audio_projects "github.com/magichourhq/magic-hour-go/resources/v1/audio_projects"
)

func TestDelete204SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	err := client.V1.AudioProjects.Delete(audio_projects.DeleteRequest{
		Id: "cuid-example",
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
	res, err := client.V1.AudioProjects.Get(audio_projects.GetRequest{
		Id: "cuid-example",
	})

	if err != nil {
		t.Fatalf("TestGet200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
