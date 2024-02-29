package logic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {

	claims := Claims{
		Email:  "test",
		Role:   "admin",
		Expire: 0,
	}
	token, err := GenerateJwt(&claims, "secret")
	assert.Equal(t, nil, err)

	out, err := ParseJwt(token, "secret")
	assert.Equal(t, nil, err)

	out.Expire = -1

	token, err = GenerateJwt(out, "secret")
	assert.Equal(t, nil, err)

	out, err = ParseJwt(token, "secret")
	assert.NotEqual(t, nil, err)
	t.Log(out)
}
