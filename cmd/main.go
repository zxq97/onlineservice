package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"onlineservice/conf"
	"onlineservice/rpc/online/pb"
	"onlineservice/server"
)

var (
	onlineConf *conf.Conf
	err        error
)

func main() {
	onlineConf, err = conf.LoadYaml(conf.OnlineConfPath)
	if err != nil {
		panic(err)
	}

	err = server.InitService(onlineConf)
	if err != nil {
		panic(err)
	}

	etcdRegistry := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = onlineConf.Etcd.Addr
	})

	service := micro.NewService(
		micro.Name(onlineConf.Grpc.Name),
		micro.Address(onlineConf.Grpc.Addr),
		micro.Registry(etcdRegistry),
	)
	service.Init()
	err = online_service.RegisterOnlineServerHandler(
		service.Server(),
		new(server.OnlineService),
	)
	if err != nil {
		panic(err)
	}
	err = service.Run()
	if err != nil {
		panic(err)
	}
}
