package lockstate

import "strconv"

// Value identifies the lock state of a windows session.
type Value uint32

// Windows session lock states.
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/ns-wtsapi32-wtsinfoex_level1_w
const (
	Locked   Value = 0          // WTS_SESSIONSTATE_LOCK
	Unlocked Value = 1          // WTS_SESSIONSTATE_UNLOCK
	Unknown  Value = 4294967295 // WTS_SESSIONSTATE_UNKNOWN
)

// String returns a string representation of the lock state.
func (v Value) String() string {
	switch v {
	case Locked:
		return "Locked"
	case Unlocked:
		return "Unlocked"
	case Unknown:
		return "Lock State Unknown"
	default:
		return "Unknown Session Lock State (" + strconv.Itoa(int(v)) + ")"
	}
}
