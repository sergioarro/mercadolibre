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

    1. make tidy
    2. make local // run all containers
    3. make migrate_up // Dependencia make local (contenedor de postgres) Crea tabla necesaria en DB. Esta la opción manual mas abajo
    4. make run

    make down-local // stop and rm dockers containers
    make deps-cleancache // go clean -modcache
    Opcional
    Si se necesita volver a generar la base de datos deben eliminar la carpeta del proyecto "pgdata"
    make down-db // Luego volver a subir el proyecto y aplicar paso 3 make migrate_up

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

    Ademas de agregar a la base de datos el satelite que no tiene registrado este retorna y valida si hay un mensaje completo y su posicion y si aun no la tiene retorna 404 Not enough data

    REQUEST OK:
    {
      "distance": 180.0,
      "message": ["","","un","","super secreto"]
    }

    RESPONSE 404:
    {
        "status": 404,
        "error": "Not enough data"
    }
    
    RESPONSE OK:
    {
        "position": {
            "x": -471.6609125,
            "y": 1525.764225
        },
        "message": "este es un mensaje privado"
    }


GET http://localhost:5001/api/v1/location/topsecret_split/
