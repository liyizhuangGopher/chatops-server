package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type Config struct {
	zrpc.RpcServerConf

	// 引入consul
	Consul consul.Conf

	// mysql配置
	Mysql struct {
		DataSource string
	}
	Secret string
}
