// +build windows

package wtsapi

import (
	"unicode/utf16"
	"unsafe"
)

func utf16PointerToString(ptr *uint16) string {
	if ptr == nil {
		return ""
	}

	s := ((*[0xffff]uint16)(unsafe.Pointer(ptr)))[0:]
	for i, v := range s {
		if v == 0 {
			s = s[0:i:i]
			break
		}
	}

	return string(utf16.Decode(s))
}
