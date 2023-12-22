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

### Improvements and other ideas/concerns:
- The project's folder structure could be better aligned with Go best practices.
- More unit tests need to be added.
- More consistent error handling is needed, as not all cases might be covered and proper actions taken.
- Graceful shutdown could be added.
- There are some hardcoded values that are left as-is for testing purposes.

## License
[MIT](https://choosealicense.com/licenses/mit/)