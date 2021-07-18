# go-for-it

## setting
```
cd src
go mod init go_test  # --> go.mod
go mod tidy          # --> go.sum
```

## run
```
cd ..
docker compose up -d
```

## add library
### Example of adding mysql
1. Fix Dockerfile.
```
  go get -u github.com/go-sql-driver/mysql
```
2. Run `go mod tidy`.
3. Use in addition to import.
```
import (
	_ "github.com/go-sql-driver/mysql"
)
```
