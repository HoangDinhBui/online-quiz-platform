- ---

# 🧠 Online Quiz Platform – Dockerized Setup

## 📦 Tổng quan

Đây là một hệ thống **thi trắc nghiệm trực tuyến** được xây dựng với kiến trúc microservices, bao gồm:

* **Frontend:** ReactJS
* **Backend:** Go + Gin + MongoDB + Redis
* **Database:** MongoDB
* **Caching:** Redis
* **Reverse Proxy & Rate Limiting:** NGINX
* **Monitoring:** Prometheus & Grafana

---

## 🛠️ Cấu trúc hệ thống

```text
.
├── backend/               # Source code backend (Go)
├── frontend/              # Source code frontend (ReactJS)
├── nginx/
│   └── nginx.conf         # Cấu hình NGINX reverse proxy
├── prometheus/
│   └── prometheus.yml     # Cấu hình Prometheus metrics
└── docker-compose.yml     # Tổ chức toàn bộ container dịch vụ
```

---

## 🚀 Cách chạy dự án

### 1. Yêu cầu

* Docker
* Docker Compose

### 2. Khởi động toàn bộ hệ thống:

```bash
docker-compose up --build
```

Toàn bộ hệ thống sẽ khởi chạy trên các cổng:

| Dịch vụ       | Cổng Host | Container | Mô tả                                                          |
| ------------- | --------- | --------- | -------------------------------------------------------------- |
| Frontend      | `5173`    | `5173`    | ReactJS dev server                                             |
| Backend       | `8080`    | `8080`    | API chính (Go + Gin)                                           |
| Backend2      | `8081`    | `8080`    | Bản backend thứ 2 cho load balancing                           |
| NGINX Gateway | `8000`    | `80`      | Reverse Proxy ([http://localhost:8000](http://localhost:8000)) |
| MongoDB       | `27017`   | `27017`   | Database chính                                                 |
| Redis         | `6379`    | `6379`    | Cache + pub/sub                                                |
| Prometheus    | `9090`    | `9090`    | Theo dõi metrics backend                                       |
| Grafana       | `3001`    | `3000`    | Giao diện hiển thị dashboard                                   |

---

## 🌐 Truy cập hệ thống

* **Frontend Web App:** [http://localhost:8000](http://localhost:8000)
* **API:** [http://localhost:8000/api](http://localhost:8000/api)
* **Prometheus:** [http://localhost:9090](http://localhost:9090)
* **Grafana:** [http://localhost:3001](http://localhost:3001)

  * Tài khoản mặc định: `admin / admin`

---

## 🔒 Cấu hình đặc biệt

* **NGINX:** Load balancing giữa 2 backend (`backend` và `backend2`), giới hạn rate `10r/s` theo IP.
* **Redis:** Sử dụng để cache dữ liệu `classes`, `questions`, tăng hiệu năng.
* **MongoDB:** Lưu toàn bộ dữ liệu người dùng, câu hỏi, kết quả làm bài.
* **Prometheus + Grafana:** Theo dõi hiệu suất backend qua `/metrics`.

---

## 📝 Ghi chú

* Nếu cần xóa toàn bộ dữ liệu MongoDB/Redis:

  ```bash
  docker volume rm <tên volume>
  # ví dụ:
  docker volume rm online-quiz-platform_mongodb_data
  docker volume rm online-quiz-platform_redis_data
  ```

* Trong môi trường production, bạn nên thay `localhost`, giới hạn CORS và dùng Docker secrets cho biến môi trường.

---

## 📧 Tác giả

* 👨‍💻 **HoangDinhBui** – [GitHub](https://github.com/HoangDinhBui)
* 📚 Dự án: **Online Quiz Platform**

---
