# Pub/sub message broker
## About
A pub/sub message broker server that communicates via QUIC.

## 3rd party packages:
- [quic-go](https://pkg.go.dev/github.com/lucas-clemente/quic-go) - a QUIC implementation in pure Go

## Task description
Your task is to write an application which acts as a pub/sub message broker server and communicates via QUIC (QUIC ).

Server specifications:

- Accepts QUIC connections on 2 ports
- Publisher port
- Subscriber port
- The server notifies publishers if a subscriber has connected
- If no subscribers are connected, the server must inform the publishers
- The server sends any messages received from publishers to all connected subscribers

Your code will be evaluated based on it’s readability, maintainability, robustness, documentation and test coverage. Make sure to use git for version control and separate your code chunks into meaningful commits, since your git skills will also be evaluated.

## Usage
### Build image:
```
docker build -t broker .
```

### Run image:
```
docker run -p 6666:6666/udp -p 6667:6667/udp broker
```

Please adjust the ports according to the needs.

### Other information
- The TLS generating function was copied from the examples in the quic-go library, which are located [here](https://github.com/quic-go/quic-go/blob/d3c5f389d44797108a1bee7e06d5b92434c26d6d/example/echo/echo.go#L99C39-L99C39).
- For server testing I've wrote a basic QUIC client, source code can be found [here](https://github.com/austgal/pubsub_client).

## License
[MIT](https://choosealicense.com/licenses/mit/)