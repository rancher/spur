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

type float64Value struct {
	ref *float64
	set bool
}

func newFloat64Value(val float64, p *float64) *float64Value {
	*p = val
	return &float64Value{p, false}
}

func (v *float64Value) init() {
	if v.ref == nil {
		v.ref = new(float64)
	}
}

func (v *float64Value) New() interface{} { return *new(float64) }

func (v *float64Value) Type() string { return "float64" }

func (v *float64Value) Elem() string { return "float64" }

func (v *float64Value) IsSlice() bool { return false }

func (v *float64Value) IsSet() bool { return v.set }

func (v *float64Value) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *float64Value) Assign(value interface{}) {
	v.init()
	*v.ref = value.(float64)
	v.set = true
}

func (v *float64Value) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *float64Value) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *float64Value) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]float64), elem.(float64))
	return slice
}

func (v *float64Value) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]float64)
	return val, ok
}

func (v *float64Value) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(float64)
	return val, ok
}

func (v *float64Value) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.(float64))
	return string(r), err
}

func (v *float64Value) Deserialize(x string) (interface{}, error) {
	val := new(float64)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *float64Value) String() string {
	return valueString(v)
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string) {
	f.Var(newFloat64Value(value, p), name, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func (f *FlagSet) Float64(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64Var(p, name, value, usage)
	return p
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func Float64Var(p *float64, name string, value float64, usage string) {
	CommandLine.Float64Var(p, name, value, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func Float64(name string, value float64, usage string) *float64 {
	return CommandLine.Float64(name, value, usage)
}
