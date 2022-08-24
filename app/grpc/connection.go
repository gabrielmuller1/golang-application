package grpc

import (
	"sync"

	"google.golang.org/grpc"
)

var (
	srvMu   sync.Mutex
	poolSrv = map[string]*Pool{}
)

// Set pool connections
func SetPool(p *Pool, srv string) {
	if _, ok := poolSrv[srv]; !ok {
		srvMu.Lock()
		defer srvMu.Unlock()
		poolSrv[srv] = p
	}
}

// Get client connection
func GetConnGRPC(poolManager *PoolManager, srv string) (*grpc.ClientConn, Done, error) {
	if _, ok := poolSrv[srv]; !ok {
		srvMu.Lock()
		defer srvMu.Unlock()
		if _, ok := poolSrv[srv]; !ok {
			p, _, e := poolManager.NewPool(srv)
			if e != nil {
				return nil, func() {}, e
			}
			poolSrv[srv] = p
		}
	}
	return poolSrv[srv].Get()
}
