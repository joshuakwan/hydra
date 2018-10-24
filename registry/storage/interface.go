package storage

import (
	"context"
)

// DestroyFunc is to destroy any resources used by the storage returned in Create() together.
type DestroyFunc func()

// Storage defines the interface of a storage client
type Storage interface {
	Create(ctx context.Context, key string, data []byte, ttl int64) error
	Delete(ctx context.Context, key string) error
	Update(ctx context.Context, key string, data []byte) error
	Get(ctx context.Context, key string) ([]byte, error)
	List(ctx context.Context, key string) ([][]byte, error)
	Watch(ctx context.Context, key string) (Watcher, error)
	Close() error
}

// Watcher defines a watcher interface
type Watcher interface {
	Stop()
	ResultChan() <-chan WatchEvent
}

// EventType defines the type of watch event
type EventType string

const (
	// Created indicates an object is created
	Created EventType = "CREATED"
	// Deleted indicates an object is deleted
	Deleted EventType = "DELETED"
	// Updated indicates an object is updated
	Updated EventType = "UPDATED"
	// Error indicates an error occurs
	Error EventType = "ERROR"
)

// WatchEvent defines events during a watch
type WatchEvent struct {
	Type EventType
	Data []byte
}
