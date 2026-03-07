package node

import "fmt"

// Server mengelola koneksi Peer-to-Peer (P2P) antar node L1
type Server struct {
NodeID  string
Port    string
Peers   []string
}

// Start menghidupkan node P2P
func (s *Server) Start() {
fmt.Printf("🚀 QuantumNode [%s] starting on port %s...\n", s.NodeID, s.Port)
// Logika gRPC dan P2P discovery akan diimplementasikan di sini
}
