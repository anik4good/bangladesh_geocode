
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


Got it ✅ I’ll extend your README to include **K3s (Kubernetes)** deployment instructions right after the Docker Compose section. I’ll keep it consistent with your style and add an example manifest + `kompose` usage. Here’s the updated part:

---

## 🚀 K3s Deployment Option 1

You can also deploy **GeoAPIBD** on a lightweight Kubernetes cluster such as **K3s**.

### 1. Generate Kubernetes Manifests from Docker Compose

If you already have the `docker-compose.yaml`, you can convert it to Kubernetes YAMLs using **kompose**:

```bash
kompose convert -f docker-compose.yaml -o k3s_deploy/
```

This will generate Kubernetes manifests (Deployments, Services, PVCs) inside the `k3s_deploy/` directory.

### 2. Apply the Manifests on K3s

```powershell
kubectl apply -f k3s_deploy/
deployment.apps/app created
service/app created
persistentvolumeclaim/db-data created
deployment.apps/db created
service/db created
persistentvolumeclaim/logs created
PS C:\Users\anik\Desktop\GO\bangladesh_geocode> k get pods
NAME                   READY   STATUS    RESTARTS   AGE
app-6997c7d995-8brs2   0/1     Pending   0          7s
db-689886f96-hz6rk     0/1     Pending   0          7s
PS C:\Users\anik\Desktop\GO\bangladesh_geocode> k get pods
NAME                   READY   STATUS         RESTARTS   AGE
app-6997c7d995-8brs2   1/1     Running        0          29s
db-689886f96-hz6rk     0/1     Running   0          29s


```


## 🚀 K3s Deployment Option 2

All required Kubernetes manifests are already available inside the [`k3s_deploy/`](./k3s_deploy) folder.

---

## 🚀 Deployment Steps

1. **Apply all manifests**
```sh
   kubectl apply -f k3s_deploy/
```

2. **Verify deployments & services**
```sh
   kubectl get pods
   kubectl get svc
```

---

## 🔑 Image Pull Secrets (if using private registry)

If your Docker image (`anik4good/bangladesh_geocode:latest`) is private, create a secret:

```sh
kubectl create secret docker-registry regcred \
  --docker-server=docker.io \
  --docker-username=<your-docker-username> \
  --docker-password=<your-docker-password> \
  --docker-email=<your-email>
```

Then patch `app-deployment.yaml` to include:

```yaml
spec:
  template:
    spec:
      imagePullSecrets:
      - name: regcred
```

If your image is **public**, you can skip this step.

---

## 🌍 Access the Application

The app is exposed via **NodePort** (`30080`):

```sh
kubectl get svc app
```

Example output:

```
NAME   TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
app    NodePort   10.43.114.105   <none>        1552:30080/TCP   5m
```

Now you can access it from any K3s node:

```sh
curl http://<node-ip>:30080/api/divisions
```

Example:

```sh
curl http://192.168.100.162:30080/api/divisions
```

---

## 🛠 Debugging

* Check logs:

  ```sh
  kubectl logs deploy/app
  ```
* Shell into pod:

  ```sh
  kubectl exec -it deploy/app -- sh
  ```
* Inspect service:

  ```sh
  kubectl describe svc app
  ```

---

## 🧹 Cleanup

To remove everything:

```sh
kubectl delete -f k3s_deploy/
```





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
