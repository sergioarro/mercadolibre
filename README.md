# mercadolibre
Operación Fuego de Quasar

Han Solo ha sido recientemente nombrado General de la Alianza Rebelde y busca dar un gran golpe contra el Imperio Galáctico para reavivar la llama de la resistencia.

El servicio de inteligencia rebelde ha detectado un llamado de auxilio de una nave portacarga imperial a la deriva en un campo de asteroides. El manifiesto de la nave es ultra clasificado, pero se rumorea que transporta raciones y armamento para una legión entera.

### Golang [Clean Architecture] REST API 🚀

#### 👨‍💻 Full list what has been used:
* [echo](https://github.com/labstack/echo) - Web framework
* [sqlx](https://github.com/jmoiron/sqlx) - Extensions to database/sql.
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
* [viper](https://github.com/spf13/viper) - Go configuration with fangs
* [go-redis](https://github.com/go-redis/redis) - Type-safe Redis client for Golang
* [zap](https://github.com/uber-go/zap) - Logger
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation
* [jwt-go](https://github.com/dgrijalva/jwt-go) - JSON Web Tokens (JWT)
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
* [minio-go](https://github.com/minio/minio-go) - AWS S3 MinIO Client SDK for Go
* [testify](https://github.com/stretchr/testify) - Testing toolkit
* [gomock](https://github.com/golang/mock) - Mocking framework
* [CompileDaemon](https://github.com/githubnemo/CompileDaemon) - Compile daemon for Go
* [Docker](https://www.docker.com/) - Docker

### Local development usage:
    make tidy
    make local // run all containers
    make migrate_up // Dependencia make local (contenedor de postgres) Crea modelo necesario en Data Base
    make run
    make down-local // stop and rm dockers containers
    make deps-cleancache // go clean -modcache

## Initial tables in postgres
cat ./migrations/01_create_initial_tables.up.sql | docker exec -i api_postgesql psql -U postgres -d satellite_db




http://localhost:5001/api/v1/health

http://localhost:5001/api/v1/location/topsecret

http://localhost:5001/api/v1/location/topsecret_split/:satellite_name