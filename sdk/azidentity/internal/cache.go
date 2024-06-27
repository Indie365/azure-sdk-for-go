// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package internal

import (
	"sync"

	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/cache"
)

// Cache represents a persistent token cache
type Cache struct {
	// impl is a pointer so a Cache can carry persistent state across copies
	impl *impl
}

// impl is a Cache's private implementation
type impl struct {
	// factory constructs implementations
	factory func(bool) (cache.ExportReplace, error)
	// cae and noCAE are previously constructed implementations
	cae, noCAE cache.ExportReplace
	// l synchronizes around cae and noCAE
	l *sync.RWMutex
}

func (i *impl) exportReplace(cae bool) (cache.ExportReplace, error) {
	if i == nil {
		// zero-value Cache: return a nil ExportReplace and MSAL will cache in memory
		return nil, nil
	}
	var (
		err error
		xr  cache.ExportReplace
	)
	i.l.RLock()
	xr = i.cae
	if !cae {
		xr = i.noCAE
	}
	i.l.RUnlock()
	if xr != nil {
		return xr, nil
	}
	i.l.Lock()
	defer i.l.Unlock()
	// double-check because another goroutine may have constructed the cache while this one locked
	if cae {
		if i.cae == nil {
			if xr, err = i.factory(cae); err == nil {
				i.cae = xr
			}
		}
		return i.cae, err
	}
	if i.noCAE == nil {
		if xr, err = i.factory(cae); err == nil {
			i.noCAE = xr
		}
	}
	return i.noCAE, err
}

// NewCache is the constructor for Cache. It takes a factory instead of an instance
// because it doesn't know whether the Cache will store both CAE and non-CAE tokens
// (these must be stored separately).
func NewCache(factory func(cae bool) (cache.ExportReplace, error)) Cache {
	return Cache{&impl{factory: factory, l: &sync.RWMutex{}}}
}

// ExportReplace returns an implementation satisfying MSAL's ExportReplace interface.
// It's an internal function instead of an exported method so packages in azidentity
// and azidentity/cache can call it without exporting it to applications. "cae"
// controls whether the desired implementation will store CAE tokens.
func ExportReplace(c Cache, cae bool) (cache.ExportReplace, error) {
	return c.impl.exportReplace(cae)
}
