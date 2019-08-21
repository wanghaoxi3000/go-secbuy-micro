# Payment Service

This is the payment service

Generated with

```
micro new github.com/wanghaoxi3000/go-secbuy-mirco/payment-srv --namespace=go.micro.secbuy --alias=payment --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.secbuy.srv.payment
- Type: srv
- Alias: order

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./payment-srv
```

Build a docker image
```
make docker
```
