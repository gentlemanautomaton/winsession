// +build windows

package wtsapi

import (
	"time"
	"unsafe"

	"github.com/gentlemanautomaton/winsession/connstate"
	"github.com/gentlemanautomaton/winsession/lockstate"
)

const (
	sessionInfoHeaderSize = unsafe.Sizeof(sessionInfoHeader{})
	sessionInfoLevel1Size = unsafe.Sizeof(sessionInfoLevel1{})
)

// SessionInfo holds detailed information about a session in windows terminal
// server. It contains data extracted from the WTSINFOEXW windows API structure.
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/ns-wtsapi32-wtsinfoexw
type SessionInfo struct {
	SessionID               uint32
	ConnState               connstate.Value
	LockState               lockstate.Value
	WindowStation           string
	UserName                string
	UserDomain              string
	LogonTime               time.Time
	ConnectTime             time.Time
	DisconnectTime          time.Time
	LastInputTime           time.Time
	CurrentTime             time.Time
	IncomingBytes           uint32
	OutgoingBytes           uint32
	IncomingFrames          uint32
	OutgoingFrames          uint32
	IncomingCompressedBytes uint32
	OutgoingCompressedBytes uint32
}

// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/ns-wtsapi32-wtsinfoexw
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/ns-wtsapi32-wtsinfoex_level_w
type sessionInfo struct {
	sessionInfoHeader
	sessionInfoLevel1
}

type sessionInfoHeader struct {
	Level uint32
}

type sessionInfoLevel1 struct {
	SessionID               uint32
	ConnState               connstate.Value
	LockState               lockstate.Value
	WindowStation           [winStationNameLength + 1]uint16
	UserName                [userNameLength + 1]uint16
	UserDomain              [domainLength + 1]uint16
	LogonTime               int64
	ConnectTime             int64
	DisconnectTime          int64
	LastInputTime           int64
	CurrentTime             int64
	IncomingBytes           uint32
	OutgoingBytes           uint32
	IncomingFrames          uint32
	OutgoingFrames          uint32
	IncomingCompressedBytes uint32
	OutgoingCompressedBytes uint32
}
