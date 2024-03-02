package main

import (
	"context"
	"fmt"
	"rsapi/auth/rpc/token/tokenclient"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func TestTokenRpc(t *testing.T) {
	cli := tokenclient.NewToken(zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{Hosts: []string{"127.0.0.1:2379"}, Key: "token.rpc"},
	}))
	// test Generate
	token, err := cli.Generate(context.TODO(), &tokenclient.GenerateReq{
		Email:  "test",
		Expire: 5,
	})
	assert.Equal(t, nil, err)
	// test verify
	verifyRes, err := cli.Verify(context.TODO(), &tokenclient.VerifyReq{
		Token:    token.Token,
		Refreash: false,
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, token.Token, verifyRes.Token)

	time.Sleep(1 * time.Second)
	verifyRes, err = cli.Verify(context.TODO(), &tokenclient.VerifyReq{
		Token:    token.Token,
		Refreash: true,
	})
	assert.Equal(t, nil, err)
	fmt.Println(token.Token)
	fmt.Println(verifyRes.Token)
	assert.NotEqual(t, token.Token, verifyRes.Token)
	// test invalid
	_, err = cli.Invalid(context.TODO(), &tokenclient.InvalidReq{Email: &token.Token})
	assert.Equal(t, nil, err)
}
