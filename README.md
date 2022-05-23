# mercadolibre
Operación Fuego de Quasar

Han Solo ha sido recientemente nombrado General de la Alianza Rebelde y busca dar un gran golpe contra el Imperio Galáctico para reavivar la llama de la resistencia.

El servicio de inteligencia rebelde ha detectado un llamado de auxilio de una nave portacarga imperial a la deriva en un campo de asteroides. El manifiesto de la nave es ultra clasificado, pero se rumorea que transporta raciones y armamento para una legión entera.

## Pre-requisitos
Este proyecto requiere Go >= 1.18.1 Puede descargarse en el siguiente link:
[Downloads - The Go Programming Language (golang.org)](https://golang.org/dl/)

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

## Instalación

Clonar el repositorio

    git clone https://github.com/sergioarro/mercadolibre.git

### Local development usage:
    Para correr y levantar de manera rapida el proyecto en local solo se debe seguir estos pasos :

    make tidy
    make local // run all containers
    make migrate_up // Dependencia make local (contenedor de postgres) Crea tabla necesaria en DB. Esta la opción manual mas abajo
    make run

    make down-local // stop and rm dockers containers
    make deps-cleancache // go clean -modcache

## Initial tables in postgres manual
cat ./migrations/01_create_initial_tables.up.sql | docker exec -i api_postgesql psql -U postgres -d satellite_db

## End-points
GET http://localhost:5001/api/v1/health
   
POST http://localhost:5001/api/v1/location/topsecret
 
     REQUEST : 
        {
            "satellites": [
                {
                    "name": "kenobi",
                    "distance": 150.0,
                    "message": ["este","","","mensaje",""]
                },
                {
                    "name": "skywalker",
                    "distance": 115.5,
                    "message": ["","es","","","privado"]
                },
                {
                    "name": "sato",
                    "distance": 142.7,
                    "message": ["este","","un","",""]
                }

            ]
        }

POST http://localhost:5001/api/v1/location/topsecret_split/:satellite_name

    REQUEST :
    {
      "distance": 180.0,
      "message": ["","","un","","super secreto"]
    }


GET http://localhost:5001/api/v1/location/topsecret_split/