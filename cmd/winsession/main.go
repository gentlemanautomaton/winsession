package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gentlemanautomaton/winsession"
	"github.com/gentlemanautomaton/winsession/connstate"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage:\n\t%s [<server>...]\n", os.Args[0])
	}
	flag.Parse()
	if flag.NArg() == 0 {
		if hostname, err := os.Hostname(); err == nil {
			fmt.Printf("--------\n%s\n--------\n", hostname)
		}
		list("")
	} else {
		for _, target := range flag.Args() {
			fmt.Printf("--------\n%s\n--------\n", target)
			list(target)
		}
	}
}

func list(target string) {
	host, err := winsession.Open(target)
	if err != nil {
		fmt.Printf("Unable to connect to host: %v", err)
		return
	}
	defer host.Close()

	sessions, err := host.Sessions(winsession.CollectSessionInfo, winsession.CollectClientInfo)
	if err != nil {
		fmt.Printf("Failed to enumerate sessions on %s: %v\n", target, err)
		return
	}

	for _, session := range sessions {
		// Session details
		if info := session.Info; !info.IsZero() {
			// Header
			line := fmt.Sprintf("Session %d", session.ID)
			if info.WindowStation != "" {
				line = fmt.Sprintf("%s: %s", line, info.WindowStation)
			}
			line = fmt.Sprintf("%s: %s (%s)", line, session.State, info.LockState)
			if user := info.User(); user != "" {
				line = fmt.Sprintf("%s: %s", line, user)
			}
			fmt.Println(line)

			// Timing information
			var now time.Time
			if !info.CurrentTime.IsZero() {
				now = info.CurrentTime
			} else {
				now = time.Now()
			}
			if !info.CurrentTime.IsZero() {
				fmt.Printf("  Current Time:    %v (%s)\n", info.CurrentTime, makeDuration(info.CurrentTime, now))
			}
			if !info.LogonTime.IsZero() {
				fmt.Printf("  Logon Time:      %v (%s)\n", info.LogonTime, makeDuration(info.LogonTime, now))
			}
			if !info.ConnectTime.IsZero() {
				fmt.Printf("  Connect Time:    %v (%s)\n", info.ConnectTime, makeDuration(info.ConnectTime, now))
			}
			switch session.State {
			case connstate.Active:
				if !info.LastInputTime.IsZero() {
					fmt.Printf("  Last Input Time: %v (%s)\n", info.LastInputTime, makeDuration(info.LastInputTime, now))
				}
			default:
				if !info.DisconnectTime.IsZero() {
					fmt.Printf("  Disconnect Time: %v (%s)\n", info.DisconnectTime, makeDuration(info.DisconnectTime, now))
				}
			}
		} else {
			// Header
			fmt.Printf("Session %d: %s: %s\n", session.ID, session.WindowStation, session.State)
		}

		// RDP details
		if client := session.Client; !client.IsZero() {
			if computer := client.Computer(); computer != "" {
				if addr := client.Address; addr != "" {
					computer = fmt.Sprintf("%s <%s>", computer, addr)
				}
				line := fmt.Sprintf("  RDP Client:      %s", computer)
				var aspects []string
				if client.HRes != 0 && client.VRes != 0 {
					aspects = append(aspects, fmt.Sprintf("%dx%d", client.HRes, client.VRes))
					aspects = append(aspects, client.ColorDepth.String())
				}
				if client.EncryptionLevel != 0 {
					aspects = append(aspects, fmt.Sprintf("Encryption Level %s", client.EncryptionLevel))
				}
				if len(aspects) > 0 {
					line = fmt.Sprintf("%s (%s)", line, strings.Join(aspects, ", "))
				}
				fmt.Println(line)
			}
		}
	}
}

func makeDuration(when, now time.Time) time.Duration {
	// There are no durations in the future
	if when.After(now) {
		return 0
	}

	// Subtract the time from now
	d := now.Sub(when)

	// Round to seconds
	d /= time.Second
	d *= time.Second

	return d
}
