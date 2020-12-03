package cache

import (
	"github.com/codingXiang/go-orm"
	"github.com/codingXiang/go-orm/v2/redis"
)

type Handler struct {
	dbClient *orm.Orm
	redisClient *redis.RedisClient
}

func (h *Handler) List(map[string]interface{}) {

}