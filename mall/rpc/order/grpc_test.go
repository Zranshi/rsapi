package main

import (
	"context"
	"fmt"
	"rsapi/mall/rpc/order/orderclient"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func TestOrderRpc(t *testing.T) {
	cli := orderclient.NewOrder(zrpc.MustNewClient(zrpc.RpcClientConf{Etcd: discov.EtcdConf{Hosts: []string{"127.0.0.1:2379"}, Key: "order.rpc"}}))
	createRes, err := cli.Create(context.TODO(), &orderclient.CreateReq{UserId: 57, GoodsId: 29, Count: 5})
	assert.Equal(t, nil, err)

	id := createRes.Id

	detailRes, err := cli.Detail(context.TODO(), &orderclient.DetailReq{Id: id})
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(5), detailRes.Item.Count)

	listRes, err := cli.List(context.TODO(), &orderclient.ListReq{UserId: 57})
	assert.Equal(t, nil, err)
	assert.GreaterOrEqual(t, len(listRes.Items), 1)
	fmt.Println(len(listRes.Items))

	cancelRes, err := cli.Cancel(context.TODO(), &orderclient.CancelReq{Id: id})
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(5), cancelRes.Item.Count)
	assert.Equal(t, int64(2), cancelRes.Item.Status)

	completeRes, err := cli.Complete(context.TODO(), &orderclient.CompleteReq{Id: id})
	assert.Equal(t, nil, err)
	assert.Equal(t, int64(5), completeRes.Item.Count)
	assert.Equal(t, int64(1), completeRes.Item.Status)

	deleteres, err := cli.Delete(context.TODO(), &orderclient.DeleteReq{Id: id})
	assert.Equal(t, nil, err)
	assert.Equal(t, id, deleteres.Item.Id)
}
