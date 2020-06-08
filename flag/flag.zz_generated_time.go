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

type timeValue struct {
	ref *time.Time
	set bool
}

func newTimeValue(val time.Time, p *time.Time) *timeValue {
	*p = val
	return &timeValue{p, false}
}

func (v *timeValue) init() {
	if v.ref == nil {
		v.ref = new(time.Time)
	}
}

func (v *timeValue) New() interface{} { return *new(time.Time) }

func (v *timeValue) Type() string { return "time.Time" }

func (v *timeValue) Elem() string { return "time.Time" }

func (v *timeValue) IsSlice() bool { return false }

func (v *timeValue) IsSet() bool { return v.set }

func (v *timeValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *timeValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.(time.Time)
	v.set = true
}

func (v *timeValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *timeValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *timeValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]time.Time), elem.(time.Time))
	return slice
}

func (v *timeValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]time.Time)
	return val, ok
}

func (v *timeValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(time.Time)
	return val, ok
}

func (v *timeValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.(time.Time))
	return string(r), err
}

func (v *timeValue) Deserialize(x string) (interface{}, error) {
	val := new(time.Time)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *timeValue) String() string {
	return valueString(v)
}

// TimeVar defines a time.Time flag with specified name, default value, and usage string.
// The argument p points to a time.Time variable in which to store the value of the flag.
func (f *FlagSet) TimeVar(p *time.Time, name string, value time.Time, usage string) {
	f.Var(newTimeValue(value, p), name, usage)
}

// Time defines a time.Time flag with specified name, default value, and usage string.
// The return value is the address of a time.Time variable that stores the value of the flag.
func (f *FlagSet) Time(name string, value time.Time, usage string) *time.Time {
	p := new(time.Time)
	f.TimeVar(p, name, value, usage)
	return p
}

// TimeVar defines a time.Time flag with specified name, default value, and usage string.
// The argument p points to a time.Time variable in which to store the value of the flag.
func TimeVar(p *time.Time, name string, value time.Time, usage string) {
	CommandLine.TimeVar(p, name, value, usage)
}

// Time defines a time.Time flag with specified name, default value, and usage string.
// The return value is the address of a time.Time variable that stores the value of the flag.
func Time(name string, value time.Time, usage string) *time.Time {
	return CommandLine.Time(name, value, usage)
}
