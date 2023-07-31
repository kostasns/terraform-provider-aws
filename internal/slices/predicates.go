// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package slices

// PredicateEquals returns a Predicate that evaluates to true if the predicate's argument equals `v`.
func PredicateEquals[T comparable](v T) Predicate[T] {
	return func(x T) bool {
		return x == v
	}
}
