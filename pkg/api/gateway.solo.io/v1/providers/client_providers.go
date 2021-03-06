// Code generated by skv2. DO NOT EDIT.

package v1

import (
	gateway_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
  The intention of these providers are to be used for Mocking.
  They expose the Clients as interfaces, as well as factories to provide mocked versions
  of the clients when they require building within a component.

  See package `github.com/solo-io/skv2/pkg/multicluster/register` for example
*/

// Provider for GatewayClient from Clientset
func GatewayClientFromClientsetProvider(clients gateway_solo_io_v1.Clientset) gateway_solo_io_v1.GatewayClient {
	return clients.Gateways()
}

// Provider for Gateway Client from Client
func GatewayClientProvider(client client.Client) gateway_solo_io_v1.GatewayClient {
	return gateway_solo_io_v1.NewGatewayClient(client)
}

type GatewayClientFactory func(client client.Client) gateway_solo_io_v1.GatewayClient

func GatewayClientFactoryProvider() GatewayClientFactory {
	return GatewayClientProvider
}

type GatewayClientFromConfigFactory func(cfg *rest.Config) (gateway_solo_io_v1.GatewayClient, error)

func GatewayClientFromConfigFactoryProvider() GatewayClientFromConfigFactory {
	return func(cfg *rest.Config) (gateway_solo_io_v1.GatewayClient, error) {
		clients, err := gateway_solo_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.Gateways(), nil
	}
}

// Provider for RouteTableClient from Clientset
func RouteTableClientFromClientsetProvider(clients gateway_solo_io_v1.Clientset) gateway_solo_io_v1.RouteTableClient {
	return clients.RouteTables()
}

// Provider for RouteTable Client from Client
func RouteTableClientProvider(client client.Client) gateway_solo_io_v1.RouteTableClient {
	return gateway_solo_io_v1.NewRouteTableClient(client)
}

type RouteTableClientFactory func(client client.Client) gateway_solo_io_v1.RouteTableClient

func RouteTableClientFactoryProvider() RouteTableClientFactory {
	return RouteTableClientProvider
}

type RouteTableClientFromConfigFactory func(cfg *rest.Config) (gateway_solo_io_v1.RouteTableClient, error)

func RouteTableClientFromConfigFactoryProvider() RouteTableClientFromConfigFactory {
	return func(cfg *rest.Config) (gateway_solo_io_v1.RouteTableClient, error) {
		clients, err := gateway_solo_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.RouteTables(), nil
	}
}

// Provider for VirtualServiceClient from Clientset
func VirtualServiceClientFromClientsetProvider(clients gateway_solo_io_v1.Clientset) gateway_solo_io_v1.VirtualServiceClient {
	return clients.VirtualServices()
}

// Provider for VirtualService Client from Client
func VirtualServiceClientProvider(client client.Client) gateway_solo_io_v1.VirtualServiceClient {
	return gateway_solo_io_v1.NewVirtualServiceClient(client)
}

type VirtualServiceClientFactory func(client client.Client) gateway_solo_io_v1.VirtualServiceClient

func VirtualServiceClientFactoryProvider() VirtualServiceClientFactory {
	return VirtualServiceClientProvider
}

type VirtualServiceClientFromConfigFactory func(cfg *rest.Config) (gateway_solo_io_v1.VirtualServiceClient, error)

func VirtualServiceClientFromConfigFactoryProvider() VirtualServiceClientFromConfigFactory {
	return func(cfg *rest.Config) (gateway_solo_io_v1.VirtualServiceClient, error) {
		clients, err := gateway_solo_io_v1.NewClientsetFromConfig(cfg)
		if err != nil {
			return nil, err
		}
		return clients.VirtualServices(), nil
	}
}
