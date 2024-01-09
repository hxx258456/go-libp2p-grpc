package go_libp2p_grpc

import (
	"context"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	manet "github.com/multiformats/go-multiaddr/net"
	"io"
	"net"
)

var _ net.Listener = (*listener)(nil)

type listener struct {
	h        host.Host
	streamCh chan network.Stream
	ctx      context.Context
	cancel   context.CancelFunc
}

func (l listener) Accept() (net.Conn, error) {
	select {
	case <-l.ctx.Done():
		return nil, io.EOF
	case s := <-l.streamCh:
		return &Conn{Stream: s}, nil
	}
}

func (l listener) Close() error {
	l.cancel()
	return nil
}

func (l listener) Addr() net.Addr {
	addrs := l.h.Network().ListenAddresses()
	if len(addrs) > 0 {
		for _, a := range addrs {
			na, err := manet.ToNetAddr(a)
			if err == nil {
				return na
			}
		}
	}
	return fakeAddr()
}
