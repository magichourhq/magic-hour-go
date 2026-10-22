package types

import (
	json "encoding/json"

	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiMusicGeneratorCreateBody
type V1AiMusicGeneratorCreateBody struct {
	// Requested track duration in seconds, from 15 to 300. Determines credits charged.
	Duration int `json:"duration"`
	// Give your audio a custom name for easy identification.
	Name                 nullable.Nullable[string]         `json:"name,omitempty"`
	Style                V1AiMusicGeneratorCreateBodyStyle `json:"style"`
	AdditionalProperties map[string]interface{}            `json:"-"`
}

func (x *V1AiMusicGeneratorCreateBody) UnmarshalJSON(data []byte) error {
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
	delete(additionalPropsMap, "duration")
	delete(additionalPropsMap, "name")
	delete(additionalPropsMap, "style")

	// convert additional props back to bytes later casting to
	additionalPropsBytes, additionalPropsMarshalErr := json.Marshal(additionalPropsMap)
	if additionalPropsMarshalErr != nil {
		return additionalPropsMarshalErr
	}

	type _V1AiMusicGeneratorCreateBody V1AiMusicGeneratorCreateBody // type alias avoids infinite recursion
	var unmarshaled _V1AiMusicGeneratorCreateBody

	unmarshalErr := json.Unmarshal(data, &unmarshaled)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	err := json.Unmarshal(additionalPropsBytes, &unmarshaled.AdditionalProperties)
	if err != nil {
		return err
	}

	*x = V1AiMusicGeneratorCreateBody(unmarshaled)

	return nil
}

func (x V1AiMusicGeneratorCreateBody) MarshalJSON() ([]byte, error) {
	// this type might have additional properties of type
	// to handle this scenario we must flatten these properties while marshalling
	type _V1AiMusicGeneratorCreateBody V1AiMusicGeneratorCreateBody // type alias avoids infinite recursion
	var toMarshal = _V1AiMusicGeneratorCreateBody(x)

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
