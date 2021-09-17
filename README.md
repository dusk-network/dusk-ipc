# dusk-ipc

Inter-Process Communication framework for composing applications as discrete
modules connected by an event bus

## About

Within the repository for the Go reference implementation in
[github.com/dusk-network/dusk-blockchain](https://github.com/dusk-network/dusk-ipc)
there already exists a inter-module messaging system called
[eventbus](https://github.com/dusk-network/dusk-blockchain/tree/master/pkg/util/nativeutils/eventbus)
that already natively communicates with binary encoded (protobuf) data, or the
wire format, as the case may be, and thus lends itself to being easily
interfaced over a binary transport such as a filesystem pipe.

To the extent this already interfaces with node subsystems, these subsystems
then can be removed from the core in-process runtime and operate as independent
processes, connected to the eventbus this way. There may be some bits and pieces
that haven't already been integrated this way, and in general the entire node
application will be broken down so more of these RPC API interfaces will be
added in each one and their runtime controlled by the main management system
that is implemented in this library.

It is simple to conceive how especially the indexes for search API calls will be
connected together, a block database with a flat filesystem, an
accounts/addresses index, transaction index, all can be easily refactored to
query downstream instead of an in-process call, and in fact some of them already
using eventbus don't need any modification except to interpose the transport and
isolate the component in an independent runtime.
