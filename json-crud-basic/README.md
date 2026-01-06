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

# Architecture Design & Pattern Explanation

Dokumen ini menjelaskan alasan dibalik keputusan arsitektur yang digunakan dalam project ini (Repository, Service, dan Handler pattern), standar implementasi teknis, serta alur data (*data flow*).

## 1. Repository & Service Pattern
Pada layer Repository dan Service, kita menggunakan pendekatan **Interface** dan **Struct** yang dihubungkan dengan **Receiver Function**.

### Mengapa Menggunakan Interface?
Ketika `Main` atau layer di atasnya memanggil Repository/Service, kita mengembalikan sebuah **Interface**, bukan Struct secara langsung.

* **Abstraction (Abstraksi):** Agar pemanggil (*caller*) hanya mengetahui **"apa"** yang bisa dilakukan (method apa yang tersedia), tanpa perlu tahu **"bagaimana"** caranya dilakukan (detail implementasi).
* **Decoupling:** Memudahkan penggantian implementasi di masa depan tanpa merusak kode di layer lain.
* **Testability:** Memudahkan proses *Unit Testing* karena Interface mudah di-*mock*.

### Mengapa Mengimplementasikan dengan Struct?
Meskipun interface bisa diimplementasikan oleh tipe data apa saja, *best practice*-nya adalah menggunakan **Struct**.

* **Dependency Injection:** Kita membutuhkan tempat untuk menyimpan dependency eksternal, seperti koneksi Database (`db *sql.DB`) atau library pihak ketiga. Struct memungkinkan kita menyimpan *state* ini sebagai *field*.
* **Receiver Function:** Di Go, untuk memenuhi kontrak Interface, kita menempelkan method-method tersebut ke Struct menggunakan *Receiver Function*.

> **Pola Implementasi:**
> 1. `Main` menginjeksi DB ke `Struct`.
> 2. `Struct` menyimpan DB tersebut.
> 3. `Struct` memiliki method (via `Receiver Function`) yang menggunakan DB tersebut.
> 4. Method tersebut memenuhi syarat `Interface`.

## 2. Handler Pattern
Berbeda dengan Repository dan Service, **Handler** (Controller) berada di layer terluar. Umumnya Handler tidak wajib mengembalikan Interface, karena:

* **Entry Point:** Handler tidak di-*inject* ke layer lain, melainkan dipanggil langsung oleh Router/Main.
* **Consumer:** Tugas utama Handler adalah mengonsumsi Service. Ia bertugas menerima request, memvalidasi input, dan memformat response.

## 3. Application Flow (Alur Data)

Alur aplikasi berjalan secara *top-down* dari request masuk hingga ke database:

1.  **Handler (Presentation Layer):**
    * Menerima HTTP Request masuk.
    * Melakukan *binding* dan memvalidasi **DTO** (Data Transfer Object).
    * Jika valid, data dikirim ke Service.
2.  **Service (Business Logic Layer):**
    * Menerima data bersih dari Handler.
    * Menjalankan logika bisnis (kalkulasi, validasi logic, transaksi, dll).
    * Memanggil method yang ada di Interface Repository.
3.  **Repository (Data Access Layer):**
    * Menerima perintah dari Service.
    * Menggunakan koneksi DB (yang disimpan di Struct) untuk melakukan query ke database.
    * Mengembalikan hasil data ke Service -> Handler -> Client.

---

## Appendix: Analogi Konsep (The Restaurant)

Bagian ini bertujuan untuk membantu pemahaman konsep teknis di atas menggunakan analogi sebuah **Restoran Profesional**.

### Peran (Components)
* **Struct = "Koki & Tas Perkakasnya"**
    Identitas pekerja yang menyimpan alat-alat penting.
    *(Contoh: Struct Repository menyimpan koneksi DB ibarat Koki menyimpan Pisau & Kunci Gudang).*
* **Function New (Constructor) = "Kontrak & Serah Terima Alat"**
    Saat manager merekrut pegawai, ia memberikan alat kerjanya.
    *(Contoh: `NewRepository(db)` memberikan koneksi DB ke Struct agar siap kerja).*
* **Receiver Function = "Skill Menggunakan Alat"**
    Skill memotong daging hanya bisa dilakukan jika Koki sedang memegang pisau.
    *(Contoh: Method `FindUser()` menempel pada Struct karena butuh koneksi DB yang ada di dalam Struct).*
* **Interface = "Job Description"**
    Pemilik restoran hanya peduli "Siapa yang BISA MASAK", tidak peduli siapa nama orangnya.
    *(Contoh: Service memanggil Repository lewat Interface, agar fleksibel jika implementasi DB berubah).*

### Alur Kerja (Flow)
1.  **Handler = Pelayan (Waiter)**
    Menerima pesanan pelanggan, cek menu (validasi DTO). Pelayan tidak memasak, hanya meneruskan pesanan valid ke Dapur.
2.  **Service = Kepala Koki (Chef)**
    Menerima pesanan, memikirkan resep & racikan (Business Logic). Koki tidak mengambil bahan sendiri ke gudang, dia menyuruh staff gudang.
3.  **Repository = Penjaga Gudang**
    Menerima perintah ambil bahan. Dia punya kunci gudang (DB Connection). Dia membuka gudang, mengambil bahan (Query SQL), dan menyerahkannya ke Koki.
