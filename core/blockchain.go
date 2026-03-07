package core

import (
"errors"
"fmt"
"sync"
)

// Blockchain menyimpan state dari seluruh rantai L1
type Blockchain struct {
Chain      []*Block
State      *WorldState
Difficulty int
mu         sync.RWMutex
}

// AddBlock memvalidasi dan menambahkan blok baru ke dalam rantai
func (bc *Blockchain) AddBlock(newBlock *Block) error {
bc.mu.Lock()
defer bc.mu.Unlock()

lastBlock := bc.Chain[len(bc.Chain)-1]

// Validasi Kriptografi
if newBlock.PrevHash != lastBlock.Hash {
return errors.New("🚨 L1 Consensus Failed: Invalid Previous Hash")
}

if newBlock.Hash != newBlock.GenerateHash() {
return errors.New("🚨 L1 Consensus Failed: Block Hash is tampered")
}

bc.Chain = append(bc.Chain, newBlock)
fmt.Printf("🧱 Block #%d added to QuantumChain L1\n", newBlock.Index)
return nil
}
