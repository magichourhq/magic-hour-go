package test_face_detection_client

import (
	fmt "fmt"
	testing "testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	face_detection "github.com/magichourhq/magic-hour-go/resources/v1/face_detection"
	types "github.com/magichourhq/magic-hour-go/types"
)

func TestGet200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.FaceDetection.Get(face_detection.GetRequest{
		Id: "string",
	})

	if err != nil {
		t.Fatalf("TestGet200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}

func TestCreate200SuccessAllParams(t *testing.T) {
	// Success test using all required and optional
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)
	res, err := client.V1.FaceDetection.Create(face_detection.CreateRequest{
		Assets: types.V1FaceDetectionCreateBodyAssets{
			TargetFilePath: "api-assets/id/1234.png",
		},
		ConfidenceScore: nullable.NewValue(0.5),
	})

	if err != nil {
		t.Fatalf("TestCreate200SuccessAllParams - failed making request with error: %#v", err)
	}

	fmt.Printf("response - %#v\n", res)
}
