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

type uintValue struct {
	ref *uint
	set bool
}

func newUintValue(val uint, p *uint) *uintValue {
	*p = val
	return &uintValue{p, false}
}

func (v *uintValue) init() {
	if v.ref == nil {
		v.ref = new(uint)
	}
}

func (v *uintValue) New() interface{} { return *new(uint) }

func (v *uintValue) Type() string { return "uint" }

func (v *uintValue) Elem() string { return "uint" }

func (v *uintValue) IsSlice() bool { return false }

func (v *uintValue) IsSet() bool { return v.set }

func (v *uintValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *uintValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.(uint)
	v.set = true
}

func (v *uintValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *uintValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *uintValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]uint), elem.(uint))
	return slice
}

func (v *uintValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]uint)
	return val, ok
}

func (v *uintValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(uint)
	return val, ok
}

func (v *uintValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.(uint))
	return string(r), err
}

func (v *uintValue) Deserialize(x string) (interface{}, error) {
	val := new(uint)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *uintValue) String() string {
	return valueString(v)
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string) {
	f.Var(newUintValue(value, p), name, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag.
func (f *FlagSet) Uint(name string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVar(p, name, value, usage)
	return p
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func UintVar(p *uint, name string, value uint, usage string) {
	CommandLine.UintVar(p, name, value, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint variable that stores the value of the flag.
func Uint(name string, value uint, usage string) *uint {
	return CommandLine.Uint(name, value, usage)
}
