# NetSeg API

## 📌 Description
**NetSeg API** is a REST API service for managing network segments and hosts, built with **Golang**, **Echo**, and **PostgreSQL**.

### 🔧 **Features**:
- Manage network segments (`NetSegment`): create, update, and retrieve segments.
- Manage hosts (`Host`): add, update, and retrieve hosts.
- Validate IP addresses to ensure they belong to the assigned subnet.
- Data validation before saving to the database.

## 🚀 **Getting Started**

### 1️⃣ **Clone the repository**
```sh
git clone https://github.com/dr-aw/netseg-api.git
cd netseg-api
```

### 2️⃣ **Create a `.env` file**
Before running the project, create a **`.env`** file in the root directory and specify the database connection parameters:
```ini
POSTGRES_USER=postgres
POSTGRES_PASSWORD=secret
POSTGRES_DB=netseg
POSTGRES_HOST=db
POSTGRES_PORT=5432
```

### 3️⃣ **Run using Docker Compose**
```sh
docker-compose up --build
```

### 4️⃣ **The API will be available at:**
```
http://localhost:8080/api/v1
```

## 🔥 **API Usage**

### 📍 **Network Segments**

#### ➕ **Create a Network Segment**
```http
POST /api/v1/segments
```
📌 **Example request body:**
```json
{
  "name": "Office Network",
  "cidr": "192.168.1.0/24",
  "dhcp": true,
  "max_hosts": 50
}
```

#### 📄 **Retrieve all Network Segments**
```http
GET /api/v1/segments
```

### 📍 **Hosts**

#### ➕ **Add a Host**
```http
POST /api/v1/hosts
```
📌 **Example request body:**
```json
{
  "ip_address": "192.168.1.10",
  "mac": "00:1A:2B:3C:4D:5E",
  "status": "online",
  "segment_id": 1
}
```

#### 📄 **Retrieve all Hosts**
```http
GET /api/v1/hosts
```

## 🛠 **Tech Stack**
- **Golang** (Echo, GORM)
- **PostgreSQL**
- **Docker, Docker Compose**
- **Clean Architecture**

## 📌 **To-Do**
- Implement unit and integration tests
- Consider adding caching for quick uniqueness checks of IP addresses within a subnet and MAC addresses

## 💡 **Author**
[Konstantin Gunbin](https://github.com/dr-aw)

## 📜 **License**
This project is licensed under the **MIT License**.