//: Copyright Verizon Media
//: Licensed under the terms of the Apache 2.0 License. See LICENSE file in the project root for terms.

package vssh

import (
	"sync"
)

var clientsShardNum = 10

type clients []*clientsShard
type clientsShard struct {
	clients map[string]*clientAttr
	sync.RWMutex
}

func (c *clients) getShard(key string) uint { _ = "STUB: not implemented"; return 0 }

func newClients() clients { _ = "STUB: not implemented"; return *new(clients) }

func (c clients) add(client *clientAttr) { _ = "STUB: not implemented"; return }

func (c clients) del(key string) { _ = "STUB: not implemented"; return }

func (c clients) get(key string) (*clientAttr, bool) { _ = "STUB: not implemented"; return nil, false }

func (c clients) enum() chan *clientAttr { _ = "STUB: not implemented"; return nil }
