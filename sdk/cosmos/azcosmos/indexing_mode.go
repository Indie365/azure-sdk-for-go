// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

// IndexingMode defines the supported indexing modes in the Azure Cosmos DB service.
type IndexingMode int

const (
	// IndexingModeConsistent Index is updated synchronously with a create, update or delete operation.
	IndexingModeConsistent IndexingMode = 0
	// No index is provided.
	IndexingModeNone IndexingMode = 2
)

// Returns a list of available consistency levels
func IndexingModeValues() []IndexingMode {
	return []IndexingMode{IndexingModeConsistent, IndexingModeNone}
}

func (c IndexingMode) ToPtr() *IndexingMode {
	return &c
}
