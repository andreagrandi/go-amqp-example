# go-amqp-example
Example of publisher and consumer of RabbitMQ messages in Golang

# Requirements

The following instructions assume that you have **Go** correctly installed (and $GOPATH set correctly) and **RabbitMQ** installed with default settings.

# Getting the code

```
go get github.com/andreagrandi/go-amqp-example
```

The code will be in **$GOPATH/src/github.com/andreagrandi/go-amqp-example**

# Running the code

Run the publisher:

```
cd $GOPATH/src/github.com/andreagrandi/go-amqp-example/publisher
go run publisher.go
```

Run the consumer:

```
cd $GOPATH/src/github.com/andreagrandi/go-amqp-example/consumer
go run consumer.go
```
