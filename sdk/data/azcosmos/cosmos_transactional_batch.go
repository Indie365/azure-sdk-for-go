// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// TransactionalBatch is a batch of operations to be executed in a single transaction.
// See https://docs.microsoft.com/azure/cosmos-db/sql/transactional-batch
type TransactionalBatch struct {
	partitionKey PartitionKey
	container    *ContainerClient
	operations   []batchOperation
}

// CreateItem adds a create operation to the batch.
func (b *TransactionalBatch) CreateItem(item []byte, o *TransactionalBatchItemOptions) {
	b.operations = append(b.operations,
		batchOperationCreate{
			operationType: "Create",
			partitionKey:  b.partitionKey,
			resourceBody:  item})
}

// DeleteItem adds a delete operation to the batch.
func (b *TransactionalBatch) DeleteItem(itemId string, o *TransactionalBatchItemOptions) {
	b.operations = append(b.operations,
		batchOperationDelete{
			operationType: "Delete",
			partitionKey:  b.partitionKey,
			id:            itemId})
}

// DeleteItem adds a delete operation to the batch.
func (b *TransactionalBatch) ReplaceItem(itemId string, item []byte, o *TransactionalBatchItemOptions) {
	if o == nil {
		o = &TransactionalBatchItemOptions{}
	}
	b.operations = append(b.operations,
		batchOperationReplace{
			operationType: "Replace",
			partitionKey:  b.partitionKey,
			id:            itemId,
			resourceBody:  item,
			ifMatch:       o.IfMatchEtag})
}

// DeleteItem adds a delete operation to the batch.
func (b *TransactionalBatch) UpsertItem(item []byte, o *TransactionalBatchItemOptions) {
	if o == nil {
		o = &TransactionalBatchItemOptions{}
	}
	b.operations = append(b.operations,
		batchOperationUpsert{
			operationType: "Upsert",
			partitionKey:  b.partitionKey,
			resourceBody:  item,
			ifMatch:       o.IfMatchEtag})
}

// DeleteItem adds a delete operation to the batch.
func (b *TransactionalBatch) ReadItem(itemId string, o *TransactionalBatchItemOptions) {
	b.operations = append(b.operations,
		batchOperationRead{
			operationType: "Read",
			partitionKey:  b.partitionKey,
			id:            itemId})
}

// Execute executes the transactional batch.
func (b *TransactionalBatch) Execute(ctx context.Context, o *TransactionalBatchOptions) (TransactionalBatchResponse, error) {
	if len(b.operations) == 0 {
		return TransactionalBatchResponse{}, errors.New("no operations in batch")
	}

	if o == nil {
		o = &TransactionalBatchOptions{}
	}

	h := headerOptionsOverride{
		partitionKey: &b.partitionKey,
	}

	// If contentResponseOnWrite is not enabled at the client level the
	// service will not even send a batch response payload
	// Instead we should automatically enforce contentResponseOnWrite for all
	// batch requests whenever at least one of the item operations requires a content response (read operation)
	enableContentResponseOnWriteForReadOperations := true
	for _, op := range b.operations {
		if op.getOperationType() == operationTypeRead {
			h.enableContentResponseOnWrite = &enableContentResponseOnWriteForReadOperations
			break
		}
	}

	operationContext := pipelineRequestOptions{
		resourceType:          resourceTypeDocument,
		resourceAddress:       b.container.link,
		headerOptionsOverride: &h}

	path, err := generatePathForNameBased(resourceTypeDocument, operationContext.resourceAddress, true)
	if err != nil {
		return TransactionalBatchResponse{}, err
	}

	azResponse, err := b.container.database.client.sendBatchRequest(
		path,
		ctx,
		b.operations,
		operationContext,
		o,
		nil)
	if err != nil {
		return TransactionalBatchResponse{}, err
	}

	return newTransactionalBatchResponse(azResponse)
}

type batchOperation interface {
	getOperationType() operationType
}

type batchOperationCreate struct {
	operationType string
	partitionKey  PartitionKey
	resourceBody  []byte
}

func (b batchOperationCreate) getOperationType() operationType {
	return operationTypeCreate
}

