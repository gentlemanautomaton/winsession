// +build windows

package winsession

import (
	"strings"

	"github.com/gentlemanautomaton/winsession/connstate"
)

// A StringMatcher is a function that matches strings
type StringMatcher func(string) bool

// MatchAny returns true if any of the filters match the session.
//
// MatchAny returns true if no filters are provided.
func MatchAny(filters ...Filter) Filter {
	return func(session Session) bool {
		if len(filters) == 0 {
			return true
		}
		for _, filter := range filters {
			if filter(session) {
				return true
			}
		}
		return false
	}
}

// MatchAll returns true if all of the filters match the session.
//
// MatchAll returns true if no filters are provided.
func MatchAll(filters ...Filter) Filter {
	return func(session Session) bool {
		for _, filter := range filters {
			if !filter(session) {
				return false
			}
		}
		return true
	}
}

// MatchID returns a filter that matches a session ID.
func MatchID(id ID) Filter {
	return func(session Session) bool {
		return session.ID == id
	}
}

// MatchState returns a filter that matches any of the given states.
func MatchState(states ...connstate.Value) Filter {
	return func(session Session) bool {
		for _, state := range states {
			if session.State == state {
				return true
			}
		}
		return false
	}
}

// MatchUser returns a filter that matches a session user name.
func MatchUser(matcher StringMatcher) Filter {
	return func(session Session) bool {
		name := session.Info.UserName
		if domain := session.Info.UserDomain; domain != "" {
			name = domain + `\` + name
		}
		return matcher(name)
	}
}

// EqualsUser returns a filter that matches a session user name case-insensitively.
func EqualsUser(name string) Filter {
	return MatchUser(func(candidate string) bool {
		return strings.EqualFold(candidate, name)
	})
}
