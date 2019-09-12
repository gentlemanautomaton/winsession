// +build windows

package winsession

// Collection holds interim processing information while collecting sessions.
type Collection struct {
	Host     *Host
	Sessions []Session
	Excluded []bool // Excluded[i] corresponds to Sessions[i]
}

// A CollectionOption is capable of applying its settings to a collection.
type CollectionOption interface {
	Apply(*Collection)
}

// A MergableCollectionOption is capable of merging with the next option in
// the list. Some options implement this interface when bulk operations
// can be performed more efficiently.
type MergableCollectionOption interface {
	Merge(next CollectionOption) (merged CollectionOption, ok bool)
}
