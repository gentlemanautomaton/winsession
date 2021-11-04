// +build windows

package wtsapi

import (
	"fmt"
	"syscall"
	"unsafe"
	"math"

	"github.com/gentlemanautomaton/winsession/infoclass"
	"golang.org/x/sys/windows"
)

var (
	modwtsapi32 = windows.NewLazySystemDLL("wtsapi32.dll")

	procWTSOpenServerEx            = modwtsapi32.NewProc("WTSOpenServerExW")
	procWTSCloseServer             = modwtsapi32.NewProc("WTSCloseServer")
	procWTSFreeMemory              = modwtsapi32.NewProc("WTSFreeMemory")
	procWTSEnumerateSessions       = modwtsapi32.NewProc("WTSEnumerateSessionsW")
	procWTSQuerySessionInformation = modwtsapi32.NewProc("WTSQuerySessionInformationW")
	procWTSQueryUserToken          = modwtsapi32.NewProc("WTSQueryUserToken")
)

// OpenServer opens a connection to the windows terminal server with the given
// server name. It calls the WTSOpenServerExW windows API function.
//
// If name is empty it connects to the local terminal server instance.
//
// It is the caller's responsibility to close the returned handle when
// finished with it by calling CloseServer().
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/nf-wtsapi32-wtsopenserverexw
func OpenServer(name string) (server syscall.Handle, err error) {
	var sname *uint16
	if name != "" {
		sname, err = syscall.UTF16PtrFromString(name)
		if err != nil {
			return syscall.InvalidHandle, err
		}
	}

	r0, _, e := syscall.Syscall(
		procWTSOpenServerEx.Addr(),
		1,
		uintptr(unsafe.Pointer(sname)),
		0,
		0)

	server = syscall.Handle(r0)

	// Note: The API doesn't officially have a means of indicating failure.
	// Preliminary testing suggests that the API produces a non-zero handle
	// even when provided a bogus server name. This test might be fruitless.
	if server == syscall.InvalidHandle {
		if e != 0 {
			err = syscall.Errno(e)
		} else {
			err = syscall.EINVAL
		}
	}

	// TODO: Find a way to detect failure?

	return
}

// CloseServer closes a connection to windows terminal server.
// It calls the WTSCloseServer windows API function.
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/nf-wtsapi32-wtscloseserver
func CloseServer(server syscall.Handle) (err error) {
	_, _, e := syscall.Syscall(
		procWTSCloseServer.Addr(),
		1,
		uintptr(server),
		0,
		0)
	if e != 0 {
		return syscall.Errno(e)
	}
	// Note: The API doesn't officially have a means of indicating failure.
	return nil
}

// EnumerateSessions returns a slice of session information for the requested
// server. It calls the WTSEnumerateSessionsW windows API function.
//
// To efficiently query the local terminal server, specify Local when calling
// this function.
//
//https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/nf-wtsapi32-wtsenumeratesessionsw
func EnumerateSessions(server syscall.Handle) (sessions []SessionEntry, err error) {
	var ptr unsafe.Pointer
	var count uint32

	r0, _, e := syscall.Syscall6(
		procWTSEnumerateSessions.Addr(),
		5,
		uintptr(server),
		0,
		1,
		uintptr(unsafe.Pointer(&ptr)),
		uintptr(unsafe.Pointer(&count)),
		0)
	if r0 == 0 {
		return nil, syscall.Errno(e)
	}
	defer freeMemory(ptr)

	// Cast the pointer to an unbounded array and then take a slice of
	// suitable size from it
	list := ((*[math.MaxInt/unsafe.Sizeof(sessionEntry{})]sessionEntry)(ptr))[0:count:count]

	sessions = make([]SessionEntry, 0, count)
	for _, s := range list {
		sessions = append(sessions, SessionEntry{
			SessionID:     s.SessionID,
			WindowStation: utf16PointerToString(s.WindowStation),
			State:         s.State,
		})
	}

	return sessions, nil
}

// QueryUserName returns the name of the user for the given server and
// session ID.
func QueryUserName(server syscall.Handle, sessionID uint32) (userName string, err error) {
	var buf [512]byte
	data, err := QuerySessionData(server, sessionID, infoclass.UserName, buf[:])
	if err != nil {
		return "", err
	}
	return utf16BytesToString(data), nil
}

// QueryUserDomain returns the domain of the user for the given server and
// session ID.
func QueryUserDomain(server syscall.Handle, sessionID uint32) (userName string, err error) {
	var buf [512]byte
	data, err := QuerySessionData(server, sessionID, infoclass.DomainName, buf[:])
	if err != nil {
		return "", err
	}
	return utf16BytesToString(data), nil
}

