package integration

import (
	v1 "github.com/juanfont/headscale/gen/go/headscale/v1"
)

type ControlServer interface {
	Shutdown() error
	Execute(command []string) (string, error)
	GetHealthEndpoint() string
	GetEndpoint() string
	WaitForReady() error
	CreateNamespace(namespace string) error
	CreateAuthKey(namespace string, reusable bool, ephemeral bool) (*v1.PreAuthKey, error)
	ListMachinesInNamespace(namespace string) ([]*v1.Machine, error)
	GetCert() []byte
	GetHostname() string
	GetIP() string
}
