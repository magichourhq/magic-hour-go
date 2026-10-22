package types

import (
	json "encoding/json"

	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiMusicGeneratorCreateBodyStyle
type V1AiMusicGeneratorCreateBodyStyle struct {
	// Generate music without vocals. When true, lyrics is ignored.
	Instrumental nullable.Nullable[bool] `json:"instrumental,omitempty"`
	// Custom lyrics. Omit or leave blank for AI-generated lyrics when instrumental is false.
	Lyrics nullable.Nullable[string] `json:"lyrics,omitempty"`
	// Describe the song, including genre, mood, and instruments.
	Prompt               string                 `json:"prompt"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

func (x *V1AiMusicGeneratorCreateBodyStyle) UnmarshalJSON(data []byte) error {
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
	delete(additionalPropsMap, "instrumental")
	delete(additionalPropsMap, "lyrics")
	delete(additionalPropsMap, "prompt")

	// convert additional props back to bytes later casting to
	additionalPropsBytes, additionalPropsMarshalErr := json.Marshal(additionalPropsMap)
	if additionalPropsMarshalErr != nil {
		return additionalPropsMarshalErr
	}

	type _V1AiMusicGeneratorCreateBodyStyle V1AiMusicGeneratorCreateBodyStyle // type alias avoids infinite recursion
	var unmarshaled _V1AiMusicGeneratorCreateBodyStyle

	unmarshalErr := json.Unmarshal(data, &unmarshaled)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	err := json.Unmarshal(additionalPropsBytes, &unmarshaled.AdditionalProperties)
	if err != nil {
		return err
	}

	*x = V1AiMusicGeneratorCreateBodyStyle(unmarshaled)

	return nil
}

func (x V1AiMusicGeneratorCreateBodyStyle) MarshalJSON() ([]byte, error) {
	// this type might have additional properties of type
	// to handle this scenario we must flatten these properties while marshalling
	type _V1AiMusicGeneratorCreateBodyStyle V1AiMusicGeneratorCreateBodyStyle // type alias avoids infinite recursion
	var toMarshal = _V1AiMusicGeneratorCreateBodyStyle(x)

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
