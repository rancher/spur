// Copyright 2020 Rancher Labs, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flag

import (
	"reflect"
	"strconv"
	"time"
)

// ToStringFunc is the function definition for converting types to strings
type ToStringFunc = func(interface{}) (string, bool)

// FromStringFunc is the function definition for converting strings to types
type FromStringFunc = func(string) (interface{}, error)

// ToStringMap provides a mapping of type to string conversion function
var ToStringMap = map[string]ToStringFunc{}

// FromStringMap provides a mapping of string to type conversion function
var FromStringMap = map[string]FromStringFunc{}

// TimeLayouts provides a list of layouts to attempt when converting time strings
var TimeLayouts = []string{
	time.RFC3339Nano,
	time.RFC3339,
	time.UnixDate,
	time.RubyDate,
	time.ANSIC,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.StampNano,
	time.StampMicro,
	time.StampMilli,
	time.Stamp,
	time.Kitchen,
}

// ToString is a convenience function for converting types to strings as defined in ToStringMap
func ToString(value interface{}) (string, bool) {
	if toString := ToStringMap[reflect.TypeOf(value).String()]; toString != nil {
		return toString(value)
	}
	return "", false
}

// FromString is a convenience function for converting strings to types as defined in FromStringMap
func FromString(someType string, value string) (result interface{}, err error) {
	err = errParse
	if fromString := FromStringMap[someType]; fromString != nil {
		result, err = fromString(value)
		err = numError(err)
	}
	return
}

func init() {
	initToStringFuncs()
	initFromStringFuncs()
}

func initToStringFuncs() {
	ToStringMap["string"] = func(value interface{}) (string, bool) {
		return value.(string), true
	}
	ToStringMap["bool"] = func(value interface{}) (string, bool) {
		return strconv.FormatBool(value.(bool)), true
	}
	ToStringMap["int"] = func(value interface{}) (string, bool) {
		return strconv.Itoa(value.(int)), true
	}
	ToStringMap["int64"] = func(value interface{}) (string, bool) {
		return strconv.FormatInt(value.(int64), 10), true
	}
	ToStringMap["uint"] = func(value interface{}) (string, bool) {
		return strconv.FormatUint(uint64(value.(uint)), 10), true
	}
	ToStringMap["uint64"] = func(value interface{}) (string, bool) {
		return strconv.FormatUint(value.(uint64), 10), true
	}
	ToStringMap["float64"] = func(value interface{}) (string, bool) {
		return strconv.FormatFloat(value.(float64), 'g', -1, 64), true
	}
	ToStringMap["time.Duration"] = func(value interface{}) (string, bool) {
		return value.(time.Duration).String(), true
	}
	ToStringMap["time.Time"] = func(value interface{}) (string, bool) {
		if len(TimeLayouts) > 0 {
			return value.(time.Time).Format(TimeLayouts[0]), true
		}
		return "", false
	}
}

func initFromStringFuncs() {
	FromStringMap["string"] = func(s string) (interface{}, error) {
		return string(s), nil
	}
	FromStringMap["bool"] = func(s string) (interface{}, error) {
		if s == "" {
			s = "false"
		}
		v, err := strconv.ParseBool(s)
		return bool(v), err
	}
	FromStringMap["int"] = func(s string) (interface{}, error) {
		v, err := strconv.ParseInt(s, 0, strconv.IntSize)
		return int(v), err
	}
	FromStringMap["int64"] = func(s string) (interface{}, error) {
		v, err := strconv.ParseInt(s, 0, 64)
		return int64(v), err
	}
	FromStringMap["uint"] = func(s string) (interface{}, error) {
		v, err := strconv.ParseUint(s, 0, strconv.IntSize)
		return uint(v), err
	}
	FromStringMap["uint64"] = func(s string) (interface{}, error) {
		v, err := strconv.ParseUint(s, 0, 64)
		return uint64(v), err
	}
	FromStringMap["float64"] = func(s string) (interface{}, error) {
		v, err := strconv.ParseFloat(s, 64)
		return float64(v), err
	}
	FromStringMap["time.Duration"] = func(s string) (interface{}, error) {
		if v, err := time.ParseDuration(s); err == nil {
			return time.Duration(v), nil
		}
		return nil, errParse
	}
	FromStringMap["time.Time"] = func(s string) (interface{}, error) {
		if v, err := strconv.ParseInt(s, 0, 64); err == nil {
			return time.Unix(v, 0), nil
		}
		for _, layout := range TimeLayouts {
			if v, err := time.Parse(layout, s); err == nil {
				return time.Time(v), nil
			}
		}
		return nil, errParse
	}
}

type genericValue interface {
	New() interface{}
	Elem() string
	IsSlice() bool
	IsSet() bool
	Assign(value interface{})
	Get() interface{}
	Append(interface{}, interface{}) interface{}
	ConvertSlice(interface{}) (interface{}, bool)
	ConvertElem(interface{}) (interface{}, bool)
	Serialize(interface{}) (string, error)
	Deserialize(x string) (interface{}, error)
}

func valueSet(v genericValue, value interface{}) error {
	if v.IsSlice() && !v.IsSet() {
		// If this is a slice and has not already been set then
		// clear any existing value
		v.Assign(v.New())
	}
	val, err := valueConvert(v, value)
	if err != nil {
		return err
	}
	v.Assign(val)
	return nil
}

func valueConvert(v genericValue, value interface{}) (interface{}, error) {
	// Try deserializing as string
	if s, ok := value.(string); ok {
		if val, err := v.Deserialize(s); err == nil {
			return val, nil
		}
	}
	// Convert an element
	elem, err := valueConvertElem(v, value)
	if !v.IsSlice() {
		// Return value and error if not a slice
		return elem, err
	}
	// Try converting as slice
	if val, ok := v.ConvertSlice(value); ok {
		return val, nil
	}
	// If no error from converting element return appended value
	if err == nil {
		return v.Append(v.Get(), elem), nil
	}
	// Try evaluating value as a slice of interfaces
	otherValue, ok := value.([]interface{})
	if !ok {
		return nil, errParse
	}
	// Create a new slice and append each converted element
	slice := v.New()
	for _, other := range otherValue {
		elem, err := valueConvertElem(v, other)
		if err != nil {
			return nil, err
		}
		slice = v.Append(slice, elem)
	}
	return slice, nil
}

func valueConvertElem(v genericValue, value interface{}) (interface{}, error) {
	// Try converting as an element type
	if val, ok := v.ConvertElem(value); ok {
		return val, nil
	}
	// Get our value as a string
	s, ok := value.(string)
	if !ok {
		if s, ok = ToString(value); !ok {
			return nil, errParse
		}
	}
	// Return a new value from the string
	return FromString(v.Elem(), s)
}

func valueString(v genericValue) string {
	if s, ok := ToString(v.Get()); ok {
		return s
	}
	if s, err := v.Serialize(v.Get()); err == nil {
		return s
	}
	panic(errParse)
}
