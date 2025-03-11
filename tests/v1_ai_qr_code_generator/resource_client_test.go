
package test_ai_qr_code_generator_client
import (
testing "testing"
sdk "github.com/magichourhq/magic-hour-go/client"
ai_qr_code_generator "github.com/magichourhq/magic-hour-go/resources/v1/ai_qr_code_generator"
types "github.com/magichourhq/magic-hour-go/types"
fmt "fmt"
)


func TestCreate200SuccessDefault(t *testing.T) {
// Success test for Default body
client := sdk.NewClient(
sdk.WithBearerAuth("API_TOKEN"),
sdk.WithEnv(sdk.MockServer),
)
res, err := client.V1.AiQrCodeGenerator.Create(ai_qr_code_generator.CreateRequest {
Content: "https://magichour.ai",
Style: types.PostV1AiQrCodeGeneratorBodyStyle {
ArtStyle: "Watercolor",
},
})

if err != nil {
t.Fatalf("TestCreate200SuccessDefault - failed making request with error: %#v", err)
}

fmt.Printf("response - %#v\n", res)
}
