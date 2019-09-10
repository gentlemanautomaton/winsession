package infoclass

// Windows session information classes. Used by
// wtsapi.querySessionInformation.
const (
	InitialProgram     = 0  // WTSInitialProgram
	ApplicationName    = 1  // WTSApplicationName
	WorkingDirectory   = 2  // WTSWorkingDirectory
	OEMID              = 3  // WTSOEMId, not used
	SessionID          = 4  // WTSSessionId
	UserName           = 5  // WTSUserName
	StationName        = 6  // WTSWinStationName
	DomainName         = 7  // WTSDomainName
	ConnectionState    = 8  // WTSConnectState
	ClientBuildNumber  = 9  // WTSClientBuildNumber
	ClientName         = 10 // WTSClientName
	ClientDirectory    = 11 // WTSClientDirectory
	ClientProductID    = 12 // WTSClientProductId
	ClientHardwareID   = 13 // WTSClientHardwareId
	ClientAddress      = 14 // WTSClientAddress
	ClientDisplay      = 15 // WTSClientDisplay
	ClientProtocolType = 16 // WTSClientProtocolType
	IdleTime           = 17 // WTSIdleTime
	LogonTime          = 18 // WTSLogonTime
	IncomingBytes      = 19 // WTSIncomingBytes
	OutgoingBytes      = 20 // WTSOutgoingBytes
	IncomingFrames     = 21 // WTSIncomingFrames
	OutgoingFrames     = 22 // WTSOutgoingFrames
	ClientInfo         = 23 // WTSClientInfo
	SessionInfo        = 24 // WTSSessionInfo
	SessionInfoEx      = 25 // WTSSessionInfoEx
	ConfigInfo         = 26 // WTSConfigInfo
	ValidationInfo     = 27 // WTSValidationInfo
	SessionAddressV4   = 28 // WTSSessionAddressV4
	IsRemoveSession    = 29 // WTSIsRemoteSession
)
