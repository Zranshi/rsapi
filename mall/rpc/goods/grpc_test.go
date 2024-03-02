package main

import (
	"context"
	"rsapi/mall/rpc/goods/goodsclient"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func TestGoodsRpc(t *testing.T) {
	cli := goodsclient.NewGoods(zrpc.MustNewClient(zrpc.RpcClientConf{Etcd: discov.EtcdConf{Hosts: []string{"127.0.0.1:2379"}, Key: "goods.rpc"}}))

	in := &goodsclient.AddReq{
		Name:        "test goods",
		Description: "test description",
		Images:      []string{"test1.png", "test2.png"},
	}

	addRes, err := cli.Add(context.TODO(), in)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, addRes.Id)

	in.Name = "test goods2"
	addRes, err = cli.Add(context.TODO(), in)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, addRes.Id)

	id := addRes.Id
	newDescription := "changed description"
	_, err = cli.Edit(context.TODO(), &goodsclient.EditReq{Id: id, Description: &newDescription})
	assert.Equal(t, nil, err)

	detailRes, err := cli.Detail(context.TODO(), &goodsclient.DetailReq{Id: id})
	assert.Equal(t, nil, err)
	assert.Equal(t, newDescription, detailRes.Item.Description)

	listRes, err := cli.List(context.TODO(), &goodsclient.ListReq{})
	assert.Equal(t, nil, err)
	assert.GreaterOrEqual(t, len(listRes.Items), 2)

	_, err = cli.Delete(context.TODO(), &goodsclient.DeleteReq{Id: id})
	assert.Equal(t, nil, err)

	_, err = cli.Detail(context.TODO(), &goodsclient.DetailReq{Id: id})
	assert.NotEqual(t, nil, err)
}
