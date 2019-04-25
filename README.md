# tezos-watcher [![GoDoc](https://godoc.org/github.com/bobheadxi/tezos-watcher?status.svg)](https://godoc.org/github.com/bobheadxi/tezos-watcher) [![Build Status](https://travis-ci.com/bobheadxi/tezos-watcher.svg?branch=master)](https://travis-ci.com/bobheadxi/tezos-watcher) [![codecov](https://codecov.io/gh/bobheadxi/tezos-watcher/branch/master/graph/badge.svg)](https://codecov.io/gh/bobheadxi/tezos-watcher)

Monitor your Tezos node.

## Installing

To install the CLI, use the following `go get` command or check the `releases` page.

```bash
$> go get github.com/bobheadxi/tezos-watcher/cmd/tezos-watcher
```

## Usage

### CLI

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

### Library

```go
w, _ := watcher.New(watcher.ConnectOpts{
  Host: "127.0.0.1", Port: "8732"})

quit := make(chan struct{})
defer close(quit)
statusCh, errCh := w.WatchBlock(watcher.BlockOptions{
  Chain: "main",
  Block: "head",
  Rate:  1 * time.Second,
}, quit)

for {
  select {
  case err := <-errCh:
    if err != nil {
      return
    }
  case event := <-statusCh:
    output, _ := json.Marshal(event)
    println(string(output))
  }
}
```

## Development

```bash
$> make          # install dependencies, install tezos-watcher CLI
$> make node     # start up Tezos sandbox node
$> make test     # execute tests against sandbox node
```

## Notes

This project was written as part of a pre-task for a job.
