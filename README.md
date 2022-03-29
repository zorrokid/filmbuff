# Running with docker-compose

To build environment 

    docker-compose build

To run environment with docker-compose

    docker-compose up -d

To tear down

    docker-compose down

# Build docker container image

To build docker container image

    docker build -t "film-db-rest-api" .

To run docker container image

    docker run -dp 8080:8080 film-db-rest-api

# To test api

List movies

    curl localhost:8080/movies

Add movie

    curl -v localhost:8080/movies -d '{"name": "Some movie"}