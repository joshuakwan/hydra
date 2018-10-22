package storage

// TODO: use transaction
// func (c *storageClient) DeleteObject(ctx context.Context, key string) (int64, error) {
// 	dresp, err := c.client.Delete(ctx, config.GetStorageRoot()+key)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return dresp.Deleted, err
// }

// func (c *storageClient) UpdateObject(ctx context.Context, key, val string) error {
// 	key = config.GetStorageRoot() + key
// 	txnResp, err := c.client.KV.Txn(ctx).If(
// 		found(key),
// 	).Then(
// 		clientv3.OpPut(key, val),
// 	).Commit()

// 	if err != nil {
// 		return err
// 	}
// 	if !txnResp.Succeeded {
// 		return fmt.Errorf("key %s not exist", key)
// 	}
// 	return nil
// }

// func (c *storageClient) GetObjects(ctx context.Context, key string) ([][]byte, error) {
// 	getResp, err := c.client.KV.Get(ctx, config.GetStorageRoot()+key, clientv3.WithPrefix())
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(getResp.Kvs) == 0 {
// 		return nil, fmt.Errorf("key %s not found", key)
// 	}

// 	var results [][]byte
// 	for _, kv := range getResp.Kvs {
// 		results = append(results, kv.Value)
// 	}

// 	return results, nil
// }

// func (c *storageClient) Watch(ctx context.Context, key string) <-chan string {
// 	out := make(chan string)

// 	go func() {
// 		rch := c.client.Watch(ctx, config.GetStorageRoot()+key, clientv3.WithPrefix())
// 		for wresp := range rch {
// 			for _, ev := range wresp.Events {
// 				out <- fmt.Sprintf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
// 			}
// 		}
// 		close(out)
// 	}()

// 	return out
// }
