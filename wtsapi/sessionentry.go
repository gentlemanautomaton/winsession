// +build windows

package wtsapi

import "github.com/gentlemanautomaton/winsession/connstate"

// SessionEntry holds data information about a session in windows terminal
// server. It contains data extracted from the WTS_SESSION_INFOW windows
// API structure.
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/ns-wtsapi32-wts_session_infow
type SessionEntry struct {
	SessionID     uint32
	WindowStation string
	State         connstate.Value
}

type sessionEntry struct {
	SessionID     uint32
	WindowStation *uint16
	State         connstate.Value
}
