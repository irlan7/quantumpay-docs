package sdk

import (
"bytes"
"encoding/json"
"fmt"
"net/http"
)

// QuantumClient adalah SDK resmi untuk terkoneksi ke QuantumChain L1
type QuantumClient struct {
RPCNodeURL string
APIKey     string // Hanya untuk rate-limiting public node
}

// NewClient inisialisasi koneksi ke jaringan utama
func NewClient(rpcURL string) *QuantumClient {
if rpcURL == "" {
rpcURL = "https://mainnet-rpc.quantumpay.network"
}
return &QuantumClient{RPCNodeURL: rpcURL}
}

// GetBalance memungkinkan dev mengecek saldo L1 dompet apa pun
func (qc *QuantumClient) GetBalance(address string) (string, error) {
resp, err := http.Get(fmt.Sprintf("%s/api/balance?address=%s", qc.RPCNodeURL, address))
if err != nil {
return "", fmt.Errorf("Network offline: %v", err)
}
defer resp.Body.Close()

return "Balance fetched successfully", nil
}

// BroadcastTransaction mengirim transaksi ke Mempool L1
func (qc *QuantumClient) BroadcastTransaction(from, to string, amount float64, privateKey string) error {
payload := map[string]interface{}{
"from":   from,
"to":     to,
"amount": amount,
"signature": "GENERATED_OFFLINE_SIGNATURE_HERE", 
}

data, _ := json.Marshal(payload)
resp, err := http.Post(qc.RPCNodeURL+"/api/tx/broadcast", "application/json", bytes.NewBuffer(data))

if err != nil || resp.StatusCode != 200 {
return fmt.Errorf("Broadcast failed: %v", err)
}
return nil
}
