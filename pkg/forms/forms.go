package forms

import (
	"strings"
)

type NewSnippet struct {
	Title    string
	Content  string
	Expires  string
	Failures map[string]string
}

func (ns *NewSnippet) Valid() bool {
	ns.Failures = make(map[string]string)

	if len(strings.TrimSpace(ns.Title)) == 0 {
		ns.Failures["Title"] = "Title is required"
	} else if len(strings.TrimSpace(ns.Title)) > 50 {
		ns.Failures["Title"] = "Title cannot be longer than 50 characters"
	}

	if len(strings.TrimSpace(ns.Content)) == 0 {
		ns.Failures["Content"] = "Content is required"
	}

	permitted := map[string]bool{"1209600": true, "604800": true, "86400": true}
	if len(strings.TrimSpace(ns.Expires)) == 0 {
		ns.Failures["Expires"] = "Expiry time is required"
	} else if !permitted[ns.Expires] {
		ns.Failures["Expires"] = "Expiry time must be 1209600, 604800 or 86400 seconds"
	}

	return len(ns.Failures) == 0
}
