package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/rogerogers/dingtalk-ops/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDingtalkRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Cache       *redis.Client
	CacheConfig map[string]string
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Data{
		Cache:       rdb,
		CacheConfig: map[string]string{"prefix": c.Redis.Prefix},
	}, cleanup, nil
}
