# Blogging platform API

RESTful API for creating blog posts. It has CRUD operations and developed in Golang using Gorilla Mux for routig and MongoDB as database. Containerized with Docker and Docker Compose.

## Requirements
- Docker
- Docker Compose
- Golang 1.23.0

## Installation 
1. Clone repository
```sh
git clone github.com/AdlbkA/blogging.git
cd blogging
```
2. Install Go dependencies:
```sh
go mod download
```
3. Create ```.env``` file in root directory and add following variables:
```env
MONGOCONN="mongodb://db:27017"
DBNAME="Blogging"
DBCOLLECTION="posts"
```

## Running
### Using Docker
1. Build and run Docker containers: 
```shell
docker compose up --build
```
2. API will be available at ```http://localhost:8080```

### Without Docker
1. Start MongoDB and create ```Blogging``` database and ```posts``` collection.
2. Run the application:
```sh
go run main.go
```
3. API will be available at ```http://localhost:8080```

## API Endpoints
### Posts
- #### Create Post
  - Method: ```POST```
  - URL: ```/post```
  - Body:
    ```json
    {
    "title": "Post title",
    "content": "Post content",
    "category": "Post category",
    "tags": ["Post", "Tags"]
    }
    ```
  - Response: ```201 Created```
- #### Get All Posts
    - Method: ```GET```
    - URL: ```/post```
    - Response: ```200 OK```
- #### Get Post by ID
    - Method: ```GET```
    - URL: ```/post/{id}```
    - Response: ```200 OK```
- #### Update Post
    - Method: ```PUT```
    - URL: ```/post/{id}```
    - Body:
      ```json
      {
       "title": "Post title2",
       "content": "Post content2",
       "category": "Post category2",
       "tags": ["Post", "Tags", "2"]
      }
      ```
    - Response: ```200 OK```
- #### Delete Post
    - Method: ```DELETE```
    - URL: ```/post/{id}```
    - Response: ```200 OK```
## Environment Variables
- ```MONGOCONN```: MongoDB URI.
- ```DBNAME```: Database name.
- ```DBCOLLECTION```: Database collection.