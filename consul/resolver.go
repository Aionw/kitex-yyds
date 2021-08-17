package consul

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consulapi "github.com/hashicorp/consul/api"
)

type ConsulResolver struct {
	ConsulClient *consulapi.Client
}

type ConsulInstance struct {
	addr net.TCPAddr
}

func (i *ConsulInstance) Address() net.Addr {
	return &i.addr
}

func (i *ConsulInstance) Weight() int {
	return 10
}

func (i *ConsulInstance) Tag(key string) (value string, exist bool) {
	return "", true
}

func (r *ConsulResolver) Target(ctx context.Context, target rpcinfo.EndpointInfo) (description string) {
	return target.ServiceName()
}

// Resolve returns a list of instances for the given description of a target.
func (r *ConsulResolver) Resolve(ctx context.Context, desc string) (discovery.Result, error) {
	service, _, err := r.ConsulClient.Catalog().Service(desc, "", nil)
	instances := make([]discovery.Instance, 0, len(service))
	if err != nil {
		return discovery.Result{}, nil
	}
	for _, s := range service {
		instances = append(instances, &ConsulInstance{
			addr: net.TCPAddr{
				IP:   net.ParseIP(s.Address),
				Port: s.ServicePort,
			},
		})
	}
	return discovery.Result{
		Cacheable: true,
		CacheKey:  desc,
		Instances: instances,
	}, nil
}

// Diff computes the difference between two results.
// When `next` is cacheable, the Change should be cacheable, too. And the `Result` field's CacheKey in
// the return value should be set with the given cacheKey.
func (r *ConsulResolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.Change{}, false
}

// Name returns the name of the resolver.
func (r *ConsulResolver) Name() string {
	return "consul"
}
