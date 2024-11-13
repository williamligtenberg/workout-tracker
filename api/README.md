# workout-tracker API

## Start the API
  ```bash
  cd api
  go run .
  ```
## Test POST request
```bash
Invoke-WebRequest -Uri "http://localhost:8080/users/1" -Headers @{ "Authorization" = "Token" } -Method "POST"
```