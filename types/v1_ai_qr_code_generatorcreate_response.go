package types

// Success
type V1AiQrCodeGeneratorcreateResponse struct {
	// The frame cost of the image generation
	FrameCost int `json:"frame_cost"`
	// Unique ID of the image. This value can be used in the [get image project API](/api/tag/image-projects/get/v1/image-projects/{id}) to fetch additional details such as status
	Id string `json:"id"`
}
