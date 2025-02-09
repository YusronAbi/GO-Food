# Sistem Pemesanan Makanan Online

Sistem Pemesanan Makananan Online yang memungkinkan pengguna untuk menjelajahi produk, membuat pesanan, dan melakukan pembayaran secara efisien. 

## Fitur

- **Pendaftaran dan Login Pengguna**: Pengguna dapat mendaftar dan masuk ke sistem untuk mengakses fitur pemesanan.
- **Menjelajahi Produk**: Pengguna dapat melihat daftar produk yang tersedia dengan detail seperti nama, deskripsi, harga, dan stok.
- **Keranjang Belanja**: Pengguna dapat menambahkan produk ke keranjang belanja sebelum melakukan pemesanan.
- **Membuat Pesanan**: Pengguna dapat membuat pesanan yang mencakup beberapa item.
- **Rincian Item Pesanan**: Setiap pesanan menyimpan rincian produk yang dipesan, termasuk jumlah dan harga.
- **Pembayaran**: Pengguna dapat melakukan pembayaran untuk menyelesaikan transaksi.
- **Pelacakan Status Pesanan**: Pengguna dapat melacak status pesanan mereka.

## Struktur Database

Berikut adalah struktur tabel yang digunakan dalam sistem ini:

1. **Users**
   - `id`: Identifikasi unik pengguna
   - `name`: Nama pengguna
   - `email`: Alamat email (unik)
   - `password`: Kata sandi pengguna
   - `phone_number`: Nomor telepon pengguna
   - `address`: Alamat pengguna
   - `created_at`: Tanggal dan waktu pembuatan akun
   - `updated_at`: Tanggal dan waktu pembaruan akun

2. **Products**
   - `id`: Identifikasi unik produk
   - `name`: Nama produk
   - `description`: Deskripsi produk
   - `price`: Harga produk
   - `stock`: Jumlah stok produk
   - `created_at`: Tanggal dan waktu pembuatan produk
   - `updated_at`: Tanggal dan waktu pembaruan produk

3. **Pesanan**
   - `id`: Identifikasi unik pesanan
   - `user_id`: ID pengguna yang membuat pesanan
   - `order_date`: Tanggal dan waktu pesanan dibuat
   - `status`: Status pesanan (misalnya, "pending", "completed")
   - `total_amount`: Jumlah total pesanan
   - `created_at`: Tanggal dan waktu pembuatan pesanan
   - `updated_at`: Tanggal dan waktu pembaruan pesanan

4. **Item Pesanan**
   - `id`: Identifikasi unik item pesanan
   - `pesanan_id`: ID pesanan yang terkait
   - `product_id`: ID produk yang dipesan
   - `quantity`: Jumlah produk yang dipesan
   - `price`: Harga per unit produk
   - `created_at`: Tanggal dan waktu pembuatan item pesanan
   - `updated_at`: Tanggal dan waktu pembaruan item pesanan

5. **Payments**
   - `id`: Identifikasi unik pembayaran
   - `pesanan_id`: ID pesanan yang dibayar
   - `payment_date`: Tanggal dan waktu pembayaran dilakukan
   - `amount`: Jumlah yang dibayarkan
   - `payment_method`: Metode pembayaran yang digunakan
   - `status`: Status pembayaran (misalnya, "completed", "failed")
   - `created_at`: Tanggal dan waktu pembuatan pembayaran
   - `updated_at`: Tanggal dan waktu pembaruan pembayaran

## Diagram Hubungan Entitas (ERD)

Berikut adalah diagram hubungan entitas (ERD) untuk sistem pemesanan online:
![Image](https://github.com/user-attachments/assets/ebb8116f-00b3-4e80-b84b-1e185a92bf11)


## Instalasi

1. Clone repositori ini ke mesin lokal Anda:
   ```bash
   git clone https://github.com/username/repo-name.git
