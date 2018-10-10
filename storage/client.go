package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/joshuakwan/hydra/config"

	"go.etcd.io/etcd/clientv3"
)

func newEtcdV3Client() (Client, DestroyFunc, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   config.GetStorageEndpoints(),
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, nil, err
	}

	_, cancel := context.WithCancel(context.Background())
	destroyFunc := func() {
		cancel()
		client.Close()
	}

	return &storageClient{client}, destroyFunc, nil

}

type storageClient struct {
	client *clientv3.Client
}

// CreateObject creates an object in etcd
// TODO:  WithLease
func (c *storageClient) CreateObject(ctx context.Context, key, val string) error {
	key = config.GetStorageRoot() + key

	txnResp, err := c.client.KV.Txn(ctx).If(
		notFound(key),
	).Then(
		clientv3.OpPut(key, val),
	).Commit()

	if err != nil {
		return err
	}
	if !txnResp.Succeeded {
		return fmt.Errorf("key %s exists", key)
	}
	return nil
}

// TODO: use transaction
func (c *storageClient) DeleteObject(ctx context.Context, key string) (int64, error) {
	dresp, err := c.client.Delete(ctx, config.GetStorageRoot()+key)
	if err != nil {
		return 0, err
	}
	return dresp.Deleted, err
}

func (c *storageClient) UpdateObject(ctx context.Context, key, val string) error {
	key = config.GetStorageRoot() + key
	txnResp, err := c.client.KV.Txn(ctx).If(
		found(key),
	).Then(
		clientv3.OpPut(key, val),
	).Commit()

	if err != nil {
		return err
	}
	if !txnResp.Succeeded {
		return fmt.Errorf("key %s not exist", key)
	}
	return nil
}

func (c *storageClient) GetObject(ctx context.Context, key string) ([]byte, error) {
	getResp, err := c.client.KV.Get(ctx, config.GetStorageRoot()+key)
	if err != nil {
		return nil, err
	}

	if len(getResp.Kvs) == 0 {
		return nil, fmt.Errorf("key %s not found", key)
	}

	return getResp.Kvs[0].Value, nil
}

func (c *storageClient) Close() error {
	return c.client.Close()
}

func notFound(key string) clientv3.Cmp {
	return clientv3.Compare(clientv3.ModifiedRevision(key), "=", 0)
}

func found(key string) clientv3.Cmp {
	return clientv3.Compare(clientv3.ModifiedRevision(key), ">", 0)
}
