package core

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"reflect"
	"github.com/magichourhq/magic-hour-go/nullable"
)

// Utility to open file that you are certain exists
func MustOpenFile(path string) os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
	}

	return *file
}

// Handles adding files, fields, or arrays of each to a form data writer
func AddToFormDataWriter(writer *multipart.Writer, field string, value interface{}) error {
	reflectVal := reflect.ValueOf(value)
	kind := reflectVal.Kind()
	if kind == reflect.Array || kind == reflect.Slice {
		for i := 0; i < reflectVal.Len(); i++ {
			item := reflectVal.Index(i).Interface()
			if err := AddToFormDataWriter(writer, field, item); err != nil {
				return err
			}
		}
	} else if file, ok := value.(os.File); ok {
		addFileToFormDataWriter(writer, field, file)
	} else if nullableLike, ok := nullable.IsNullableInterface(value); ok {
		if nullableVal, err := nullableLike.InterfaceValue(); err == nil {
			AddToFormDataWriter(writer, field, nullableVal)
		}
	} else {
		addFieldToFormDataWriter(writer, field, value)
	}

	return nil
}

// Adds non-file to form data writer
func addFieldToFormDataWriter(writer *multipart.Writer, field string, value interface{}) error {
	label, err := writer.CreateFormField(field)
	if err != nil {
		return err
	}
	label.Write([]byte(FmtStringParam(value)))
	return nil
}

// Adds file to form data writer
func addFileToFormDataWriter(writer *multipart.Writer, field string, file os.File) error {
	part, err := writer.CreateFormFile(field, path.Base(file.Name()))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, &file)
	if err != nil {
		return err
	}

	return nil
}
