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

type durationValue struct {
	ref *time.Duration
	set bool
}

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
	*p = val
	return &durationValue{p, false}
}

func (v *durationValue) init() {
	if v.ref == nil {
		v.ref = new(time.Duration)
	}
}

func (v *durationValue) New() interface{} { return *new(time.Duration) }

func (v *durationValue) Type() string { return "time.Duration" }

func (v *durationValue) Elem() string { return "time.Duration" }

func (v *durationValue) IsSlice() bool { return false }

func (v *durationValue) IsSet() bool { return v.set }

func (v *durationValue) IsBoolFlag() bool { return v.Elem() == "bool" }

func (v *durationValue) Assign(value interface{}) {
	v.init()
	*v.ref = value.(time.Duration)
	v.set = true
}

func (v *durationValue) Get() interface{} {
	v.init()
	return *v.ref
}

func (v *durationValue) Set(value interface{}) error {
	return valueSet(v, value)
}

func (v *durationValue) Append(slice interface{}, elem interface{}) interface{} {
	slice = append(slice.([]time.Duration), elem.(time.Duration))
	return slice
}

func (v *durationValue) ConvertSlice(value interface{}) (interface{}, bool) {
	val, ok := value.([]time.Duration)
	return val, ok
}

func (v *durationValue) ConvertElem(value interface{}) (interface{}, bool) {
	val, ok := value.(time.Duration)
	return val, ok
}

func (v *durationValue) Serialize(value interface{}) (string, error) {
	r, err := json.Marshal(value.(time.Duration))
	return string(r), err
}

func (v *durationValue) Deserialize(x string) (interface{}, error) {
	val := new(time.Duration)
	err := yaml.Unmarshal([]byte(x), val)
	return *val, err
}

func (v *durationValue) String() string {
	return valueString(v)
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.Var(newDurationValue(value, p), name, usage)
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, value, usage)
	return p
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	CommandLine.DurationVar(p, name, value, usage)
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
func Duration(name string, value time.Duration, usage string) *time.Duration {
	return CommandLine.Duration(name, value, usage)
}
