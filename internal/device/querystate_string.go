// Code generated by "stringer -type QueryState -linecomment"; DO NOT EDIT.

package device

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[QueryNone-0]
	_ = x[QueryStarted-1]
	_ = x[QuerySuccess-2]
	_ = x[QueryFailed-3]
}

const _QueryState_name = "nonestartedsuccessfailed"

var _QueryState_index = [...]uint8{0, 4, 11, 18, 24}

func (i QueryState) String() string {
	if i >= QueryState(len(_QueryState_index)-1) {
		return "QueryState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _QueryState_name[_QueryState_index[i]:_QueryState_index[i+1]]
}