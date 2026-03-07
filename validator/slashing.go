package validator

import "fmt"

// Konstanta Hukuman
const SlashingPenalty = 0.05 // Potong 5% dari aset yang di-stake jika curang

// Slash mengeksekusi hukuman pemotongan dana pada validator nakal (Double Signing)
func Slash(validatorAddress string, reason string) {
fmt.Printf("⚖️ SLASHING EXECUTED: Validator %s penalized for %s!\n", validatorAddress, reason)
// Logika pemotongan saldo dan pembekuan (Jailing) node
}

// Jail membekukan validator agar tidak bisa memproduksi blok
func Jail(validatorAddress string) {
fmt.Printf("🔒 Validator %s is now JAILED.\n", validatorAddress)
}
