package rpc

import (
"encoding/json"
"net/http"
)

// RPCHandler menangani permintaan dari Web3 Wallet atau SDK
type RPCHandler struct {
// Terkoneksi dengan State Machine
}

// GetBalanceAPI adalah endpoint untuk SDK `client.GetBalance()`
func (rpc *RPCHandler) GetBalanceAPI(w http.ResponseWriter, r *http.Request) {
address := r.URL.Query().Get("address")
if address == "" {
http.Error(w, "Address required", http.StatusBadRequest)
return
}

// Format balasan standar RPC Node
json.NewEncoder(w).Encode(map[string]interface{}{
"jsonrpc": "2.0",
"result": map[string]string{
"address": address,
"balance": "0", // Data asli akan ditarik dari core.WorldState
},
"id": 1,
})
}
