# QuantumPay (QTM) – Model Ekonomi Kedaulatan 50 Tahun

**Versi:** v10.8 (Pembaruan Umur Panjang - Dikunci)  
**Jaringan:** QuantumPay Mainnet  
**ID Rantai (Chain ID):** 77077  
**Status:** Mainnet Aktif – Ekonomi & Suplai Maksimal Dibekukan (*Frozen*)  

---

## 1. Ringkasan Eksekutif
QuantumPay (QTM) adalah koin kedaulatan asli (*native sovereign coin*) dari blockchain Layer-1 QuantumPay. Koin ini direkayasa secara fundamental bukan sekadar sebagai alat tukar biasa, melainkan sebagai **Infrastruktur Digital Tingkat Institusi** yang dirancang untuk bertahan melampaui siklus ekonomi 50 tahun ke depan.

Fungsi mutlak QTM di dalam jaringan meliputi:
* **Alat Tukar** untuk transaksi yang terdesentralisasi dan berdaulat.
* **Biaya Gas (*Gas Fee*)** untuk komputasi jaringan dan eksekusi *Smart Contract* / QRC-20.
* **Mekanisme Perutean Pajak L1** (Pemotongan otomatis PPN & PPh kripto melalui *Treasury Split*).
* **Aset Jaminan (*Staking Asset*)** untuk konsensus *Permissionless BFT* & *Proof-of-Stake* yang berkeamanan tinggi.
* **Insentif Ekonomi** untuk validator yang mendistribusikan keamanan secara geografis.

---

## 2. Spesifikasi Token
Berbeda dengan token EVM standar, QTM menggunakan presisi atomik 8-desimal (Arsitektur Anti-Kehilangan Dana) agar selaras sempurna dengan standar akuntansi mata uang fiat dan sistem *Central Bank Digital Currency* (CBDC).

| Parameter | Nilai Ketetapan |
| :--- | :--- |
| **Simbol** | QTM |
| **Desimal** | 8 (Presisi Atomik) |
| **Suplai Maksimal** | 210.000.000 QTM (Batas Keras / *Hardcapped*) |
| **Tipe Emisi** | Kurva Deflasioner *Halving* (~2,85 Tahun) |
| **Penjualan Token** | Tidak Ada (Tanpa ICO, IEO, IDO, atau Penjualan Publik) |

---

## 3. 5 Pilar Distribusi Kedaulatan (Suplai Maksimal 210 Juta)
Untuk memastikan jaringan ini kuat secara politik dan ekonomi, batas absolut 210.000.000 QTM dibagi menjadi 5 Pilar Kedaulatan. 

Pada saat **Genesis (Blok 0)**, 120.000.000 QTM dicetak untuk didistribusikan ke Pilar 2, 3, 4, dan 5. Sisa 90.000.000 QTM (Pilar 1) dikunci secara matematis di dalam protokol untuk ditambang secara perlahan selama 50 tahun.

| Pilar | Alokasi | Jumlah (QTM) | Tujuan & Eksekusi |
| :--- | :--- | :--- | :--- |
| **1. Validator & Keamanan Jaringan** | 42,8% | 90.000.000 | Hadiah penambangan algoritmik selama 50 tahun. Dicetak secara ketat melalui konsensus (Awal 2.5 QTM, dibagi dua setiap 18 Juta blok). |
| **2. Cadangan Kas Negara & Institusi** | 19,0% | 40.000.000 | Dicadangkan untuk Bank Sentral, BUMN, dan Konsorsium Institusi untuk menjalankan *Sovereign Nodes* dan bantalan likuiditas aset CBDC/RWA. |
| **3. Komunitas & Ekosistem Publik** | 23,8% | 50.000.000 | Mendorong adopsi akar rumput melalui hibah *Developer*, insentif *Merchant*, subsidi lintas batas, dan *Airdrop* komunitas. |
| **4. Tim Inti & Yayasan Pendiri** | 9,5% | 20.000.000 | Didedikasikan untuk pendiri dan insinyur inti. **Dikunci ketat (*Vested*) selama 5-10 tahun** via *Smart Contract* untuk menjamin keamanan pasar. |
| **5. Likuiditas Bursa & Jembatan (*Bridge*)** | 4,9% | 10.000.000 | Penyediaan likuiditas awal untuk menembus Bursa Global Tier-1 (CEX) dan jembatan lintas rantai (*Cross-chain EVM*). |

