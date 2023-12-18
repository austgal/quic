Your task is to write an application which acts as a pub/sub message broker server and communicates via QUIC (QUIC ).


Server specifications:

Accepts QUIC connections on 2 ports

Publisher port

Subscriber port

The server notifies publishers if a subscriber has connected

If no subscribers are connected, the server must inform the publishers

The server sends any messages received from publishers to all connected subscribers

Your code will be evaluated based on it’s readability, maintainability, robustness, documentation and test coverage. Make sure to use git for version control and separate your code chunks into meaningful commits, since your git skills will also be evaluated.