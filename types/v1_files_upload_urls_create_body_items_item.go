package types

// V1FilesUploadUrlsCreateBodyItemsItem
type V1FilesUploadUrlsCreateBodyItemsItem struct {
	// The extension of the file to upload. Do not include the dot (.) before the extension. Possible extensions are mp4,m4v,mov,webm,mp3,mpeg,wav,aac,aiff,flac,png,jpg,jpeg,webp,avif,jp2,tiff,bmp
	Extension string `json:"extension"`
	// The type of asset to upload. Possible types are video, audio, image
	Type V1FilesUploadUrlsCreateBodyItemsItemTypeEnum `json:"type"`
}
