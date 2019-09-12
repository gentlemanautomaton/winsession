// +build windows

package winsession_test

import (
	"fmt"

	"github.com/gentlemanautomaton/winsession"
)

func ExampleHost_Sessions() {
	sessions, err := winsession.Local.Sessions(
		winsession.Include(winsession.MatchID(0)),
		winsession.CollectSessionInfo)
	if err != nil {
		fmt.Printf("Failed to retrieve session list: %v\n", err)
		return
	}

	for _, session := range sessions {
		fmt.Printf("Session %d: %s (%s)\n", session.ID, session.WindowStation, session.Info.LockState)
	}

	// Output:
	// Session 0: Services (Unlocked)
}
