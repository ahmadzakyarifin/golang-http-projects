# Go HTTP CRUD (Modular Architecture)

Project ini adalah implementasi aplikasi web sederhana untuk manajemen data Mahasiswa (CRUD) menggunakan **Golang** (Native `net/http`), **MySQL**, dan **HTML template (TAILWIND CSS)**

Aplikasi ini tidak sekadar "jalan", tetapi dibangun dengan **Pola Arsitektur Terpisah (Layered Architecture)** untuk mensimulasikan standar industri yang rapi, mudah dites, dan mudah dikembangkan.

---

## 🧠 Konsep & Alasan Arsitektur

Project ini menghindari penulisan kode "Spaghetti" (semua numpuk di `main.go`). Berikut alasan logis di balik struktur folder ini:

### 1. Kenapa Database dipisah? (`/database`)
* **Alasan Utama:** *Testability & Mocking*.
* **Penjelasan:** Dengan memisahkan inisialisasi koneksi database dari logika bisnis, kita bisa melakukan **Unit Testing** tanpa harus menyentuh database asli. Kita bisa menggantinya dengan database palsu (Mock) saat testing agar tes berjalan cepat dan aman.

### 2. Kenapa pakai Model/Struct? (`/model`)
* **Alasan Utama:** *Scalability & Consistency*.
* **Penjelasan:** Kita **TIDAK** menggunakan variabel satu per satu di handler (contoh: `func(nama, nim, alamat, tgl_lahir, ...)`). Itu akan menyulitkan jika data bertambah.
* Dengan `Struct`, kita membungkus data dalam satu "wadah" objek. Jika ada penambahan kolom baru (misal: `NoHP`), kita cukup update Struct-nya saja tanpa merusak fungsi-fungsi yang menggunakannya.

### 3. Kenapa Handler & Routes dipisah?
* **Routes (`/routes`):** Bertindak sebagai "Resepsionis" yang mengatur url mana masuk ke meja mana. Jika ingin ganti URL endpoint, cukup ubah di sini tanpa mengganggu logika.
* **Handler (`/handler`):** Bertindak sebagai "Pelayan". Fokus hanya menerima input, memproses ke database, dan mengembalikan HTML/JSON. Tidak perlu memikirkan konfigurasi server.

---

## 🔮 Pengembangan Selanjutnya (To-Do)

Saat ini logic SQL masih menyatu di dalam Handler. Untuk pengembangan ke depan agar lebih *Clean Architecture*:

* [ ] Implementasi **Repository Pattern** (Memisahkan Query SQL dari Handler).
* [ ] Menambahkan **Interface** untuk layer Service/Repository agar lebih mudah di-Mocking.
* [ ] Implementasi Unit Testing.