// QueryClientInfo returns information about the client connected to the
// requested session.
func QueryClientInfo(server syscall.Handle, sessionID uint32) (ClientInfo, error) {
	var buf [clientInfoSize]byte
	data, err := QuerySessionData(server, sessionID, infoclass.ClientInfo, buf[:])
	if err != nil {
		return ClientInfo{}, err
	}

	raw := (*clientInfo)(unsafe.Pointer(&data[0]))
	return ClientInfo{
		ComputerName:            syscall.UTF16ToString(raw.ComputerName[:]),
		ComputerDomain:          syscall.UTF16ToString(raw.ComputerDomain[:]),
		UserName:                syscall.UTF16ToString(raw.UserName[:]),
		WorkDirectory:           syscall.UTF16ToString(raw.WorkDirectory[:]),
		InitialProgram:          syscall.UTF16ToString(raw.InitialProgram[:]),
		EncryptionLevel:         raw.EncryptionLevel,
		ClientAddressFamily:     raw.ClientAddressFamily,
		ClientAddress:           clientAddressToString(raw.ClientAddress[:], raw.ClientAddressFamily),
		HRes:                    raw.HRes,
		VRes:                    raw.VRes,
		ColorDepth:              raw.ColorDepth,
		ClientDirectory:         syscall.UTF16ToString(raw.ClientDirectory[:]),
		ClientBuildNumber:       raw.ClientBuildNumber,
		OutputBufferCountHost:   raw.OutputBufferCountHost,
		OutputBufferCountClient: raw.OutputBufferCountClient,
		OutputBufferLength:      raw.OutputBufferLength,
		DeviceID:                syscall.UTF16ToString(raw.DeviceID[:]),
	}, nil
}

// QuerySessionInfo returns detailed information for the requested session.
func QuerySessionInfo(server syscall.Handle, sessionID uint32) (SessionInfo, error) {
	var buf [sessionInfoLevel1Size]byte
	data, err := QuerySessionData(server, sessionID, infoclass.SessionInfoEx, buf[:])
	if err != nil {
		return SessionInfo{}, err
	}

	header := (*sessionInfoHeader)(unsafe.Pointer(&data[0]))

	switch header.Level {
	default:
		return SessionInfo{}, fmt.Errorf("unknown session info level: %d", header.Level)
	case 1:
		raw := (*sessionInfo)(unsafe.Pointer(&data[0]))
		return SessionInfo{
			SessionID:               raw.SessionID,
			ConnState:               raw.ConnState,
			LockState:               raw.LockState,
			WindowStation:           syscall.UTF16ToString(raw.WindowStation[:]),
			UserName:                syscall.UTF16ToString(raw.UserName[:]),
			UserDomain:              syscall.UTF16ToString(raw.UserDomain[:]),
			LogonTime:               timeFrom100NsecUTC(raw.LogonTime),
			ConnectTime:             timeFrom100NsecUTC(raw.ConnectTime),
			DisconnectTime:          timeFrom100NsecUTC(raw.DisconnectTime),
			LastInputTime:           timeFrom100NsecUTC(raw.LastInputTime),
			CurrentTime:             timeFrom100NsecUTC(raw.CurrentTime),
			IncomingBytes:           raw.IncomingBytes,
			OutgoingBytes:           raw.OutgoingBytes,
			IncomingFrames:          raw.IncomingFrames,
			OutgoingFrames:          raw.OutgoingFrames,
			IncomingCompressedBytes: raw.IncomingCompressedBytes,
			OutgoingCompressedBytes: raw.OutgoingCompressedBytes,
		}, nil
	}
}

// QuerySessionData returns raw session data for the requested server, session
// and information class. It calls the WTSQuerySessionInformationW
// windows API function.
//
// This is a low-level function that returns data as a slice of bytes. If a
// non-nil buffer of sufficient size is provided, the returned data will be
// sliced from the buffer.
//
// To efficiently query the local terminal server, specify Local when calling
// this function.
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/nf-wtsapi32-wtsquerysessioninformationw
func QuerySessionData(server syscall.Handle, sessionID, infoClass uint32, buffer []byte) (data []byte, err error) {
	var ptr unsafe.Pointer
	var size uint32

	r0, _, e := syscall.Syscall6(
		procWTSQuerySessionInformation.Addr(),
		5,
		uintptr(server),
		uintptr(sessionID),
		uintptr(infoClass),
		uintptr(unsafe.Pointer(&ptr)),
		uintptr(unsafe.Pointer(&size)),
		0)
	if r0 == 0 {
		return nil, syscall.Errno(e)
	}
	defer freeMemory(ptr)

	// Allocate data only if buffer is too small
	if int(size) <= len(buffer) {
		data = buffer[:size]
	} else {
		data = make([]byte, size)
	}

	// Cast the pointer to an unbounded array and then take a slice of
	// suitable size from it
	raw := ((*[math.MaxInt]byte)(ptr))[0:size:size]

	// Copy the data to Go-backed memory
	copy(data, raw)

	return data, nil
}

// QueryUserToken returns the primary access token of the user currently
// logged on to the given local session ID. It calls the WTSQueryUserToken
// windows API function.
//
// This function is highly restricted and will only succeed for applications
// running as LocalSystem.
//
// It is the caller's responsibility to close the token handle when finished
// with it by calling syscall.CloseHandle().
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/nf-wtsapi32-wtsqueryusertoken
func QueryUserToken(sessionID uint32) (token syscall.Token, err error) {
	r0, _, e := syscall.Syscall(
		procWTSQueryUserToken.Addr(),
		2,
		uintptr(sessionID),
		uintptr(unsafe.Pointer(&token)),
		0)
	if r0 == 0 {
		err = syscall.Errno(e)
	}
	return
}

// freeMemory releases memory allocated by previous WTS function calls.
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/nf-wtsapi32-wtsfreememory
func freeMemory(addr unsafe.Pointer) {
	syscall.Syscall(
		procWTSFreeMemory.Addr(),
		1,
		uintptr(addr),
		0,
		0)
}
