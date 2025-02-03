package nullable

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
)

// Nullable supports 3 common JSON states that don't have built-in support in Go and is
// build for JSON marshalling / unmarshalling with the encoding/json library. The three state
// are:
//
// 1. field is not set / the key does not exist (undefined)
// 2. field is explicitly set to `null`
// 3. field is explicitly set to a value of type T
//
// In the event a JSON field should be both nullable and optional, be sure
// to include the `json:"omitempty"` tag for that field so this generic type will work
// as intended.
//
// The type uses map[bool]T under the hood rather than a struct
// as encoding/json's omitempty tag does not function with embedded
// structs: https://www.sohamkamani.com/golang/omitempty/#values-that-cannot-be-omitted
// Under the hood states:
// - map[true]T => Value is set and should be present in the final marshaled JSON
// - map[false]T => Null is set and should be present in the final marshled JSON
// - Zero val or nil =>  map means the field is undefined and should *not* be present in the final marshaled JSON
type Nullable[T any] map[bool]T
type NullableLike interface {
	IsNull() bool
	IsUndefined() bool
	InterfaceValue() (interface{}, error)
}

// ----- Constructors -----

// Constructor of a `Nullable` with a given value
func NewValue[T any](t T) Nullable[T] {
	var n Nullable[T]
	n.Set(t)
	return n
}

// Constructor of a `Nullable` with an explict value of null
func NewNull[T any]() Nullable[T] {
	var n Nullable[T]
	n.SetNull()
	return n
}

// ----- Helper ------

func IsNullableInterface(v interface{}) (NullableLike, bool) {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr || val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	if val.Type().Implements(reflect.TypeOf((*NullableLike)(nil)).Elem()) {
		return val.Interface().(NullableLike), true
	}

	return nil, false
}

// ----- Methods -----

// Sets an explicit value
func (n *Nullable[T]) Set(value T) {
	*n = map[bool]T{true: value}
}

// Is the value explicitly null (rather than undefined)
func (n Nullable[T]) IsNull() bool {
	_, isNull := n[false]
	return isNull
}

// Set the value for marshalling explicitly to `null`
func (n *Nullable[T]) SetNull() {
	var zero T
	*n = map[bool]T{false: zero}
}

// Is the value not explicitly set to `null` or a value
func (t Nullable[T]) IsUndefined() bool {
	return len(t) == 0
}

// Explicitly set the Nullable to undefined
func (t *Nullable[T]) SetUndefined() {
	*t = map[bool]T{}
}

// Return the underlying value if set. Will return an error if the Nullable is set to `null` explicitly
// or if it is undefined
func (n Nullable[T]) Value() (T, error) {
	var zero T
	if n.IsNull() {
		return zero, errors.New("value is null")
	} else if n.IsUndefined() {
		return zero, errors.New("value is undefined")
	}
	return n[true], nil
}

// Return the underlying value if set as an interface{}. Helpful for generic functions where
// the underlying Nullable T is not known.
// Will return an error if the Nullable is set to `null` explicitly or if it is undefined
func (n Nullable[T]) InterfaceValue() (interface{}, error) {
	var zero T
	if n.IsNull() {
		return zero, errors.New("value is null")
	} else if n.IsUndefined() {
		return zero, errors.New("value is undefined")
	}

	return n[true], nil
}

// ----- encoding/json implementations -----

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	// Encode explicity `null` if the value is set
	if n.IsNull() {
		return []byte("null"), nil
	}

	// if field was unspecified, and thus contains a zero value for n[true] and `omitempty` is set on the field's tags,
	// `json.Marshal` will automatically omit this field

	// otherwise: we have a value, so marshal it
	return json.Marshal(n[true])
}

func (t *Nullable[T]) UnmarshalJSON(data []byte) error {
	// if field is unspecified, UnmarshalJSON won't be called

	// if field explictily present as `null`
	if bytes.Equal(data, []byte("null")) {
		t.SetNull()
		return nil
	}
	// we have an actual value, unmarshal it
	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	t.Set(v)
	return nil
}
