package watcher

import (
	"time"

	"github.com/bobheadxi/tezos-watcher/tezos"
)

// Header denotes metadata from a tezos RPC endpoint
type Header struct {
	Level     int       `json:"level"`
	Timestamp time.Time `json:"timestamp"`
}

// ChainHeadStatus is a simplified overview of blockchain status
type ChainHeadStatus struct {
	ChainID string `json:"chain_id"`
	Hash    string `json:"hash"`
	Header  Header `json:"header"`
	Baker   string `json:"baker"`
	Level   struct {
		Level                int  `json:"level"`
		LevelPosition        int  `json:"level_position"`
		Cycle                int  `json:"cycle"`
		CyclePosition        int  `json:"cycle_position"`
		VotingPeriod         int  `json:"voting_period"`
		VotingPeriodPosition int  `json:"voting_period_position"`
		ExpectedCommitment   bool `json:"expected_commitment"`
	} `json:"level"`
	ConsumedGas string `json:"consumed_gas"`
}

// NewChainHeadStatus parses a tezosChainBlockStatusResponse into a simplified overview
func NewChainHeadStatus(r tezos.ChainBlockStatusResponse) ChainHeadStatus {
	return ChainHeadStatus{
		ChainID:     r.ChainID,
		Hash:        r.Hash,
		Header:      Header{r.Header.Level, r.Header.Timestamp},
		Baker:       r.Metadata.Baker,
		Level:       r.Metadata.Level,
		ConsumedGas: r.Metadata.ConsumedGas,
	}
}