---

## 4. Kebijakan Moneter & Siklus Halving 50 Tahun
QuantumPay menerapkan jadwal emisi deflasioner yang presisi dan dapat diprediksi untuk menjamin umur jaringan dan kelangkaan absolut koin.

* **Hadiah Blok Awal:** 2.5 QTM per blok.
* **Waktu Blok:** ~5 Detik.
* **Interval Halving:** Setiap 18.000.000 blok (Sekitar 2,85 tahun).
* **Umur Panjang Jaringan:** Kurva emisi akan berhenti secara matematis tepat sebelum menyentuh 90.000.000 QTM pada Tahun ke-50. Setelah itu, Validator akan hidup murni dari *Gas Fee* Jaringan dan Pendapatan Pajak Protokol L1.

---

## 5. Utilitas L1 Unik: Mesin Pajak & Kepatuhan
QTM adalah koin Layer-1 pertama di dunia yang memiliki sistem **Pemisahan Kas (*Treasury Split*)** dan **Hak Veto Kedaulatan** yang ditanamkan langsung secara *native*:
* **Perpajakan Otomatis:** Jaringan mengeksekusi pemisahan atomik saat transaksi terjadi, merutekan biaya operasional (misal 0,5%) ke Kas Platform, dan menghitung pajak negara (misal PPN 11%, PPh Kripto 0,21%) langsung ke Dompet Pajak Negara.
* **Resistensi Sensor vs. Kepatuhan Hukum:** Koin ritel/publik (seperti *Utility/Memecoins*) dijamin 100% kebal dari sensor. Namun, jaringan menyediakan fungsi `FREEZE` bawaan yang khusus ditujukan untuk aset berlabel `REGULATED` (seperti CBDC atau Emas Digital RWA) demi mematuhi otoritas nasional.

---

## 6. Ekosistem Validator & Penalti
Validator adalah tulang punggung pertahanan jaringan QuantumPay.
* **Staking:** Dibutuhkan minimal jaminan (misal: 10.000 QTM) untuk mengoperasikan *Full Validator Node*.
* **Hadiah (*Rewards*):** Didistribusikan secara proporsional berdasarkan bobot jaminan, waktu *online* (uptime), dan kejujuran dalam mencetak blok.
* **Pemotongan Paksa (*Slashing*):** Jaringan mempekerjakan robot patroli otomatis. Validator akan menghadapi penyitaan QTM yang dijaminkan jika terbukti melakukan:
  * Tanda tangan ganda (*Double signing*) / Menciptakan *Fork* palsu.
  * Mati / *Offline* dalam jangka waktu yang lama tanpa pemberitahuan.

---

## 7. Posisi Regulasi & Penyangkalan Hukum (*Disclaimer*)
* **Tanpa Penjualan Token:** QTM tidak pernah dijual melalui skema pengumpulan dana publik (ICO).
* **Tanpa Janji Keuntungan:** QTM secara murni adalah **koin utilitas jaringan** yang diwajibkan untuk membayar komputasi jaringan (*gas*), mengeksekusi *smart contract*, dan mengamankan L1.
* **Kepatuhan Institusi:** Arsitektur ini dirancang untuk mematuhi praktik terbaik tata kelola *blockchain* global dan menyediakan wadah uji coba (*sandbox*) yang aman dan transparan bagi aset digital terbitan pemerintah.
* *Sangkalan: QuantumPay (QTM) bukanlah produk investasi, sekuritas, atau instrumen keuangan.*

***Akhir dari ECONOMICS_ID.md***
