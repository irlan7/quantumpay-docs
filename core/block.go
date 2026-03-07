package core

import (
"crypto/sha256"
"encoding/hex"
"fmt"
"time"
)

// Block mewakili satu blok data yang sudah divalidasi oleh konsensus
type Block struct {
Index        uint64
Timestamp    int64
Transactions []*Transaction
PrevHash     string
Hash         string
Validator    string
}

// GenerateHash membuat cryptographic hash dari isi blok (Sovereign Security)
func (b *Block) GenerateHash() string {
record := fmt.Sprintf("%d%d%s%s", b.Index, b.Timestamp, b.PrevHash, b.Validator)
h := sha256.New()
h.Write([]byte(record))
return hex.EncodeToString(h.Sum(nil))
}

// NewBlock mencetak blok baru ke dalam jaringan
func NewBlock(prevBlock *Block, txs []*Transaction, validator string) *Block {
block := &Block{
Index:        prevBlock.Index + 1,
Timestamp:    time.Now().UnixNano(),
Transactions: txs,
PrevHash:     prevBlock.Hash,
Validator:    validator,
}
block.Hash = block.GenerateHash()
return block
}
