package validator

import "fmt"

// ValidatorInfo menyimpan data institusi yang menjaga jaringan
type ValidatorInfo struct {
WalletAddress string
StakedAmount  uint64
IsActive      bool
Uptime        float64
}

// DelegateStake memungkinkan user menitipkan QTM ke Validator
func DelegateStake(delegator string, validator string, amount uint64) {
fmt.Printf("🔐 Wallet %s delegated %d QTM to Validator %s\n", delegator, amount, validator)
// Logika penguncian (lock-up) di State Machine
}
