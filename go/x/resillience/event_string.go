// Code generated by "stringer -type Event -trimprefix=Event"; DO NOT EDIT.

package resillience

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EventNew-0]
	_ = x[EventClose-1]
	_ = x[EventExpired-2]
}

const _Event_name = "NewCloseExpired"

var _Event_index = [...]uint8{0, 3, 8, 15}

func (i Event) String() string {
	if i < 0 || i >= Event(len(_Event_index)-1) {
		return "Event(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Event_name[_Event_index[i]:_Event_index[i+1]]
}