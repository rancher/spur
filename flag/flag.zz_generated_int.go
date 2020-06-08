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

type intValue struct {
	ref *int
	set bool
}

func newIntValue(val int, p *int) *intValue {
	*p = val
	return &intValue{p, false}
}

func (v *intValue) init() {
	if v.ref == nil {
		v.ref = new(int)
	}
}

func (v *intValue) New() interface{} { return *new(int) }

func (v *intValue) Type() string { return "int" }

func (v *intValue) Elem() string { return "int" }

func (v *intValue) IsSlice() bool { return false }

func (v *intValue) IsSet() bool { return v.set }

func (v *intValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *intValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.(int)
	v.set = true
}

func (v *intValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *intValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *intValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]int), elem.(int))
	return slice
}

func (v *intValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]int)
	return val, ok
}

func (v *intValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(int)
	return val, ok
}

func (v *intValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.(int))
	return string(r), err
}

func (v *intValue) Deserialize(x string) (interface{}, error) {
	val := new(int)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *intValue) String() string {
	return valueString(v)
}

// IntVar defines a int flag with specified name, default value, and usage string.
// The argument p points to a int variable in which to store the value of the flag.
func (f *FlagSet) IntVar(p *int, name string, value int, usage string) {
	f.Var(newIntValue(value, p), name, usage)
}

// Int defines a int flag with specified name, default value, and usage string.
// The return value is the address of a int variable that stores the value of the flag.
func (f *FlagSet) Int(name string, value int, usage string) *int {
	p := new(int)
	f.IntVar(p, name, value, usage)
	return p
}

// IntVar defines a int flag with specified name, default value, and usage string.
// The argument p points to a int variable in which to store the value of the flag.
func IntVar(p *int, name string, value int, usage string) {
	CommandLine.IntVar(p, name, value, usage)
}

// Int defines a int flag with specified name, default value, and usage string.
// The return value is the address of a int variable that stores the value of the flag.
func Int(name string, value int, usage string) *int {
	return CommandLine.Int(name, value, usage)
}
