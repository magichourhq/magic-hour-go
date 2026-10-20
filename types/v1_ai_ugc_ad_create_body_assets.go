package types

import (
	json "encoding/json"
)

// V1AiUgcAdCreateBodyAssets
type V1AiUgcAdCreateBodyAssets struct {
	// Prepared image to animate. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	ImageFilePath        string                 `json:"image_file_path"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

func (x *V1AiUgcAdCreateBodyAssets) UnmarshalJSON(data []byte) error {
	// this type might have additional properties of type interface{}
	// to handle this scenario we unmarshal twice, once for the named object
	// properties, and once for the additional properties

	// sort additional properties from named properties
	var additionalPropsMap map[string]interface{}
	additionalPropsUnmarshalErr := json.Unmarshal(data, &additionalPropsMap)
	if additionalPropsUnmarshalErr != nil {
		return additionalPropsUnmarshalErr
	}

	// remove any known keys from additional properties
	delete(additionalPropsMap, "image_file_path")

	// convert additional props back to bytes later casting to
	additionalPropsBytes, additionalPropsMarshalErr := json.Marshal(additionalPropsMap)
	if additionalPropsMarshalErr != nil {
		return additionalPropsMarshalErr
	}

	type _V1AiUgcAdCreateBodyAssets V1AiUgcAdCreateBodyAssets // type alias avoids infinite recursion
	var unmarshaled _V1AiUgcAdCreateBodyAssets

	unmarshalErr := json.Unmarshal(data, &unmarshaled)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	err := json.Unmarshal(additionalPropsBytes, &unmarshaled.AdditionalProperties)
	if err != nil {
		return err
	}

	*x = V1AiUgcAdCreateBodyAssets(unmarshaled)

	return nil
}

func (x V1AiUgcAdCreateBodyAssets) MarshalJSON() ([]byte, error) {
	// this type might have additional properties of type
	// to handle this scenario we must flatten these properties while marshalling
	type _V1AiUgcAdCreateBodyAssets V1AiUgcAdCreateBodyAssets // type alias avoids infinite recursion
	var toMarshal = _V1AiUgcAdCreateBodyAssets(x)

	marshalled, err := json.Marshal(toMarshal)
	if err != nil {
		return nil, err
	}

	var objMap map[string]interface{}
	err = json.Unmarshal(marshalled, &objMap)
	if err != nil {
		return nil, err
	}

	for key, value := range x.AdditionalProperties {
		objMap[key] = value
	}

	return json.Marshal(objMap)
}
