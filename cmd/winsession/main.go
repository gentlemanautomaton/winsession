package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/gentlemanautomaton/winsession/wtsapi"
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
	var server syscall.Handle
	if target == "" {
		server = wtsapi.Local
	} else {
		var err error
		server, err = wtsapi.OpenServer(target)
		if err != nil {
			fmt.Printf("Error: OpenServer: %v", err)
			return
		}
		defer wtsapi.CloseServer(server)
	}

	sessions, err := wtsapi.EnumerateSessions(server)
	if err != nil {
		fmt.Printf("Error: EnumerateSessions: %v\n", err)
		return
	}
	for _, session := range sessions {
		line := fmt.Sprintf("Session %d: %s: %s", session.ID, session.StationName, session.State)
		if userName, _ := wtsapi.QueryUserName(server, session.ID); userName != "" {
			if domainName, _ := wtsapi.QueryUserDomain(server, session.ID); domainName != "" {
				line = fmt.Sprintf("%s (%s\\%s)", line, domainName, userName)
			} else {
				line = fmt.Sprintf("%s (%s)", line, userName)
			}
		}
		fmt.Println(line)
	}
}
