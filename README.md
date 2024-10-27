## E-commerce REST API in Go 

Golang Ecommerce Project based on [Tiago's Complete Backend API in Golang](https://youtu.be/7VLmLOiQ3ck?si=oEIeB-1h5U_chUcJ)

## Setup Environment

- Edit your environment at .env.example file then rename to .env


## Setup local Database

- Install [Mysql on Local Machine](https://dev.mysql.com/doc/mysql-shell/8.4/en/mysql-shell-install.html)
- Create Database
- On project directory terminal run
```sh
make migrate up
```

## Setup with Makefile

- On terminal run
```sh
make run
```


## Setup with Docker

- On terminal run

```sh
sudo docker build .

sudo docker run -d --rm --env-file .env --network host {Docker Image ID}
```

## Setup with Docker Compose

- On terminal run

```sh
sudo docker compose up -d
```
