package loadbalancer

import (
	"log"
	"net/http"
	"github.com/bikz007/rr-loadbalancer-golang/simpleserver"
)

type LoadBalancer struct {
	port string
	roundRobinCount int
	servers []simpleserver.Server
}

func (lb *LoadBalancer) getNextAvailableServer() simpleserver.Server {
	resultServer := lb.servers[lb.roundRobinCount % len(lb.servers)]
	for !resultServer.IsAlive() {
		lb.roundRobinCount++;
		resultServer = lb.servers[lb.roundRobinCount % len(lb.servers)]
	}
	lb.roundRobinCount++
	return resultServer
}

func (lb *LoadBalancer) GetPort() string {
	return lb.port
}

func (lb *LoadBalancer) ServeProxy(resp http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()
	log.SetPrefix("[INFO]")
	log.Printf("Forwarding request to address %q\n",targetServer.Address())
	targetServer.Serve(resp,req)
}

func CreateNewLoadBalancer(port string, servers []simpleserver.Server) *LoadBalancer  {
	return & LoadBalancer{
		port: port,
		roundRobinCount: 0,
		servers: servers,
	}
}