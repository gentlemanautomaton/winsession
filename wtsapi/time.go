// +build windows

package wtsapi

import (
	"time"
)

// timeFrom100NsecUTC converts 100 nanosecond intervals since January 1
// 1601 into a time.Time value.
func timeFrom100NsecUTC(nsec int64) time.Time {
	// Let zero values continue to be zero values
	if nsec == 0 {
		return time.Time{}
	}

	// Change starting time to the Epoch (00:00:00 UTC, January 1, 1970)
	nsec -= 116444736000000000

	// Convert into nanoseconds
	nsec *= 100

	// Convert into time
	return time.Unix(0, nsec)
}
