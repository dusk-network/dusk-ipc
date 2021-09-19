# dusk-ipc

Inter-Process Communication framework for composing applications as discrete
modules connected by an event bus

## About

 **TODO:**

  This will composit the EventBus type and extend its methods with shadowing to
  implement a local broker for client processes that mirrors requests to the
  IPCBus by proxying protocol messages of interest back and forth over the 
  connection (ReaderWriterCloser interface implementation) which is created 
  by running a command that attaches via stdin/out and exposes interrupt and 
  kill methods for use in interrupt processing of the application, including 
  shutting down if the connection is closed, so the process terminates when 
  the controller terminates regardless of whether it formally signals this 
  (shutdown signal allows cleanup of network and pipe buffers).

