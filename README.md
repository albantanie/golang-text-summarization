# Artificial Intelligence - Text Summarization

Proyek ini adalah aplikasi web sederhana yang memungkinkan pengguna untuk merangkum teks menggunakan model AI dari Hugging Face. Aplikasi ini menggunakan API dari Hugging Face untuk melakukan tugas peringkasan teks.

## Fitur

- Input teks melalui textarea.
- Tombol untuk memulai proses peringkasan.
- Animasi loading saat proses peringkasan berlangsung.
- Output ringkasan teks ditampilkan di bawah input.

## Prasyarat

Pastikan Anda memiliki:
- [Go](https://golang.org/doc/install) terinstal di sistem Anda.
- Token API dari [Hugging Face](https://huggingface.co/).

## Instalasi

1. Clone repositori ini:
    ```bash
    git clone https://github.com/albantanie/golang-text-summarization
    cd saad
    ```

2. Instal dependensi Go:
    ```bash
    go mod tidy
    ```

3. Tambahkan token Hugging Face Anda di kode Go:
    ```go
    ic := huggingface.NewInferenceClient("YOUR-HUGGINGFACE-TOKEN")
    ```

4. Jalankan server Go:
    ```bash
    go run main.go
    ```

5. Buka `index.html` di browser Anda.

## Struktur Proyek

- `main.go`: Kode backend untuk memanggil API Hugging Face dan melakukan peringkasan teks.
- `index.html`: Halaman web yang berisi antarmuka pengguna.
- `style.css`: (Opsional) Gaya tambahan untuk halaman web.
- `README.md`: Dokumen ini.

## Penggunaan
-  Masukkan teks yang ingin dirangkum ke dalam textarea.
-  Klik tombol "Summarize".
-  Tunggu proses peringkasan selesai, dan hasil ringkasan akan ditampilkan di bawah textarea.

## Kontribusi
-  Kontribusi sangat diterima! Harap fork repositori ini dan ajukan pull request untuk perbaikan, fitur baru, atau perbaikan bug.