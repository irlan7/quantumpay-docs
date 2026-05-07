# 🌌 QuantumPay Official Documentation (QTM)
**The Sovereign Post-Quantum Layer-1 Blockchain for the Global Economy.**

[![QuantumPay Mainnet](https://img.shields.io/badge/Network-Mainnet_Live-success?style=for-the-badge)](#)
[![Version](https://img.shields.io/badge/Release-v1.0.0--core-orange?style=for-the-badge)](https://github.com/irlan7/quantumpay-docs/releases/tag/v1.0.0-core)
[![Post-Quantum](https://img.shields.io/badge/Security-Kyber1024_%2B_Dilithium3-purple?style=for-the-badge)](#)
[![License](https://img.shields.io/badge/License-MIT-gray?style=for-the-badge)](#)

QuantumPay adalah infrastruktur blockchain L1 pertama dari Asia Tenggara yang dirancang khusus untuk mengamankan **Aset Dunia Nyata (RWA)** dan kedaulatan digital nasional. Menggunakan arsitektur hibrida mutakhir (**Go, Rust, dan Solidity**), QuantumPay siap menyerap likuiditas global melalui *Omnichain Gateway*.

---

## 🌐 Network Identity (Single Source of Truth)
Parameter resmi jaringan untuk verifikasi sinkronisasi L1:

| Parameter | Value |
| :--- | :--- |
| **Chain ID** | `77077` [FROZEN] |
| **Genesis Fingerprint** | `0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a` |
| **State Root** | `0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a` |
| **Network Phase** | `Mainnet-Alpha v2.0 (GEN 4.5 ACTIVE)` |
| **Binary Release** | `v1.0.0-core (Linux AMD64)` |

---

## 🛠️ Run a Validator Node (Quickstart v1.0.0)

Kami merekomendasikan penggunaan *binary executable* resmi untuk menjamin integritas konsensus dan performa maksimal. Jaringan saat ini dibatasi maksimal **100 Validator Global**.

### 1. Persyaratan Perangkat Keras (Minimum)
* **CPU:** 2 Cores (Intel/AMD)
* **RAM:** 4GB (Optimized usage: ~5%)
* **OS:** Ubuntu 22.04 LTS / 24.04 LTS

### 2. Langkah Instalasi (Binary Method)
Unduh paket mesin L1 QuantumPay langsung dari rilis resmi:

```bash
# 1. Unduh Binary Tarball
wget [https://github.com/irlan7/quantumpay-docs/releases/download/v1.0.0-core/qtm-validator-linux-amd64.tar.gz](https://github.com/irlan7/quantumpay-docs/releases/download/v1.0.0-core/qtm-validator-linux-amd64.tar.gz)

# 2. Ekstrak Peti Kemas
tar -xvzf qtm-validator-linux-amd64.tar.gz

# 3. Masuk ke direktori
cd qtm-validator

# 4. Konfigurasi & Jalankan Mesin
./start-validator.sh
3. Verifikasi Keaslian Genesis
Gunakan perintah ini untuk memastikan Node Anda terhubung ke rantai (Gen 4.5) yang sah:

Bash
pm2 logs qp-node --lines 100 | grep "Genesis Block Created"
# Expected: 0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a
🏗️ Hybrid Architecture & Ecosystem
Core Ledger: Arsitektur ganda menggunakan efisiensi PebbleDB dan analitik SQL tingkat lanjut dari CockroachDB.

Smart Contracts: Kompatibilitas penuh dengan EVM (Solidity) untuk penerbitan token QRC-20.

Omnichain Gateway: Mendukung aset institusional termasuk BTC, ETH, SOL, XRP, USDT, USDC, dan XLM.

🏦 Institutional & Exchange Partners
QuantumPay didukung oleh entitas hukum resmi di Indonesia (PT) dan dirancang untuk kepatuhan Bappebti.

API/RPC Integration: Silakan merujuk pada direktori rpc/ untuk dokumentasi endpoint bursa.

QuantumPay Foundation | Bridging Institutional Finance and Web3 Liquidity.

🌐 Website: quantumpaychain.org

📧 Contact: contact@quantumpaychain.org (Partnership & VC)

𝕏 Twitter: @quantumpaychain

Copyright © 2026 QuantumPay Network. All Rights Reserved.
