# My Bookmarks API

## Stack used:
- https://gofiber.io/
- https://gorm.io/
- https://github.com/uber-go/zap
- Postgresql

## Run Postgresql DB in docker

```
docker run --name bookmarksdb -e POSTGRES_PASSWORD=mysecret -e POSTGRES_DB=bookmarks -p 5432:5432 -d postgres
docker container start bookmarksdb
```

### Remove cpntainer
```
docker rm bookinggormdb
```

## Github 

https://github.com/gandra/my-bookmarks-api