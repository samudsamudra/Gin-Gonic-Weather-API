# Gin Gonic Weather API

**Gin Gonic Weather API** adalah proyek RESTful API yang dibangun menggunakan bahasa pemrograman **Golang** dan framework **Gin Gonic**. Proyek ini menyediakan fitur untuk mendapatkan informasi cuaca dari layanan OpenWeatherMap dan memungkinkan pengguna untuk menyimpan lokasi favorit.

## Fitur
- **Informasi Cuaca**: Mendapatkan data cuaca real-time berdasarkan nama kota.
- **Lokasi Favorit**: Menyimpan dan mengambil lokasi favorit pengguna.
- **Sistem Login**: Menggunakan autentikasi JWT untuk memastikan keamanan data pengguna.

## Teknologi yang Digunakan
- **Golang** dengan framework **Gin Gonic**.
- **MySQL** sebagai basis data untuk menyimpan data pengguna dan lokasi favorit.
- **OpenWeatherMap API** untuk data cuaca.

## Cara Menjalankan
1. **Clone repository ini**:
   ```bash
   git clone https://github.com/samudsamudra/Gin-Gonic-Weather-API
   cd gin-gonic-weather-api


   Setup database:

Buat database MySQL dengan nama weather_api.
Jalankan script berikut untuk membuat tabel:
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE favorites (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    location VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
Pasang dependensi:

go mod tidy
Jalankan aplikasi:

go run main.go
Dokumentasi Endpoint
Register: Mendaftarkan pengguna baru.

POST /register
Body:
{
  "username": "user1",
  "password": "password123"
}
Login: Login dan mendapatkan token JWT.

POST /login
Body:
{
  "username": "user1",
  "password": "password123"
}
Cuaca: Mendapatkan informasi cuaca.

GET /weather?location=Malang
Header:
Authorization: Bearer <jwt-token>
Lokasi Favorit:

Tambah: POST /favorites Body:

{
  "location": "Malang"
}
Lihat Semua: GET /favorites


