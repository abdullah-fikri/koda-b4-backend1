# REST API Gin

Project ini adalah contoh sederhana implementasi **CRUD (Create, Read, Update, Delete)** menggunakan framework **Gin** di bahasa pemrograman Go.

API ini menyimpan data user (tanpa database), dan memanfaatkan penggunaan:

- `GET` untuk mengambil data
- `GET by id` untuk mengambil data by id
- `POST` untuk menambah data
- `PATCH` untuk mengupdate data
- `DELETE` untuk menghapus data

---

## 🚀 Cara Menjalankan

1. GET /users

Menampilkan seluruh data user.

Request

GET /users HTTP/1.1
Host: localhost:8081

Response

{
"success": true,
"message": "Nama list",
"data": [
{ "id": "1", "name": "Fiki" },
{ "id": "2", "name": "Anggi" }
]
}

2. GET /users/:id

Menampilkan data user berdasarkan ID.

Request

GET /users/1 HTTP/1.1
Host: localhost:8081

Response

{
"success": true,
"message": "User ditemukan",
"data": {
"id": "1",
"name": "Fiki"
}
}

3. POST /users

Menambahkan data user baru.

Request

POST /users HTTP/1.1
Host: localhost:8081
Content-Type: application/x-www-form-urlencoded

name=fiki3&id=3

4. PATCH /users/:id

Mengupdate data user berdasarkan ID.

Request

PATCH /users/2 HTTP/1.1
Host: localhost:8081
Content-Type: application/x-www-form-urlencoded

name=fiki122&id=2

5. DELETE /users/:id

Menghapus user berdasarkan ID.

Request

DELETE /users/1 HTTP/1.1
Host: localhost:8081
