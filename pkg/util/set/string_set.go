// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package set

import (
	"strings"
)

// StringSet is a string set.
type StringSet map[string]struct{}

// NewStringSet builds a string set.
func NewStringSet(ss ...string) StringSet {
	set := make(StringSet, len(ss))
	for _, s := range ss {
		set.Insert(s)
	}
	return set
}

// Exist checks whether `val` exists in `s`.
func (s StringSet) Exist(val string) bool {
	_, ok := s[val]
	return ok
}

// Insert inserts `val` into `s`.
func (s StringSet) Insert(val string) {
	s[val] = struct{}{}
}

// Intersection returns the intersection of two sets
func (s StringSet) Intersection(rhs StringSet) StringSet {
	newSet := NewStringSet()
	for elt := range s {
		if rhs.Exist(elt) {
			newSet.Insert(elt)
		}
	}
	return newSet
}

// IntersectionWithLower returns the intersection of two sets with different case of string.
func (s StringSet) IntersectionWithLower(rhs StringSet, toLower bool) StringSet {
	newSet := NewStringSet()
	for origElt := range rhs {
		var elt string
		if toLower {
			elt = strings.ToLower(origElt)
		} else {
			elt = strings.ToUpper(origElt)
		}
		if s.Exist(elt) {
			newSet.Insert(origElt)
		}
	}
	return newSet
}

// Count returns the number in Set s.
func (s StringSet) Count() int {
	return len(s)
}

// Empty returns whether s is empty.
func (s StringSet) Empty() bool {
	return len(s) == 0
}

// Clear clears the set.
func (s StringSet) Clear() {
	clear(s)
}

// IterateWith iterate items in StringSet and pass it to `fn`.
func (s StringSet) IterateWith(fn func(string)) {
	for k := range s {
		fn(k)
	}
}
