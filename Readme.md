# Sistem Pelacakan Aktivitas Pengguna (User Activity Tracking System)

Dokumen ini berfungsi sebagai panduan teknis dan dokumentasi untuk penyelesaian Tes Teknis Backend Engineer. Proyek ini bertujuan untuk membangun sistem pelacakan aktivitas pengguna yang mampu menangani lalu lintas tinggi, menjamin konsistensi data, dan menerapkan standar keamanan API.

Sesuai dengan persyaratan bonus, solusi ini diimplementasikan dalam dua bahasa pemrograman yang berbeda untuk menunjukkan fleksibilitas dan kedalaman teknis, yaitu **Go (Golang)** dan **TypeScript (NestJS)**.
## Tinjauan Arsitektur

Sistem ini dirancang sebagai layanan backend RESTful yang mencatat dan menganalisis penggunaan API oleh klien. Sistem ini menerapkan strategi caching untuk performa tinggi dan mekanisme pemrosesan batch untuk penulisan log ke basis data guna mengurangi beban pada operasi penulisan sinkron.

## Teknologi yang Digunakan

Proyek ini memanfaatkan teknologi berikut:

- **Bahasa Pemrograman:**
    - Go (Versi 1.20+) dengan Framework Fiber.
    - TypeScript (Node.js) dengan Framework NestJS.
- **Basis Data:** PostgreSQL.
- **Cache & Antrian:** Redis (digunakan untuk caching respons dan pembatasan laju permintaan).
- **Kontainerisasi:** Docker dan Docker Compose.
- **Pengujian Beban:** k6.

## Fitur Utama

Implementasi ini mencakup persyaratan inti dan persyaratan bonus sebagai berikut:

- **Manajemen Klien:**
    - Pendaftaran klien baru dengan pembuatan kunci API (API Key) yang aman.
    - Enkripsi data sensitif pada penyimpanan.
- **Pencatatan Aktivitas (Logging):**
    - Endpoint untuk mencatat aktivitas API (hit).
    - Pemrosesan asinkron untuk penulisan log guna performa tinggi.
- **Analitik Penggunaan:**
    - Pengambilan data total permintaan harian per klien (7 hari terakhir).
    - Peringkat 3 klien teratas berdasarkan permintaan dalam 24 jam terakhir.
    - Implementasi caching pada endpoint analitik untuk mengurangi beban basis data.
- **Keamanan:**
    - **Autentikasi JWT:** Token JWT digunakan untuk mengamankan endpoint sensitif.
    - **Pembatasan Laju (Rate Limiting):** Mencegah penyalahgunaan API dengan membatasi jumlah permintaan per klien.
    - **Daftar Putih IP (IP Whitelisting):** (Fitur Bonus) Membatasi akses endpoint tertentu hanya untuk alamat IP yang terdaftar.
    - **Perlindungan SQL Injection:** Menggunakan parameterisasi kueri (pada Go) dan ORM Prisma (pada NestJS).

## Struktur Direktori

- backend_go/: Implementasi solusi menggunakan bahasa Go dan framework Fiber.
- backend_nest/: Implementasi solusi menggunakan TypeScript dan framework NestJS.
- k6/: Skrip pengujian beban menggunakan k6.
- openapi.yaml: Spesifikasi OpenAPI untuk dokumentasi endpoint.
- collection.json: Koleksi Postman untuk pengujian API.

## Prasyarat

Sebelum menjalankan aplikasi, pastikan perangkat Anda telah terinstal:

- Docker dan Docker Compose.
- Go (jika menjalankan manual).
- Node.js dan NPM (jika menjalankan manual).
- PostgreSQL dan Redis (jika tidak menggunakan Docker Compose).

## Panduan Instalasi dan Menjalankan Aplikasi

Anda dapat memilih untuk menjalankan implementasi Go atau NestJS. Disarankan menggunakan Docker untuk kemudahan pengaturan lingkungan.

### Persiapan Variabel Lingkungan

Salin berkas contoh konfigurasi di masing-masing direktori proyek:

**Untuk Go:**

cp backend_go/.env.example backend_go/.env

**Untuk NestJS:**

cp backend_nest/.env.example backend_nest/.env

Sesuaikan nilai di dalam berkas .env dengan konfigurasi lokal Anda jika diperlukan.

### Menggunakan Docker (Disarankan)

Setiap proyek memiliki berkas docker-compose.yml masing-masing.

**Menjalankan Versi Go:**

cd backend_go  
docker-compose up --build -d

**Menjalankan Versi NestJS:**

cd backend_nest  
docker-compose up --build -d

Aplikasi akan berjalan pada port yang didefinisikan dalam variabel lingkungan (standar: 3000 atau 8080).

### Menjalankan Layanan Go Secara Manual

- Pastikan PostgreSQL dan Redis sudah berjalan.
- Masuk ke direktori: cd backend_go
- Unduh dependensi: go mod tidy
- Jalankan aplikasi: go run cmd/server/main.go
- Atau gunakan Air untuk pengembangan (hot-reload): air

### Menjalankan Layanan NestJS Secara Manual

- Pastikan PostgreSQL dan Redis sudah berjalan.
- Masuk ke direktori: cd backend_nest
- Instal dependensi: npm install
- Lakukan migrasi basis data: npx prisma migrate dev
- Jalankan aplikasi: npm run start:dev

## Dokumentasi API

Dokumentasi lengkap mengenai endpoint, parameter permintaan, dan format respons tersedia dalam berkas openapi.yaml di direktori utama.
[POSTMAN API DOCS](https://documenter.getpostman.com/view/38037610/2sBXVckXmV)

Anda juga dapat mengimpor berkas collection.json ke dalam aplikasi Postman untuk menguji endpoint secara langsung.

**Daftar Endpoint Utama:**

- POST /api/register: Mendaftarkan klien baru.
- POST /api/logs: Mencatat aktivitas API.
- GET /api/usage/daily: Mendapatkan statistik harian.
- GET /api/usage/top: Mendapatkan klien dengan aktivitas tertinggi.

## Pengujian Beban (Load Testing)

Untuk memastikan sistem mampu menangani lalu lintas tinggi dan memverifikasi mekanisme pembatasan laju (rate limiting), skrip pengujian k6 telah disediakan.

Cara menjalankan pengujian:

- Pastikan k6 telah terinstal pada sistem Anda.
- Jalankan perintah berikut dari direktori root:

k6 run k6/rate-limit-test.js

Hasil pengujian akan menampilkan metrik performa, latensi, dan status keberhasilan permintaan untuk memvalidasi stabilitas sistem di bawah beban kerja.
