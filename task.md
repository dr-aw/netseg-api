# Go Developer Test Assignment

We are looking for an experienced Go Developer who is not only proficient with the technical stack but also capable of writing well-structured, readable, and maintainable code following best modern practices. To evaluate your skills, we invite you to complete the following test assignment.

---

## 📌 Assignment Requirements

You need to develop a **Go-based HTTP API** that manages two entities:  
- **NetSegment** (Network segment/subnet)
- **Host** (Host/node connected to a subnet)  

The application should also include a **database** for storing these entities.  
The solution must be hosted on a **public Git repository** (e.g., GitHub, GitLab, Bitbucket).  
To submit your solution, provide a link to your repository.

---

## 📌 Implementation Requirements

### **1. Language**
- The project must be implemented in **Go (Golang)**.
- The code must be **structured, readable, and maintainable**.
- It is recommended to follow **Clean Architecture** and **Domain-Driven Design (DDD)** principles.

### **2. Database**
- The preferred **database** is **PostgreSQL**.
- Alternatively, **MySQL, MariaDB, or MS SQL** may be used, but **PostgreSQL is recommended**.

---

## 📌 Database Entities

### **NetSegment (Network Segment)**
- **Name** of the subnet.
- **CIDR block** (network prefix and subnet mask).
- **DHCP** enabled/disabled flag.
- **Max Hosts** – maximum number of hosts allowed in the subnet.
- **CreatedAt** – timestamp of creation.
- **UpdatedAt** – timestamp of the last modification.

### **Host (Connected Node)**
- **IP Address** of the host.
- **MAC Address** of the host.
- **Status** (Online/Offline).
- **CreatedAt** – timestamp of creation.
- **UpdatedAt** – timestamp of the last modification.

📌 **You can add additional fields if necessary.**  
📌 **You must select appropriate data types for all required fields.**

---

## 📌 Data Validation & Consistency Requirements

The data in the database must **always remain consistent**, ensuring:
- **CIDR blocks must be unique** across all network segments.
- **The number of hosts in a subnet must not exceed** the `max_hosts` value.
- **IP addresses must be valid for the subnet** they belong to.
- **IP addresses must be unique** within their subnet.
- **MAC addresses must be globally unique**.

📌 The system **is the only application working with the database**.  
📌 The **database schema must be generated from the code** (**Code-First** approach).  
📌 The ORM **must be GORM** (other ORMs or query builders are allowed but not preferred).

---

## 📌 HTTP API Requirements

The application should expose a **RESTful API**.  
📌 The preferred framework for REST API is **Echo** (alternative frameworks are allowed but not recommended).

### **1. Network Segments**
- **POST** `/segments` – Create a new network segment.
- **PUT** `/segments/:id` – Update an existing network segment.
- **GET** `/segments` – Retrieve all network segments.

### **2. Hosts**
- **POST** `/hosts` – Create a new host.
- **PUT** `/hosts/:id` – Update an existing host.
- **GET** `/hosts` – Retrieve all hosts.

### **General API Rules**
- Endpoints must follow **RESTful principles**.
- Use appropriate **HTTP status codes** (`200 OK`, `404 Not Found`, `400 Bad Request`, etc.).
- API requests and responses must be in **JSON format**.

---

## 📌 Architecture

The code **must be well-structured**, ensuring clear separation of concerns:
- **API Layer** (Handlers, Routing).
- **Business Logic Layer** (Services / Use Cases).
- **Database Layer** (Repositories).

---

## 📌 Logging

Logging is **optional but recommended**.  
Even simple console logging is acceptable.

---

## 📌 Evaluation Criteria

We will assess your solution based on:
1. **Technical Requirements Compliance**  
2. **Code Quality** – readability, structure, adherence to Go style guide  
3. **Architecture Decisions**  
4. **Correct API Implementation** – REST principles compliance  
5. **Data Consistency Rules** – ensuring all constraints are enforced  
6. **Error Handling** – clear and meaningful error messages  
7. **Git Best Practices** – proper commit history and repository management  

---

## 📌 Estimated Completion Time

- The estimated completion time for this task is **one working day**.  
- If you are unable to implement all requirements, submit what you have completed.  

This test aims to **demonstrate your expertise in Go, API development, database integration, and software architecture**.  

🔥 **Good luck, and we look forward to your submission!** 🚀
