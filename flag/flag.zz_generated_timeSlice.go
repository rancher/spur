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

type timeSliceValue struct {
	ref *[]time.Time
	set bool
}

func newTimeSliceValue(val []time.Time, p *[]time.Time) *timeSliceValue {
	*p = val
	return &timeSliceValue{p, false}
}

func (v *timeSliceValue) init() {
	if v.ref == nil {
		v.ref = new([]time.Time)
	}
}

func (v *timeSliceValue) New() interface{} { return *new([]time.Time) }

func (v *timeSliceValue) Type() string { return "[]time.Time" }

func (v *timeSliceValue) Elem() string { return "time.Time" }

func (v *timeSliceValue) IsSlice() bool { return true }

func (v *timeSliceValue) IsSet() bool { return v.set }

func (v *timeSliceValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *timeSliceValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.([]time.Time)
	v.set = true
}

func (v *timeSliceValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *timeSliceValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *timeSliceValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]time.Time), elem.(time.Time))
	return slice
}

func (v *timeSliceValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]time.Time)
	return val, ok
}

func (v *timeSliceValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(time.Time)
	return val, ok
}

func (v *timeSliceValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.([]time.Time))
	return string(r), err
}

func (v *timeSliceValue) Deserialize(x string) (interface{}, error) {
	val := new([]time.Time)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *timeSliceValue) String() string {
	return valueString(v)
}

// TimeSliceVar defines a []time.Time flag with specified name, default value, and usage string.
// The argument p points to a []time.Time variable in which to store the value of the flag.
func (f *FlagSet) TimeSliceVar(p *[]time.Time, name string, value []time.Time, usage string) {
	f.Var(newTimeSliceValue(value, p), name, usage)
}

// TimeSlice defines a []time.Time flag with specified name, default value, and usage string.
// The return value is the address of a []time.Time variable that stores the value of the flag.
func (f *FlagSet) TimeSlice(name string, value []time.Time, usage string) *[]time.Time {
	p := new([]time.Time)
	f.TimeSliceVar(p, name, value, usage)
	return p
}

// TimeSliceVar defines a []time.Time flag with specified name, default value, and usage string.
// The argument p points to a []time.Time variable in which to store the value of the flag.
func TimeSliceVar(p *[]time.Time, name string, value []time.Time, usage string) {
	CommandLine.TimeSliceVar(p, name, value, usage)
}

// TimeSlice defines a []time.Time flag with specified name, default value, and usage string.
// The return value is the address of a []time.Time variable that stores the value of the flag.
func TimeSlice(name string, value []time.Time, usage string) *[]time.Time {
	return CommandLine.TimeSlice(name, value, usage)
}
