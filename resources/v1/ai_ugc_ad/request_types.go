package ai_ugc_ad

import (
	json "encoding/json"

	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	Assets types.V1AiUgcAdCreateBodyAssets `json:"assets"`
	// Video duration in seconds. Supported values by model:
	// * kling-3.0: 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15 seconds.
	// * ltx-2.5: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60 seconds.
	// * veo3.1: 4, 6, 8, 15, 22, 29, 36, 43, 50, 57, 64, 71, 78, 85, 92, 99, 106, 113, 120, 127, 134, 141 seconds.
	// Free plans are limited to 5 seconds unless a credit pack has been purchased.
	EndSeconds float64 `json:"end_seconds"`
	// Defaults to kling-3.0 on paid plans and ltx-2.5 on free plans.
	Model nullable.Nullable[types.V1AiUgcAdCreateBodyModelEnum] `json:"model,omitempty"`
	// Defaults to 720p on paid plans and 480p on free plans. Unsupported model/resolution combinations and resolutions unavailable for your plan are rejected.
	Resolution nullable.Nullable[types.V1AiUgcAdCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// The words the actor should speak. Character limits apply after trimming:
	// * kling-3.0: 2500 characters.
	// * ltx-2.5: 1804 characters.
	// * veo3.1: 2500 characters.
	// LTX's dialogue wrapper counts toward its model prompt limit.
	Script               string                 `json:"script"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

func (x *CreateRequest) UnmarshalJSON(data []byte) error {
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
	delete(additionalPropsMap, "assets")
	delete(additionalPropsMap, "end_seconds")
	delete(additionalPropsMap, "model")
	delete(additionalPropsMap, "resolution")
	delete(additionalPropsMap, "script")

	// convert additional props back to bytes later casting to
	additionalPropsBytes, additionalPropsMarshalErr := json.Marshal(additionalPropsMap)
	if additionalPropsMarshalErr != nil {
		return additionalPropsMarshalErr
	}

	type _CreateRequest CreateRequest // type alias avoids infinite recursion
	var unmarshaled _CreateRequest

	unmarshalErr := json.Unmarshal(data, &unmarshaled)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	err := json.Unmarshal(additionalPropsBytes, &unmarshaled.AdditionalProperties)
	if err != nil {
		return err
	}

	*x = CreateRequest(unmarshaled)

	return nil
}

func (x CreateRequest) MarshalJSON() ([]byte, error) {
	// this type might have additional properties of type
	// to handle this scenario we must flatten these properties while marshalling
	type _CreateRequest CreateRequest // type alias avoids infinite recursion
	var toMarshal = _CreateRequest(x)

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
