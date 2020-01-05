# Partitioner

## Descrition

Some time ago I saw an architecture solution using Kafka and a kind of "pipeline" process to solve a problem of concurrent updates to the database. The solution uses the approach of split the input commands coming from a topic, between partition keys x consumers group that guarantee that a message with a specific key would always be processed on the same instance.
Thus the idea behind _partioner_ is to build a tool to make easier and lighter a solution for the problem of concurrency. Or maybe, this is just more one solution for queueing. 😂

- an `instance` controls the partitioning of messages
- the partitioning is made by a key that is associated to a client when configuration refresh
- the association of keys to clients is maintained by the instance (the sharing of configuration between instances ~was not considerate for a first version~ will be implemented using [Raft](https://godoc.org/github.com/hashicorp/raft))


### Features
- [x] Base API
- [ ] Base Server
- [ ] HTTP REST API
- [ ] Console CLI
- [ ] Consul (with [Raft](https://godoc.org/github.com/hashicorp/raft))