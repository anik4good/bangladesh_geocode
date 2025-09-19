# Codebase Analysis

## 1. Application Overview for new jobs
- Go application using Echo framework
- Provides geographic data API for Bangladesh (divisions, districts, upazilas, unions)
- Uses GORM for database operations
- Structured logging and rate limiting

## 2. Key Components
- Main application (main.go)
- API routes (routes/api.go)
- Division controller (app/http/controllers/DivisionController.go)
- Metrics middleware (middleware/metrics.go)
- Prometheus monitoring
- Grafana visualization

## 3. Architecture
```mermaid
graph TD
    A[Client] --> B[Echo Web Framework]
    B --> C[Rate Limiter]
    B --> D[Logger]
    B --> E[Metrics Middleware]
    B --> F[Division Controller]
    F --> G[GORM Database]
    E --> H[Prometheus]
    H --> I[Grafana]
```
![alt text](image.png)
## 4. Deployment
- Containerized using Docker
- Prometheus for metrics collection
- Grafana for visualization
- Custom Prometheus configuration
