package consul

import (
	consulapi "github.com/hashicorp/consul/api"
)

var Client *consulapi.Client

func Init() {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	var err error
	Client, err = consulapi.NewClient(config)
	if err != nil {
		panic(err)
	}
}

func Register(svcName string, port int) {
	registration := new(consulapi.AgentServiceRegistration)
	registration.Name = svcName
	registration.Port = port

	err := Client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}

