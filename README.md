# Banking Service API

Aplikasi REST API sederhana untuk layanan perbankan dengan fitur dasar: registrasi nasabah, menabung, penarikan dana, dan cek saldo.

## Fitur

- Registrasi nasabah baru dengan validasi NIK dan nomor HP
- Menambah saldo rekening (menabung)
- Penarikan dana dengan validasi saldo
- Pengecekan saldo rekening
- Logging terstruktur untuk setiap operasi

## Teknologi

- Golang dengan Echo framework
- PostgreSQL sebagai database
- GORM sebagai ORM
- Docker dan Docker Compose untuk deployment
- Structured logging

## Struktur Aplikasi

Aplikasi ini menggunakan arsitektur clean dengan pemisahan layer:

- **API Layer** - Handler untuk request HTTP
- **Service Layer** - Business logic
- **Repository Layer** - Akses database
- **Model Layer** - Struktur data

## API Endpoints

### 1. Registrasi Nasabah

```
POST /daftar
```

Request:
```json
{
  "nama": "Budi Santoso",
  "nik": "1234567890123456",
  "no_hp": "081234567890"
}
```

Response (200):
```json
{
  "no_rekening": "0000000001"
}
```

### 2. Menabung

```
POST /tabung
```

Request:
```json
{
  "no_rekening": "0000000001",
  "nominal": 50000
}
```

Response (200):
```json
{
  "saldo": 50000
}
```

### 3. Penarikan Dana

```
POST /tarik
```

Request:
```json
{
  "no_rekening": "0000000001",
  "nominal": 25000
}
```

Response (200):
```json
{
  "saldo": 25000
}
```

### 4. Cek Saldo

```
GET /saldo/{no_rekening}
```

Response (200):
```json
{
  "saldo": 25000
}
```

## Cara Menjalankan

### Menggunakan Docker Compose

1. Clone repository ini:
   ```bash
   git clone https://github.com/Fahrizal0112/banking-service.git
   cd banking-service
   ```

2. Jalankan dengan Docker Compose:
   ```bash
   docker compose up -d
   ```

3. Aplikasi tersedia di http://localhost:8080

### Pengembangan Lokal

1. Clone repository dan install dependencies:
   ```bash
   git clone https://github.com/Fahrizal0112/banking-service.git
   cd banking-service
   go mod download
   ```

2. Siapkan PostgreSQL di local atau container:
   ```bash
   docker run -d --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=bank -p 5432:5432 postgres:15
   ```

3. Copy .env.example ke .env dan sesuaikan:
   ```bash
   cp .env.example .env
   ```

4. Jalankan aplikasi:
   ```bash
   go run cmd/api/main.go
   ```

## Konfigurasi

Aplikasi menggunakan kombinasi environment variables dan command line flags:

### Environment Variables:
- `DB_HOST` - Host database PostgreSQL
- `DB_PORT` - Port database
- `DB_USER` - Username database
- `DB_PASSWORD` - Password database
- `DB_NAME` - Nama database

### Command Line Flags:
- `-host` - Host untuk server API (default: "0.0.0.0")
- `-port` - Port untuk server API (default: 8080)

## Testing

Contoh testing dengan curl:

```bash
# Registrasi nasabah
curl -X POST http://localhost:8080/daftar \
  -H "Content-Type: application/json" \
  -d '{"nama":"Muchammad Fahrizal","nik":"1234567890123456","no_hp":"081234567890"}'

# Menabung
curl -X POST http://localhost:8080/tabung \
  -H "Content-Type: application/json" \
  -d '{"no_rekening":"0000000001","nominal":50000}'

# Cek saldo
curl http://localhost:8080/saldo/0000000001
```

## Struktur Database

Aplikasi menggunakan satu tabel utama:

**Tabel Nasabah**
- `id` - Primary key
- `nama` - Nama nasabah
- `nik` - Nomor Induk Kependudukan (unique)
- `no_hp` - Nomor handphone (unique)
- `no_rekening` - Nomor rekening (unique)
- `saldo` - Saldo rekening