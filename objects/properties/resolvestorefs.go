// Copyright 2015-2020 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package properties

// ----------------------------------------------------------------------
// Define Types
// ----------------------------------------------------------------------

/*
ResolvesToRefsProperty -
*/
type ResolvesToRefsProperty struct {
	ResolvesToRefs []string `json:"resolves_to_refs,omitempty"`
}
