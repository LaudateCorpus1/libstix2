// Copyright 2017 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package resources

// ----------------------------------------------------------------------
// Public Create Functions
// ----------------------------------------------------------------------

// NewAPIRoot - This function will create a new TAXII API Root object.
func NewAPIRoot() APIRootType {
	var obj APIRootType
	return obj
}

// NewCollection - This function will create a new TAXII Collection object.
func NewCollection() CollectionType {
	var obj CollectionType
	return obj
}

// NewCollections - This function will create a new TAXII Collections object.
func NewCollections() CollectionsType {
	var obj CollectionsType
	return obj
}

// NewDiscovery - This function will create a new TAXII Discovery object.
func NewDiscovery() DiscoveryType {
	var obj DiscoveryType
	return obj
}

// NewManifest - This function will create a new TAXII Manifest object.
// func NewManifest() ManifestType {
// 	return manifest.New()
// }

// NewStatus - This function will create a new TAXII Status object.
// func NewStatus() StatusType {
// 	return status.New()
// }
