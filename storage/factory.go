package storage

import (
	"context"
	"fmt"

	"github.com/joshuakwan/hydra/config"
	"go.etcd.io/etcd/clientv3"
)

// DestroyFunc is to destroy any resources used by the storage returned in Create() together.
type DestroyFunc func()

//type DeserializeFunc func([]byte) interface{}

// PutResponse encapsulates the response of a PUT op
type PutResponse struct {
	*clientv3.PutResponse
}

// GetResponse encapsulates the response of a GET op
type GetResponse struct {
	*clientv3.GetResponse
}

// DeleteResponse encapsulates the response of a DELETE op
type DeleteResponse struct {
	*clientv3.DeleteResponse
}

// Client defines the interface to interact with a storage backend
type Client interface {
	CreateObject(ctx context.Context, key, val string) error
	GetObject(ctx context.Context, key string) ([]byte, error)
	DeleteObject(ctx context.Context, key string) (int64, error)
	UpdateObject(ctx context.Context, key, val string) error
	Close() error
}

// CreateClient as a factory to create a storage backend client upon the configuration
func CreateClient() (Client, DestroyFunc, error) {
	switch config.GetStorageType() {
	case "etcdv3":
		return newEtcdV3Client()
	default:
		return nil, nil, fmt.Errorf("unknown storage type: %s", config.GetStorageType())
	}
}
