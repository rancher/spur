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

type name__Value struct {
	ref *Type__
	set bool
}

func newTitle__Value(val Type__, p *Type__) *name__Value {
	*p = val
	return &name__Value{p, false}
}

func (v *name__Value) init() {
	if v.ref == nil {
		v.ref = new(Type__)
	}
}

func (v *name__Value) New() interface{} { return *new(Type__) }

func (v *name__Value) Type() string { return "Type__" }

func (v *name__Value) Elem() string { return "Elem__" }

func (v *name__Value) IsSlice() bool { return IsSlice__ }

func (v *name__Value) IsSet() bool { return v.set }

func (v *name__Value) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *name__Value) Assign(value interface{}) {
	v.init()
	*v.ref = value.(Type__)
	v.set = true
}

func (v *name__Value) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *name__Value) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *name__Value) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]Elem__), elem.(Elem__))
	return slice
}

func (v *name__Value) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]Elem__)
	return val, ok
}

func (v *name__Value) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(Elem__)
	return val, ok
}

func (v *name__Value) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.(Type__))
	return string(r), err
}

func (v *name__Value) Deserialize(x string) (interface{}, error) {
	val := new(Type__)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *name__Value) String() string {
	return valueString(v)
}

// Title__Var defines a Type__ flag with specified name, default value, and usage string.
// The argument p points to a Type__ variable in which to store the value of the flag.
func (f *FlagSet) Title__Var(p *Type__, name string, value Type__, usage string) {
	f.Var(newTitle__Value(value, p), name, usage)
}

// Title__ defines a Type__ flag with specified name, default value, and usage string.
// The return value is the address of a Type__ variable that stores the value of the flag.
func (f *FlagSet) Title__(name string, value Type__, usage string) *Type__ {
	p := new(Type__)
	f.Title__Var(p, name, value, usage)
	return p
}

// Title__Var defines a Type__ flag with specified name, default value, and usage string.
// The argument p points to a Type__ variable in which to store the value of the flag.
func Title__Var(p *Type__, name string, value Type__, usage string) {
	CommandLine.Title__Var(p, name, value, usage)
}

// Title__ defines a Type__ flag with specified name, default value, and usage string.
// The return value is the address of a Type__ variable that stores the value of the flag.
func Title__(name string, value Type__, usage string) *Type__ {
	return CommandLine.Title__(name, value, usage)
}
