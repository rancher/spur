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

type int64SliceValue struct {
	ref *[]int64
	set bool
}

func newInt64SliceValue(val []int64, p *[]int64) *int64SliceValue {
	*p = val
	return &int64SliceValue{p, false}
}

func (v *int64SliceValue) init() {
	if v.ref == nil {
		v.ref = new([]int64)
	}
}

func (v *int64SliceValue) New() interface{} { return *new([]int64) }

func (v *int64SliceValue) Type() string { return "[]int64" }

func (v *int64SliceValue) Elem() string { return "int64" }

func (v *int64SliceValue) IsSlice() bool { return true }

func (v *int64SliceValue) IsSet() bool { return v.set }

func (v *int64SliceValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *int64SliceValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.([]int64)
	v.set = true
}

func (v *int64SliceValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *int64SliceValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *int64SliceValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]int64), elem.(int64))
	return slice
}

func (v *int64SliceValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]int64)
	return val, ok
}

func (v *int64SliceValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(int64)
	return val, ok
}

func (v *int64SliceValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.([]int64))
	return string(r), err
}

func (v *int64SliceValue) Deserialize(x string) (interface{}, error) {
	val := new([]int64)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *int64SliceValue) String() string {
	return valueString(v)
}

// Int64SliceVar defines a []int64 flag with specified name, default value, and usage string.
// The argument p points to a []int64 variable in which to store the value of the flag.
func (f *FlagSet) Int64SliceVar(p *[]int64, name string, value []int64, usage string) {
	f.Var(newInt64SliceValue(value, p), name, usage)
}

// Int64Slice defines a []int64 flag with specified name, default value, and usage string.
// The return value is the address of a []int64 variable that stores the value of the flag.
func (f *FlagSet) Int64Slice(name string, value []int64, usage string) *[]int64 {
	p := new([]int64)
	f.Int64SliceVar(p, name, value, usage)
	return p
}

// Int64SliceVar defines a []int64 flag with specified name, default value, and usage string.
// The argument p points to a []int64 variable in which to store the value of the flag.
func Int64SliceVar(p *[]int64, name string, value []int64, usage string) {
	CommandLine.Int64SliceVar(p, name, value, usage)
}

// Int64Slice defines a []int64 flag with specified name, default value, and usage string.
// The return value is the address of a []int64 variable that stores the value of the flag.
func Int64Slice(name string, value []int64, usage string) *[]int64 {
	return CommandLine.Int64Slice(name, value, usage)
}
