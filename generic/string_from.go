// Copyright 2020 Rancher Labs, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package generic

import (
	"strconv"
	"time"
)

func init() {
	FromStringMap["string"] = func(s string, ptr interface{}) error {
		Set(ptr, s)
		return nil
	}
	FromStringMap["bool"] = func(s string, ptr interface{}) error {
		if s == "" {
			s = "false"
		}
		v, err := strconv.ParseBool(s)
		Set(ptr, bool(v))
		return err
	}
	FromStringMap["int"] = func(s string, ptr interface{}) error {
		v, err := strconv.ParseInt(s, 0, strconv.IntSize)
		Set(ptr, int(v))
		return err
	}
	FromStringMap["int64"] = func(s string, ptr interface{}) error {
		v, err := strconv.ParseInt(s, 0, 64)
		Set(ptr, int64(v))
		return err
	}
	FromStringMap["uint"] = func(s string, ptr interface{}) error {
		v, err := strconv.ParseUint(s, 0, strconv.IntSize)
		Set(ptr, uint(v))
		return err
	}
	FromStringMap["uint64"] = func(s string, ptr interface{}) error {
		v, err := strconv.ParseUint(s, 0, 64)
		Set(ptr, uint64(v))
		return err
	}
	FromStringMap["float64"] = func(s string, ptr interface{}) error {
		v, err := strconv.ParseFloat(s, 64)
		Set(ptr, float64(v))
		return err
	}
	FromStringMap["time.Duration"] = func(s string, ptr interface{}) error {
		if v, err := time.ParseDuration(s); err == nil {
			Set(ptr, time.Duration(v))
			return nil
		}
		return errParse
	}
	FromStringMap["time.Time"] = func(s string, ptr interface{}) error {
		if v, err := strconv.ParseInt(s, 0, 64); err == nil {
			Set(ptr, time.Unix(v, 0))
			return nil
		}
		for _, layout := range TimeLayouts {
			if v, err := time.Parse(layout, s); err == nil {
				Set(ptr, time.Time(v))
				return nil
			}
		}
		return errParse
	}
}
