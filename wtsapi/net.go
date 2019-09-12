// +build windows

package wtsapi

import (
	"strconv"
	"strings"

	"golang.org/x/sys/windows"
)

func clientAddressToString(addr []uint16, family uint32) string {
	// TODO: Support IPV6 addresses
	switch family {
	case windows.AF_INET:
		var parts []string
		for _, value := range addr {
			if value == 0 {
				break
			}
			parts = append(parts, strconv.Itoa(int(value)))
		}
		return strings.Join(parts, ".")
	default:
		return ""
	}
}
