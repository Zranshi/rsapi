package test

import (
	"context"
	"rsapi/mall/rpc/cart/cartclient"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func TestCartRpc(t *testing.T) {
	cli := cartclient.NewCart(zrpc.MustNewClient(zrpc.RpcClientConf{Etcd: discov.EtcdConf{Hosts: []string{"127.0.0.1:2379"}, Key: "cart.rpc"}}))
	_, err := cli.Push(context.Background(), &cartclient.PushReq{
		UserId:  57,
		GoodsId: 29,
		Count:   5,
	})
	assert.Equal(t, nil, err)

	pushRes, err := cli.Push(context.Background(), &cartclient.PushReq{
		UserId:  57,
		GoodsId: 29,
		Count:   6,
	})
	assert.Equal(t, nil, err)

	listRes, err := cli.List(context.TODO(), &cartclient.ListReq{})
	assert.Equal(t, nil, err)

	for _, item := range listRes.Items {
		if item.Id == pushRes.Id {
			assert.Equal(t, 6, item.Count)
		}
	}

	cli.Pop(context.TODO(), &cartclient.PopReq{Id: pushRes.Id, Count: 5})

	listRes, err = cli.List(context.TODO(), &cartclient.ListReq{})
	assert.Equal(t, nil, err)

	for _, item := range listRes.Items {
		if item.Id == pushRes.Id {
			assert.Equal(t, 6, item.Count)
		}
	}
}
