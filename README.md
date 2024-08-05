# sanbercode-golang-batch-58-julia
# Book Management API

API ini dibuat menggunakan Go dan Gin Gonic untuk mengelola data buku dan kategori.

## Fitur

- Menambahkan buku baru
- Memperbarui buku berdasarkan ID
- Menghapus buku berdasarkan ID
- Mendapatkan daftar semua buku
- Mendapatkan detail buku berdasarkan ID
- Menambahkan kategori baru
- Mendapatkan daftar semua kategori
- Mendapatkan detail kategori berdasarkan ID

## Menjalankan Aplikasi

Aplikasi akan berjalan di http://localhost:8080.

- GET (/api/books)
    Mendapatkan Daftar Semua Buku.
- GET (/api/books/:id)
    Mendapatkan Detail Buku Berdasarkan ID
- POST (/api/books)
    Menambahkan Buku Baru
- PUT (/api/books/:id)
    Memperbarui Buku Berdasarkan ID
- DELETE (/api/books/:id)
    Menghapus Buku Berdasarkan ID

- GET (/api/categories)
    Mendapatkan Daftar Semua Categories.
- GET (/api/categories/:id)
    Mendapatkan Detail Categories Berdasarkan ID.
- POST (/api/categories)
    Menambahkan Categories Baru.
- PUT (/api/categories/:id)
    Memperbarui Categories Berdasarkan ID.
- DELETE (/api/categories/:id)
    Menghapus Categories Berdasarkan ID.

Jalankan aplikasi menggunakan perintah berikut:

```bash
go run main.go
```

## License

Dokumentasi ini memberikan panduan lengkap untuk menjalankan, dan menggunakan API. Anda dapat menyesuaikan informasi sesuai kebutuhan proyek Anda.