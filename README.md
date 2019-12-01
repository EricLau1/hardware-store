# Golang API REST with MySQL & Docker

- Install
  * Go
  * Docker
  * docker-compose

- Usage: 

```bash
    cd $GOPATH/src

    go get -u github.com/EricLau1/hardware-store

    cd hardware-store

    mkdir data

    sudo docker-compose up -d
```

## Set up test database

```bash
    docker exec -it mysql_docker bash -l

    mysql -u root -p
```

- Copy and paste the sql script: `supertest_data.sql`

## Api endpoint inside the container:

- localhost:5000/products

## Api endpoint out of container:

```bash
    cd $GOPATH/src/hardware-store

    go build -o main .

    ./main -p=8080
```
- localhost:8080/products
