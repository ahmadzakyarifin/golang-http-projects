# JSON CRUD API dengan Golang
**(Handler – Service – Repository Architecture)**

---

## 📖 Pendahuluan

Project ini dibuat sebagai **latihan dan fondasi pemahaman backend modern** menggunakan **Golang** dengan pendekatan **REST API berbasis JSON**, bukan HTML templating.

README ini bukan sekadar dokumentasi teknis, tapi **menjelaskan ALASAN (The "Why")** di balik setiap keputusan arsitektur. Tujuannya agar tidak sekadar menghafal kode, tapi memahami pola pikir *Software Engineer*.

---

## 1. JSON vs HTML Template
**Kenapa project ini pakai JSON, bukan HTML templating?**

### 🏪 Analogi: Cara Penyajian Makanan
Bayangkan kamu punya restoran.

**A. HTML Template (Monolith)**
* Kamu masak di dapur.
* Makanan langsung disajikan di piring restoran ke meja nomor 5.
* **Keterbatasan:** Pelanggan harus makan di tempat. Kamu tidak bisa kirim makanan ini ke orang di luar restoran.
* **Teknis:** Backend dan UI menyatu.

**B. JSON API (Decoupled)**
* Kamu masak di dapur.
* Makanan dikemas rapi dalam kotak (JSON).
* **Kebebasan:** Kotak ini bisa dimakan di restoran, dibawa pulang (Takeaway), dikirim lewat Ojol (Mobile App), atau dikirim ke kantor sebelah (Integration).
* **Teknis:** Backend hanya mengurus data. UI bebas menggunakan apa saja (React, Vue, Android, iOS).

---

## 2. Apa itu REST API & HTTP Method?

REST API adalah standar komunikasi antar sistem. Kita menggunakan **HTTP Method** sebagai kata kerja untuk memberi tahu server apa yang harus dilakukan.

### 🏢 Analogi: Kantor Administrasi

| HTTP Method | Aksi di Kantor | Fungsi Teknis |
|:-----------:|----------------|---------------|
| **GET** | Melihat papan pengumuman | Ambil data |
| **POST** | Mengisi formulir pendaftaran baru | Buat data baru |
| **PUT** | Mengganti seluruh isi formulir lama | Update seluruh data |
| **PATCH** | Mengoreksi salah ketik saja | Update sebagian data |
| **DELETE** | Menyobek/membakar dokumen | Hapus data |

**Kenapa HTML biasa cuma punya GET & POST?**
Karena browser secara standar hanya mendukungan "Lihat Halaman" (GET) dan "Kirim Form" (POST). Namun, JSON API tidak terikat browser, jadi kita bisa menggunakan semua metode di atas agar lebih spesifik.

---

## 3. Routes (Jalur Akses)
**Kenapa routes dipisah, tidak langsung di `main.go`?**

### 🏠 Analogi: Panel Listrik Rumah
Bayangkan sebuah rumah besar.

* **Tanpa Routes Terpisah:** Semua kabel dari kulkas, TV, AC, dan Lampu disambung langsung ke satu colokan di ruang tamu (`main.go`). Kabel jadi ruwet, rawan korsleting, dan susah diperbaiki.
* **Dengan Routes:** Setiap ruangan punya panel sendiri. Panel dapur, panel kamar, panel ruang tamu. Jika lampu kamar mati, kita tahu panel mana yang harus dicek tanpa mematikan kulkas.

---

## 4. Arsitektur: Handler – Service – Repository
**Ini adalah jantung dari project ini. Kenapa kode harus dibagi tiga? Kenapa tidak satu file saja?**

Mari kita gunakan analogi **Operasional Restoran Bintang Lima**.

### 🍽️ The Story: "Restoran Bintang Lima vs Warung Kaki Lima"

Jika kita menulis semua kode di `main.go`, itu seperti **Warung Kaki Lima**. Satu orang melakukan semuanya: mencatat pesanan, memasak, belanja ke pasar, dan mencuci piring. Cepat untuk skala kecil, tapi "pingsan" jika pelanggan membludak.

Project ini menggunakan gaya **Restoran Profesional**:

