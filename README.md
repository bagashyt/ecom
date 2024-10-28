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

- Require **Setup Environment** and **Setup local Database**
- On terminal run
```sh
make run
```


## Setup with Docker

- Require **Setup Environment** and **Setup local Database**
- On terminal run

```sh
sudo docker build .

sudo docker run -d --rm --env-file .env --network host {Docker Image ID}
```

## Setup with Docker Compose

- Require **Setup Environment**
- On terminal run

```sh
sudo docker compose up -d
```

## Test Api

You can test consume API with Postman or other app.

- Register
End-point: [POST] `http://localhost:8080/v2/register`
Body:
```json
{
    "firstName":"Adam",
    "lastName":"Sanders",
    "email":"adam@mail.com",
    "password":"some-secret-password"
}
```

- Login
End-point: [POST] `http://localhost:8080/v2/login`
Body:
```json
{
    "email":"adam@mail.com",
    "password":"some-secret-password"
}
```

- Get List Products
End-point: [GET] `http://localhost:8080/v2/products`

- Create Product
End-point: [POST] `http://localhost:8080/v2/products`
Headers: "Authorization" "token-generated-from-login"
Body:
```json
{
 "name": "Shirt",
  "description": "Oxford Shirt",
  "image": "image.png",
  "price": 48,
  "quantity": 100
}
```

- Order Chart
End-point: [POST] `http://localhost:8080/v2/cart/checkout`
Headers: "Authorization" "token-generated-from-login"
Body:
```json
{
  "items": [
    {
      "productID":1,
      "quantity": 2
    },
    {
      "productID":2,
      "quantity": 3
    }
  ]
}
```