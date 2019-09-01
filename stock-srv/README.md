# Stock Service

This is the Stock service

Generated with

```
micro new github.com/wanghaoxi3000/go-secbuy-micro/stock-srv --namespace=go.micro.secbuy --alias=stock --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.secbuy.srv.stock
- Type: srv
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
./stock-srv
```

Build a docker image
```
make docker
```


## prepare
```
go get -u github.com/micro/protoc-gen-micro
```


## Test
```
 micro --registry=mdns call go.micro.secbuy.srv.stock StockService.CreateCommodity '{"name":"商品","count":10,"sale":0}'
```

```
micro --registry=mdns call go.micro.secbuy.srv.stock StockService.GetCommodity '{"id":5}'
```