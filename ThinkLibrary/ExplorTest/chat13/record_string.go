// Code generated by "stringer -type=status -output=record_string.go"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[none-0]
	_ = x[res-1]
	_ = x[energy-2]
}

const _status_name = "noneresenergy"

var _status_index = [...]uint8{0, 4, 7, 13}

func (i status) String() string {
	if i < 0 || i >= status(len(_status_index)-1) {
		return "status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _status_name[_status_index[i]:_status_index[i+1]]
}
