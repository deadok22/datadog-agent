// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build linux

package probe

import (
	"context"
	"time"

	"github.com/DataDog/ebpf/manager"
)

type reOrdererNodePool struct {
	root *reOrdererNode
}

func (p *reOrdererNodePool) alloc() *reOrdererNode {
	node := p.root
	if node != nil && node.timestamp == 0 {
		p.root = node.next
		return node
	}

	return &reOrdererNode{}
}

func (p *reOrdererNodePool) free(node *reOrdererNode) {
	node.timestamp = 0

	if p.root == nil {
		p.root = node
	} else {
		node.next = p.root
		p.root = node
	}
}

func newReOrdererNodePool(size int) *reOrdererNodePool {
	return &reOrdererNodePool{}
}

type reOrdererList struct {
	root *reOrdererNode
	tail *reOrdererNode
	size uint64
}

type reOrdererNode struct {
	timestamp uint64
	data      []byte
	next      *reOrdererNode
	prev      *reOrdererNode
}

func (l *reOrdererList) append(node *reOrdererNode) {
	l.size++

	if l.root == nil {
		l.root = node
		l.tail = node

		return
	}

	var prev *reOrdererNode

	curr := l.tail
	for curr != nil {
		if node.timestamp >= curr.timestamp {
			if prev != nil {
				prev.prev = node
			} else {
				l.tail = node
			}
			node.next = prev
			curr.next = node
			node.prev = curr

			return
		}

		prev = curr
		curr = curr.prev
	}

	l.root.prev = node
	node.next = l.root
	l.root = node
}

func (l *reOrdererList) pop(tm uint64) *reOrdererNode {
	if l.tail == nil {
		return nil
	}

	node := l.tail
	l.tail = node.prev

	return node
}

type ReOrdererOpts struct {
	QueueSize  uint64        // size of the chan where the perf data are pushed
	WindowSize uint64        // number of element to keep for orderering
	Delay      time.Duration // delay to wait before handle an element outside of the window in millisecond
}

// ReOrderer defines an event re-orderer
type ReOrderer struct {
	queue            chan []byte
	handler          func(data []byte)
	list             *reOrdererList
	pool             *reOrdererNodePool
	resolveTimestamp func(t uint64) time.Time
	timestampGetter  func(data []byte) (uint64, error)
	opts             ReOrdererOpts
}

// Start event handler loop
func (r *ReOrderer) Start(ctx context.Context) {
	ticker := time.NewTicker(r.opts.Delay)
	defer ticker.Stop()

	consume := func(limit uint64) {
		curr := r.list.root
		for curr != nil && r.list.size > limit {
			r.handler(curr.data)

			next := curr.next

			r.pool.free(curr)

			curr = next
			r.list.size--
		}

		r.list.root = curr
		if curr == nil {
			r.list.tail = nil
		} else {
			curr.prev = nil
		}
	}

	for {
		select {
		case data := <-r.queue:
			tm, err := r.timestampGetter(data)
			if err != nil {
				continue
			}

			node := r.pool.alloc()
			node.timestamp = tm
			node.data = data

			r.list.append(node)

			consume(r.opts.WindowSize)
		case now := <-ticker.C:
			curr := r.list.tail
			if curr == nil {
				continue
			}

			tm := r.resolveTimestamp(curr.timestamp)
			if now.Sub(tm) < r.opts.Delay {
				continue
			}

			consume(0)
		case <-ctx.Done():
			return
		}
	}
}

// HandleEvent handle event form perf ring
func (e *ReOrderer) HandleEvent(CPU int, data []byte, perfMap *manager.PerfMap, manager *manager.Manager) {
	e.queue <- data
}

// NewReOrderer returns a new ReOrderer
func NewReOrderer(handler func([]byte), tsg func(data []byte) (uint64, error), rts func(t uint64) time.Time, opts ReOrdererOpts) *ReOrderer {
	return &ReOrderer{
		queue:            make(chan []byte, opts.QueueSize),
		handler:          handler,
		list:             &reOrdererList{},
		pool:             &reOrdererNodePool{},
		timestampGetter:  tsg,
		resolveTimestamp: rts,
		opts:             opts,
	}
}
