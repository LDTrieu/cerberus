package utils

import (
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Deal with variadic functions in Go.
func DoubleSlice(s interface{}) []interface{} {
	v := reflect.ValueOf(s)
	items := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		items[i] = v.Index(i).Interface()
	}
	return items
}

// Get value from interface.
// Return default value if key not exists or key exists with zero value.
func GetFromInterface(src map[string]interface{}, key string, defaultValue interface{}) interface{} {
	value, exists := src[key]
	if !exists || IsZero(value) {
		return defaultValue
	}
	return value
}

// Get value from interface.
// Return default value if key not exists or key exists with zero value.
func GetFromInterfaceV2[T any](src map[string]interface{}, key string, defaultValue T) T {
	value, exists := src[key]
	if !exists || IsZero(value) {
		return defaultValue
	}
	return value.(T)
}

// Get value from interface.
// Return default value if key not exists or key exists with zero value (optional).
//
// NOTE: when checkZeroValue is true, it will check if the value is zero value.
// When checkZeroValue is false, it only check if the key exists.
func GetFromInterfaceV3[T any](src map[string]interface{}, key string, defaultValue T, checkZeroValue ...bool) T {
	value, exists := src[key]
	if !exists || (IsZero(value) && len(checkZeroValue) > 0 && checkZeroValue[0]) {
		return defaultValue
	}
	return value.(T)
}

// Detect whether a value is the zero value for its type.
func IsZero(p interface{}) bool {
	v := reflect.ValueOf(p)
	if v.Kind() == reflect.Slice {
		return v.Len() == 0
	}
	return !v.IsValid() || reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

func DecodeModel(input map[string]interface{}, result interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata: nil,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			toTimeHookFunc()),
		Result:  result,
		TagName: "json",
	})
	if err != nil {
		return err
	}

	// nolint
	if err := decoder.Decode(input); err != nil {
		return err
	}
	return err
}

func toTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		switch f.Kind() {
		case reflect.String:
			return time.Parse(time.RFC3339, data.(string))
		case reflect.Float64:
			return time.Unix(0, int64(data.(float64))*int64(time.Millisecond)), nil
		case reflect.Int64:
			return time.Unix(0, data.(int64)*int64(time.Millisecond)), nil
		default:
			return data, nil
		}
		// Convert it by parsing
	}
}
