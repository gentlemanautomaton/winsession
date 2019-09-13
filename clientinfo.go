// +build windows

package winsession

import (
	"github.com/gentlemanautomaton/winsession/colordepth"
	"github.com/gentlemanautomaton/winsession/encryptionlevel"
)

// ClientInfo holds detailed information about a remote client connected to
// a windows session via RDP.
type ClientInfo struct {
	ComputerName    string
	ComputerDomain  string
	EncryptionLevel encryptionlevel.Value
	AddressFamily   uint32
	Address         string
	HRes            uint16 // Horizontal resolution (pixels)
	VRes            uint16 // Vertical resolution (pixels)
	ColorDepth      colordepth.Value
	BuildNumber     uint32
}

// Computer returns the client computer in the form DOMAIN\NAME.
func (info ClientInfo) Computer() string {
	switch {
	case info.ComputerName == "":
		return ""
	case info.ComputerDomain == "":
		return info.ComputerName
	default:
		return info.ComputerDomain + `\` + info.ComputerName
	}
}

// IsZero returns true if the client info is unset.
func (info ClientInfo) IsZero() bool {
	return info == ClientInfo{}
}
