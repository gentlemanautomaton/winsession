// +build windows

package winsession

import (
	"time"

	"github.com/gentlemanautomaton/winsession/lockstate"
)

// SessionInfo holds detailed information about a windows session.
type SessionInfo struct {
	LockState      lockstate.Value
	WindowStation  string
	UserName       string
	UserDomain     string
	LogonTime      time.Time
	ConnectTime    time.Time
	DisconnectTime time.Time
	LastInputTime  time.Time
	CurrentTime    time.Time
}

// IsZero returns true if the session info is unset.
func (info SessionInfo) IsZero() bool {
	return info == SessionInfo{}
}
