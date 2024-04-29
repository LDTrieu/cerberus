package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ArrayJsonNumberToString(a []json.Number, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func ArrayIntToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func ArrayInt64ToString(a []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func ArrayStringToString(a []string, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func StringToArrayString(str, delim string, trim bool) []string {
	var result []string
	for _, v := range strings.Split(str, delim) {
		if trim {
			v = strings.TrimSpace(v)
		}
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

func StringToArrayInt(str, delim string) []int {
	var result []int
	for _, v := range strings.Split(str, delim) {
		if v != "" {
			i, _ := strconv.Atoi(v)
			result = append(result, i)
		}
	}
	return result
}

func StringToArrayInt64(str string, delim string) []int64 {
	var result []int64
	for _, v := range strings.Split(str, delim) {
		if v != "" {
			i, _ := strconv.ParseInt(v, 10, 64)
			result = append(result, i)
		}
	}
	return result
}

func StringToBitwiseInt64(str string) int64 {
	var result int64
	for _, v := range strings.Split(str, ",") {
		if v != "" {
			i, _ := strconv.ParseInt(v, 10, 64)
			result |= i
		}
	}
	return result
}
func BitwiseToString(bitwise int64) string {
	var result []string
	for i := 0; i < 64; i++ {
		if bitwise&(1<<uint(i)) > 0 {
			result = append(result, strconv.Itoa(i))
		}
	}
	return strings.Join(result, ",")
}
func DistinctStr(str string, delim string) string {
	return ArrayStringToString(Distinct(StringToArrayString(str, ",", false)), ",")
}

func ToPointer[T any](i T) *T {
	return &i
}

type parsable interface {
	int | int32 | int64 | float32 | float64 | string
}

// Parse[T](a interface{}) (T,error)

// Number to String OR String to Number
func Parse[T parsable](in interface{}) (out T, err error) {
	_in := reflect.TypeOf(in)
	_out := reflect.TypeOf(&out).Elem()

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error parsing %v to %v with err %v", reflect.TypeOf(in), reflect.TypeOf(out).Elem(), r)
		}
	}()

	switch _in.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		switch _out.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64:
			reflect.ValueOf(&out).Elem().SetInt(reflect.ValueOf(in).Int())
		case reflect.Float32, reflect.Float64:
			reflect.ValueOf(&out).Elem().SetFloat(float64(reflect.ValueOf(in).Int()))
		case reflect.String:
			reflect.ValueOf(&out).Elem().SetString(strconv.FormatInt(reflect.ValueOf(in).Int(), 10))
		}
	case reflect.Float32, reflect.Float64:
		switch _out.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64:
			reflect.ValueOf(&out).Elem().SetInt(int64(reflect.ValueOf(in).Float()))
		case reflect.Float32, reflect.Float64:
			reflect.ValueOf(&out).Elem().SetFloat(reflect.ValueOf(in).Float())
		case reflect.String:
			reflect.ValueOf(&out).Elem().SetString(strconv.FormatFloat(reflect.ValueOf(in).Float(), 'f', -1, 64))
		}
	case reflect.String:
		switch _out.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(reflect.ValueOf(in).String(), 10, 64)
			if err != nil {
				return out, err
			}
			reflect.ValueOf(&out).Elem().SetInt(int64(x))
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(reflect.ValueOf(in).String(), 64)
			if err != nil {
				return out, err
			}
			reflect.ValueOf(&out).Elem().SetFloat(float64(x))
		case reflect.String:
			reflect.ValueOf(&out).Elem().SetString(reflect.ValueOf(in).String())
		}
	}
	return out, nil
}
