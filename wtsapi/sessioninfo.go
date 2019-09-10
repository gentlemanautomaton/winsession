// +build windows

package wtsapi

import "github.com/gentlemanautomaton/winsession/sessionstate"

// SessionInfo holds basic information about a session in windows terminal
// server.
type SessionInfo struct {
	ID          uint32
	StationName string
	State       sessionstate.State
}

type rawSessionInfo struct {
	ID          uint32
	StationName *uint16
	State       sessionstate.State
}
