package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joshuakwan/hydra/codec"

	"github.com/joshuakwan/hydra/config"
	"go.etcd.io/etcd/clientv3"
)

type storageClient struct {
	client       *clientv3.Client
	leaseManager *leaseManager
	watcher      *watcher
}

// NewStorage initialize a new storage client
func NewStorage() (Storage, DestroyFunc, error) {
	switch config.GetStorageType() {
	case "etcdv3":
		return newEtcdV3Client()
	default:
		return nil, nil, fmt.Errorf("unknown storage type: %s", config.GetStorageType())
	}
}

func newEtcdV3Client() (Storage, DestroyFunc, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   config.GetStorageEndpoints(),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, nil, err
	}

	_, cancel := context.WithCancel(context.Background())
	destroyFunc := func() {
		log.Printf("start destroyFunc\n")
		cancel()
		client.Close()
	}

	return &storageClient{
		client:       client,
		leaseManager: newLeaseManager(client),
		watcher:      &watcher{client: client},
	}, destroyFunc, nil
}

func (c *storageClient) ttlOpt(ctx context.Context, ttl int64) clientv3.OpOption {
	var leaseID clientv3.LeaseID
	if ttl == 0 {
		leaseID = clientv3.LeaseID(0)
	} else {
		leaseID = c.leaseManager.GetLease(ctx, ttl)
	}
	return clientv3.WithLease(leaseID)
}

// CreateObject creates an object in etcd
func (c *storageClient) Create(ctx context.Context, key string, data []byte, ttl int64) error {
	key = config.GetStorageRoot() + key
	opOpt := c.ttlOpt(ctx, ttl)

	log.Printf("start the txn to create %s\n", key)
	txnResp, err := c.client.KV.Txn(ctx).If(
		notFound(key),
	).Then(
		clientv3.OpPut(key, string(data), opOpt),
	).Commit()

	if err != nil {
		return err
	}
	if !txnResp.Succeeded {
		return fmt.Errorf("key %s exists", key)
	}
	return nil
}

func (c *storageClient) Delete(ctx context.Context, key string) error {
	key = config.GetStorageRoot() + key
	delResp, err := c.client.KV.Delete(ctx, key)
	if err != nil {
		return err
	}
	if delResp.Deleted == 0 {
		return fmt.Errorf("key %s not found, nothing deleted", key)
	}
	return nil
}

func (c *storageClient) Update(ctx context.Context, key string, data []byte) error {
	key = config.GetStorageRoot() + key

	txnResp, err := c.client.KV.Txn(ctx).If(
		found(key),
	).Then(
		clientv3.OpPut(key, string(data)),
	).Commit()

	if err != nil {
		return err
	}
	if !txnResp.Succeeded {
		return fmt.Errorf("key %s not exists", key)
	}
	return nil
}

func (c *storageClient) Get(ctx context.Context, key string) ([]byte, error) {
	getResp, err := c.client.KV.Get(ctx, config.GetStorageRoot()+key)
	if err != nil {
		return nil, err
	}

	if len(getResp.Kvs) == 0 {
		return nil, fmt.Errorf("key %s not found", key)
	}

	return getResp.Kvs[0].Value, nil
}

func (c *storageClient) List(ctx context.Context, key string) ([][]byte, error) {
	getResp, err := c.client.KV.Get(ctx, config.GetStorageRoot()+key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	if len(getResp.Kvs) == 0 {
		return nil, fmt.Errorf("key %s not found", key)
	}

	data := make([][]byte, len(getResp.Kvs))
	for idx, kv := range getResp.Kvs {
		data[idx] = kv.Value
	}

	return data, nil
}

func (c *storageClient) Watch(ctx context.Context, key string, codec codec.Codec) (Watcher, error) {
	return c.watcher.Watch(ctx, config.GetStorageRoot()+key, codec)
}

func (c *storageClient) Close() error {
	return c.client.Close()
}

func notFound(key string) clientv3.Cmp {
	return clientv3.Compare(clientv3.ModRevision(key), "=", 0)
}

func found(key string) clientv3.Cmp {
	return clientv3.Compare(clientv3.ModRevision(key), ">", 0)
}
