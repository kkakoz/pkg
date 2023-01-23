package redisx

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"time"
)

type unLockFunc func()

func Lock(ctx context.Context, key string, timeout time.Duration) (bool, unLockFunc, error) {
	value := uuid.NewV4().String()
	result, err := client.SetNX(ctx, key, value, timeout).Result()
	unLock := func() {
		_, err = client.Eval(ctx, "if redis.call('get',KEYS[1]) == ARGV[1] then"+
			"   return redis.call('del',KEYS[1]) "+
			"else"+
			"   return 0 "+
			"end", []string{key}, []string{value}).Result()
	}
	return result, unLock, err
}
