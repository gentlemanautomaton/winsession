// +build windows

package winsession

import (
	"syscall"

	"github.com/gentlemanautomaton/winsession/wtsapi"
)

// Host is an open connection to a local or remote computer.
// It manages an open system handle internally.
//
// Each server must be closed when it is no longer needed.
type Host struct {
	handle syscall.Handle
}

// Open opens a connection to the server with the given address.
// If addr is empty a connection to the local server will be
// established.
func Open(addr string) (*Host, error) {
	if addr == "" {
		return local, nil
	}
	h, err := wtsapi.OpenServer(addr)
	if err != nil {
		return nil, err
	}
	return &Host{handle: h}, nil
}

// Handle returns the system handle for the connection to host.
func (host *Host) Handle() syscall.Handle {
	return host.handle
}

// Close closes any system handles held by host.
func (host *Host) Close() error {
	return wtsapi.CloseServer(host.handle)
}

// Sessions returns a list of windows sessions from host.
//
// Collection options can be provided to filter the list and collect
// additional session information. Options will be evaluated in order.
//
// If a filter relies on session information gathered by one or more
// collector options, those options must be included before the filter.
func (host *Host) Sessions(options ...CollectionOption) ([]Session, error) {
	// Enumerate all sessions on the host
	entries, err := wtsapi.EnumerateSessions(host.Handle())
	if err != nil {
		return nil, err
	}

	// Form a collection
	col := Collection{
		Host:     host,
		Sessions: make([]Session, len(entries)),
		Excluded: make([]bool, len(entries)),
	}

	// Copy the enumerated session data to the collection
	for i, entry := range entries {
		session := &col.Sessions[i]
		session.ID = ID(entry.SessionID)
		session.WindowStation = entry.WindowStation
		session.State = entry.State
	}

	// Apply each collection option in order
	for i := 0; i < len(options); i++ {
		opt := options[i]

		// Merge adjacent options when possible.
		//
		// This can improve efficiency by combining several options that work
		// more efficiently as a batch.
		for i+1 < len(options) {
			mergable, ok := opt.(MergableCollectionOption)
			if !ok {
				break
			}
			merged, ok := mergable.Merge(options[i+1])
			if !ok {
				break
			}
			opt = merged
			i++
		}

		opt.Apply(&col)
	}

	// Count the number of matches
	total := 0
	for i := range col.Excluded {
		if !col.Excluded[i] {
			total++
		}
	}

	// If all sessions matched just return the original slice
	if len(col.Sessions) == total {
		return col.Sessions, nil
	}

	// Return the matches
	matched := make([]Session, 0, total)
	for i := range col.Sessions {
		if col.Excluded[i] {
			continue
		}
		matched = append(matched, col.Sessions[i])
	}
	return matched, nil
}
