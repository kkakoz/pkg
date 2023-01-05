package redisx

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type unLockFunc func()

func Lock(key string, timeout time.Duration) (bool, unLockFunc, error) {
	value := uuid.NewV4().String()
	result, err := client.SetNX(key, value, timeout).Result()
	unLock := func() {
		_, err = client.Eval("if redis.call('get',KEYS[1]) == ARGV[1] then"+
			"   return redis.call('del',KEYS[1]) "+
			"else"+
			"   return 0 "+
			"end", []string{key}, []string{value}).Result()
	}
	return result, unLock, err
}
