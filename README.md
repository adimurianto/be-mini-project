# BE Mini Project

## Setup Project

- Clone project

```bash
  git clone https://github.com/adimurianto/be-mini-project.git
  cd be-mini-project
```

- Get all dependency

```bash
  go get .
```

- Create file **.env** base on file **.env.example** and adjust the contents of the variable

- Init for generate Swagger API documentation

```bash
  swag init
```

- Init for live reload

```bash
  air init
```

- Running project

Choose one

```bash
  // if using live reload
  air

  // test locally on your system
  go run main.go

  // build locally on your system then make sure run /main.exe or /main
  go build main.go
```

- Access url http://127.0.0.1:5000/docs/index.html

  <img width="auto" alt="image" src="https://github.com/user-attachments/assets/5dbd14a8-eb04-48c7-a0a6-1e5bb47cbdb6" />

## Deploy ke GCP Menggunakan Docker

### 1. Setup VM di GCP

1. Buka [Google Cloud Console](https://console.cloud.google.com)
2. Buat VM Instance:
   - **Name**: be-mini-project
   - **Region**: us-central1 (atau yang terdekat)
   - **Machine type**: e2-micro (Free Tier)
   - **Boot disk**: Ubuntu 22.04 LTS
   - **Firewall**: Allow HTTP dan HTTPS

### 2. Install Docker di VM

```bash
sudo apt update && sudo apt install -y docker.io
```

### 3. Clone Project

```bash
git clone https://github.com/adimurianto/be-mini-project.git
cd be-mini-project
```

### 4. Buat .env

```bash
cat > .env << 'EOF'
PSQL_DB_HOST=xxxxxxxxxxxx
PSQL_DB_PORT=5432
PSQL_DB_USER=xxxxx
PSQL_DB_PASSWORD=xxxxx
PSQL_DB_NAME=xxxx
PSQL_SSL_MODE=require
SERVER_PORT=8080
EOF
```

**Catatan**: Sesuaikan dengan credentials database kamu (Neon/Railway/PostgreSQL lain).

### 5. Build & Run

```bash
sudo docker build -t be-mini-project:latest .
sudo docker run -d --name be-mini-project -p 8080:8080 be-mini-project:latest
```

### 6. Setup Firewall (jalankan dari local)

```bash
gcloud compute firewall-rules create allow-8080 --allow tcp:8080 --source-ranges 0.0.0.0/0
```

### 7. Akses Aplikasi

```
http://[EXTERNAL_IP]:8080/docs/index.html
```

---

## Folder Structure

Here are the folders to consider when adding a new endpoint

```
├── controllers/
│   ├── user_controller.go
│   └── ...
├── models/
│   ├── user_model.go
│   └── ...
├── routers/
|   ├── groups/
│       ├── user.go
|       └── ...
|   └── ...
```
