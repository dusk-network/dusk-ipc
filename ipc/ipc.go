package ipc

import (
	"github.com/dusk-network/dusk-blockchain/pkg/p2p/wire/message"
	"github.com/dusk-network/dusk-blockchain/pkg/util/nativeutils/eventbus"
	"google.golang.org/grpc"
)

// compile time check that eventbus.Listener interface is implemented
var _ eventbus.Listener = &Listener{}

// Listener implements the eventbus.Listener
type Listener struct {
	srv *grpc.Server
}

// NewListener returns a new IPCListener
func NewListener(srv *grpc.Server) *Listener {
	if srv == nil {
		panic("without grpc server there is no ipc")
	}
	cs := &Listener{
		srv: srv,
	}

	return cs
}

func (il *Listener) Notify(message message.Message) error {
	panic("implement me")
}

func (il *Listener) Close() {
	panic("implement me")
}
