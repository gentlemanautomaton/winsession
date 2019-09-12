// +build windows

package winsession

import (
	"github.com/gentlemanautomaton/winsession/connstate"
)

// Session holds information about a windows session.
type Session struct {
	ID            ID
	WindowStation string
	State         connstate.Value
	Info          SessionInfo
	Client        ClientInfo
}
