# Go Skeleton

is lightweight project directory structure in golang, which has a 4 basic layers handler (http), service, repository, and model. This project implement automation dependecy injection and graceful shutdown to handle resource released properly.

## Package Included
1. [Dependency Injection](https://godoc.org/github.com/facebookgo/inject)
2. [Message Broker](https://github.com/nsqio/go-nsq)
3. [Configuration](https://github.com/spf13/viper) (will change it later)
4. [Router](https://github.com/gorilla/mux)