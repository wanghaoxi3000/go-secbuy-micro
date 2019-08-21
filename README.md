## 环境准备
### Postgres
```
docker run --name postgres-dev -d --restart=always -p 5432:5432 -v /var/docker/postgres:/var/lib/postgresql/data -e POSTGRES_PASSWORD=postgres123 postgres:11
```

```
docker run --name adminer -d --restart=always -p 10080:8080 adminer
```

```
docker exec -it postgres-dev bash

su potgres
createuser -P -e dev
psql -U postgres
CREATE DATABASE secbuy OWNER dev;
GRANT ALL PRIVILEGES ON DATABASE secbuy to dev;
```

### go
安装 [protoc compiler](https://github.com/google/protobuf/releases)
```
go get -u -v github.com/micro/protoc-gen-micro
go get -u github.com/micro/protobuf/protoc-gen-go
```