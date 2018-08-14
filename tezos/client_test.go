package tezos

import (
	"testing"
)

func testAddress() string {
	h, p := GetTestNodeParams()
	return h + ":" + p
}

func TestNew(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"invalid address", args{"123.456.789"}, true},
		{"valid address", args{testAddress()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.addr != tt.args.addr {
				t.Errorf("wanted %s, got %s", got.addr, tt.args.addr)
				return
			}
		})
	}
}

func TestClient_Get(t *testing.T) {
	type args struct {
		endpoint string
		output   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"invalid request", args{"/i/am/hungry", &ChainBlockStatusResponse{}}, true},
		{"valid request", args{"/chains/main/blocks/head", &ChainBlockStatusResponse{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := New(testAddress())
			if err != nil {
				t.Errorf("unexpected New() error = %v", err)
				return
			}
			if err = c.Get(tt.args.endpoint, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("Client.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
