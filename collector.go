// +build windows

package winsession

import (
	"sync"

	"github.com/gentlemanautomaton/winsession/wtsapi"
)

// A Collector is a collection option that collects additional information
// about a session.
//
// Information will only be collected for sessions that have not been
// excluded by previous filtering options.
type Collector int

const (
	// CollectSessionInfo enables collection of detailed information about
	// each session.
	CollectSessionInfo Collector = 1 << iota

	// CollectClientInfo enables collection of detailed information about
	// remote desktop protocol clients.
	CollectClientInfo
)

// Contains returns true if c contains b.
func (c Collector) Contains(b Collector) bool {
	return c&b == b
}

// Apply applies the collector to the collection.
func (c Collector) Apply(col *Collection) {
	if c == 0 {
		return
	}
	server := col.Host.Handle()

	var wg sync.WaitGroup
	wg.Add(len(col.Sessions))

	for i := range col.Sessions {
		if col.Excluded[i] {
			wg.Done()
			continue
		}
		go func(i int) {
			defer wg.Done()
			session := &col.Sessions[i]

			if c.Contains(CollectSessionInfo) {
				if info, err := wtsapi.QuerySessionInfo(server, uint32(session.ID)); err == nil {
					session.Info = SessionInfo{
						LockState:      info.LockState,
						WindowStation:  info.WindowStation,
						UserName:       info.UserName,
						UserDomain:     info.UserDomain,
						LogonTime:      info.LogonTime,
						ConnectTime:    info.ConnectTime,
						DisconnectTime: info.DisconnectTime,
						LastInputTime:  info.LastInputTime,
						CurrentTime:    info.CurrentTime,
					}
				}
			}

			if c.Contains(CollectClientInfo) {
				if info, err := wtsapi.QueryClientInfo(server, uint32(session.ID)); err == nil {
					session.Client = ClientInfo{
						ComputerName:    info.ComputerName,
						ComputerDomain:  info.ComputerDomain,
						EncryptionLevel: info.EncryptionLevel,
						AddressFamily:   info.ClientAddressFamily,
						Address:         info.ClientAddress,
						HRes:            info.HRes,
						VRes:            info.VRes,
						ColorDepth:      info.ColorDepth,
						BuildNumber:     info.ClientBuildNumber,
					}
				}
			}
		}(i)
	}

	wg.Wait()
}

// Merge attempts to merge the collector with the next option. It returns true
// if successful.
func (c Collector) Merge(next CollectionOption) (merged CollectionOption, ok bool) {
	n, ok := next.(Collector)
	if !ok {
		return nil, false
	}
	return c | n, true
}
