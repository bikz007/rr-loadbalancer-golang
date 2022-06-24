package simpleserver

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/bikz007/rr-loadbalancer-golang/utils"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(resp http.ResponseWriter, req *http.Request)
}

type simpleServer struct {
	addr string
	proxy *httputil.ReverseProxy
}

func CreateNewSimpleServer(address string) *simpleServer {
	serverUrl, err := url.Parse(address)
	utils.HandleErr(err)

	return &simpleServer{
		addr: address,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func (s *simpleServer) Address() string { return s.addr}
func (s *simpleServer) IsAlive() bool { return true}
func (s *simpleServer) Serve(resp http.ResponseWriter, req *http.Request) { 
	s.proxy.ServeHTTP(resp,req)
}