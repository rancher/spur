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

type boolSliceValue struct {
	ref *[]bool
	set bool
}

func newBoolSliceValue(val []bool, p *[]bool) *boolSliceValue {
	*p = val
	return &boolSliceValue{p, false}
}

func (v *boolSliceValue) init() {
	if v.ref == nil {
		v.ref = new([]bool)
	}
}

func (v *boolSliceValue) New() interface{} { return *new([]bool) }

func (v *boolSliceValue) Type() string { return "[]bool" }

func (v *boolSliceValue) Elem() string { return "bool" }

func (v *boolSliceValue) IsSlice() bool { return true }

func (v *boolSliceValue) IsSet() bool { return v.set }

func (v *boolSliceValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *boolSliceValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.([]bool)
	v.set = true
}

func (v *boolSliceValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *boolSliceValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *boolSliceValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]bool), elem.(bool))
	return slice
}

func (v *boolSliceValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]bool)
	return val, ok
}

func (v *boolSliceValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(bool)
	return val, ok
}

func (v *boolSliceValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.([]bool))
	return string(r), err
}

func (v *boolSliceValue) Deserialize(x string) (interface{}, error) {
	val := new([]bool)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *boolSliceValue) String() string {
	return valueString(v)
}

// BoolSliceVar defines a []bool flag with specified name, default value, and usage string.
// The argument p points to a []bool variable in which to store the value of the flag.
func (f *FlagSet) BoolSliceVar(p *[]bool, name string, value []bool, usage string) {
	f.Var(newBoolSliceValue(value, p), name, usage)
}

// BoolSlice defines a []bool flag with specified name, default value, and usage string.
// The return value is the address of a []bool variable that stores the value of the flag.
func (f *FlagSet) BoolSlice(name string, value []bool, usage string) *[]bool {
	p := new([]bool)
	f.BoolSliceVar(p, name, value, usage)
	return p
}

// BoolSliceVar defines a []bool flag with specified name, default value, and usage string.
// The argument p points to a []bool variable in which to store the value of the flag.
func BoolSliceVar(p *[]bool, name string, value []bool, usage string) {
	CommandLine.BoolSliceVar(p, name, value, usage)
}

// BoolSlice defines a []bool flag with specified name, default value, and usage string.
// The return value is the address of a []bool variable that stores the value of the flag.
func BoolSlice(name string, value []bool, usage string) *[]bool {
	return CommandLine.BoolSlice(name, value, usage)
}
