# GolangFinal

## Content
- [Diagrams](#diagrams)
- [How to run locally](#how-to-run-locally)
- [Auth Service Endpoints](#api-endpoints-auth-service)
- [Blogging Service Endpoints](#api-endpoints-blogging-service)

## Diagrams
ERD Diagram:
![ERD](ERD.png)

Sequence Diagram:

![Sequence](Sequence_Diagram.png)


## How to run locally

1. Create database in postgres with values:
```postgresql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username varchar,
    password varchar
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title varchar,
    content varchar,
    author_id int references users (id) on delete cascade 
)
```
2. Create ```.env``` file in the AuthService ```internal/app/config``` directory with values:
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
And ```.env``` in BloggingService ```internal/app/config``` 
```env
HTTP_CONFIG_ADDR=":8081"
HTTP_REQUEST_TIMEOUT=60
POSTGRES_HOST="postgres"  #Place "localhost" if running locally, leave "postgres" if running with docker
POSTGRES_PORT="5432"
POSTGRES_USER="postgres" #Place your user if running locally, leave "postgres" if running with docker
POSTGRES_PASSWORD="postgres" #Place your password if running locally, leave "postgres" if running with docker
POSTGRES_DB="postgres" #Place your dbname if running locally, leave "postgres" if running with docker
POSTGRES_SSL_MODE="disable"
```

3. In ```BloggingService/pkg/middlewares``` change authServiceUrl to localhost when running locally. Input name of a container of Auth Service when running in Docker

4. In root directory start server by running ```go run ./cmd/app/main.go```

API will be available at http://localhost:8080 (Auth Service) and http://localhost:8081 (Blogging Service)

## API Endpoints Auth Service
- #### Register user
  - Method: ```POST```
  - URL: ```/register```
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
    "username": "Test123",
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
    "username": "Test"
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlRlc3QifQ.1Pf_RQJP7UeXyLfAoM-abeA4M5IOj7ntPBG_mEfXML4"
    }
    ```
## API Endpoints Blogging Service
- #### Create Post
  - Method: ```POST```
  - URL: ```/posts```
  - Body:
  ```json
    {
    "title": "title",
    "content": "content",
    "author_id": 1
    }
   ```
  - Response: ```200 OK```
  - Response body:
    ```json
    {
    "user": {
      "id": 1,
      "title": "title",
      "content": "content",
      "author_id": 1
      }
    }
    ```
- #### Delete Post
  - Method: ```DELETE```
  - URL: ```/posts/:id```
  - Response: ```200 OK```
  - Response body:
    ```json
    {
    "post": {
        "id": 0,
        "title": "",
        "content": "",
        "author_id": 0
      }
    }
    ```
- #### Update Post
  - Method: ```PUT```
  - URL: ```/posts/:id```
  - Response: ```200 OK```
  - Body:
  ```json
    {
    "title": "title2",
    "content": "content2",
    "author_id": 2
    }
  ```
  - Response body:
    ```json
    {
    "post": {
        "id": 1,
        "title": "smth2",
        "content": "smth3",
        "author_id": 1
      }
    }
    ```
- #### Get All Posts
  - Method: ```GET```
  - URL: ```/posts```
  - Response: ```200 OK```
  - Response body:
    ```json
    {
    "posts": [
        {
            "id": 1,
            "title": "smth",
            "content": "smth",
            "author_id": 1
        },
        {
            "id": 2,
            "title": "smth2",
            "content": "smth2",
            "author_id": 1
        },
        {
            "id": 3,
            "title": "smth2",
            "content": "smth2",
            "author_id": 1
        }
      ]
    }
    ```