
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
```bash
export SERVER_URL=https://bdgeo.root2tech.com
```

---

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
