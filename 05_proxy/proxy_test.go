package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserControllerProxy_Login(t *testing.T) {
	uc := &UserController{}
	ucp := NewUserControllerProxy(uc)
	err := ucp.Login("test", "123")
	assert.Equal(t, nil, err)
}

func TestUserControllerProxy_Register(t *testing.T) {
	uc := &UserController{}
	ucp := NewUserControllerProxy(uc)
	err := ucp.Register("test", "123")
	assert.Equal(t, nil, err)
}
