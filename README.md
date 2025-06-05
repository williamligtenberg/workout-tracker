# workout-tracker API

## Start the API
  ```bash
  cd api
  go run .
  ```

## API endpoints
### Creating an account
 Payload:
  ``` bash
  {"username": "username", "email": "email", "password": "password"}
  ```
  Response:
  ``` bash
  {"id": "id", "payload": "User created successfully", "success": "true"}
  ```
`Status: 201`
### Deleting an account

  Payload:
  ``` bash
  {"username": "username", "email": "email", "password": "password"}
  ```
  Response:
  ``` bash
  {"id": "id", "payload": "User created successfully", "success": "true"}
  ```
`Status: 201`