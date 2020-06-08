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

type float64SliceValue struct {
	ref *[]float64
	set bool
}

func newFloat64SliceValue(val []float64, p *[]float64) *float64SliceValue {
	*p = val
	return &float64SliceValue{p, false}
}

func (v *float64SliceValue) init() {
	if v.ref == nil {
		v.ref = new([]float64)
	}
}

func (v *float64SliceValue) New() interface{} { return *new([]float64) }

func (v *float64SliceValue) Type() string { return "[]float64" }

func (v *float64SliceValue) Elem() string { return "float64" }

func (v *float64SliceValue) IsSlice() bool { return true }

func (v *float64SliceValue) IsSet() bool { return v.set }

func (v *float64SliceValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *float64SliceValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.([]float64)
	v.set = true
}

func (v *float64SliceValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *float64SliceValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *float64SliceValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]float64), elem.(float64))
	return slice
}

func (v *float64SliceValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]float64)
	return val, ok
}

func (v *float64SliceValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(float64)
	return val, ok
}

func (v *float64SliceValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.([]float64))
	return string(r), err
}

func (v *float64SliceValue) Deserialize(x string) (interface{}, error) {
	val := new([]float64)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *float64SliceValue) String() string {
	return valueString(v)
}

// Float64SliceVar defines a []float64 flag with specified name, default value, and usage string.
// The argument p points to a []float64 variable in which to store the value of the flag.
func (f *FlagSet) Float64SliceVar(p *[]float64, name string, value []float64, usage string) {
	f.Var(newFloat64SliceValue(value, p), name, usage)
}

// Float64Slice defines a []float64 flag with specified name, default value, and usage string.
// The return value is the address of a []float64 variable that stores the value of the flag.
func (f *FlagSet) Float64Slice(name string, value []float64, usage string) *[]float64 {
	p := new([]float64)
	f.Float64SliceVar(p, name, value, usage)
	return p
}

// Float64SliceVar defines a []float64 flag with specified name, default value, and usage string.
// The argument p points to a []float64 variable in which to store the value of the flag.
func Float64SliceVar(p *[]float64, name string, value []float64, usage string) {
	CommandLine.Float64SliceVar(p, name, value, usage)
}

// Float64Slice defines a []float64 flag with specified name, default value, and usage string.
// The return value is the address of a []float64 variable that stores the value of the flag.
func Float64Slice(name string, value []float64, usage string) *[]float64 {
	return CommandLine.Float64Slice(name, value, usage)
}
