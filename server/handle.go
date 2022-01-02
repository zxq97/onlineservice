package server

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
	"onlineservice/conf"
	"onlineservice/rpc/online/pb"
)

type OnlineService struct {
}

var (
	mcCli    *memcache.Client
	redisCli redis.Cmdable
	dbCli    *gorm.DB
)

func InitService(config *conf.Conf) error {
	var err error
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)
	mcCli = conf.GetMC(config.MC.Addr)
	redisCli = conf.GetRedisCluster(config.RedisCluster.Addr)
	//dbCli, err = conf.GetGorm(fmt.Sprintf(conf.MysqlAddr, config.Mysql.User, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.DB))
	return err
}

func (os *OnlineService) StartUp(ctx context.Context, req *online_service.OnlineRequest, res *online_service.EmptyResponse) error {
	err := startUp(ctx, req.Uid)
	return err
}

func (os *OnlineService) Shutdown(ctx context.Context, req *online_service.OnlineRequest, res *online_service.EmptyResponse) error {
	err := shutdown(ctx, req.Uid)
	return err
}

func (os *OnlineService) GetOnlineAll(ctx context.Context, req *online_service.EmptyRequest, res *online_service.OnlineResponse) error {
	uids, cnt, err := getOnlineAll(ctx)
	if err != nil {
		return err
	}
	res.Uids = uids
	res.Count = cnt
	return nil
}
