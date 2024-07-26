package config

import (
	"google.golang.org/grpc"

	aTransport "benneighbour.com/a_service/transport"
	bTransport "benneighbour.com/b_service/transport"
	//bTransport "benneighbour.com/c_service/transport"
	//bTransport "benneighbour.com/d_service/transport"
	//bTransport "benneighbour.com/f_service/transport"
)

var Services = []struct {
	Name     string
	Register func(*grpc.Server)
}{
	{Name: "a_transport", Register: aTransport.Register},
	{Name: "b_transport", Register: bTransport.Register},
	//{Name: "c_transport", Register: cTransport.Register},
	//{Name: "d_transport", Register: dTransport.Register},
	//{Name: "e_transport", Register: eTransport.Register},
	//{Name: "f_transport", Register: fTransport.Register},
	//etc...
}
