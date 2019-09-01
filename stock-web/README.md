# Stock Service

This is the Stock service

Generated with

```
micro new github.com/wanghaoxi3000/go-secbuy-micro/stock-web --namespace=go.micro.secbuy --alias=stock --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.secbuy.web.stock
- Type: web
- Alias: stock

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
./stock-web
```

Build a docker image
```
make docker
```