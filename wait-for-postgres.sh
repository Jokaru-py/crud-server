#!/bin/sh
# wait-for-postgres.sh

set -e

host="pg_db"
port="5432"
shift
cmd="$@"

>&2 echo "!!!!!!!! Check pg_db for available !!!!!!!!"

until PGPASSWORD=$DB_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
exec $cmd