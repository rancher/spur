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

type int64Value struct {
	ref *int64
	set bool
}

func newInt64Value(val int64, p *int64) *int64Value {
	*p = val
	return &int64Value{p, false}
}

func (v *int64Value) init() {
	if v.ref == nil {
		v.ref = new(int64)
	}
}

func (v *int64Value) New() interface{} { return *new(int64) }

func (v *int64Value) Type() string { return "int64" }

func (v *int64Value) Elem() string { return "int64" }

func (v *int64Value) IsSlice() bool { return false }

func (v *int64Value) IsSet() bool { return v.set }

func (v *int64Value) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *int64Value) Assign(value interface{}) {
	v.init()
	*v.ref = value.(int64)
	v.set = true
}

func (v *int64Value) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *int64Value) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *int64Value) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]int64), elem.(int64))
	return slice
}

func (v *int64Value) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]int64)
	return val, ok
}

func (v *int64Value) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(int64)
	return val, ok
}

func (v *int64Value) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.(int64))
	return string(r), err
}

func (v *int64Value) Deserialize(x string) (interface{}, error) {
	val := new(int64)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *int64Value) String() string {
	return valueString(v)
}

// Int64Var defines a int64 flag with specified name, default value, and usage string.
// The argument p points to a int64 variable in which to store the value of the flag.
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string) {
	f.Var(newInt64Value(value, p), name, usage)
}

// Int64 defines a int64 flag with specified name, default value, and usage string.
// The return value is the address of a int64 variable that stores the value of the flag.
func (f *FlagSet) Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64Var(p, name, value, usage)
	return p
}

// Int64Var defines a int64 flag with specified name, default value, and usage string.
// The argument p points to a int64 variable in which to store the value of the flag.
func Int64Var(p *int64, name string, value int64, usage string) {
	CommandLine.Int64Var(p, name, value, usage)
}

// Int64 defines a int64 flag with specified name, default value, and usage string.
// The return value is the address of a int64 variable that stores the value of the flag.
func Int64(name string, value int64, usage string) *int64 {
	return CommandLine.Int64(name, value, usage)
}
