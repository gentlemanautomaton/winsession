package connstate

import "strconv"

// Value identifies a windows session connection state.
type Value uint32

// Windows session connection states.
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/ne-wtsapi32-wts_connectstate_class
const (
	Active       Value = 0 // WTSActive
	Connected    Value = 1 // WTSConnected
	Connecting   Value = 2 // WTSConnectQuery
	Shadow       Value = 3 // WTSShadow
	Disconnected Value = 4 // WTSDisconnected
	Idle         Value = 5 // WTSIdle
	Listen       Value = 6 // WTSListen
	Reset        Value = 7 // WTSReset
	Down         Value = 8 // WTSDown
	Init         Value = 9 // WTSInit
)

// String returns a string representation of the connection state.
func (v Value) String() string {
	switch v {
	case Active:
		return "Active"
	case Connected:
		return "Connected"
	case Connecting:
		return "Connecting"
	case Shadow:
		return "Shadow"
	case Disconnected:
		return "Disconnected"
	case Idle:
		return "Idle"
	case Listen:
		return "Listen"
	case Reset:
		return "Reset"
	case Down:
		return "Down"
	case Init:
		return "Init"
	default:
		return "Unknown Session Connection State (" + strconv.Itoa(int(v)) + ")"
	}
}
