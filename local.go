// +build windows

package winsession

import "github.com/gentlemanautomaton/winsession/wtsapi"

var local = &Host{
	handle: wtsapi.Local,
}

// Local represents the local host.
var Local = local
