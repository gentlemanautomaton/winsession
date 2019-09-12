// +build windows

package wtsapi

// Constants used internally by the wtsapi library.
const (
	winStationNameLength = 32  // WINSTATIONNAME_LENGTH
	domainLength         = 17  // DOMAIN_LENGTH
	userNameLength       = 20  // USERNAME_LENGTH
	clientNameLength     = 20  // CLIENTNAME_LENGTH
	clientAddressLength  = 30  // CLIENTADDRESS_LENGTH
	maxPath              = 260 // MAX_PATH
)
