package core

import (
	json "encoding/json"
	fmt "fmt"
	url "net/url"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

func FmtStringParam(value interface{}) string {
	if value == nil {
		return "null"
	} else if intVal, ok := value.(int); ok {
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

func AddQueryParam(queryParams url.Values, paramName string, value interface{}, style string, explode bool) {
	if style == "form" {
		addFormQueryParam(queryParams, paramName, value, explode)
	} else if style == "spaceDelimited" {
		addSpaceDelimitedQueryParam(queryParams, paramName, value, explode)
	} else if style == "pipeDelimited" {
		addPipeDelimitedQueryParam(queryParams, paramName, value, explode)
	} else if style == "deepObject" {
		addDeepObjQueryParam(queryParams, paramName, value, explode)
	} else {
		panic(fmt.Sprintf("query param style '%s' not implemented", style))
	}

}

func addFormQueryParam(queryParams url.Values, paramName string, value interface{}, explode bool) {
	v := reflect.ValueOf(value)
	// handle pointers
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Map:
		// explode form maps should be encoded like /users?key0=val0&key1=val1
		// the input param name will be omitted
		// non-explode form maps should be encoded like /users?id=key0,val0,key1,val1
		var chunks []string
		for _, mapKey := range v.MapKeys() {
			mapVal := FmtStringParam(v.MapIndex(mapKey).Interface())

			chunks = append(chunks, mapKey.String(), mapVal)

			if explode {
				queryParams.Add(mapKey.String(), mapVal)
			}
		}

		if !explode && len(chunks) > 0 {
			queryParams.Add(paramName, strings.Join(chunks, ","))
		}

	case reflect.Struct:
		// structs that are part of a query param must implement json marshaling
		// marshal then unmarshal back to an interface to process.
		jsonData, err := json.Marshal(value)
		if err == nil {
			var jsonInterface interface{}
			err = json.Unmarshal(jsonData, &jsonInterface)

			if err == nil {
				addFormQueryParam(queryParams, paramName, jsonInterface, explode)
				return
			}
		}

		fmt.Printf("Failed converting complex struct into native map/primitive: %v", err)
		return

	case reflect.Slice, reflect.Array:
		if explode {
			// explode form lists should be encoded like /users?id=3&id=4&id=5
			for i := 0; i < v.Len(); i++ {
				queryParams.Add(paramName, FmtStringParam(v.Index(i).Interface()))
			}
		} else {
			// non-explode form lists should be encoded like /users?id=3,4,5
			var items []string
			for i := 0; i < v.Len(); i++ {
				items = append(items, FmtStringParam(v.Index(i).Interface()))
			}
			if len(items) > 0 {
				queryParams.Add(paramName, strings.Join(items, ","))
			}
		}
	default:
		queryParams.Add(paramName, FmtStringParam(value))
	}
}

func addSpaceDelimitedQueryParam(queryParams url.Values, paramName string, value interface{}, explode bool) {
	v := reflect.ValueOf(value)
	// handle pointers
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if (v.Kind() == reflect.Array || v.Kind() == reflect.Slice) && !explode {
		// non-explode spaceDelimited lists should be encoded like /users?id=3%204%205
		var items []string
		for i := 0; i < v.Len(); i++ {
			items = append(items, FmtStringParam(v.Index(i).Interface()))
		}

		if len(items) > 0 {
			queryParams.Add(paramName, strings.Join(items, " "))
		}
	} else {
		// according to the docs, spaceDelimited + explode=false only effects lists,
		// all other encodings are marked as n/a or are the same as `form` style
		// fall back on form style as it is the default for query params
		addFormQueryParam(queryParams, paramName, value, explode)
	}
}

func addPipeDelimitedQueryParam(queryParams url.Values, paramName string, value interface{}, explode bool) {
	v := reflect.ValueOf(value)
	// handle pointers
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if (v.Kind() == reflect.Array || v.Kind() == reflect.Slice) && !explode {
		// non-explode pipeDelimited lists should be encoded like /users?id=3|4|5
		var items []string
		for i := 0; i < v.Len(); i++ {
			items = append(items, FmtStringParam(v.Index(i).Interface()))
		}

		if len(items) > 0 {
			queryParams.Add(paramName, strings.Join(items, "|"))
		}
	} else {
		// according to the docs, pipeDelimited + explode=false only effects lists,
		// all other encodings are marked as n/a or are the same as `form` style
		// fall back on form style as it is the default for query params
		addFormQueryParam(queryParams, paramName, value, explode)
	}
}

func addDeepObjQueryParam(queryParams url.Values, paramName string, value interface{}, explode bool) {
	v := reflect.ValueOf(value)
	// handle pointers
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice:
		encodeDeepObjectKey(queryParams, paramName, value)
	default:
		// according to the docs, deepObject style only applies to
		// object encodes, encodings for primitives & arrays are listed as n/a,
		// fall back on form style as it is the default for query params
		addFormQueryParam(queryParams, paramName, value, explode)
	}

}

func encodeDeepObjectKey(queryParams url.Values, key string, value interface{}) {
	v := reflect.ValueOf(value)
	// handle pointers
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Map:
		for _, mapKey := range v.MapKeys() {
			encodeDeepObjectKey(
				queryParams,
				fmt.Sprintf("%s[%s]", key, mapKey.String()),
				v.MapIndex(mapKey).Interface(),
			)
		}

	case reflect.Struct:
		// structs that are part of a query param must implement json marshaling
		// marshal then unmarshal back to an interface to process.
		jsonData, err := json.Marshal(value)
		if err == nil {
			var jsonInterface interface{}
			err = json.Unmarshal(jsonData, &jsonInterface)

			if err == nil {
				encodeDeepObjectKey(queryParams, key, jsonInterface)
				return
			}
		}

		fmt.Printf("Failed converting complex struct into native map/primitive: %v", err)
		return

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			encodeDeepObjectKey(queryParams, fmt.Sprintf("%s[%d]", key, i), v.Index(i).Interface())
		}
	default:
		queryParams.Add(key, FmtStringParam(value))
	}

}

// Encodes any struct that supports json encodeing to url values
func FormUrlEncodedBody(value interface{}, styleMap map[string]string, explodeMap map[string]bool) (*strings.Reader, error) {
	v := reflect.ValueOf(value)
	// handle pointers
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Map:
		formValues := url.Values{}
		for _, mapKey := range v.MapKeys() {
			key := mapKey.String()
			style, styleOk := styleMap[key]
			if !styleOk {
				style = "form"
			}
			explode, explodeOk := explodeMap[key]
			if !explodeOk {
				explode = style == "form"
			}

			AddQueryParam(formValues, key, v.MapIndex(mapKey).Interface(), style, explode)
		}

		bodyBuf := strings.NewReader(formValues.Encode())
		return bodyBuf, nil

	case reflect.Struct:
		// structs that are part of a form url encoded body must implement json marshaling
		// marshal then unmarshal back to an interface to process.
		jsonData, err := json.Marshal(value)
		if err == nil {
			var jsonInterface interface{}
			err = json.Unmarshal(jsonData, &jsonInterface)

			if err == nil {
				return FormUrlEncodedBody(jsonInterface, styleMap, explodeMap)
			}
		}
		return &strings.Reader{}, err

	default:
		return &strings.Reader{}, fmt.Errorf("x-www-form-urlencoded data must be a map or a struct at the top level")
	}
}
