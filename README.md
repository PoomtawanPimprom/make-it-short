# 🚀 make-it-short

A simple URL Shortener application that allows you to shorten long URLs into short links with redirect functionality and click tracking.  
Built with **Vue 3 + TailwindCSS** on the frontend and **Go (Gin + GORM + PostgreSQL)** on the backend.

---

## 📌 Features

- ✅ Shorten long URLs into short links  
- ✅ Redirect to the original URL  
- ✅ Track click counts  
- ✅ REST API for frontend or external clients  
- 🎯 Easily extendable with custom alias, dashboard, or authentication  

---

## 🛠 Tech Stack

### Frontend
- Vue 3 + Vite  
- TailwindCSS  
- Axios  

### Backend
- Go + Gin (Web Framework)  
- GORM (ORM)  
- PostgreSQL (Database)  

---

## 📂 Project Structure
short-url/
├── backend/ # Go + Gin + GORM + PostgreSQL
│ ├── main.go
│ ├── database/
│ ├── models/
│ ├── controllers/
│ └── routes/
└── frontend/ # Vue 3 + Tailwind
├── src/
│ ├── App.vue
│ ├── components/
│ └── api/
└── tailwind.config.js

---

## ⚡ API Endpoints

### Create Short URL
**POST** `/shorten`

**Request**
```
json
{
  "url": "https://example.com/long/path"
}
```
**Response**
```
{
  "short_url": "http://localhost:8080/abc123"
}
```
### Redirect to Original URL
**GET** `/:code`
**Example:**
```
http://localhost:8080/abc123 → Redirects to https://example.com/long/path
```
---
### Setup 

## Backend
```
cd backend
go mod tidy
go run main.go
```
## Frontend
```
cd frontend
npm install
npm run dev
```
