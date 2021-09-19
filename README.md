# dusk-ipc

Inter-Process Communication framework for composing applications as discrete
modules connected by an event bus

## About

### In one very long sentence:

  This will composit the EventBus type and extend its methods with shadowing to
  implement a local broker for client processes that mirrors requests to the
  IPCBus by proxying protocol messages of interest back and forth over the 
  connection (ReaderWriterCloser interface implementation) which is created 
  by running a command that attaches via stdin/out and exposes interrupt and 
  kill methods for use in interrupt processing of the application, including 
  shutting down if the connection is closed, so the process terminates when 
  the controller terminates regardless of whether it formally signals this 
  (shutdown signal allows cleanup of network and pipe buffers).

### The Junction and the Driver

Because this package is an implementation of an EventBus, an internal 
process message routing system, this 'bus' motif is riffed for the client 
and server, to make less generic names, the server is the Junction and the 
client is the Driver.

The analogy is that messages are like passengers and the driver takes on 
passengers at the Junction in both directions, like a bus interchange, or a 
short word, Junction, which also conveys the connectivity that it provides. 

The Driver extends the network to a new process. Then, a network transport 
in place of the filesystem pipe of the default initial implementation could 
be added and this could be called a Road as it carries in either direction. 
Actually, the filesystem pipe could be called `Road` and you have local 
`Road`s and `Highway` for the network version. If allowance can be made to 
expand the coordinate space from 8 to 16 bits, some very very interesting 
architectures could be made for connecting services together in a redundant 
and distributed way, just add a blockchain based identity system and traffic 
accounting market.

## Roadmap

1. Basic implementation with attaching using stdio pipe in Go.
2. Port eventbus and dusk-ipc repos to Rust.
3. Break dusk-blockchain components down into IPCBus Services in Go.
4. Port them one by one into Rust, simplifying also repository collaboration 
   due to the small scope.
