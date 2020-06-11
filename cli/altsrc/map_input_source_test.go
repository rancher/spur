package altsrc

import (
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	inputSource := &MapInputSource{
		file: "test",
		valueMap: map[interface{}]interface{}{
			"thing_of_duration_type": time.Minute,
			"thing_of_string_type":   "1m",
			"thing_of_int_type":      1000,
		},
	}
	d, ok := inputSource.Get("thing_of_duration_type")
	expect(t, time.Minute, d)
	expect(t, true, ok)
	d, ok = inputSource.Get("thing_of_string_type")
	expect(t, "1m", d)
	expect(t, true, ok)
	d, ok = inputSource.Get("thing_of_int_type")
	expect(t, 1000, d)
	expect(t, true, ok)
	d, ok = inputSource.Get("thing_of_no_type")
	expect(t, nil, d)
	expect(t, false, ok)
}
