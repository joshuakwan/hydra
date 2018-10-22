package storage

// Client defines the interface to interact with a storage backend
// type Client interface {
// 	CreateObject(ctx context.Context, key, val string, ttl int64) error
// 	GetObject(ctx context.Context, key string) ([]byte, error)
// 	GetObjects(ctx context.Context, key string) ([][]byte, error)
// 	DeleteObject(ctx context.Context, key string) (int64, error)
// 	UpdateObject(ctx context.Context, key, val string) error
// 	Watch(ctx context.Context, key string) <-chan string
// 	Close() error
// }

// CreateClient as a factory to create a storage backend client upon the configuration

// func CloseClient() {
// 	destroyFunc()
// }
