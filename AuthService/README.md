# GolangFinal

## How to run locally

1. Create database in postgres with values:
```postgresql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username varchar,
    password varchar,
    role varchar
)
```
2. Create ```.env``` file in the ```internal/app/config``` directory with values:
```env
HTTP_ADDR=":8080"
HTTP_REQUEST_TIMEOUT=60
POSTGRES_HOST="postgres" #Place "localhost" if running locally, leave "postgres" if running with docker
POSTGRES_PORT="5432"
POSTGRES_USER="postgres" #Place your user if running locally, leave "postgres" if running with docker 
POSTGRES_PASSWORD="postgres" #Place your password if running locally, leave "postgres" if running with docker
POSTGRES_DB="postgres" #Place your dbname if running locally, leave "postgres" if running with docker
POSTGRES_SSL_MODE="disable"
JWT_SECRET_KEY=<YOUR_SECRET_KEY>
```
3. In root directory start server by running ```go run ./cmd/app/main.go```

API will be available at http://localhost:8080

## API Endpoints
- #### Register user
  - Method: ```POST```
  - URL: ```/register```
  - Body:
  ```json
    {
    "username": "Username",
    "password": "Password",
    "role": "Author, Student"
    }
   ```
  - Response: ```200 OK```
  - Response body:
    ```json
    {
    "user": {
    "username": "Test123",
    "role": "Student"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlRlc3QxMjMifQ.LQWBueEQFGcjaEayZ9JFT5rEwXgMpzk_ppm2ttYYoaw"
    }
    ```
- ### Login user
    - Method: ```POST```
    - URL: ```/login```
    - Body:
  ```json
    {
    "username": "Username",
    "password": "Password"
    }
   ```
    - Response: ```200 OK```
    - Response body: 
    ```json
    {
    "user": {
    "username": "Test",
    "role": "Author"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlRlc3QifQ.1Pf_RQJP7UeXyLfAoM-abeA4M5IOj7ntPBG_mEfXML4"
    }
    ```

