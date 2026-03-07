package network

import (
"fmt"
"sync"
)

// Peer mewakili node lain di dalam jaringan QuantumChain
type Peer struct {
ID      string
Address string
IsSuper bool // Supernode / Validator
}

// P2PManager mengatur komunikasi antar node
type P2PManager struct {
Peers map[string]*Peer
mu    sync.RWMutex
}

// BroadcastBlock menyebarkan blok baru ke seluruh jaringan L1
func (net *P2PManager) BroadcastBlock(blockHash string) {
fmt.Printf("📡 Broadcasting Block [%s] to %d peers...\n", blockHash, len(net.Peers))
// Logika Gossip Protocol (seperti Tendermint) akan diimplementasikan di sini
}

// BroadcastTx menyebarkan transaksi dari Mempool ke node lain
func (net *P2PManager) BroadcastTx(txID string) {
fmt.Printf("📡 Broadcasting Tx [%s] to network...\n", txID)
}
