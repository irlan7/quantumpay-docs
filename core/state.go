package core

import (
"errors"
"sync"
)

// Konstanta L1
const OneQTM = 100_000_000

// Account menyimpan data saldo dasar jaringan (Native Coin)
type Account struct {
Address string
Balance uint64
Nonce   uint64
}

// Transaction mewakili instruksi mutasi di QuantumChain
type Transaction struct {
ID    []byte
From  string
To    string
Value uint64
Nonce uint64
Data  []byte // Payload QRC-20 atau Kontrak Pintar
}

// WorldState adalah struktur data absolut dari blockchain
type WorldState struct {
Accounts map[string]*Account
mu       sync.RWMutex
}

// NewWorldState menginisialisasi Genesis Block
func NewWorldState() *WorldState {
return &WorldState{
Accounts: make(map[string]*Account),
}
}

// GetAccount mengambil data dompet secara transparan
func (ws *WorldState) GetAccount(address string) *Account {
ws.mu.RLock()
defer ws.mu.RUnlock()
if acc, exists := ws.Accounts[address]; exists {
return acc
}
// Buat dompet baru jika belum ada
newAcc := &Account{Address: address, Balance: 0, Nonce: 0}
ws.Accounts[address] = newAcc
return newAcc
}

// ApplyTransaction mengeksekusi perpindahan saldo dengan validasi ketat
func (ws *WorldState) ApplyTransaction(tx *Transaction) error {
ws.mu.Lock()
defer ws.mu.Unlock()

sender := ws.GetAccount(tx.From)
receiver := ws.GetAccount(tx.To)

// Validasi Syarat Konsensus
if sender.Balance < tx.Value {
return errors.New("L1 REVERT: Saldo tidak mencukupi")
}
if tx.Nonce <= sender.Nonce {
return errors.New("L1 REVERT: Nonce transaksi kedaluwarsa (Replay Attack Protection)")
}

// Mutasi Ledger
sender.Balance -= tx.Value
receiver.Balance += tx.Value
sender.Nonce = tx.Nonce

return nil
}
