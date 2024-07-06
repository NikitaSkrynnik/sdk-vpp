// Copyright (c) 2020-2023 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ifindex allows storing interface_types.InterfaceIndex stored in per Connection.Id metadata
package ifindex

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/edwarnicke/genericsync"
	"github.com/networkservicemesh/govpp/binapi/interface_types"
)

var servermap genericsync.Map[string, interface_types.InterfaceIndex]
var clientmap genericsync.Map[string, interface_types.InterfaceIndex]

var print = func() {
	for {
		time.Sleep(time.Second * 10)

		fmt.Println("SERVER MAP")
		servermap.Range(func(key string, value interface_types.InterfaceIndex) bool {
			fmt.Printf("\tkey: %v, value: %v\n", key, value)
			return true
		})

		fmt.Println("CLIENT MAP")
		clientmap.Range(func(key string, value interface_types.InterfaceIndex) bool {
			fmt.Printf("\tkey: %v, value: %v\n", key, value)
			return true
		})
	}
}

var once sync.Once

type key struct{}

// Store sets the interface_types.InterfaceIndex stored in per Connection.Id metadata.
func Store(ctx context.Context, isClient bool, swIfIndex interface_types.InterfaceIndex) {
	//metadata.Map(ctx, isClient).Store(key{}, swIfIndex)

	once.Do(func() {
		go print()
	})

	id := connIDFromCtx(ctx)
	if id == "" {
		return
	}

	if isClient {
		clientmap.Store(id, swIfIndex)
		return
	}

	servermap.Store(id, swIfIndex)
}

// Delete deletes the interface_types.InterfaceIndex stored in per Connection.Id metadata
func Delete(ctx context.Context, isClient bool) {

	once.Do(func() {
		go print()
	})

	id := connIDFromCtx(ctx)
	if id == "" {
		return
	}

	if isClient {
		clientmap.Delete(id)
		return
	}

	servermap.Delete(id)
}

// Load returns the interface_types.InterfaceIndex stored in per Connection.Id metadata, or nil if no
// value is present.
// The ok result indicates whether value was found in the per Connection.Id metadata.
func Load(ctx context.Context, isClient bool) (value interface_types.InterfaceIndex, ok bool) {

	once.Do(func() {
		go print()
	})

	id := connIDFromCtx(ctx)
	if id == "" {
		return
	}

	if isClient {
		return clientmap.Load(id)
	}

	return servermap.Load(id)
}

// LoadOrStore returns the existing interface_types.InterfaceIndex stored in per Connection.Id metadata if present.
// Otherwise, it stores and returns the given nterface_types.InterfaceIndex.
// The loaded result is true if the value was loaded, false if stored.
func LoadOrStore(ctx context.Context, isClient bool, swIfIndex interface_types.InterfaceIndex) (value interface_types.InterfaceIndex, ok bool) {

	once.Do(func() {
		go print()
	})

	id := connIDFromCtx(ctx)
	if id == "" {
		return
	}

	if isClient {
		return clientmap.LoadOrStore(id, swIfIndex)
	}

	return servermap.LoadOrStore(id, swIfIndex)
}

// LoadAndDelete deletes the interface_types.InterfaceIndex stored in per Connection.Id metadata,
// returning the previous value if any. The loaded result reports whether the key was present.
func LoadAndDelete(ctx context.Context, isClient bool) (value interface_types.InterfaceIndex, ok bool) {

	once.Do(func() {
		go print()
	})

	id := connIDFromCtx(ctx)
	if id == "" {
		return
	}

	if isClient {
		return clientmap.LoadAndDelete(id)
	}

	return servermap.LoadAndDelete(id)
}

func WithConnID(ctx context.Context, connID string) context.Context {
	return context.WithValue(ctx, key{}, connID)
}

func connIDFromCtx(ctx context.Context) string {
	val, _ := ctx.Value(key{}).(string)
	return val
}
