//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package blockblob

import "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated"

// BlockListType defines values for BlockListType
type BlockListType = generated.BlockListType

const (
	BlockListTypeCommitted   BlockListType = "committed"
	BlockListTypeUncommitted BlockListType = "uncommitted"
	BlockListTypeAll         BlockListType = "all"
)

// PossibleBlockListTypeValues returns the possible values for the BlockListType const type.
func PossibleBlockListTypeValues() []BlockListType {
	return []BlockListType{
		BlockListTypeCommitted,
		BlockListTypeUncommitted,
		BlockListTypeAll,
	}
}