#### 1. Handler = Pelayan (Waiter)
* **Tugas:** Menghadapi Tamu (Client/Frontend).
* **Kerjaannya:**
    * Menerima pesanan (Request).
    * Cek apakah pesanan masuk akal (Validasi dasar: "Maaf Pak, kami tidak jual ban mobil").
    * Meneruskan pesanan ke dapur.
    * Mengantar makanan jadi ke meja (Response JSON).
* **Aturan:** Pelayan **dilarang** masuk dapur untuk ikut memotong bawang. Tugasnya hanya melayani tamu.

#### 2. Service = Koki Kepala (Head Chef)
* **Tugas:** Mengolah bahan menjadi hidangan (Business Logic).
* **Kerjaannya:**
    * Menerima bon pesanan dari Pelayan.
    * Menjalankan resep: "Nasi Goreng = Nasi + Telur + Kecap".
    * Jika stok habis, dia berteriak ke gudang: "Ambilkan telur!".
* **Aturan:** Koki tidak perlu tahu di mana telur dibeli (Pasar A atau B). Dia hanya tahu cara memasak.

#### 3. Repository = Petugas Gudang (Storekeeper)
* **Tugas:** Mengambil dan menyimpan barang (Database Access).
* **Kerjaannya:**
    * Hapal mati letak semua bahan (SQL Query).
    * Saat Koki minta telur, dia ambilkan dari rak database.
* **Aturan:** Petugas gudang **dilarang** menemui tamu. Bayangkan kalau orang gudang yang kotor langsung menyajikan makanan? (Security risk).

---

## 5. DTO vs Model
**Kenapa struktur datanya ada dua? Kelihatannya mirip, kenapa harus dipisah?**

### 📦 The Story: "Paket Belanja Online"

Bayangkan kamu membeli HP dari toko online.

#### A. Model (Barang di Gudang)
Ini adalah bentuk data asli di Database (Gudang). HP di gudang punya banyak tempelan:
* Stiker *barcode* gudang.
* Catatan modal harga beli (Rahasia dapur!).
* Catatan "Barang Retur".
* Kolom `Password_Hash`, `Created_At`, `Is_Deleted`.

#### B. DTO (Paket yang Dikirim)
Ini adalah data yang dikirim ke Client (Pembeli). Sebelum dikirim, barang gudang harus di-**mapping**:
* Stiker gudang dilepas.
* Catatan modal dibuang (supaya pembeli tidak tahu untung kita).
* Dibungkus rapi.
* Hanya menyisakan: Unit HP dan Harga Jual.

**Kesimpulan:**
* **Model:** Struktur tabel database (isi lengkap + sensitif).
* **DTO (Data Transfer Object):** Struktur JSON response (isi bersih + aman).
* Kita memisahkan keduanya agar **password user, field internal, dan data sensitif tidak bocor ke publik.**

---

## 6. Config & Environment
**Kenapa config dipisah?**
Agar kita bisa mengganti setting tanpa mengubah kode.
* Di Laptop (Local): Pakai DB `localhost`.
* Di Server (Production): Pakai DB `192.168.x.x`.
Dengan memisahkan config, kita cukup ubah file `.env`, tidak perlu bongkar kodingan inti.

---

## 7. Apa itu Mock?
**Mock = Database Palsu.**

Bayangkan Koki (Service) sedang latihan resep baru. Dia tidak perlu menunggu Petugas Gudang (Repo) beneran datang. Dia bisa latihan pakai bahan mainan dulu.
* Membuat testing jadi **cepat** (tidak perlu connect DB asli).
* **Aman** (data asli tidak terhapus saat testing).

---

## 8. Catatan Teknis Kecil

### Kenapa `LastInsertId()` penting?
Saat pelanggan mengambil nomor antrian, nomor itu adalah identitasnya. `LastInsertId()` mengambil ID yang baru saja dibuat oleh database (AUTO_INCREMENT) agar kita bisa mengembalikan data lengkap ke user segera setelah mereka create data.

### Kenapa DELETE tidak mengembalikan data?
Karena datanya sudah hilang. Cukup berikan konfirmasi "Berhasil dihapus". Mengirimkan balik data yang baru saja dimusnahkan itu mubazir bandwidth.

---
