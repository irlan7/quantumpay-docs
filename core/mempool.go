package core

import "sync"

// Mempool adalah ruang tunggu untuk transaksi yang belum masuk ke dalam blok
type Mempool struct {
Pending []*Transaction
mu      sync.Mutex
}

// NewMempool menginisialisasi mempool baru
func NewMempool() *Mempool {
return &Mempool{
Pending: make([]*Transaction, 0),
}
}

// AddTransaction memasukkan transaksi baru ke ruang tunggu
func (m *Mempool) AddTransaction(tx *Transaction) {
m.mu.Lock()
defer m.mu.Unlock()
m.Pending = append(m.Pending, tx)
}

// GetAndClear mengambil semua transaksi untuk divalidasi dan mengosongkan mempool
func (m *Mempool) GetAndClear() []*Transaction {
m.mu.Lock()
defer m.mu.Unlock()
txs := m.Pending
m.Pending = make([]*Transaction, 0)
return txs
}
