// +build windows

package winsession

// A Filter returns true if it matches a session.
type Filter func(Session) bool

// Include is an inclusion filter.
type Include Filter

// Apply applies the inclusion filter to the collection.
func (include Include) Apply(col *Collection) {
	if include == nil {
		return
	}

	for i := range col.Sessions {
		if !col.Excluded[i] {
			col.Excluded[i] = !include(col.Sessions[i])
		}
	}
}

// Exclude is an exclusion filter.
type Exclude Filter

// Apply applies the exclusion filter to the collection.
func (exclude Exclude) Apply(col *Collection) {
	if exclude == nil {
		return
	}

	for i := range col.Sessions {
		if !col.Excluded[i] {
			col.Excluded[i] = exclude(col.Sessions[i])
		}
	}
}
