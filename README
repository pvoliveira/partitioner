#Partioner

##Descrition

Some time ago I saw an architecture solution using Kafka and a kind of "pipeline" process to solve a problem of concurrent updates to the database. The solution uses the approach of split the input commands coming from a topic, between partition keys x consumers group that guarantee that a message with a specific key would always be processed on the same instance, all time that a message with that key occurs.
Thus the idea behind _partioner_ is to build a tool to make easier and lighter a solution for the problem of concurrency. Or maybe, this is just more one solution for queue. 😂
