package go_libp2p_grpc

import "net"

// fakeAddr returns a dummy address.
func fakeAddr() net.Addr {
	remoteIp := net.ParseIP("127.1.0.1")
	return &net.TCPAddr{IP: remoteIp, Port: 0}
}
