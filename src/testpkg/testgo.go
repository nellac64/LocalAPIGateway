package testpkg

import "github.com/nellac64/common-sdk/etcd/discovery"

func TestFunc() *discovery.Service {
	return &discovery.Service{
		Name:     "my service",
		IP:       "10.0.0.1",
		Port:     "1234",
		Protocol: "1234",
	}
}
