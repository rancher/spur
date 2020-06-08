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

type uintSliceValue struct {
	ref *[]uint
	set bool
}

func newUintSliceValue(val []uint, p *[]uint) *uintSliceValue {
	*p = val
	return &uintSliceValue{p, false}
}

func (v *uintSliceValue) init() {
	if v.ref == nil {
		v.ref = new([]uint)
	}
}

func (v *uintSliceValue) New() interface{} { return *new([]uint) }

func (v *uintSliceValue) Type() string { return "[]uint" }

func (v *uintSliceValue) Elem() string { return "uint" }

func (v *uintSliceValue) IsSlice() bool { return true }

func (v *uintSliceValue) IsSet() bool { return v.set }

func (v *uintSliceValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *uintSliceValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.([]uint)
	v.set = true
}

func (v *uintSliceValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *uintSliceValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *uintSliceValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]uint), elem.(uint))
	return slice
}

func (v *uintSliceValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]uint)
	return val, ok
}

func (v *uintSliceValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(uint)
	return val, ok
}

func (v *uintSliceValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.([]uint))
	return string(r), err
}

func (v *uintSliceValue) Deserialize(x string) (interface{}, error) {
	val := new([]uint)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *uintSliceValue) String() string {
	return valueString(v)
}

// UintSliceVar defines a []uint flag with specified name, default value, and usage string.
// The argument p points to a []uint variable in which to store the value of the flag.
func (f *FlagSet) UintSliceVar(p *[]uint, name string, value []uint, usage string) {
	f.Var(newUintSliceValue(value, p), name, usage)
}

// UintSlice defines a []uint flag with specified name, default value, and usage string.
// The return value is the address of a []uint variable that stores the value of the flag.
func (f *FlagSet) UintSlice(name string, value []uint, usage string) *[]uint {
	p := new([]uint)
	f.UintSliceVar(p, name, value, usage)
	return p
}

// UintSliceVar defines a []uint flag with specified name, default value, and usage string.
// The argument p points to a []uint variable in which to store the value of the flag.
func UintSliceVar(p *[]uint, name string, value []uint, usage string) {
	CommandLine.UintSliceVar(p, name, value, usage)
}

// UintSlice defines a []uint flag with specified name, default value, and usage string.
// The return value is the address of a []uint variable that stores the value of the flag.
func UintSlice(name string, value []uint, usage string) *[]uint {
	return CommandLine.UintSlice(name, value, usage)
}
