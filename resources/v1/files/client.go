package files

import (
	"fmt"
	http "net/http"
	"os"
	"path/filepath"
	"strings"

	sdkcore "github.com/magichourhq/magic-hour-go/core"
	upload_urls "github.com/magichourhq/magic-hour-go/resources/v1/files/upload_urls"
	"github.com/magichourhq/magic-hour-go/types"
)

type Client struct {
	coreClient *sdkcore.CoreClient
	UploadUrls *upload_urls.Client
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient: coreClient,
		UploadUrls: upload_urls.NewClient(coreClient),
	}

	return &client
}

// UploadFile uploads a local file to Magic Hour storage and returns the file path for use in API calls
//
// This function takes a local file path, determines the file type based on the extension,
// requests an upload URL, uploads the file, and returns the file path that can be used
// in API calls for image_file_path, video_file_path, or audio_file_path fields.
//
// Example:
//
//	filePath, err := client.V1.Files.UploadFile("/path/to/your/image.jpg")
//	if err != nil {
//		// handle error
//	}
//	// Use filePath in API calls
//
// Supported file extensions:
// - Video: mp4, m4v, mov, webm
// - Audio: mp3, mpeg, wav, aac, aiff, flac
// - Image: png, jpg, jpeg, webp, avif, jp2, tiff, bmp
func (c *Client) UploadFile(localFilePath string, reqModifiers ...RequestModifier) (string, error) {
	// Check if file exists
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", localFilePath)
	}

	// Get file extension
	ext := strings.TrimPrefix(filepath.Ext(localFilePath), ".")
	if ext == "" {
		return "", fmt.Errorf("file must have an extension: %s", localFilePath)
	}

	// Determine file type based on extension
	fileType, err := c.determineFileType(ext)
	if err != nil {
		return "", err
	}

	// Create upload URL request
	request := upload_urls.CreateRequest{
		Items: []types.V1FilesUploadUrlsCreateBodyItemsItem{
			{
				Extension: ext,
				Type:      fileType,
			},
		},
	}

	// Get upload URL
	response, err := c.UploadUrls.Create(request, reqModifiers...)
	if err != nil {
		return "", fmt.Errorf("failed to get upload URL: %w", err)
	}

	if len(response.Items) == 0 {
		return "", fmt.Errorf("no upload URL received")
	}

	uploadItem := response.Items[0]

	// Upload the file
	err = c.uploadFileToURL(localFilePath, uploadItem.UploadUrl)
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// Return the file path for use in API calls
	return uploadItem.FilePath, nil
}

// determineFileType determines the file type based on the file extension
func (c *Client) determineFileType(extension string) (types.V1FilesUploadUrlsCreateBodyItemsItemTypeEnum, error) {
	ext := strings.ToLower(extension)

	// Video extensions
	videoExts := []string{"mp4", "m4v", "mov", "webm"}
	for _, videoExt := range videoExts {
		if ext == videoExt {
			return types.V1FilesUploadUrlsCreateBodyItemsItemTypeEnumVideo, nil
		}
	}

	// Audio extensions
	audioExts := []string{"mp3", "mpeg", "wav", "aac", "aiff", "flac"}
	for _, audioExt := range audioExts {
		if ext == audioExt {
			return types.V1FilesUploadUrlsCreateBodyItemsItemTypeEnumAudio, nil
		}
	}

	// Image extensions
	imageExts := []string{"png", "jpg", "jpeg", "webp", "avif", "jp2", "tiff", "bmp"}
	for _, imageExt := range imageExts {
		if ext == imageExt {
			return types.V1FilesUploadUrlsCreateBodyItemsItemTypeEnumImage, nil
		}
	}

	return "", fmt.Errorf("unsupported file extension: %s", extension)
}

// uploadFileToURL uploads a file to the given pre-signed URL
func (c *Client) uploadFileToURL(localFilePath, uploadURL string) error {
	// Open the file
	file, err := os.Open(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create PUT request
	req, err := http.NewRequest("PUT", uploadURL, file)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %w", err)
	}

	// Execute the request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode >= 300 {
		return fmt.Errorf("upload failed with status %d", resp.StatusCode)
	}

	return nil
}