// MarshalJSON implements the json.Marshaler interface
func (b batchOperationCreate) MarshalJSON() ([]byte, error) {
	partitionKeyAsString, err := b.partitionKey.toJsonString()
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBufferString("{")
	buffer.WriteString(fmt.Sprintf("\"operationType\":\"%s\"", b.operationType))
	buffer.WriteString(fmt.Sprintf(",\"partitionKey\":\"%s\"", partitionKeyAsString))
	buffer.WriteString(",\"resourceBody\":")
	buffer.Write(b.resourceBody)
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

type batchOperationDelete struct {
	operationType string
	id            string
	partitionKey  PartitionKey
}

func (b batchOperationDelete) getOperationType() operationType {
	return operationTypeDelete
}

// MarshalJSON implements the json.Marshaler interface
func (b batchOperationDelete) MarshalJSON() ([]byte, error) {
	partitionKeyAsString, err := b.partitionKey.toJsonString()
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBufferString("{")
	buffer.WriteString(fmt.Sprintf("\"operationType\":\"%s\"", b.operationType))
	buffer.WriteString(fmt.Sprintf(",\"partitionKey\":\"%s\"", partitionKeyAsString))
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

type batchOperationReplace struct {
	operationType string
	ifMatch       *azcore.ETag
	partitionKey  PartitionKey
	id            string
	resourceBody  []byte
}

func (b batchOperationReplace) getOperationType() operationType {
	return operationTypeReplace
}

// MarshalJSON implements the json.Marshaler interface
func (b batchOperationReplace) MarshalJSON() ([]byte, error) {
	partitionKeyAsString, err := b.partitionKey.toJsonString()
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBufferString("{")
	buffer.WriteString(fmt.Sprintf("\"operationType\":\"%s\"", b.operationType))
	if b.ifMatch != nil {
		buffer.WriteString(",\"ifMatch\":")
		etag, err := json.Marshal(b.ifMatch)
		if err != nil {
			return nil, err
		}
		buffer.Write(etag)
	}

	buffer.WriteString(fmt.Sprintf(",\"partitionKey\":\"%s\"", partitionKeyAsString))
	buffer.WriteString(fmt.Sprintf(",\"id\":\"%s\"", b.id))
	buffer.WriteString(",\"resourceBody\":")
	buffer.Write(b.resourceBody)
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

type batchOperationUpsert struct {
	operationType string
	ifMatch       *azcore.ETag
	partitionKey  PartitionKey
	resourceBody  []byte
}

func (b batchOperationUpsert) getOperationType() operationType {
	return operationTypeUpsert
}

// MarshalJSON implements the json.Marshaler interface
func (b batchOperationUpsert) MarshalJSON() ([]byte, error) {
	partitionKeyAsString, err := b.partitionKey.toJsonString()
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBufferString("{")
	buffer.WriteString(fmt.Sprintf("\"operationType\":\"%s\"", b.operationType))
	if b.ifMatch != nil {
		buffer.WriteString(",\"ifMatch\":")
		etag, err := json.Marshal(b.ifMatch)
		if err != nil {
			return nil, err
		}
		buffer.Write(etag)
	}

	buffer.WriteString(fmt.Sprintf(",\"partitionKey\":\"%s\"", partitionKeyAsString))
	buffer.WriteString(",\"resourceBody\":")
	buffer.Write(b.resourceBody)
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

type batchOperationRead struct {
	operationType string
	partitionKey  PartitionKey
	id            string
}

func (b batchOperationRead) getOperationType() operationType {
	return operationTypeRead
}

// MarshalJSON implements the json.Marshaler interface
func (b batchOperationRead) MarshalJSON() ([]byte, error) {
	partitionKeyAsString, err := b.partitionKey.toJsonString()
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBufferString("{")
	buffer.WriteString(fmt.Sprintf("\"operationType\":\"%s\"", b.operationType))
	buffer.WriteString(fmt.Sprintf(",\"partitionKey\":\"%s\"", partitionKeyAsString))
	buffer.WriteString(fmt.Sprintf(",\"id\":\"%s\"", b.id))
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}
