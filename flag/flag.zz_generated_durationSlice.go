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

type durationSliceValue struct {
	ref *[]time.Duration
	set bool
}

func newDurationSliceValue(val []time.Duration, p *[]time.Duration) *durationSliceValue {
	*p = val
	return &durationSliceValue{p, false}
}

func (v *durationSliceValue) init() {
	if v.ref == nil {
		v.ref = new([]time.Duration)
	}
}

func (v *durationSliceValue) New() interface{} { return *new([]time.Duration) }

func (v *durationSliceValue) Type() string { return "[]time.Duration" }

func (v *durationSliceValue) Elem() string { return "time.Duration" }

func (v *durationSliceValue) IsSlice() bool { return true }

func (v *durationSliceValue) IsSet() bool { return v.set }

func (v *durationSliceValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *durationSliceValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.([]time.Duration)
	v.set = true
}

func (v *durationSliceValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *durationSliceValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *durationSliceValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]time.Duration), elem.(time.Duration))
	return slice
}

func (v *durationSliceValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]time.Duration)
	return val, ok
}

func (v *durationSliceValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(time.Duration)
	return val, ok
}

func (v *durationSliceValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.([]time.Duration))
	return string(r), err
}

func (v *durationSliceValue) Deserialize(x string) (interface{}, error) {
	val := new([]time.Duration)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *durationSliceValue) String() string {
	return valueString(v)
}

// DurationSliceVar defines a []time.Duration flag with specified name, default value, and usage string.
// The argument p points to a []time.Duration variable in which to store the value of the flag.
func (f *FlagSet) DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string) {
	f.Var(newDurationSliceValue(value, p), name, usage)
}

// DurationSlice defines a []time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a []time.Duration variable that stores the value of the flag.
func (f *FlagSet) DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration {
	p := new([]time.Duration)
	f.DurationSliceVar(p, name, value, usage)
	return p
}

// DurationSliceVar defines a []time.Duration flag with specified name, default value, and usage string.
// The argument p points to a []time.Duration variable in which to store the value of the flag.
func DurationSliceVar(p *[]time.Duration, name string, value []time.Duration, usage string) {
	CommandLine.DurationSliceVar(p, name, value, usage)
}

// DurationSlice defines a []time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a []time.Duration variable that stores the value of the flag.
func DurationSlice(name string, value []time.Duration, usage string) *[]time.Duration {
	return CommandLine.DurationSlice(name, value, usage)
}
