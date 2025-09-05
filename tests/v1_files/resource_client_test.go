package test_files_client

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	sdk "github.com/magichourhq/magic-hour-go/client"
)

func TestUploadFileFunctionExists(t *testing.T) {
	// Test that the UploadFile function exists and has the correct signature
	client := sdk.NewClient(
		sdk.WithBearerAuth("API_TOKEN"),
		sdk.WithEnv(sdk.MockServer),
	)

	// Test that the function exists and can be called (will fail at runtime with mock server)
	// but should not fail due to missing function or wrong signature
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("UploadFile function not found or has wrong signature: %v", r)
		}
	}()

	// This will fail at runtime because we're using a mock server, but it should not panic
	// due to function signature issues
	_, _ = client.V1.Files.UploadFile("/nonexistent/file.jpg")
}

func TestSupportedFileExtensions(t *testing.T) {
	client := sdk.NewClient(sdk.WithEnv(sdk.MockServer))

	// Test that supported extensions don't immediately fail with "unsupported extension" error
	supportedExtensions := []string{
		"mp4", "m4v", "mov", "webm", // Video
		"mp3", "mpeg", "wav", "aac", "aiff", "flac", // Audio
		"png", "jpg", "jpeg", "webp", "avif", "jp2", "tiff", "bmp", // Image
	}

	for _, extension := range supportedExtensions {
		t.Run(fmt.Sprintf("supported_extension_%s", extension), func(t *testing.T) {
			tempFile := createTempFile(t, extension)
			defer os.Remove(tempFile)

			_, err := client.V1.Files.UploadFile(tempFile)

			// We expect an error (due to mock server), but NOT an "unsupported extension" error
			if err != nil {
				errorMsg := err.Error()
				if strings.Contains(errorMsg, "unsupported file extension") {
					t.Errorf("Extension %s should be supported, but got error: %s", extension, errorMsg)
				}
			}
		})
	}
}

func TestUnsupportedFileExtension(t *testing.T) {
	client := sdk.NewClient(sdk.WithEnv(sdk.MockServer))

	tempFile := createTempFile(t, "xyz")
	defer os.Remove(tempFile)

	_, err := client.V1.Files.UploadFile(tempFile)

	if err == nil {
		t.Error("Expected error for unsupported file extension, but got none")
	}

	if !strings.Contains(err.Error(), "unsupported file extension: xyz") {
		t.Errorf("Expected 'unsupported file extension' error, got: %s", err.Error())
	}
}

func TestUploadFileWithNonexistentFile(t *testing.T) {
	client := sdk.NewClient(sdk.WithEnv(sdk.MockServer))

	_, err := client.V1.Files.UploadFile("/nonexistent/file.jpg")

	if err == nil {
		t.Error("Expected error for nonexistent file, but got none")
	}

	expectedError := "file does not exist: /nonexistent/file.jpg"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}
}

func TestUploadFileWithNoExtension(t *testing.T) {
	client := sdk.NewClient(sdk.WithEnv(sdk.MockServer))

	// Create a temp file without extension
	tempFile := createTempFile(t, "")
	defer os.Remove(tempFile)

	_, err := client.V1.Files.UploadFile(tempFile)

	if err == nil {
		t.Error("Expected error for file without extension, but got none")
	}

	expectedError := "file must have an extension: " + tempFile
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}
}

// Helper function to create a temporary file with the given extension
func createTempFile(t *testing.T, extension string) string {
	tempDir := t.TempDir()
	filename := "test_file"
	if extension != "" {
		filename = filename + "." + extension
	}

	tempFile := filepath.Join(tempDir, filename)

	file, err := os.Create(tempFile)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer file.Close()

	// Write some dummy content
	_, err = file.WriteString("test content")
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	return tempFile
}
