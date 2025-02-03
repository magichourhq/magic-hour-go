package core

import (
	json "encoding/json"
	fmt "fmt"
	url "net/url"
	strconv "strconv"
	strings "strings"
)

func FmtStringParam(value interface{}) string {
	if intVal, ok := value.(int); ok {
		return strconv.Itoa(intVal)
	} else if int64Val, ok := value.(int64); ok {
		return strconv.FormatInt(int64Val, 10)
	} else if float64Val, ok := value.(float64); ok {
		return strconv.FormatFloat(float64Val, 'f', -1, 64)
	} else if strVal, ok := value.(string); ok {
		return strVal
	} else if stringerVal, ok := value.(fmt.Stringer); ok {
		return stringerVal.String()
	}

	// try jsonifying
	if marshaler, ok := value.(json.Marshaler); ok {
		bytesVal, err := marshaler.MarshalJSON()
		if err == nil {
			var rawVal interface{}
			if err := json.Unmarshal(bytesVal, &rawVal); err == nil {
				// strip quotes if the json marshaled result is a string
				if rawStrVal, ok := rawVal.(string); ok {
					return rawStrVal
				} else {
					return string(bytesVal)
				}
			}
		}
	}

	// fallback on debug formatting
	return fmt.Sprintf("%v", value)
}

// AddQueryParam is a generic function to add a query parameter to url.Values,
// supporting both exploding and non-exploding behavior.
func AddQueryParam(queryParams url.Values, paramName string, value interface{}, explode bool) {
	if sliceValue, ok := value.([]string); ok {
		if explode {
			// If explode is true, add each element of the slice as a separate parameter
			for _, v := range sliceValue {
				queryParams.Add(paramName, v)
			}
		} else {
			// If explode is false, join the slice elements with a comma and add as a single parameter
			queryParams.Add(paramName, strings.Join(sliceValue[:], ","))
		}
	} else {
		// If explode is false, or value is not a slice, add the value as a single parameter
		queryParams.Add(paramName, FmtStringParam(value))
	}
}

// Encodes any struct that supports json encodeing to url values
func EncodeUrlParams(val interface{}) (url.Values, error) {
	// convert to json string then map[string]interface{} to
	// ensure all name mappings/undefined/null values are handled correctly
	jsonData, err := json.Marshal(val)
	if err != nil {
		return url.Values{}, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return url.Values{}, err
	}

	values := url.Values{}
	for key, value := range data {
		values.Set(key, FmtStringParam(value))
	}

	return values, nil
}
