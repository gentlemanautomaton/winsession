// +build windows

package wtsapi

import (
	"unsafe"

	"github.com/gentlemanautomaton/winsession/colordepth"
	"github.com/gentlemanautomaton/winsession/encryptionlevel"
)

const (
	clientInfoSize = unsafe.Sizeof(clientInfo{})
)

// ClientInfo holds information about a remote client with a session in
// windows terminal server. It contains data extracted from the WTSCLIENTW
// windows API structure.
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/ns-wtsapi32-wtsclientw
type ClientInfo struct {
	ComputerName            string
	ComputerDomain          string
	UserName                string
	WorkDirectory           string
	InitialProgram          string
	EncryptionLevel         encryptionlevel.Value
	ClientAddressFamily     uint32
	ClientAddress           string
	HRes                    uint16 // Horizontal resolution (pixels)
	VRes                    uint16 // Vertical resolution (pixels)
	ColorDepth              colordepth.Value
	ClientDirectory         string
	ClientBuildNumber       uint32
	OutputBufferCountHost   uint16
	OutputBufferCountClient uint16
	OutputBufferLength      uint16
	DeviceID                string
}

type clientInfo struct {
	ComputerName            [clientNameLength + 1]uint16
	ComputerDomain          [domainLength + 1]uint16
	UserName                [userNameLength + 1]uint16
	WorkDirectory           [maxPath + 1]uint16
	InitialProgram          [maxPath + 1]uint16
	EncryptionLevel         encryptionlevel.Value
	ClientAddressFamily     uint32
	ClientAddress           [clientAddressLength + 1]uint16
	HRes                    uint16
	VRes                    uint16
	ColorDepth              colordepth.Value
	ClientDirectory         [maxPath + 1]uint16
	ClientBuildNumber       uint32
	ClientHardwareID        uint32 // Reserved
	ClientProductID         uint16 // Reserved
	OutputBufferCountHost   uint16
	OutputBufferCountClient uint16
	OutputBufferLength      uint16
	DeviceID                [maxPath + 1]uint16
}
