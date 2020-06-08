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

type boolValue struct {
	ref *bool
	set bool
}

func newBoolValue(val bool, p *bool) *boolValue {
	*p = val
	return &boolValue{p, false}
}

func (v *boolValue) init() {
	if v.ref == nil {
		v.ref = new(bool)
	}
}

func (v *boolValue) New() interface{} { return *new(bool) }

func (v *boolValue) Type() string { return "bool" }

func (v *boolValue) Elem() string { return "bool" }

func (v *boolValue) IsSlice() bool { return false }

func (v *boolValue) IsSet() bool { return v.set }

func (v *boolValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *boolValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.(bool)
	v.set = true
}

func (v *boolValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *boolValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *boolValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]bool), elem.(bool))
	return slice
}

func (v *boolValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]bool)
	return val, ok
}

func (v *boolValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(bool)
	return val, ok
}

func (v *boolValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.(bool))
	return string(r), err
}

func (v *boolValue) Deserialize(x string) (interface{}, error) {
	val := new(bool)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *boolValue) String() string {
	return valueString(v)
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
	f.Var(newBoolValue(value, p), name, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVar(p, name, value, usage)
	return p
}

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func BoolVar(p *bool, name string, value bool, usage string) {
	CommandLine.BoolVar(p, name, value, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Bool(name string, value bool, usage string) *bool {
	return CommandLine.Bool(name, value, usage)
}
