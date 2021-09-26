# MongoDB CLI Study
This repository has the purpose of studying some Golang concepts and MongoDB driver usage like:
- Basic Go concepts like `context` and project structure
- CLI using Cobra with commands and subcommands
- MongoDB <-> Go communication and mapping

## Usage
Start the mongodb using `docker-compose`

```
docker-compose up -d
```

Run the program and get some cli info
```
go run main.go --help
```

## Available commands
mongodb-cli
- create
- user
  - create
  - delete
  - read-all
  - find <id>
  - update <id>