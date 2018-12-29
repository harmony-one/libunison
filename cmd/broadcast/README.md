## Benchmark test for coopcast

#### What is Coopcast
Coopcast is a fast, scalable and fault resilient network protocol to deliver messages to a group of nodes in the network. 
The sender uses RaptorQ encoder (RFC6330) to encode the message and broadcasts small pieces of the encoded message (called Symbol) to its neighbor peers. Each peer who receives the Symbol will also responsible to relay it to its own neighbor peers, that's why we call it Coop(erate) Cast. 

##### Fast
Coopcast uses UDP protocol to send the packets, which reduce the TCP handshake and response time.

##### Scalability
The sender doesn't need to send the whole message to all it's neighbor peers, instead, it just sends different symbols to different peers and the RaptorQ decoder will guarantee the receiver can recover the original message. 


##### Fault tolerance
Actually, the RaptorQ decoder guarantees the receiver to recover from any K (a number determined by RaptorQ decoder parameters and message size) received Symbols with high probability 99%; or from any K+1 Symbols with 99.99% probability; or from any K+2 Symbols with 99.9999% etc.
At the same time, the RaptorQ encoder is able to deliver practically infinite number ( of Symbols. It means the network is robust to network failure as well as Byzantine nodes as long as majority of the nodes are honest and online.


#### Example of build, local test and clean
The following example is only for test purpose on local computer. It doesn't consider the network condition as well as Byzantine nodes. The real benchmark should be conducted on real network (possibly across multiple regions) with each node on separate machine. For local test, the recommended file size is less or equal than 1MB. 

##### Dependency
It depends on the following library:
[go-raptorq](https://github.com/harmony-one/go-raptorq)


##### Run example
###### Build go executable
./build.sh

###### Generate network topology (generate a network of 5 nodes, fully connected)
./generate_configs.sh  graph1.txt

###### Start 4 server nodes (0,1,2,3) waiting for receiving messages
./start_server 5  [coopcast|manycast]

###### node 4 will send file (test.txt) to other peers
./send_file.sh 5 test.txt [coopcast|manycast]

###### Kill background servers
./killserver.sh




