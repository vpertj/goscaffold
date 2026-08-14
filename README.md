# goscaffold
Components used
fiber web framework
https://docs.gofiber.io/

https://learnku.com/docs/gofiber/2.x
```
swag: generates Swagger docs
https://github.com/swaggo/swag documentation guide
```

go install github.com/swaggo/swag/cmd/swag@latest
```
air: used for hot-reload development
go install github.com/air-verse/air@latest
```
```
Database access component xorm db.GetOrm()
https://xorm.io/zh/
```
```
Logging component logrus
https://github.com/sirupsen/logrus
```
```
Database schema generation (requires cgo, install the gcc toolchain)
go install xorm.io/reverse@latest
Generate with:
reverse -f mydbtables.yml
```
```
Config file access config.GetConfig()
https://github.com/jinzhu/configor

```
```
lancet utility library
https://github.com/duke-git/lancet
```
```
resty HTTP client library
https://github.com/go-resty/resty
```
