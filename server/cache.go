package server

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"onlineservice/global"
	"onlineservice/util/cast"
	"onlineservice/util/constant"
)

const (
	RedisKeySOnlineUser = "online_service_user_%v"
)

func cacheStartUp(ctx context.Context, uid int64) error {
	idx := uid % constant.OnlineUIDMod
	key := fmt.Sprintf(RedisKeySOnlineUser, idx)
	err := redisCli.SAdd(ctx, key, uid).Err()
	if err != nil {
		global.ExcLog.Printf("ctx %v cacheStartUp uid %v err %v", ctx, uid, err)
	}
	return err
}

func cacheShutdown(ctx context.Context, uid int64) error {
	idx := uid % constant.OnlineUIDMod
	key := fmt.Sprintf(RedisKeySOnlineUser, idx)
	err := redisCli.SRem(ctx, key, uid).Err()
	if err != nil {
		global.ExcLog.Printf("ctx %v cacheShutdown uid %v err %v", ctx, uid, err)
	}
	return err
}

func cacheGetOnlineAll(ctx context.Context) ([]int64, int64, error) {
	uids := make([]int64, 0)
	var cnt int64
	for i := 0; i < constant.OnlineUIDMod; i++ {
		key := fmt.Sprintf(RedisKeySOnlineUser, i)
		val, err := redisCli.SMembers(ctx, key).Result()
		cnt += int64(len(val))
		if err != nil && err != redis.Nil {
			global.ExcLog.Printf("ctx %v cacheGetOnlineAll key %v err %v", ctx, key, err)
			return nil, 0, err
		}
		for _, v := range val {
			uids = append(uids, cast.ParseInt(v, 0))
		}
	}
	return uids, cnt, nil
}
