<h3>Backend</h3>

---

> Start container

docker compose up --build

> Run migration

docker build -t migrator ./migrator

source .env

docker run --network host migrator -path=/migrations/ -database "$POSTGRES_URL" up

**ALL CREDIENTIALS ARE IN .env**
