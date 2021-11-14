module github.com/cnson19700/user_service

go 1.16

require (
	github.com/cnson19700/auth_service v0.0.0-20211114030454-c5a85d117d0d
	github.com/cnson19700/pkg v0.0.0-20211018031701-377617a7e12d
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/joho/godotenv v1.4.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo/v4 v4.6.1
	github.com/pkg/errors v0.9.1
	github.com/pressly/goose v2.7.0+incompatible
	gorm.io/driver/mysql v1.1.3
	gorm.io/gorm v1.22.3
	gorm.io/plugin/dbresolver v1.1.0
)

require (
	github.com/cnson19700/book_service v0.0.0-20211113142847-42b4e4d73c93
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211113001501-0c823b97ae02 // indirect
	google.golang.org/genproto v0.0.0-20211112145013-271947fe86fd // indirect
	google.golang.org/grpc v1.42.0
)
