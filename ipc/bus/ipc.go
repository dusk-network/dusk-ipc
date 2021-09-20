package bus

import (
	"github.com/dusk-network/dusk-blockchain/pkg/p2p/wire/message"
	"github.com/dusk-network/dusk-blockchain/pkg/p2p/wire/topics"
	"github.com/dusk-network/dusk-blockchain/pkg/util/nativeutils/eventbus"
)

type IPC struct {
	*eventbus.EventBus
}

// eventbus.Subscriber implementation

func (ib *IPC) Subscribe(topic topics.Topic, listener eventbus.Listener,
) uint32 {
	panic("implement me")
}

func (ib *IPC) Unsubscribe(topic topics.Topic, u uint32) {
	panic("implement me")
}

// eventbus.Publisher implementation

func (ib *IPC) Publish(topic topics.Topic, message message.Message) []error {
	panic("implement me")
}

// eventbus.Multicaster implementation

func (ib *IPC) AddDefaultTopic(topic topics.Topic) {
	panic("implement me")
}

func (ib *IPC) SubscribeDefault(listener eventbus.Listener) uint32 {
	panic("implement me")
}
