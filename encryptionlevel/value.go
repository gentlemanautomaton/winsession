package encryptionlevel

import "strconv"

// Value identifies a remote desktop services encryption level.
type Value byte

// RDP encryption levels.
//
// https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpbcgr/f1c7c93b-94cc-4551-bb90-532a0185246a
const (
	None             Value = 0 // ENCRYPTION_LEVEL_NONE
	Low              Value = 1 // ENCRYPTION_LEVEL_LOW
	ClientCompatible Value = 2 // ENCRYPTION_LEVEL_CLIENT_COMPATIBLE
	High             Value = 3 // ENCRYPTION_LEVEL_HIGH
	FIPS             Value = 4 // ENCRYPTION_LEVEL_FIPS
)

// String returns a string representation of the encryption level.
func (v Value) String() string {
	switch v {
	case None:
		return "None"
	case Low:
		return "Low"
	case ClientCompatible:
		return "Client Compatible"
	case High:
		return "High"
	case FIPS:
		return "FIPS Compliant"
	default:
		return "Unknown Encryption Level (" + strconv.Itoa(int(v)) + ")"
	}
}
