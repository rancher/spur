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

type uint64SliceValue struct {
	ref *[]uint64
	set bool
}

func newUint64SliceValue(val []uint64, p *[]uint64) *uint64SliceValue {
	*p = val
	return &uint64SliceValue{p, false}
}

func (v *uint64SliceValue) init() {
	if v.ref == nil {
		v.ref = new([]uint64)
	}
}

func (v *uint64SliceValue) New() interface{} { return *new([]uint64) }

func (v *uint64SliceValue) Type() string { return "[]uint64" }

func (v *uint64SliceValue) Elem() string { return "uint64" }

func (v *uint64SliceValue) IsSlice() bool { return true }

func (v *uint64SliceValue) IsSet() bool { return v.set }

func (v *uint64SliceValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *uint64SliceValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.([]uint64)
	v.set = true
}

func (v *uint64SliceValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *uint64SliceValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *uint64SliceValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]uint64), elem.(uint64))
	return slice
}

func (v *uint64SliceValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]uint64)
	return val, ok
}

func (v *uint64SliceValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(uint64)
	return val, ok
}

func (v *uint64SliceValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.([]uint64))
	return string(r), err
}

func (v *uint64SliceValue) Deserialize(x string) (interface{}, error) {
	val := new([]uint64)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *uint64SliceValue) String() string {
	return valueString(v)
}

// Uint64SliceVar defines a []uint64 flag with specified name, default value, and usage string.
// The argument p points to a []uint64 variable in which to store the value of the flag.
func (f *FlagSet) Uint64SliceVar(p *[]uint64, name string, value []uint64, usage string) {
	f.Var(newUint64SliceValue(value, p), name, usage)
}

// Uint64Slice defines a []uint64 flag with specified name, default value, and usage string.
// The return value is the address of a []uint64 variable that stores the value of the flag.
func (f *FlagSet) Uint64Slice(name string, value []uint64, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64SliceVar(p, name, value, usage)
	return p
}

// Uint64SliceVar defines a []uint64 flag with specified name, default value, and usage string.
// The argument p points to a []uint64 variable in which to store the value of the flag.
func Uint64SliceVar(p *[]uint64, name string, value []uint64, usage string) {
	CommandLine.Uint64SliceVar(p, name, value, usage)
}

// Uint64Slice defines a []uint64 flag with specified name, default value, and usage string.
// The return value is the address of a []uint64 variable that stores the value of the flag.
func Uint64Slice(name string, value []uint64, usage string) *[]uint64 {
	return CommandLine.Uint64Slice(name, value, usage)
}
