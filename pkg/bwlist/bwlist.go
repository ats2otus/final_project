package bwlist

import (
	"net"
	"sync"
)

type BWList interface {
	Contains(ip net.IP) bool
	Append(subnet *net.IPNet)
	Remove(subnet *net.IPNet)
}

type bwList struct {
	mu   sync.RWMutex
	list map[string]*net.IPNet
}

func (l *bwList) Contains(ip net.IP) bool {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for _, subnet := range l.list {
		if subnet.Contains(ip) {
			return true
		}
	}
	return false
}

func (l *bwList) Append(subnet *net.IPNet) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.list[subnet.String()]; ok {
		return
	}
	l.list[subnet.String()] = subnet
}

func (l *bwList) Remove(subnet *net.IPNet) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.list[subnet.String()]; !ok {
		return
	}
	delete(l.list, subnet.String())
}

func New(subnets ...*net.IPNet) BWList {
	list := bwList{
		list: make(map[string]*net.IPNet, len(subnets)),
	}
	for _, subnet := range subnets {
		list.Append(subnet)
	}
	return &list
}
