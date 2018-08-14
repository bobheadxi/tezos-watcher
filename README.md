# tezos-watcher [![GoDoc](https://godoc.org/github.com/bobheadxi/tezos-watcher?status.svg)](https://godoc.org/github.com/bobheadxi/tezos-watcher) [![Build Status](https://travis-ci.org/bobheadxi/tezos-watcher.svg?branch=master)](https://travis-ci.org/bobheadxi/tezos-watcher) [![Coverage Status](https://coveralls.io/repos/github/bobheadxi/tezos-watcher/badge.svg?branch=master)](https://coveralls.io/github/bobheadxi/tezos-watcher?branch=master)

Monitor your Tezos node.

## Installing

```bash
$> go get github.com/bobheadxi/tezos-watcher/cmd/tezos-watcher
```

## Usage

Get the [Tezos Alphanet up and running](http://tezos.gitlab.io/betanet/introduction/howtoget.html). Make sure to start up a node with the RPC endpoints exposed - for example:

```bash
$> ./alphanet.sh start --rpc-port 8732
```

Optionally grab some [free Tezzies](http://tezos.gitlab.io/betanet/introduction/howtouse.html#get-free-tezzies) to start off.

You can now start watching your local node:

```bash
$> tezos-watcher watch-chain
```

The `-h` flag offers documentation on the command line tool.