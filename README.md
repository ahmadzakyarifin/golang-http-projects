# 📂 Golang HTTP Projects: Belajar CRUD & Arsitektur

Halo! 👋

Repositori ini adalah tempat saya mendokumentasikan proses belajar saya dalam membangun aplikasi backend menggunakan **Golang**.

Di sini saya bereksperimen membuat aplikasi CRUD (*Create, Read, Update, Delete*) dengan dua pendekatan berbeda untuk memahami perbedaannya:
1.  **Web Tradisional:** Menggunakan HTML Template.
2.  **REST API Modern:** Menggunakan JSON & Clean Architecture sederhana.

> **📝 Catatan:**
> Kode di sini adalah hasil latihan dan eksplorasi saya. Mungkin belum sempurna, tapi ini adalah bukti pemahaman saya terhadap konsep-konsep dasar backend development.

---

## 🗂️ Apa yang Saya Pelajari?

Repositori ini dibagi menjadi dua project utama:

### 1. HTML CRUD Basic (`/html-crud-basic`)
Di sini saya belajar cara membuat web sederhana di mana server langsung mengirimkan halaman HTML (*Server-Side Rendering*).
* **Fokus:** Membuat CRUD simpel yang bisa dilihat langsung di browser.
* **Teknologi:** `html/template`, Form Handling.
* **Struktur:** Lebih sederhana, semua logic digabung agar mudah dipahami di awal.

### 2. JSON CRUD Basic (`/json-crud-basic`)
Di sini saya menantang diri untuk membuat backend yang lebih rapi dan terstruktur (*REST API*).
* **Fokus:** Membuat API yang merespons dengan data JSON (bukan HTML).
* **Clean Architecture Simple:** Saya belajar memisahkan kode menjadi layer-layer agar tidak berantakan:
    * **Handler:** Mengurus request HTTP.
    * **Service:** Tempat logika bisnis.
    * **Repository:** Tempat query database.
    * **DTO:** Memfilter data yang keluar/masuk.

---

## 🛠️ Cara Menjalankan

Silakan masuk ke folder masing-masing project untuk melihat detail cara menjalankannya.

```bash
# Untuk melihat web HTML
cd html-crud-basic
go run main.go

# Untuk melihat API JSON
cd json-crud-basic
go run main.go