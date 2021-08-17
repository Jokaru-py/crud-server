# crud-api-server

Настройки подключения к БД должны находиться в файле config.json в директории ./congif.json
```
{   
    "server": {                         // API
        "host": "127.0.0.1",                   
        "port": "8000"                          
    },    
    "db": {                             // База данных в docker
        "name": "postgres",
        "host": "pg_db",
        "port": 5432,
        "dbname": "postgres",
        "user": "postgres",
        "password": "123"        
    }    
}
```

## Сборка и запуск в docker
```sh
# Сборка
docker-compose up --build
```

## Curl-команды для работы внутри контейнера server
```
curl 127.0.0.1:8000/api/v1/users 'GET'

curl 127.0.0.1:8000/api/v1/users 'POST' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Jhon",
  "age": "22",
  "email": "jhon@mail.ru"
}'

curl 127.0.0.1:8000/api/v1/users 'PATCH' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "Din",
  "age": "44",
  "email": "Din@mail.ru",
  "id": "1"
}'

curl 127.0.0.1:8000/api/v1/users 'DELETE' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "id": "1"
}'
```

# Postgres
Настройки создания БД должны находиться в файле .env в директории ./.env

# Миграция
Данные для миграции должны находиться в директории ./migrations/postgres