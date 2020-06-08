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

type stringValue struct {
	ref *string
	set bool
}

func newStringValue(val string, p *string) *stringValue {
	*p = val
	return &stringValue{p, false}
}

func (v *stringValue) init() {
	if v.ref == nil {
		v.ref = new(string)
	}
}

func (v *stringValue) New() interface{} { return *new(string) }

func (v *stringValue) Type() string { return "string" }

func (v *stringValue) Elem() string { return "string" }

func (v *stringValue) IsSlice() bool { return false }

func (v *stringValue) IsSet() bool { return v.set }

func (v *stringValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *stringValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.(string)
	v.set = true
}

func (v *stringValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *stringValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *stringValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]string), elem.(string))
	return slice
}

func (v *stringValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]string)
	return val, ok
}

func (v *stringValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(string)
	return val, ok
}

func (v *stringValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.(string))
	return string(r), err
}

func (v *stringValue) Deserialize(x string) (interface{}, error) {
	val := new(string)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *stringValue) String() string {
	return valueString(v)
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
	f.Var(newStringValue(value, p), name, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSet) String(name string, value string, usage string) *string {
	p := new(string)
	f.StringVar(p, name, value, usage)
	return p
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, value string, usage string) {
	CommandLine.StringVar(p, name, value, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func String(name string, value string, usage string) *string {
	return CommandLine.String(name, value, usage)
}
