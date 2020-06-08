// Copyright 2020 Rancher Labs, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flag

import (
	"encoding/json"
	"time"

	"gopkg.in/yaml.v2"
)

var _ = time.Time{}

type stringSliceValue struct {
	ref *[]string
	set bool
}

func newStringSliceValue(val []string, p *[]string) *stringSliceValue {
	*p = val
	return &stringSliceValue{p, false}
}

func (v *stringSliceValue) init() {
	if v.ref == nil {
		v.ref = new([]string)
	}
}

func (v *stringSliceValue) New() interface{} { return *new([]string) }

func (v *stringSliceValue) Type() string { return "[]string" }

func (v *stringSliceValue) Elem() string { return "string" }

func (v *stringSliceValue) IsSlice() bool { return true }

func (v *stringSliceValue) IsSet() bool { return v.set }

func (v *stringSliceValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *stringSliceValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.([]string)
	v.set = true
}

func (v *stringSliceValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *stringSliceValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *stringSliceValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]string), elem.(string))
	return slice
}

func (v *stringSliceValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]string)
	return val, ok
}

func (v *stringSliceValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(string)
	return val, ok
}

func (v *stringSliceValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.([]string))
	return string(r), err
}

func (v *stringSliceValue) Deserialize(x string) (interface{}, error) {
	val := new([]string)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *stringSliceValue) String() string {
	return valueString(v)
}

// StringSliceVar defines a []string flag with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the flag.
func (f *FlagSet) StringSliceVar(p *[]string, name string, value []string, usage string) {
	f.Var(newStringSliceValue(value, p), name, usage)
}

// StringSlice defines a []string flag with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the flag.
func (f *FlagSet) StringSlice(name string, value []string, usage string) *[]string {
	p := new([]string)
	f.StringSliceVar(p, name, value, usage)
	return p
}

// StringSliceVar defines a []string flag with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the flag.
func StringSliceVar(p *[]string, name string, value []string, usage string) {
	CommandLine.StringSliceVar(p, name, value, usage)
}

// StringSlice defines a []string flag with specified name, default value, and usage string.
// The return value is the address of a []string variable that stores the value of the flag.
func StringSlice(name string, value []string, usage string) *[]string {
	return CommandLine.StringSlice(name, value, usage)
}
