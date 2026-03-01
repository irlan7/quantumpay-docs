# 💎 QuantumPay Core (QTM)
**Sovereign Layer-1 Hybrid Blockchain Infrastructure.**

QuantumPay adalah infrastruktur blockchain Layer-1 berkinerja tinggi yang dibangun untuk kedaulatan digital nasional dan skalabilitas global. Menggunakan arsitektur **Hybrid-Ledger**, QuantumPay menggabungkan kecepatan penulisan data pada **PebbleDB** dengan kecanggihan analitik **CockroachDB**.

## 🌐 Network Identity (SSoT)
Parameter ini mendefinisikan status resmi jaringan saat ini (Single Source of Truth):

| Parameter | Value |
| :--- | :--- |
| **Chain ID** | `77077` [FROZEN] |
| **Genesis Fingerprint** | `0x3e850ed3cce6d7ae604fcd11f5aff6983f4a620dcb48f03082a0259bdb499012` |
| **State Root** | `3e850ed3cce6d7ae604fcd11f5aff6983f4a620dcb48f03082a0259bdb499012` |
| **Core Engine** | `Go-Lang (quantumpay-go-v1.1)` |
| **Network Phase** | `Mainnet-Alpha v2.0 (Hybrid Active)` |

---

## 🏗️ Arsitektur Hybrid (Pebble + SQL)
Berbeda dengan blockchain tradisional, QuantumPay memisahkan jalur data untuk efisiensi maksimal:
1.  **Write Layer (PebbleDB):** Menangani konsensus dan finalitas blok instan (< 5 detik) dengan penggunaan RAM yang sangat hemat (~5%).
2.  **Read Layer (CockroachDB):** Menyediakan layer analitik SQL untuk Explorer, Wallet, dan Exchange tanpa membebani performa node utama.

---

## ⚡ Technical Features
* **High Efficiency:** Dioptimalkan untuk berjalan pada hardware standar dengan konsumsi RAM minimal (~24.7MB per node).
* **Fast Finality:** Konfirmasi transaksi instan di bawah 5 detik.
* **Process Management:** Mendukung **PM2** untuk uptime 24/7 dan pemantauan bridge database secara real-time.

---

## 🛠️ Run a Node (Join the Decentralization)

### Hardware Requirements
* **CPU:** 2 Cores (Minimum)
* **RAM:** 4GB (Optimized usage: ~5%)
* **Storage:** 40GB SSD
* **OS:** Ubuntu 22.04 LTS / 24.04 LTS

### Installation
1.  **Clone the Repository**
    ```bash
    git clone [https://github.com/irlan7/quantumpay-go.git](https://github.com/irlan7/quantumpay-go.git)
    cd quantumpay-go
    ```
2.  **Build the Node**
    ```bash
    go build -o qtm-core ./cmd/node
    ```
3.  **Start with PM2**
    ```bash
    pm2 start ./qtm-core --name "qp-node"
    ```

### Genesis Verification
Pastikan node Anda memiliki Fingerprint yang valid:
```bash
pm2 logs qp-node --lines 100 | grep "GENESIS FINGERPRINT"


Expected: 0x3e850ed3cce6d7ae604fcd11f5aff6983f4a620dcb48f03082a0259bdb499012

🏛️ Official Genesis Allocations
Founder (Irlan): 0x6d047da4f3AB9Dda7647D8ff901f65DDa6597040

Legacy V1 : 0xB766497a96d061887CeC4aAaCFBA25676a749061

Legacy V2: 0x1c83F44cca36cb423E30940571cFc81b0fEC9A81

📡 Official Channels
Website: quantumpaychain.org

X (Twitter): @quantumpaychain

Email: quantumpaysec@gmail.com

📜 License & Vision
Proyek ini bersifat open-source di bawah MIT License. Kami mengikuti visi Satoshi Nakamoto dan Vitalik Buterin: membangun dunia yang trustless, transparan, dan permissionless di mana setiap individu dapat menjalankan node dan memverifikasi kebenaran secara mandiri.

Copyright © 2026 QuantumPay - All Rights Reserved.
