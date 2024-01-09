package go_libp2p_grpc

import (
	"context"
	"errors"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"google.golang.org/grpc"
	"net"
)

func WithP2PDialer(h host.Host, pid protocol.ID) grpc.DialOption {
	return grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		peerId, err := peer.Decode(s)
		if err != nil {
			return nil, err
		}

		if h.Network().Connectedness(peerId) != network.Connected {
			return nil, errors.New("not connected to peer")
		}

		stream, err := h.NewStream(ctx, peerId, pid)
		if err != nil {
			return nil, err
		}

		return &Conn{
			stream,
		}, nil
	})
}
