package main

import (
	"log"
	"net/http"
	"os"
	"github.com/bikz007/rr-loadbalancer-golang/simpleserver"
	"github.com/bikz007/rr-loadbalancer-golang/loadbalancer"
	"github.com/bikz007/rr-loadbalancer-golang/utils"
)

const domain = "http://127.0.0.1"

func main() {
	servers := []simpleserver.Server{
		simpleserver.CreateNewSimpleServer(domain+":"+os.Args[1]),
		simpleserver.CreateNewSimpleServer(domain+":"+os.Args[2]),
		simpleserver.CreateNewSimpleServer(domain+":"+os.Args[3]),
		simpleserver.CreateNewSimpleServer(domain+":"+os.Args[4]),
	}

	roundRobinLoadBalancer := loadbalancer.CreateNewLoadBalancer(os.Args[5], servers)

	redirectHandler := func (resp http.ResponseWriter, req *http.Request)  {
		roundRobinLoadBalancer.ServeProxy(resp,req)
	}
	http.HandleFunc("/",redirectHandler)

	log.SetPrefix("[INFO]")

	log.Printf("Serving requests at 'http://127.0.0.1:%s' \n",roundRobinLoadBalancer.GetPort())
	err := http.ListenAndServe(":"+roundRobinLoadBalancer.GetPort(), nil)

	utils.HandleErr(err)
}