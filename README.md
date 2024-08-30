
# Backend Test




## Run Locally

Clone the project

```bash
  git clone https://github.com/SyarifuanEfendi/go-animals.git
```

Go to the project directory

```bash
  cd go-animals
```

create database

```bash
  docker compose up -d --build db
```

Copy .env

```bash
  cp .env.example .env
```

Fetch Dependencies

```bash
  go mod tidy
```

Start Application

```bash
  go run main.go
```


## Run With Docker

Clone the project

```bash
  git clone https://github.com/SyarifuanEfendi/go-animals.git
```

Go to the project directory

```bash
  cd go-animals
```

Docker build

```bash
  docker compose up -d --build
```

Check Containers

```bash
  docker ps
```

## Postman

```
  Import Collection file in directory postman to Postman Application
```

## API Reference

#### Create Animals

```http
  POST /animals
```

Body Request :

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` |  |
| `class` | `string` |  |
| `legs` | `int` |  |

#### List Animals

```http
  GET /animals
```

#### List By Id

```http
  GET /animals/:id
```
#### Update Animal

```http
  PUT /animals/:id
```
Body Request :

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | |
| `class`      | `string` |  |
| `legs`      | `int` |  |

#### Delete Animal

```http
  DELETE /animals/:id
```