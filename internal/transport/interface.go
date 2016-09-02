package transport

import (
	"errors"
	"strings"
	"sync"
)

type Transport interface {
	GetName() string
	GetInstance() (Transport, error)
}

var (
	availableTransports     = make(map[string]Transport)
	availableTransportsLock sync.Mutex

	// ErrDuplicatedTransport is returned when a duplicate name is registered
	ErrDuplicatedTransport = errors.New("Transport name duplicated")
	// ErrNoSuchTransport is returned when a unexistant transport is requested
	ErrNoSuchTransport = errors.New("Specified transport type does not exist")
)

func Register(transport Transport) error {
	availableTransportsLock.Lock()
	defer availableTransportsLock.Unlock()

	name := strings.ToLower(transport.GetName())
	if _, ok := availableTransports[name]; !ok {
		availableTransports[name] = transport
		return nil
	}
	return ErrDuplicatedTransport
}

func Get(name string) (Transport, error) {
	availableTransportsLock.Lock()
	defer availableTransportsLock.Unlock()

	name = strings.ToLower(name)
	if t, ok := availableTransports[name]; ok {
		return t.GetInstance()
	}
	return nil, ErrNoSuchTransport
}
