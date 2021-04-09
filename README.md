# Gomicro Service

This is the Gomicro service

Generated with

```
micro new gomicro-demo --namespace=microdemo --type=service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: microdemo.service.gomicro
- Type: service
- Alias: gomicro

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./gomicro-service
```

Build a docker image
```
make docker
```