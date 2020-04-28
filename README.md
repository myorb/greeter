# Golang gRPC example

This is simple gRPC implementation for a simple game

- client update nessesary game numbers to server
- server receives numbers and sends it back
- both client and serever handle context errors (try to close client during send)

## Requirements

- go 1.14
- protobuf installed
- go support for protobuf installed

## Installation

### MacOS

```bash
brew install go
brew install protobuf
brew install make
```

Make sure `protoc-gen-go` added in PATH

### Linux

TBD

## Complie

```bash
make all
```

It should create two binaries `game-server` and `game-client`

## Use

Start server `./game-server` and in other terminal start `./game-client`

Client output example:

```bash
./client
2020/04/28 03:42:05 c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:05 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:1 rmax:{x:1 y:1 length:1}
2020/04/28 03:42:05 status:true
2020/04/28 03:42:05 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:2 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:05 start update
2020/04/28 03:42:05 start update 1
2020/04/28 03:42:05 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:4 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:06 start update 2
2020/04/28 03:42:06 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:8 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:08 start update 3
2020/04/28 03:42:08 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:16 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:09 start update 4
2020/04/28 03:42:09 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:32 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:10 start update 5
2020/04/28 03:42:10 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:64 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:11 start update 6
2020/04/28 03:42:11 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:128 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:12 start update 7
2020/04/28 03:42:12 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:256 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:13 start update 8
2020/04/28 03:42:13 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:512 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:14 start update 9
2020/04/28 03:42:14 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:1024 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:15 start update 10
2020/04/28 03:42:15 id:"c8783c37-644d-4719-b355-eb53872c2f27" selected:2048 rmax:{x:2 y:2 length:2}
2020/04/28 03:42:16 finished
```

Server output:

```bash
./server
2020/04/28 03:41:58 This is log form main
2020/04/28 03:42:05 start new server
2020/04/28 03:42:05 game updated c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:06 game updated c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:08 game updated c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:09 game updated c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:10 game updated c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:11 game updated c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:12 game updated c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:13 game updated c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:14 game updated c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:15 game updated c8783c37-644d-4719-b355-eb53872c2f27
2020/04/28 03:42:16 exit
```
