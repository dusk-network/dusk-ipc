# IPCBus

IPCBus contains all the common parts between the [`Junction`](../junction) and
[`Driver`](../driver) to implement a bridge to connect and extend the
[`EventBus`](https://github.com/dusk-network/dusk-blockchain/tree/master/pkg/util/nativeutils/eventbus)
to another process's own `EventBus` allowing the building of single purpose
services in an application's normally monolithic structure, simplifying
maintenance, bug fixing and porting to other languages, where there is an 
available EventBus and IPCBus implementation.
