# workout-tracker

## Run API
  ```bash
  cd api
  go run .
  ```
## Test GET request
```bash
Invoke-WebRequest -Uri "http://localhost:8080/users/1" -Headers @{ "Authorization" = "Token" }
```