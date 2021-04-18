package initialize

import (
	"gin-vue-admin/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	var err error
	if global.GVA_REDIS != nil {
		_, err = global.GVA_REDIS.Ping().Result()
		if err == nil {
			return global.GVA_REDIS
		}
	}
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
	}
	return client
}
