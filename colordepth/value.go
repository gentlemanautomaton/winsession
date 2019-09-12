package colordepth

import "strconv"

// Value identifies a color depth as used by windows terminal server.
type Value uint16

// Color depth values.
//
// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/ns-wtsapi32-wts_client_display
const (
	D4   = 1  // 4 bits per pixel
	D8   = 2  // 8 bits per pixel
	D16  = 4  // 16 bits per pixel
	D24A = 8  // 24 bits per pixel
	D15  = 16 // 15 bits per pixel
	D24B = 24 // 24 bits per pixel
	D32  = 32 // 32 bits per pixel
)

// String returns a string representation of the connection state.
func (v Value) String() string {
	switch v {
	case D4:
		return "4-bit RGB"
	case D8:
		return "8-bit RGB"
	case D16:
		return "16-bit RGB"
	case D24A:
		return "24-bit RGB"
	case D15:
		return "15-bit RGB"
	case D24B:
		return "24-bit RGB"
	case D32:
		return "32-bit RGB"
	default:
		return "Unknown Color Depth (" + strconv.Itoa(int(v)) + ")"
	}
}
