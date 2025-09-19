
# GeoAPIBD

**GeoAPIBD** is a RESTful API built with Go, providing structured geographical data for Bangladesh, including divisions, districts, upazilas, and unions. The API is designed to facilitate quick and easy access to location-based data for various applications and services.

---

## 📑 Table of Contents
- [Features](#features)
- [Endpoints](#endpoints)
- [Setup](#setup)
- [Environment Variables](#environment-variables)
- [Usage](#usage)
- [License](#license)

---

## 🌟 Features
- **Comprehensive Geographical Data**: Retrieve data at multiple levels of granularity – from divisions down to unions.
- **High-Performance API**: Developed with **Go** and the **Echo** framework to ensure efficient handling of requests.
- **Environment-Driven URLs**: Dynamic URL generation through environment variables for flexible deployments.

---

## 📌 Endpoints

### Division-Level Endpoints
- **Get All Divisions**
  `GET /api/divisions`
  Returns a list of all divisions.

- **Get a Specific Division**
  `GET /api/division/{division_name}`
  Fetches data for a division by its name.

### District-Level Endpoints
- **Get All Districts**
  `GET /api/districts`
  Returns a list of all districts.

- **Get Districts in a Division**
  `GET /api/division/{division_name}/districts`
  Retrieves districts within a specified division.

### Upazila-Level Endpoints
- **Get Upazilas in a Division**
  `GET /api/division/{division_name}/upazilas`
  Fetches upazilas within a division.

- **Get Upazilas in a District**
  `GET /api/division/{division_name}/district/{district_name}/upazilas`
  Retrieves upazilas within a specified district.

### Union-Level Endpoint
- **Get Unions in an Upazila**
  `GET /api/division/{division_name}/district/{district_name}/upazila/{upazila_name}/unions`
  Fetches unions within a specified upazila.

---

## 🚀 Setup

## 🚀 Docker Compose
```bash
version: '3'

services:
  app:
    image: anik4good/bangladesh_geocode:latest
    environment:
      - CONNECTION=user:password@tcp(db:3306)/bangladesh_geocode?charset=utf8mb4&parseTime=True&loc=Local
      - SERVER_PORT=1552
      - SERVER_URL=http://localhost:1552
    ports:
      - "8112:1552"
    restart: unless-stopped
    volumes:
      - ./logs:/app/logs
    depends_on:
      - db

  db:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: bangladesh_geocode
      MYSQL_USER: user
      MYSQL_PASSWORD: password
    ports:
      - '33061:3306'
    volumes:
      - ./db-data:/var/lib/mysql

volumes:
  db-

```

To set up the project locally, follow these steps:

```bash
# Clone the repository
git clone https://github.com/anik4good/geoapibd.git

# Change into the project directory
cd geoapibd

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

---

## 🔧 Environment Variables

The application relies on environment variables for configuration. Set the following variable:

- `SERVER_URL`: Base URL for the API.

Example:
```ini
CONNECTION=user:password@tcp(db:3306)/bangladesh_geocode?charset=utf8mb4&parseTime=True&loc=Local
SERVER_PORT=1552
SERVER_URL=http://localhost:1552
```

---

## 🗄️ Database Setup

Before running the application, you need to create the MySQL database and import the initial data.

### 1. Create the Database

You can create the database using the MySQL CLI or a GUI tool:

```bash
mysql -u root -p
```
Then run:
```sql
CREATE DATABASE bangladesh_geocode CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'user'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON bangladesh_geocode.* TO 'user'@'%';
FLUSH PRIVILEGES;
EXIT;
```

### 2. Import the SQL File

You have a file named `bangladesh_geocode.sql` in your project directory:

```bash
mysql -u user -p bangladesh_geocode < bangladesh_geocode.sql
```

This will import all the tables and data required for the API.

---

**Note:**
If you are using Docker Compose, the database will be created automatically, but you still need to import the SQL file into the running MySQL container. You can do this with:

```bash
docker exec -i <mysql_container_name> mysql -u user -p bangladesh_geocode < bangladesh_geocode.sql
```
Replace `<mysql_container_name>` with the actual name of your MySQL container (e.g.,


## 🛠️ Usage

Once the server is running, you can test the API using **curl**, **Postman**, or similar tools.

Example request to retrieve all divisions:
```bash
curl https://bdgeo.root2tech.com/api/divisions
```

---

## 📜 License

This project is licensed under the MIT License. For more details, see the [LICENSE](LICENSE) file in the repository.

---

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

---

## 📞 Contact

For any inquiries, please reach out to [your email or contact information here].

---

**GeoAPIBD** - Built with ❤️ and Go
