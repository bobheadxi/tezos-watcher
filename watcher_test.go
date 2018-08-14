package watcher

import (
	"testing"
	"time"

	"github.com/bobheadxi/tezos-watcher/tezos"
)

func TestNew(t *testing.T) {
	type args struct {
		opts ConnectOpts
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"invalid connectOpts", args{ConnectOpts{tezos.TestHost, ""}}, true},
		{"valid connectOpts", args{ConnectOpts{tezos.TestHost, tezos.TestPort}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Error("client should not be nil")
				return
			}
		})
	}
}

func TestTezosWatcher_WatchBlock(t *testing.T) {
	type args struct {
		opts BlockOptions
	}
	tests := []struct {
		name string
		args args
	}{
		{"should stream statuses", args{
			BlockOptions{Chain: "main", Block: "head", Rate: 1 * time.Second}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := New(ConnectOpts{tezos.TestHost, tezos.TestPort})
			if err != nil {
				t.Errorf("unexpected New() error = %v", err)
			}

			quit := make(chan (struct{}))
			statusCh, errCh := w.WatchBlock(tt.args.opts, quit)

			// make sure one status is received
			status := <-statusCh
			if status.Hash == "" || status.ChainID == "" {
				t.Error("expected hash and chain ID")
				return
			}

			// make sure cleanup is done properly
			close(quit)
			_, statusOpen := (<-statusCh)
			_, errorOpen := (<-errCh)
			if statusOpen || errorOpen {
				t.Error("channels did not close")
				return
			}
		})
	}
}

func TestTezosWatcher_getBlockStatus(t *testing.T) {
	type args struct {
		chain string
		block string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid chain", args{"main", "head"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := New(ConnectOpts{tezos.TestHost, tezos.TestPort})
			if err != nil {
				t.Errorf("unexpected New() error = %v", err)
			}
			got, err := w.getBlockStatus(tt.args.chain, tt.args.block)
			if (err != nil) != tt.wantErr {
				t.Errorf("TezosWatcher.getBlockStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (got.Hash == "" || got.ChainID == "") {
				t.Error("expected hash and chain ID")
				return
			}
		})
	}
}
