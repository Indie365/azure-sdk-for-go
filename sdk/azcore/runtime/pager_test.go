//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/pipeline"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/mock"
	"github.com/stretchr/testify/require"
)

type PageResponse struct {
	Values   []int `json:"values"`
	NextPage bool  `json:"next"`
}

func TestPagerSinglePage(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.AppendResponse(mock.WithStatusCode(http.StatusOK), mock.WithBody([]byte(`{"values": [1, 2, 3, 4, 5]}`)))
	pl := pipeline.NewPipeline(srv, NewRetryPolicy(testRetryOptions()))

	pager := NewPager(PageProcessor[PageResponse]{
		Do: pl.Do,
		More: func(current PageResponse) bool {
			return current.NextPage
		},
		Fetcher: func(ctx context.Context, current *PageResponse) (*policy.Request, error) {
			return NewRequest(ctx, http.MethodGet, srv.URL())
		},
		Responder: func(resp *http.Response) (PageResponse, error) {
			pr := PageResponse{}
			if err := UnmarshalAsJSON(resp, &pr); err != nil {
				return PageResponse{}, err
			}
			return pr, nil
		},
	})
	require.True(t, pager.firstPage)

	pageCount := 0
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		require.Equal(t, []int{1, 2, 3, 4, 5}, page.Values)
		require.Empty(t, page.NextPage)
		pageCount++
	}
	require.Equal(t, 1, pageCount)
	page, err := pager.NextPage(context.Background())
	require.Error(t, err)
	require.Empty(t, page)
}

func TestPagerMultiplePages(t *testing.T) {
	srv, close := mock.NewServer()
	defer close()
	srv.AppendResponse(mock.WithStatusCode(http.StatusOK), mock.WithBody([]byte(`{"values": [1, 2, 3, 4, 5], "next": true}`)))
	srv.AppendResponse(mock.WithStatusCode(http.StatusOK), mock.WithBody([]byte(`{"values": [6, 7, 8], "next": true}`)))
	srv.AppendResponse(mock.WithStatusCode(http.StatusOK), mock.WithBody([]byte(`{"values": [9, 0, 1, 2]}`)))
	pl := pipeline.NewPipeline(srv, NewRetryPolicy(testRetryOptions()))

	pageCount := 0
	pager := NewPager(PageProcessor[PageResponse]{
		Do: pl.Do,
		More: func(current PageResponse) bool {
			return current.NextPage
		},
		Fetcher: func(ctx context.Context, current *PageResponse) (*policy.Request, error) {
			if pageCount == 1 {
				require.Nil(t, current)
			} else {
				require.NotNil(t, current)
			}
			return NewRequest(ctx, http.MethodGet, srv.URL())
		},
		Responder: func(resp *http.Response) (PageResponse, error) {
			pr := PageResponse{}
			if err := UnmarshalAsJSON(resp, &pr); err != nil {
				return PageResponse{}, err
			}
			return pr, nil
		},
	})
	require.True(t, pager.firstPage)

	for pager.More() {
		pageCount++
		page, err := pager.NextPage(context.Background())
		require.NoError(t, err)
		switch pageCount {
		case 1:
			require.Equal(t, []int{1, 2, 3, 4, 5}, page.Values)
			require.True(t, page.NextPage)
		case 2:
			require.Equal(t, []int{6, 7, 8}, page.Values)
			require.True(t, page.NextPage)
		case 3:
			require.Equal(t, []int{9, 0, 1, 2}, page.Values)
			require.False(t, page.NextPage)
		}
	}
	require.Equal(t, 3, pageCount)
	page, err := pager.NextPage(context.Background())
	require.Error(t, err)
	require.Empty(t, page)
}
