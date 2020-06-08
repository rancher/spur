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

type uint64Value struct {
	ref *uint64
	set bool
}

func newUint64Value(val uint64, p *uint64) *uint64Value {
	*p = val
	return &uint64Value{p, false}
}

func (v *uint64Value) init() {
	if v.ref == nil {
		v.ref = new(uint64)
	}
}

func (v *uint64Value) New() interface{} { return *new(uint64) }

func (v *uint64Value) Type() string { return "uint64" }

func (v *uint64Value) Elem() string { return "uint64" }

func (v *uint64Value) IsSlice() bool { return false }

func (v *uint64Value) IsSet() bool { return v.set }

func (v *uint64Value) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *uint64Value) Assign(value interface{}) {
	v.init()
	*v.ref = value.(uint64)
	v.set = true
}

func (v *uint64Value) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *uint64Value) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *uint64Value) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]uint64), elem.(uint64))
	return slice
}

func (v *uint64Value) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]uint64)
	return val, ok
}

func (v *uint64Value) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(uint64)
	return val, ok
}

func (v *uint64Value) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.(uint64))
	return string(r), err
}

func (v *uint64Value) Deserialize(x string) (interface{}, error) {
	val := new(uint64)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *uint64Value) String() string {
	return valueString(v)
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.Var(newUint64Value(value, p), name, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64Var(p, name, value, usage)
	return p
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	CommandLine.Uint64Var(p, name, value, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func Uint64(name string, value uint64, usage string) *uint64 {
	return CommandLine.Uint64(name, value, usage)
}
