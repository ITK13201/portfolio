# Portfolio

This is an application for my portfolio site.
You can access it at the following link.

https://www.i-tk.dev

## Components

- frontend
  - Framework: [Next.js](https://nextjs.org/)
  - Design: [Material UI](https://mui.com/)
- backend
  - Web Framework: [Gin](https://github.com/gin-gonic/gin)
  - ORM: [ent](https://github.com/ent/ent)
  - SQL Migration tool: [goose](https://github.com/pressly/goose)

## Install

### 0. Prepare

- docker

### 1. Clone

```shell
git clone https://github.com/ITK13201/portfolio.git
cd portfolio
```

### 2. Init environment variables and create log files

```shell
./bin/init.sh
```

### 3. Build

```shell
docker compose build
```

## Usage

### Start Container

```shell
docker compose up
```

### Stop container

```shell
docker compose down
```

## For details

### Backend README

[./backend/README.md](./backend/README.md)

### Frontend README

[./frontend/README.md](./frontend/README.md)
