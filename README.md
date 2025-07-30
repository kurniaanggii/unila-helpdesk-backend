# Unila Helpdesk API

API untuk aplikasi Helpdesk Universitas Lampung yang terintegrasi dengan survei kepuasan pelanggan.

## Fitur Utama

1. **Helpdesk berbasis tiket** - Sistem manajemen tiket dengan prioritas dan status
2. **Dua jenis pengguna** - User terdaftar (mahasiswa, dosen, staf) dan guest
3. **Dua jenis tipe perangkat** - Mobile untuk role user, Web untuk admin
4. **Prioritas dan status tiket** - Rendah/sedang/tinggi, menunggu/diproses/selesai
5. **Sistem kuesioner** - Admin dapat membuat kuesioner untuk setiap kategori layanan
6. **Analisis kohort** - Laporan tren layanan dan kepuasan pelanggan
7. **Pencarian tiket** - Fitur pencarian tiket
8. **CRUD tiket** - Buat, lihat, edit, hapus tiket
9. **Mock SSO Unila** - Simulasi login SSO (dapat diganti dengan SSO asli)
10. **Notifikasi push** - Firebase Cloud Messaging untuk update status tiket dan reminder survei

## Teknologi

- **Framework**: Go dengan Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Arsitektur**: Clean Architecture Minimal
- **Notifikasi**: Firebase Cloud Messaging

## Struktur Proyek

```
unila-helpdesk-backend/
├── main.go                     # Entry point aplikasi
├── go.mod                      # Dependencies
├── .env.example               # Template konfigurasi
├── config/
│   └── config.go               # Konfigurasi aplikasi
├── database/
│   └── database.go             # Koneksi & setup database
├── models/
│   └── models.go               # Semua struct models
├── handlers/
│   ├── auth.go                 # Handler autentikasi
│   ├── ticket.go               # Handler tiket
│   ├── survey.go               # Handler survei & kuesioner
│   ├── analytics.go            # Handler analitik
|   └── notification.go         # Handler notifikasi push
├── services/
│   ├── auth.go                 # Service autentikasi
│   ├── ticket.go               # Service tiket
│   ├── survey.go               # Service survei
│   └── analytics.go            # Service analitik
├── middleware/
│   └── auth.go                 # Middleware autentikasi
├── utils/
│   ├── response.go             # Utility response
│   └── notification.go         # Utility FCM
└── routes/
    └── routes.go               # Definisi routes
```

## Database Schema

### Tabel Utama

1. **users** - Data pengguna terdaftar (NIM/NIP sebagai primary key)
2. **service_categories** - Kategori layanan (website, vclass, dll)
3. **tickets** - Data tiket helpdesk
4. **questionnaires** - Template kuesioner per kategori
5. **questions** - Pertanyaan dalam kuesioner
6. **question_options** - Pilihan jawaban untuk multiple choice
7. **survey_responses** - Header respons survei
8. **survey_answers** - Detail jawaban per pertanyaan
9. **notifications** - Notifikasi push
10. **cohort_analysis** - Data analisis kohort

### Kategori Layanan

**Perlu Login:**
- website
- jaringan internet
- siakadu
- sistem informasi
- lainnya

**Tanpa Login:**
- lupa password
- buat email unila
- buat sso unila

## Instalasi

1. **Clone repository**
```bash
git clone <repository-url>
cd unila-helpdesk-backend
```

2. **Install dependencies**
```bash
go mod tidy
```

3. **Setup database PostgreSQL**

Buat database:
```bash
# Login ke PostgreSQL sebagai user postgres
psql -U postgres

# Buat database
CREATE DATABASE unila_helpdesk;

# Keluar dari psql
\q
```

Import schema database:
```bash
# Import schema ke database
psql -U postgres -d unila_helpdesk -f database_schema.sql
```

4. **Konfigurasi environment**
```bash
cp .env.example .env
# File .env sudah dikonfigurasi dengan:
# DB_USER=postgres
# DB_PASSWORD=admin123
# DB_NAME=unila_helpdesk
```

5. **Setup Firebase (opsional)**
- Download service account key dari Firebase Console
- Simpan sebagai `firebase-key.json` di root project
- Update `FCM_KEY_PATH` di .env

