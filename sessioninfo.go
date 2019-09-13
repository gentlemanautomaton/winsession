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

// User returns the session's user in the form DOMAIN\USER.
func (info SessionInfo) User() string {
	switch {
	case info.UserName == "":
		return ""
	case info.UserDomain == "":
		return info.UserName
	default:
		return info.UserDomain + `\` + info.UserName
	}
}

// IsZero returns true if the session info is unset.
func (info SessionInfo) IsZero() bool {
	return info == SessionInfo{}
}
