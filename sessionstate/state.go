package sessionstate

import "strconv"

// State is a windows session state.
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/ne-wtsapi32-wts_connectstate_class
type State uint32

// Windows session states.
const (
	Active       State = 0 // WTSActive
	Connected    State = 1 // WTSConnected
	Connecting   State = 2 // WTSConnectQuery
	Shadow       State = 3 // WTSShadow
	Disconnected State = 4 // WTSDisconnected
	Idle         State = 5 // WTSIdle
	Listen       State = 6 // WTSListen
	Reset        State = 7 // WTSReset
	Down         State = 8 // WTSDown
	Init         State = 9 // WTSInit
)

// String returns a string representation of the state.
func (s State) String() string {
	switch s {
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
		return "Unknown State " + strconv.Itoa(int(s))
	}
}
