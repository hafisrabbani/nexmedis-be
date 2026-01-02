# Backend Engineer Technical Test – NexMedis

## Overview

Project ini merupakan implementasi backend service untuk **API Usage Tracking** dengan fokus pada:

* Scalability
* Concurrency safety
* Clean architecture
* Explicit dependency management
* Redis-first performance strategy

Service ini menyediakan kemampuan:

* Client registration berbasis API Key
* Logging API usage berfrekuensi tinggi
* Aggregasi usage harian & top client
* Rate limiting berbasis Redis
* **Real-time usage streaming menggunakan Server-Sent Events (SSE)**

---

## Tech Stack

* Language: Go
* HTTP Framework: Fiber (fasthttp)
* Database: PostgreSQL
* ORM: GORM
* Cache / Counter: Redis
* Load Testing: k6
* Hot Reload: air

---

## High Level Architecture

```text
┌───────────┐
│  Client   │
└─────┬─────┘
      │ HTTP
┌─────▼─────┐
│  Fiber    │
│  Router   │
└─────┬─────┘
      │
┌─────▼─────────────┐
│ Middleware Layer  │
│ - API Key Auth    │
│ - Rate Limit      │
│ - Usage Tracking  │
└─────┬─────────────┘
      │
┌─────▼─────────────┐
│ Service Layer     │
│ (Business Logic)  │
└─────┬─────────────┘
      │
┌─────▼─────────────┐
│ Repository Layer  │
│ - Redis (primary) │
│ - DB (fallback)   │
└─────┬─────────────┘
      │
┌─────▼─────────────┐
│ SSE Endpoint      │
│ /usage/stream     │
└───────────────────┘
```

---

## Project Structure

```text
cmd/server
  └── main.go              # application entry point

internal/
  ├── config               # app, db, redis config
  ├── module               # module-based wiring (manual DI)
  ├── delivery/
  │   ├── http              # router, handler, middleware
  │   └── sse               # Server-Sent Events delivery
  ├── repository            # data access (gorm + redis)
  ├── service               # business logic
  ├── model
  │   ├── request
  │   ├── response
  │   └── error
  └── shared                # helpers
```

---

## Design Decisions

### 1. Manual Dependency Injection (Module Pattern)

Project ini tidak menggunakan DI framework.
Sebagai gantinya digunakan **module-based wiring**:

```go
clientModule := client.Register(app, db)
usage.Register(app, db, rdb, clientModule.APIKeyMiddleware)
```

Alasan:

* Dependency eksplisit
* Mudah dibaca & direview
* Tidak ada magic / reflection
* Idiomatic Go

---

### 2. Redis-First Strategy for Usage Tracking

Usage logging menggunakan Redis sebagai **primary write path**:

* `INCR` untuk daily usage
* `ZINCRBY` untuk top client
* TTL-based expiration

Database digunakan sebagai:

* Persistent storage (fallback)
* Future batch insert

Keuntungan:

* Aman untuk high concurrency
* Minim lock & contention
* Performa jauh lebih baik dibanding DB aggregation

---

### 3. API Key Authentication

* API Key digunakan untuk identifikasi client
* Validasi dilakukan di middleware
* Client object disimpan di `fiber.Ctx.Locals`

```go
c.Locals("client", *repository.Client)
```

Pendekatan ini:

* Type-safe
* Tidak bergantung JWT di tahap awal
* Cocok untuk service-to-service authentication

---

### 4. Rate Limiting (Redis Based)

Rate limit diterapkan:

* Per client
* Per jam
* Atomic menggunakan Redis `INCR`

```text
ratelimit:{client_id}:{yyyy-mm-dd-hh}
```

Jika Redis down:

* Request tetap diizinkan
* Sistem tetap berjalan (graceful degradation)

---

### 5. Real-Time Usage Streaming (Server-Sent Events)

Project ini mengimplementasikan **Server-Sent Events (SSE)** untuk men-stream data usage secara real-time.

Endpoint:

```
GET /api/usage/stream
```

Karakteristik implementasi:

* Menggunakan pola resmi Fiber (`SetBodyStreamWriter`)
* Tanpa goroutine tambahan atau channel in-memory
* Mengandalkan error `Flush()` sebagai sinyal client disconnect
* Polling ringan ke service layer (Redis-backed)

Contoh response SSE:

```json
{
    "status": true,
    "message": "success",
    "data": [
        {
            "date": "2026-01-02",
            "total_requests": 16
        },
        {
            "date": "2026-01-01",
            "total_requests": 1004
        },
        {
            "date": "2025-12-31",
            "total_requests": 0
        },
        {
            "date": "2025-12-30",
            "total_requests": 0
        },
        {
            "date": "2025-12-29",
            "total_requests": 0
        },
        {
            "date": "2025-12-28",
            "total_requests": 0
        },
        {
            "date": "2025-12-27",
            "total_requests": 0
        }
    ]
}
```

Pendekatan ini dipilih karena:

* Stabil pada Fiber (fasthttp)
* Cocok untuk long-lived connection
* Mudah dikembangkan ke Redis Pub/Sub

---

### 6. Error Handling Strategy

* Repository mengembalikan domain error
* Service meneruskan error
* Handler melakukan mapping ke HTTP response

Contoh:

* Duplicate client → `400 Bad Request`
* Rate limit → `429 Too Many Requests`
* Auth gagal → `401 Unauthorized`

---

## API Endpoints

### Public

* `POST /api/register`
* `GET /health`

### Protected (API Key)

* `POST /api/logs`
* `GET /api/usage/daily`
* `GET /api/usage/top`
* `GET /api/usage/stream` (SSE)

---

## Performance Testing

Load testing dilakukan menggunakan **k6** dengan skenario:

* Smoke test
* High concurrency
* Rate limit validation

Ekspektasi hasil:

* Tidak ada 5xx error
* Response `200` & `429` valid
* Redis counter konsisten

---

## Future Improvements

* JWT Authorization untuk `/usage/*`
* Batch DB insert (buffered worker)
* Redis Pub/Sub untuk SSE push-based
* WebSocket alternative
* Docker & docker-compose
* OpenAPI / Swagger documentation

---

## Conclusion

Project ini dirancang dengan fokus pada:

* Scalability
* Clarity
* Maintainability

Tanpa over-engineering, sistem ini sudah siap untuk:

* Menangani high traffic
* Diskalakan secara horizontal
* Dikembangkan ke tahap production
