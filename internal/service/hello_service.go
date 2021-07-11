package service

import "errors"

type HelloService interface {
	Greeting(flag bool) (string, error)
}

type HelloServiceImpl struct {
}

func ProvideHelloService() HelloService {
	return &HelloServiceImpl{}
}

func (h *HelloServiceImpl) Greeting(flag bool) (string, error) {
	if flag {
		return "Hello there !", nil
	} else {
		return "", errors.New("expected error")
	}
}
