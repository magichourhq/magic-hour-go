
package types

// Success
type PostV1FaceSwapPhotoResponse struct {
    // The frame cost of the image generation
    FrameCost int `json:"frame_cost"`
    // Unique ID of the image. This value can be used in the [get image project API](https://docs.magichour.ai/api-reference/image-projects/get-image-details) to fetch additional details such as status
    Id string `json:"id"`
}



