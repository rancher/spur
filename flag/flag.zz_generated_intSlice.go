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

type intSliceValue struct {
	ref *[]int
	set bool
}

func newIntSliceValue(val []int, p *[]int) *intSliceValue {
	*p = val
	return &intSliceValue{p, false}
}

func (v *intSliceValue) init() {
	if v.ref == nil {
		v.ref = new([]int)
	}
}

func (v *intSliceValue) New() interface{} { return *new([]int) }

func (v *intSliceValue) Type() string { return "[]int" }

func (v *intSliceValue) Elem() string { return "int" }

func (v *intSliceValue) IsSlice() bool { return true }

func (v *intSliceValue) IsSet() bool { return v.set }

func (v *intSliceValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *intSliceValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.([]int)
	v.set = true
}

func (v *intSliceValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *intSliceValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *intSliceValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]int), elem.(int))
	return slice
}

func (v *intSliceValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]int)
	return val, ok
}

func (v *intSliceValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(int)
	return val, ok
}

func (v *intSliceValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.([]int))
	return string(r), err
}

func (v *intSliceValue) Deserialize(x string) (interface{}, error) {
	val := new([]int)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *intSliceValue) String() string {
	return valueString(v)
}

// IntSliceVar defines a []int flag with specified name, default value, and usage string.
// The argument p points to a []int variable in which to store the value of the flag.
func (f *FlagSet) IntSliceVar(p *[]int, name string, value []int, usage string) {
	f.Var(newIntSliceValue(value, p), name, usage)
}

// IntSlice defines a []int flag with specified name, default value, and usage string.
// The return value is the address of a []int variable that stores the value of the flag.
func (f *FlagSet) IntSlice(name string, value []int, usage string) *[]int {
	p := new([]int)
	f.IntSliceVar(p, name, value, usage)
	return p
}

// IntSliceVar defines a []int flag with specified name, default value, and usage string.
// The argument p points to a []int variable in which to store the value of the flag.
func IntSliceVar(p *[]int, name string, value []int, usage string) {
	CommandLine.IntSliceVar(p, name, value, usage)
}

// IntSlice defines a []int flag with specified name, default value, and usage string.
// The return value is the address of a []int variable that stores the value of the flag.
func IntSlice(name string, value []int, usage string) *[]int {
	return CommandLine.IntSlice(name, value, usage)
}
