package core

import "fmt"

// Konstanta Pajak Resmi Negara (Sesuai Regulasi Indonesia)
const (
PPNRate       = 0.11  // Pajak Pertambahan Nilai (11%)
CryptoTaxRate = 0.001 // Pajak Penghasilan Kripto Final (0.1%)
PlatformFee   = 0.005 // QuantumPay L1 Fee (0.5%)
)

// TaxAssessment adalah struktur bukti setoran pajak yang transparan
type TaxAssessment struct {
Principal   uint64 // Jumlah dasar transaksi
Fee         uint64 // Potongan platform
VAT         uint64 // Potongan PPN
CryptoTax   uint64 // Potongan PPh
TotalDebit  uint64 // Total yang dipotong dari saldo user
}

// CalculateCompliance menghitung pajak secara absolut di tingkat Layer 1
func CalculateCompliance(amountAtomic uint64) TaxAssessment {
amountFloat := float64(amountAtomic)

feeFloat := amountFloat * PlatformFee
vatFloat := feeFloat * PPNRate
cryptoTaxFloat := amountFloat * CryptoTaxRate

totalDebit := amountFloat + feeFloat + vatFloat + cryptoTaxFloat

return TaxAssessment{
Principal:  amountAtomic,
Fee:        uint64(feeFloat),
VAT:        uint64(vatFloat),
CryptoTax:  uint64(cryptoTaxFloat),
TotalDebit: uint64(totalDebit),
}
}

// PrintReceipt mencetak log transparan untuk auditor publik
func (t *TaxAssessment) PrintReceipt() {
fmt.Printf("--- L1 COMPLIANCE RECEIPT ---\n")
fmt.Printf("Principal : %d\n", t.Principal)
fmt.Printf("L1 Fee    : %d\n", t.Fee)
fmt.Printf("VAT (PPN) : %d\n", t.VAT)
fmt.Printf("PPh (0.1%): %d\n", t.CryptoTax)
fmt.Printf("Total     : %d\n", t.TotalDebit)
fmt.Printf("-----------------------------\n")
}
