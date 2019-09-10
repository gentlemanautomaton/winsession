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
	var srv syscall.Handle
	if target == "" {
		srv = wtsapi.Local
	} else {
		var err error
		srv, err = wtsapi.OpenServer(target)
		if err != nil {
			fmt.Printf("Error: OpenServer: %v", err)
			return
		}
		defer wtsapi.CloseServer(srv)
	}

	sessions, err := wtsapi.EnumerateSessions(srv)
	if err != nil {
		fmt.Printf("Error: EnumerateSessions: %v\n", err)
		return
	}
	for _, session := range sessions {
		fmt.Printf("Session %d: %s: %s\n", session.ID, session.StationName, session.State)
	}
}
