# Supplier Management System

Sistem manajemen supplier lengkap dengan Backend Golang-PostgreSQL dan Frontend SPA Laravel-Vue.js berdasarkan spesifikasi dokumen yang diberikan.

## 🏗️ Arsitektur Sistem

### Backend (Golang)
- **Framework**: Gin (REST API)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Port**: 8080
- **URL**: https://localhost:8080

### Frontend (Laravel SPA)
- **Framework**: Laravel 
- **Port**: 8000
- **URL**: https://127.0.0.1:8000

## 📊 Database Schema

### Entitas Utama:
- **Suppliers**: Data supplier utama
- **Addresses**: Alamat supplier (Head Office, Branch Office)
- **Contacts**: Kontak person supplier
- **Supplier Groups**: Grup/kategori supplier
- **Materials**: Katalog material dari supplier
- **Other Informations**: Informasi tambahan (SAP Vendor ID, NPWP, dll)
- **Order Histories**: Riwayat pesanan
- **Invoices**: Data invoice
- **Review Approvals**: Proses review dan approval
- **Workflows**: Alur kerja approval

## 🚀 Fitur Utama

### Dashboard
- Statistik supplier (Total, New, Blocked, Avg Cost)
- Statistik order (Total, On Time, Late, In Progress)
- Statistik invoice (Total, In Progress, Paid, Outstanding)

### Supplier Management
- **Supplier List**: Daftar supplier dengan filter dan pencarian
- **Supplier Detail**: Detail lengkap supplier dengan tabs:
  - Overview: Alamat, kontak, grup, informasi lain
  - Material Catalog: Katalog material supplier
  - Orders: Riwayat pesanan
  - Invoices: Data invoice
- **New/Edit Supplier**: Form tambah/edit supplier
- **Block/Unblock**: Fitur blokir/buka blokir supplier

### Review & Approvals
- **Supplier Creation**: Proses approval supplier baru
- **Supplier Assessment**: Penilaian supplier
- **Block/Unblock Supplier**: Approval blokir/buka blokir
- **Workflow Management**: Manajemen alur kerja approval

## 🔧 Instalasi dan Setup

### Prerequisites
- Go 1.21+
- PHP 8.1+
- Node.js 20+
- PostgreSQL 14+
- Composer

### Backend Setup
```bash
cd supplier-management-backend
go mod tidy
go build -o supplier-management-backend
./supplier-management-backend
```
Atau 
```bash
go run main.go
```
### Frontend Setup
```bash
cd supplier-management-frontend
composer install
npm install
npm run build
php artisan serve --host=0.0.0.0 --port=3000
```



## 📋 API Endpoints

### Dashboard
- `GET /api/v1/dashboard/stats` - Statistik dashboard

### Suppliers
- `GET /api/v1/suppliers` - List suppliers (dengan pagination & filter)
- `GET /api/v1/suppliers/{id}` - Detail supplier
- `POST /api/v1/suppliers` - Tambah supplier baru
- `PUT /api/v1/suppliers/{id}` - Update supplier
- `DELETE /api/v1/suppliers/{id}` - Hapus supplier
- `POST /api/v1/suppliers/{id}/block` - Blokir supplier
- `POST /api/v1/suppliers/{id}/unblock` - Buka blokir supplier

### Review & Approvals
- `GET /api/v1/review-approvals` - List proses approval
- `GET /api/v1/review-approvals/{id}` - Detail proses approval
- `POST /api/v1/review-approvals` - Buat proses approval baru
- `PUT /api/v1/review-approvals/{id}` - Update proses approval



## 📈 Monitoring & Logging

- Gin logging middleware
- Database query logging via GORM
- Console logging untuk debugging frontend



## 📝 Catatan Implementasi

1. **Database**: Menggunakan PostgreSQL dengan auto-migration GORM
2. **API**: RESTful API dengan response format JSON standar
3. **Frontend**: SPA menggunakan Vue.js 3 dengan Laravel sebagai backend
4. **Styling**: Custom CSS tanpa framework eksternal untuk kontrol penuh
5. **Responsiveness**: Design responsive untuk berbagai ukuran layar

## 🔄 Future Enhancements

- Authentication & Authorization
- File upload untuk logo supplier
- Export data ke Excel/PDF
- Email notifications untuk approval
- Advanced reporting & analytics
- Real-time notifications
- Audit trail untuk semua perubahan data
