# API

Welcome to the API repository! This document provides essential information on how to set up, run the application, use Swagger for API documentation, and follow the Commitizen commit convention.

## Table of Contents
- [Getting Started](#getting-started)
- [Running the Application](#running-the-application)
- [Setup env](#setup-env)
- [Using Swagger](#using-swagger)
- [Commit Convention](#commit-convention)
- [Creating Branches](#creating-branches)
- [Tree](#tree)
- [Flow](#flow)

## Getting Started

To get started with this project, follow the instructions below.


## Running the Application

To run the application in Go, follow these steps:

1. Make sure you have Go installed on your system. If not, you can download and install it from the [official Go website](https://golang.org/dl/).

2. Clone the repository:
   ```bash
   $ git clone git@github.com:erwinwahyura/go-boilerplate.git
   ```
   
## Setup env

adjust the value since env.sample for test purpose
```bash
$ cp env.sample .env 
```

## Using Swagger
Generate swagger docs, before we generate the swagger docs lets install the swaggo cli
```bash
$ go install github.com/swaggo/swag/cmd/swag@latest
```

run this command to Generate swagger docs
```bash
$ swag init -g cmd/http/main.go ./docs
```

## Commit Convention
We follow the Commitizen commit convention for version control. When making changes, please use the following format for commit messages:
- https://www.conventionalcommits.org/en/v1.0.0/


Follow the prompts to craft your commit message. Choose the appropriate commit type (e.g., feat, fix, chore, docs) and provide a concise description.

Example commit messages:
```
feat(core): Add user authentication feature
fix(api): Resolve issue with data loading
docs: Update README with new instructions
```

## Creating Branches
To create a new branch for your work, follow this convention:

- Create a branch for a new feature: git checkout -b feat/branch-name
- Create a branch for a bug fix: git checkout -b fix/branch-name

Make sure to replace branch-name with a descriptive name for your feature or bug fix.



## Tree
```
go-boilerplate
├─ app
│  ├─ database
│  │  ├─ mongodb.go
│  │  └─ postgres.go
│  ├─ handler
│  │  ├─ health.go
│  │  └─ user.go
│  ├─ middleware
│  │  └─ middleware.go
│  ├─ model
│  │  ├─ config.go
│  │  └─ user.go
│  ├─ repository
│  │  └─ user.go
│  ├─ route
│  │  └─ route.go
│  └─ service
│     ├─ healthcheck
│     │  └─ health.go
│     ├─ shared
│     │  └─ shared.go
│     └─ user
│        └─ user.go
├─ cmd
│  └─ http
│     └─ main.go
├─ config
│  ├─ config.json.sample
├─ docs
│  ├─ docs.go
│  ├─ swagger.json
│  └─ swagger.yaml
├─ go.mod
├─ go.sum
└─ utils
   └─ utils.go

```

## FLOW
- route -> handler -> service -> repository -> model
