<h3>How to run migrations</h3>

---

> Build container

docker build -t migrator ./migrator

> Run container migrations

source .env

docker run --network host migrator -path=/migrations/ -database "$POSTGRES_URL" up

> Create migrations

docker run -v ./migrator/migrations:/migrations migrator create -ext sql -dir ./migrations -seq MIGRATIONS_NAME
