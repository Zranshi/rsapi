package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"rsapi/mall/rpc/cart/cart"
	"rsapi/util"

	"rsapi/mall/rpc/cart/internal/config"
	"rsapi/mall/rpc/cart/internal/server"
	"rsapi/mall/rpc/cart/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/cart.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if c.Service.RandomPort {
		port, err := util.GetRandomPort(10000, 20000)
		if err != nil {
			panic(err)
		}
		c.ListenOn = fmt.Sprintf("127.0.0.1:%d", port)
	}

	ctx := svc.NewServiceContext(c)

	logx.DisableStat()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		cart.RegisterCartServer(grpcServer, server.NewCartServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