6. **Jalankan aplikasi**
```bash
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## API Endpoints

### Authentication
- `POST /api/v1/auth/login` - Login user
- `GET /api/v1/auth/profile` - Profil user (protected)
- `GET /api/v1/auth/validate` - Validasi token
- `POST /api/v1/auth/logout` - Logout

### Tickets
- `POST /api/v1/tickets` - Buat tiket (guest/user)
- `GET /api/v1/tickets/search` - Pencarian tiket
- `GET /api/v1/tickets/number/:number` - Cari berdasarkan nomor
- `GET /api/v1/tickets/categories` - Kategori layanan
- `GET /api/v1/tickets/my` - Tiket saya (protected)
- `GET /api/v1/tickets/:id` - Detail tiket (protected)
- `PUT /api/v1/tickets/:id` - Update tiket (protected)
- `DELETE /api/v1/tickets/:id` - Hapus tiket (protected)
- `GET /api/v1/tickets` - Semua tiket (admin)

### Surveys
- `GET /api/v1/surveys/questionnaires/category/:categoryId` - Kuesioner berdasarkan kategori
- `GET /api/v1/surveys/questionnaires/:id` - Detail kuesioner
- `POST /api/v1/surveys/submit` - Submit survei (protected)
- `GET /api/v1/surveys/tickets/:ticketId` - Survei berdasarkan tiket (protected)

**Admin only:**
- `POST /api/v1/surveys/questionnaires` - Buat kuesioner
- `GET /api/v1/surveys/questionnaires` - Semua kuesioner
- `PUT /api/v1/surveys/questionnaires/:id` - Update kuesioner
- `DELETE /api/v1/surveys/questionnaires/:id` - Hapus kuesioner
- `POST /api/v1/surveys/questions` - Buat pertanyaan
- `DELETE /api/v1/surveys/questions/:id` - Hapus pertanyaan
- `POST /api/v1/surveys/question-options` - Buat pilihan jawaban
- `DELETE /api/v1/surveys/question-options/:id` - Hapus pilihan jawaban
- `GET /api/v1/surveys/responses` - Semua respons survei
- `GET /api/v1/surveys/responses/category/:categoryId` - Respons berdasarkan kategori

### Analytics (Admin only)
- `GET /api/v1/analytics/cohort` - Analisis kohort
- `POST /api/v1/analytics/cohort/save` - Simpan analisis kohort
- `GET /api/v1/analytics/service-trends` - Tren layanan
- `GET /api/v1/analytics/ticket-status` - Statistik status tiket
- `GET /api/v1/analytics/user-entities` - Statistik entitas user
- `GET /api/v1/analytics/satisfaction-trend` - Tren kepuasan
- `GET /api/v1/analytics/top-issues` - Masalah teratas
- `GET /api/v1/analytics/resolution-time` - Waktu penyelesaian
- `GET /api/v1/analytics/dashboard` - Statistik dashboard

## Aturan Bisnis

1. **Akses Kategori Layanan**
   - User terdaftar: dapat membuat semua kategori tiket
   - Guest: hanya dapat membuat tiket kategori keanggotaan

2. **Survei Kepuasan**
   - Hanya dapat diisi untuk tiket dengan status "selesai"
   - Hanya pemilik tiket yang dapat mengisi survei
   - Setiap tiket hanya dapat diisi survei sekali

3. **Kuesioner**
   - Setiap kategori layanan memiliki kuesioner berbeda
   - Admin dapat membuat pertanyaan dengan tipe: text, multiple_choice, rating

4. **Pencarian Tiket**
   - User terdaftar: dapat melihat semua tiket setelah login
   - Guest: dapat mengakses melalui fitur pencarian

5. **Analisis Kohort**
   - Menganalisis pengguna baru vs kembali
   - Tren layanan berdasarkan periode
   - Tingkat kepuasan per kategori

6. **Notifikasi**
   - Dikirim saat status tiket berubah
   - Reminder untuk mengisi survei

## Format Data

### Login Request
```json
{
  "username": "string",
  "password": "string"
}
```

### Create Ticket Request
```json
{
  "title": "string",
  "description": "string",
  "priority": "rendah|sedang|tinggi",
  "service_category_id": 1,
  "attachment_path": "string", // untuk user terdaftar
  "guest_user_type": "dosen|mahasiswa|tendik", // untuk guest
  "guest_email": "string", // untuk guest
  "id_card_path": "string", // untuk guest
  "selfie_path": "string" // untuk guest
}
```

### Submit Survey Request
```json
{
  "ticket_id": 1,
  "questionnaire_id": 1,
  "answers": [
    {
      "question_id": 1,
      "question_option_id": 1, // untuk multiple choice
      "answer_text": "string", // untuk text
      "answer_value": 5 // untuk rating
    }
  ]
}
```

## Authentication

Aplikasi menggunakan sistem token sederhana untuk development:
- Token = User ID (NIM/NIP)
- Header: `Authorization: Bearer <user_id>`

Untuk production, dapat diganti dengan JWT atau integrasi SSO Unila yang sesungguhnya.

## Development

### Mock Data
Database akan otomatis dibuat dengan kategori layanan default saat pertama kali dijalankan.

### Testing
```bash
# Test health check
curl http://localhost:8080/health

# Test create ticket (guest)
curl -X POST http://localhost:8080/api/v1/tickets \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Test Ticket",
    "description": "Test Description",
    "priority": "sedang",
    "service_category_id": 6,
    "guest_user_type": "mahasiswa",
    "guest_email": "test@student.unila.ac.id"
  }'
```

## Deployment

1. Build aplikasi:
```bash
go build -o unila-helpdesk-backend main.go
```

2. Setup database production
3. Update konfigurasi .env
4. Jalankan binary:
```bash
./unila-helpdesk-backend
```

## Kontribusi

1. Fork repository
2. Buat feature branch
3. Commit perubahan
4. Push ke branch
5. Buat Pull Request

## Lisensi

MIT License
