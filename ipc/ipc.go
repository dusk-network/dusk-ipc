package ipc

import (
	"github.com/dusk-network/dusk-blockchain/pkg/p2p/wire/message"
	"github.com/dusk-network/dusk-blockchain/pkg/p2p/wire/topics"
	"github.com/dusk-network/dusk-blockchain/pkg/util/nativeutils/eventbus"
)

type IPCBus struct {
	*eventbus.EventBus
}

// eventbus.Subscriber implementation

func (ib *IPCBus) Subscribe(topic topics.Topic, listener eventbus.Listener,
) uint32 {
	panic("implement me")
}

func (ib *IPCBus) Unsubscribe(topic topics.Topic, u uint32) {
	panic("implement me")
}

// eventbus.Publisher implementation

func (ib *IPCBus) Publish(topic topics.Topic, message message.Message) []error {
	panic("implement me")
}

// eventbus.Multicaster implementation

func (ib *IPCBus) AddDefaultTopic(topic topics.Topic) {
	panic("implement me")
}

func (ib *IPCBus) SubscribeDefault(listener eventbus.Listener) uint32 {
	panic("implement me")
}
