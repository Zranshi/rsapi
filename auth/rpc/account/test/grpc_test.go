package test

import (
	"context"
	"errors"
	"rsapi/auth/rpc/account/accountclient"
	"rsapi/auth/rpc/account/internal"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

func TestAccount(t *testing.T) {
	// test register
	cli := accountclient.NewAccount(
		zrpc.MustNewClient(zrpc.RpcClientConf{
			Etcd: discov.EtcdConf{Hosts: []string{"127.0.0.1:2379"}, Key: "account.rpc"},
		}),
	)

	in := &accountclient.RegisterReq{
		Email: "test@test.com",
		Pwd:   "password",
		Name:  "test",
	}

	registerRes, err := cli.Register(context.TODO(), in)
	assert.Equal(t, nil, err)
	assert.Equal(t, "test@test.com", registerRes.Email)

	_, err = cli.Register(context.TODO(), in)
	assert.True(t, errors.Is(err, internal.ErrAccountAlreadyExist))

	//	test login
	loginIn := &accountclient.LoginReq{Email: "test@test.com", Pwd: "password"}

	loginRes, err := cli.Login(context.TODO(), loginIn)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", loginRes.Token)

	// test info
	infoRes, err := cli.Info(context.TODO(), &accountclient.InfoReq{Email: "test@test.com"})
	assert.Equal(t, nil, err)
	assert.Equal(t, "test", infoRes.Name)

	// test edit
	pwd, name := "changed pwd", "changed name"
	editIn := &accountclient.EditReq{Email: "test@test.com", Pwd: &pwd, Name: &name}
	_, err = cli.Edit(context.TODO(), editIn)
	assert.Equal(t, nil, err)

	// check changed field
	infoRes, err = cli.Info(context.TODO(), &accountclient.InfoReq{Email: "test@test.com"})
	assert.Equal(t, nil, err)
	assert.Equal(t, "changed name", infoRes.Name)

	//	test delete
	_, err = cli.Signout(context.TODO(), &accountclient.SignOutReq{Email: "test@test.com",
		Pwd: "changed pwd"})
	assert.Equal(t, nil, err)
}
