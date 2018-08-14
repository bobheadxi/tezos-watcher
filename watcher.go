package watcher

import (
	"fmt"
	"time"

	"github.com/bobheadxi/tezos-watcher/tezos"
)

// TezosWatcher is a tezos node watcher
type TezosWatcher struct {
	c *tezos.Client
}

// ConnectOpts defines connection options to your tezos node
type ConnectOpts struct {
	Host string
	Port string
}

// New creates a new tezos-watcher
func New(opts ConnectOpts) (*TezosWatcher, error) {
	address := opts.Host
	if opts.Port != "" {
		address += ":" + opts.Port
	}
	c, err := tezos.New(address)
	if err != nil {
		return nil, err
	}
	return &TezosWatcher{c}, nil
}

// BlockOptions defines block configuration
type BlockOptions struct {
	Chain string
	Block string
	Rate  time.Duration
}

// WatchBlock watches for changes to the node's chain via the /chains/main/blocks/head endpoint.
// It is up to the caller to stop the spawned goroutine.
func (w *TezosWatcher) WatchBlock(opts BlockOptions, quit <-chan struct{}) (<-chan ChainHeadStatus, <-chan error) {
	var (
		statusCh = make(chan ChainHeadStatus)
		errCh    = make(chan error)
		ticker   = time.NewTicker(opts.Rate)
	)
	go func() {
		for {
			select {
			case <-ticker.C:
				go func() {
					s, err := w.getBlockStatus(opts.Chain, opts.Block)
					if err != nil {
						errCh <- err
					} else {
						statusCh <- s
					}
				}()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return statusCh, errCh
}

func (w *TezosWatcher) getBlockStatus(chain, block string) (ChainHeadStatus, error) {
	var (
		resp tezos.ChainBlockStatusResponse
		req  = fmt.Sprintf("chains/%s/blocks/%s", chain, block)
	)
	if err := w.c.Get(req, &resp); err != nil {
		return ChainHeadStatus{}, nil
	}
	return NewChainHeadStatus(resp), nil
}
