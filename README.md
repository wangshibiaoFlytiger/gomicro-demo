# 若本项目给您带来收获, 还请您动动小拇指,右上角给点个赞哈,万分感谢哈哈!!!

**项目原始仓库地址: https://github.com/wangshibiaoFlytiger/gomicro-demo**  
**欢迎关注官方微信公众号www_sofineday_com,及时收到一手技术资料!**  

# 博客地址
[https://sofineday.com](https://sofineday.com)    
欢迎加作者微信`645102170`或进群共同交流, 请扫下方二维码. 请备注 sofineday😄

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