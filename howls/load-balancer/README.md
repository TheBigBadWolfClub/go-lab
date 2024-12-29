

# load-balancer

create a service cable of routing requests to servers using Load balancing algorithms

### Load balancing algorithms

The most common load balancing algorithms are:
- **Round Robin**: Simple and widely used, it distributes requests sequentially among the available servers.
- **Least Connections**: Routes traffic to the server with the least number of active connections, commonly used for balancing loads in real-time.
- **Weighted Round Robin**: An extension of Round Robin, it assigns weights to servers, allowing more powerful servers to handle more requests.
- **IP Hashing**: Ensures that the same client is always directed to the same server based on the client's IP address.